package coach

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"morning30/internal/app"
	"morning30/internal/content"
	"morning30/internal/db"
)

func setup(t *testing.T) int64 {
	t.Helper()
	d, err := db.Open(filepath.Join(t.TempDir(), "coach.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = d.Close() })
	app.DB = d
	app.SeedBuiltins()
	u, err := d.CreateUser("coach@b.com", "h", "ru", "Аня")
	if err != nil {
		t.Fatal(err)
	}
	return u.ID
}

func TestDaysAndRanges(t *testing.T) {
	uid := setup(t)
	for in, want := range map[string]string{`"all"`: "all", `"4-6"`: "4,5,6", `[1,3]`: "1,3", `7`: "7", `"1-3, 9"`: "1,2,3,9"} {
		var raw any
		_ = json.Unmarshal([]byte(in), &raw)
		d, err := args{"days": raw}.days(uid)
		if err != nil {
			t.Fatalf("%s: %v", in, err)
		}
		got := "all"
		if !d.All {
			parts := make([]string, len(d.List))
			for i, n := range d.List {
				parts[i] = strings.TrimSpace(strings.Repeat(" ", 0) + itoa(n))
			}
			got = strings.Join(parts, ",")
		}
		if got != want {
			t.Errorf("days(%s) = %s, want %s", in, got, want)
		}
	}
	if got := DayRanges([]int{5, 1, 2, 3, 7, 8, 30}); got != "1–3, 5, 7–8, 30" {
		t.Errorf("DayRanges = %q", got)
	}
}

func itoa(n int) string { b, _ := json.Marshal(n); return string(b) }

// Tools edit the current program; the change list comes from the program
// itself; undo brings everything back.
func TestToolsChangesAndUndo(t *testing.T) {
	uid := setup(t)
	before := takeSnapshot(uid)
	s := &session{userID: uid, lang: "ru"}
	day1 := app.UserProgram(uid).Days[0].Items
	from := day1[0].ID
	for _, call := range []struct{ name, args string }{
		{"replace_exercise", `{"from":"` + from + `","to":"P01","days":"1-3"}`},
		{"set_workout", `{"rounds":3,"rest_main":15}`},
		{"rename_program", `{"name":"Утро"}`},
		{"create_exercise", `{"name":"Лодочка с хлопком","part":"main","unit":"reps","description":"Лёжа на животе.","steps":["Лечь","Поднять руки и ноги","Хлопнуть"],"picture_from":"B01"}`},
	} {
		if out := s.run(call.name, call.args); strings.HasPrefix(out, "error") {
			t.Fatalf("%s: %s", call.name, out)
		}
	}
	if !s.edited || len(s.created) != 1 {
		t.Fatalf("session = %+v", s)
	}
	if out := s.run("add_exercise", `{"exercise":"`+s.created[0]+`","part":"main","days":"all","value":8}`); strings.HasPrefix(out, "error") {
		t.Fatal(out)
	}
	if out := s.run("replace_exercise", `{"from":"NOPE","to":"P01"}`); !strings.HasPrefix(out, "error") {
		t.Fatal("unknown exercise must fail")
	}
	changes := Diff(uid, "ru", before, s.created)
	kinds := map[string]Change{}
	for _, c := range changes {
		kinds[c.Kind] = c
	}
	if c := kinds["replace"]; len(c.Days) == 0 || c.Days[0] != 1 || c.Days[len(c.Days)-1] > 3 || c.To == "" || c.From == "" {
		t.Errorf("replace = %+v", c)
	}
	if c := kinds["rounds"]; c.Old != 2 || c.New != 3 {
		t.Errorf("rounds = %+v", c)
	}
	if c := kinds["rest_main"]; c.New != 15 {
		t.Errorf("rest = %+v", c)
	}
	if c := kinds["rename"]; c.To != "Утро" {
		t.Errorf("rename = %+v", c)
	}
	if c := kinds["new"]; c.To != "Лодочка с хлопком" {
		t.Errorf("new = %+v", c)
	}
	if c := kinds["add"]; len(c.Days) != 30 || c.Part != "main" {
		t.Errorf("add = %+v", c)
	}
	if err := Undo(uid, SnapshotJSON(before)); err != nil {
		t.Fatal(err)
	}
	after := takeSnapshot(uid)
	days := func(raw string) string {
		r, _ := content.ParseResolved(raw)
		b, _ := json.Marshal(r.Days)
		return string(b)
	}
	if days(after.Program) != days(before.Program) || after.Rounds != 2 || after.RestMain != before.RestMain || after.Name != before.Name {
		t.Fatalf("undo left changes: rounds %d rest %d name %q same days %v", after.Rounds, after.RestMain, after.Name, days(after.Program) == days(before.Program))
	}
}

// The prompt tells the model where it is and shows the program.
func TestSystemPromptHasContext(t *testing.T) {
	uid := setup(t)
	p := SystemPrompt(uid, "ru")
	for _, want := range []string{"morning30.com", "CURRENT program", "Русский", "Day 1", "warm-up:"} {
		if !strings.Contains(p, want) {
			t.Errorf("prompt has no %q", want)
		}
	}
}

// A whole exchange against a fake OpenAI-compatible server: a tool call, then
// the reply; the main model failing falls back to the second one.
func TestAskWithFakeModel(t *testing.T) {
	uid := setup(t)
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Model    string    `json:"model"`
			Messages []Message `json:"messages"`
			Provider struct {
				Only []string `json:"only"`
			} `json:"provider"`
		}
		_ = json.NewDecoder(r.Body).Decode(&body)
		calls.Add(1)
		if body.Model == "broken" {
			http.Error(w, "gone", http.StatusNotFound)
			return
		}
		last := body.Messages[len(body.Messages)-1]
		if last.Role == "user" {
			_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"","tool_calls":[{"id":"1","type":"function","function":{"name":"set_workout","arguments":"{\"rounds\":4}"}}]}}]}`))
			return
		}
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"Готово: четыре круга."}}],"usage":{"prompt_tokens":10,"completion_tokens":5}}`))
	}))
	defer srv.Close()
	fallback := Config{BaseURL: srv.URL, Key: "k", Model: "good"}
	cfg := Config{BaseURL: srv.URL, Key: "k", Model: "broken", Fallback: &fallback}
	res, _, err := Ask(context.Background(), cfg, uid, "ru", nil, "Хочу четыре круга")
	if err != nil {
		t.Fatal(err)
	}
	if res.Reply != "Готово: четыре круга." || app.DB.GetRounds(uid) != 4 || len(res.Changes) != 1 || res.Changes[0].Kind != "rounds" {
		t.Fatalf("res = %+v rounds %d", res, app.DB.GetRounds(uid))
	}
	if _, _, err := Ask(context.Background(), Config{}, uid, "ru", nil, "x"); err != ErrDisabled {
		t.Fatalf("no key: %v", err)
	}
}

func TestModelSpec(t *testing.T) {
	t.Setenv("AI_API_KEY", "k")
	t.Setenv("AI_MODEL", "qwen/qwen3.7-flash@alibaba")
	t.Setenv("AI_FALLBACK", "none")
	c := FromEnv()
	if c.Model != "qwen/qwen3.7-flash" || len(c.Providers) != 1 || c.Providers[0] != "alibaba" || c.Fallback != nil {
		t.Fatalf("config = %+v", c)
	}
	t.Setenv("AI_MODEL", "")
	t.Setenv("AI_FALLBACK", "")
	c = FromEnv()
	if !strings.HasPrefix(c.Model, "~deepseek/") || c.Providers[0] != "deepseek" || c.Fallback == nil || c.Fallback.Providers[0] != "alibaba" || c.Fallback.Model != "qwen/qwen3.8-flash" {
		t.Fatalf("defaults = %+v", c)
	}
}

func TestCleanReply(t *testing.T) {
	if got := cleanReply("Let me think. </think> **Готово**: три круга."); got != "Готово: три круга." {
		t.Fatalf("cleanReply = %q", got)
	}
}
