package mcp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"morning30/internal/app"
	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
)

func setupMCP(t *testing.T) *db.DB {
	t.Helper()
	d, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { d.Close() })
	app.DB = d
	app.SeedBuiltins()
	return d
}

func TestPutExerciseWritesUserDB(t *testing.T) {
	d := setupMCP(t)
	u, _ := d.CreateUser("s@b.com", "h", "en", "Sergey")
	doc, _ := json.Marshal(catalog.Doc{
		ID: "X99", Slot: "main", Unit: "reps", SVG: "P08",
		Content: map[string]catalog.Loc{"en": {Name: "Wall sit hold"}},
	})
	res := callTool("put_exercise", jsonRaw(map[string]any{
		"account":  "Sergey",
		"exercise": json.RawMessage(doc),
	}))
	if res.IsError {
		t.Fatalf("%s", res.Content[0].Text)
	}
	j, _ := d.GetCustomExercises(u.ID)
	if !strings.Contains(j, "X99") {
		t.Fatalf("custom_exercises missing X99: %s", j)
	}
	got, ok := app.UserCatalog(u.ID).Doc("X99")
	if !ok || got.SVG != "P08" {
		t.Fatalf("overlay: ok=%v svg=%q", ok, got.SVG)
	}
}

func TestSetProgramWritesUserDB(t *testing.T) {
	d := setupMCP(t)
	u, _ := d.CreateUser("i@b.com", "h", "ru", "Ilya")
	_ = d.AdvancePosition(u.ID)
	r := content.ResolveDemo()
	r.Name = "Ilya morning"
	res := callTool("set_program", jsonRaw(map[string]any{
		"account": "i@b.com",
		"program": json.RawMessage(r.Marshal()),
	}))
	if res.IsError {
		t.Fatalf("%s", res.Content[0].Text)
	}
	if app.UserProgram(u.ID).Name != "Ilya morning" {
		t.Fatalf("got %q", app.UserProgram(u.ID).Name)
	}
	if p, _ := d.Position(u.ID); p != 1 {
		t.Fatalf("position reset to %d", p)
	}
}

func TestHTTPRequiresBearer(t *testing.T) {
	setupMCP(t)
	h := Handler("secret")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, httptest.NewRequest("GET", "/api/accounts", nil))
	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("code = %d", rr.Code)
	}
	req := httptest.NewRequest("GET", "/api/accounts", nil)
	req.Header.Set("Authorization", "Bearer secret")
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("authed code = %d body=%s", rr.Code, rr.Body.String())
	}
}

func TestUserMCPTokenScoped(t *testing.T) {
	d := setupMCP(t)
	a, _ := d.CreateUser("a@b.com", "h", "en", "Ann")
	b, _ := d.CreateUser("b@b.com", "h", "en", "Bob")
	tokA, err := d.IssueMCPToken(a.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = d.IssueMCPToken(b.ID)
	h := Handler("")
	req := httptest.NewRequest("GET", "/api/accounts", nil)
	req.Header.Set("Authorization", "Bearer "+tokA)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK || !strings.Contains(rr.Body.String(), "Ann") || strings.Contains(rr.Body.String(), "Bob") {
		t.Fatalf("list self: %d %s", rr.Code, rr.Body.String())
	}
	req = httptest.NewRequest("GET", "/api/accounts/"+strconv.FormatInt(b.ID, 10)+"/program", nil)
	req.Header.Set("Authorization", "Bearer "+tokA)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusForbidden {
		t.Fatalf("cross-account: %d %s", rr.Code, rr.Body.String())
	}
	req = httptest.NewRequest("GET", "/api/accounts/me/program", nil)
	req.Header.Set("Authorization", "Bearer "+tokA)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("me: %d %s", rr.Code, rr.Body.String())
	}
}

func TestHandleRPCToolsList(t *testing.T) {
	raw := handleRPC([]byte(`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`))
	if !bytes.Contains(raw, []byte("put_exercise")) || !bytes.Contains(raw, []byte("set_program")) ||
		!bytes.Contains(raw, []byte("put_pose")) || !bytes.Contains(raw, []byte("set_exercise_poses")) {
		t.Fatalf("tools/list: %s", raw)
	}
}

func TestPutPoseAndSetExercisePoses(t *testing.T) {
	d := setupMCP(t)
	png := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="
	res := callTool("put_pose", jsonRaw(map[string]any{
		"id": "plank-high", "kind": "plank", "label": "High plank", "sex": "m", "image": png,
	}))
	if res.IsError {
		t.Fatalf("put_pose: %s", res.Content[0].Text)
	}
	res = callTool("put_pose", jsonRaw(map[string]any{
		"id": "plank-low", "kind": "plank", "label": "Low plank", "sex": "m", "image": png,
	}))
	if res.IsError {
		t.Fatalf("put_pose 2: %s", res.Content[0].Text)
	}
	res = callTool("set_exercise_poses", jsonRaw(map[string]any{
		"id": "P08", "poses": []string{"plank-high", "plank-low"},
	}))
	if res.IsError {
		t.Fatalf("set: %s", res.Content[0].Text)
	}
	ids := app.UserCatalog(0).PoseIDs("P08")
	if len(ids) != 2 {
		t.Fatalf("assigned %#v", ids)
	}
	list := callTool("list_poses", jsonRaw(map[string]any{}))
	if list.IsError || !strings.Contains(list.Content[0].Text, "plank-high") {
		t.Fatalf("list: %+v", list)
	}
	_ = d
}

func jsonRaw(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}

// tokenClient calls the HTTP API with a per-account token.
func tokenClient(t *testing.T, token string) func(method, url, body string) (int, string) {
	h := Handler("")
	return func(method, url, body string) (int, string) {
		req := httptest.NewRequest(method, url, strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		return rr.Code, rr.Body.String()
	}
}

func decode[T any](t *testing.T, body string) T {
	t.Helper()
	var v T
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		t.Fatalf("decode %T: %v in %s", v, err, body)
	}
	return v
}

// Every new feature of the app is reachable with the token from settings:
// catalog, selected programs, switching with per-program days, renaming,
// the day plan, library, similar exercises, add/replace/remove, settings
// with four figures. The token never reaches another account.
func TestTokenCoversEveryProgramFeature(t *testing.T) {
	d := setupMCP(t)
	a, _ := d.CreateUser("a@b.com", "h", "ru", "Ann")
	b, _ := d.CreateUser("b@b.com", "h", "en", "Bob")
	tok, _ := d.IssueMCPToken(a.ID)
	call := tokenClient(t, tok)
	must := func(method, url, body string) string {
		t.Helper()
		code, out := call(method, url, body)
		if code != http.StatusOK {
			t.Fatalf("%s %s: %d %s", method, url, code, out)
		}
		return out
	}

	catalog := decode[[]readyView](t, must("GET", "/api/programs", ""))
	if len(catalog) != len(content.Catalog()) || catalog[0].Key != content.DemoKey || catalog[0].Category != content.CatStart {
		t.Fatalf("catalog: %+v", catalog[:1])
	}
	one := decode[readyView](t, must("GET", "/api/programs/slim30", ""))
	if one.Program == nil || !one.Tags.Jumps || one.Minutes[0] == 0 {
		t.Fatalf("program detail: %+v", one)
	}
	if code, _ := call("GET", "/api/programs/nope", ""); code != http.StatusNotFound {
		t.Fatalf("unknown program: %d", code)
	}

	// Twelve days on the morning workout, then plank: plank starts at day 1.
	for i := 0; i < 12; i++ {
		_ = d.AdvancePosition(a.ID)
	}
	v := decode[programView](t, must("POST", "/api/accounts/me/select", `{"key":"plank30"}`))
	if v.Day != 1 || v.Program.Name != "Планка за 30 дней" || len(v.Selected) != 2 {
		t.Fatalf("select plank: day %d name %q selected %d", v.Day, v.Program.Name, len(v.Selected))
	}
	must("POST", "/api/accounts/me/selected", `{"key":"abs30"}`)
	sel := decode[[]app.SelectedView](t, must("GET", "/api/accounts/me/selected", ""))
	if len(sel) != 3 {
		t.Fatalf("selected: %+v", sel)
	}

	// Edit plank: add, replace, remove — on specific days and every day.
	must("POST", "/api/accounts/me/program/exercises", `{"exercise":"B01","part":"main","days":"today","value":25}`)
	day := decode[app.DayPlan](t, must("GET", "/api/accounts/me/day", ""))
	found := false
	for _, it := range day.Items {
		found = found || (it.ID == "B01" && it.Part == "main")
	}
	if !found || day.Day != 1 {
		t.Fatalf("added exercise missing from today: %+v", day)
	}
	sim := decode[app.Similar](t, must("GET", "/api/accounts/me/exercises/P01/similar", ""))
	if len(sim.Easier) == 0 || len(sim.Harder) == 0 {
		t.Fatalf("similar: %+v", sim)
	}
	must("POST", "/api/accounts/me/program/replace", `{"from":"C01","to":"C02","days":[1,2,3]}`)
	must("POST", "/api/accounts/me/program/remove", `{"exercise":"C09","part":"main","days":"all"}`)
	if code, out := call("POST", "/api/accounts/me/program/remove", `{"exercise":"ZZZ"}`); code != http.StatusNotFound {
		t.Fatalf("remove missing: %d %s", code, out)
	}
	must("PUT", "/api/accounts/me/program/name", "Моя планка")

	// The edited plank needs consent before removal.
	sel = decode[[]app.SelectedView](t, must("GET", "/api/accounts/me/selected", ""))
	var plankID int64
	for _, s := range sel {
		if s.Source == "plank30" {
			plankID = s.ID
			if !s.Edited || s.Name != "Моя планка" || !s.Current {
				t.Fatalf("edited plank view: %+v", s)
			}
		}
	}
	if code, _ := call("DELETE", "/api/accounts/me/selected/"+strconv.FormatInt(plankID, 10), ""); code != http.StatusConflict {
		t.Fatalf("edited program removed without consent: %d", code)
	}

	// Back to the morning workout: its day 13 is still there.
	v = decode[programView](t, must("POST", "/api/accounts/me/restore-demo", ""))
	if v.Day != 13 {
		t.Fatalf("morning workout day %d, want 13", v.Day)
	}
	// Library with filters; settings with a child figure.
	lib := decode[[]app.ExerciseInfo](t, must("GET", "/api/accounts/me/library?group=cardio&noJumps=1", ""))
	for _, e := range lib {
		if e.Group != "cardio" || e.Jumps {
			t.Fatalf("library filter leaked %+v", e)
		}
	}
	st := decode[app.Settings](t, must("PATCH", "/api/accounts/me/settings", `{"figure":"g","restMain":25,"level":2}`))
	if st.Figure != "g" || st.RestMain != 25 || st.Level != 2 || st.Lang != "ru" {
		t.Fatalf("settings: %+v", st)
	}
	if code, _ := call("PATCH", "/api/accounts/me/settings", `{"figure":"x"}`); code != http.StatusBadRequest {
		t.Fatalf("bad figure accepted: %d", code)
	}

	// Import adds a new program; the token cannot touch Bob.
	imp := content.ResolveDemo()
	imp.Name = "Агент"
	imp.Days[0].Items[0].Value = 21
	v = decode[programView](t, must("POST", "/api/accounts/me/programs/import", imp.Marshal()))
	if v.Program.Name != "Агент" || v.Day != 1 {
		t.Fatalf("import: %q day %d", v.Program.Name, v.Day)
	}
	bob := "/api/accounts/" + strconv.FormatInt(b.ID, 10)
	for _, url := range []string{bob + "/select", bob + "/selected", bob + "/program/exercises"} {
		if code, _ := call("POST", url, `{"key":"abs30","exercise":"C01","part":"main"}`); code != http.StatusForbidden {
			t.Fatalf("%s reached another account: %d", url, code)
		}
	}
	if app.HasCustomProgram(b.ID) {
		t.Fatal("another account was changed")
	}
}

// The same operations through the stdio tools, including the older names.
func TestProgramToolsOverRPC(t *testing.T) {
	d := setupMCP(t)
	u, _ := d.CreateUser("c@b.com", "h", "de", "Cara")
	_ = d.AdvancePosition(u.ID)

	raw := handleRPC([]byte(`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`))
	for _, name := range []string{"list_programs", "get_ready_program", "list_selected", "select_program", "star_program",
		"unselect_program", "rename_program", "set_program", "import_program", "get_day", "add_exercise",
		"replace_exercise", "remove_exercise", "list_exercises", "similar_exercises", "get_settings", "set_settings",
		"use_course", "list_courses", "get_course", "list_saved_programs", "restore_saved_program", "put_pose"} {
		if !bytes.Contains(raw, []byte(`"`+name+`"`)) {
			t.Fatalf("tools/list lacks %s", name)
		}
	}
	res := callTool("use_course", jsonRaw(map[string]any{"account": "c@b.com", "course": "pushups30"}))
	if res.IsError || app.UserProgram(u.ID).Name != "Liegestütze in 30 Tagen" {
		t.Fatalf("use_course: %s", res.Content[0].Text)
	}
	res = callTool("add_exercise", jsonRaw(map[string]any{"account": "c@b.com", "exercise": "A01", "part": "main", "days": 3, "value": 8}))
	if res.IsError {
		t.Fatalf("add_exercise: %s", res.Content[0].Text)
	}
	res = callTool("get_day", jsonRaw(map[string]any{"account": "c@b.com", "day": 3}))
	if res.IsError || !strings.Contains(res.Content[0].Text, `"A01"`) {
		t.Fatalf("get_day: %s", res.Content[0].Text)
	}
	res = callTool("select_program", jsonRaw(map[string]any{"account": "c@b.com", "key": "demo"}))
	if res.IsError || app.CurrentDay(u.ID) != 2 {
		t.Fatalf("back to the morning workout on day %d: %s", app.CurrentDay(u.ID), res.Content[0].Text)
	}
	res = callTool("list_saved_programs", jsonRaw(map[string]any{"account": "c@b.com"}))
	if res.IsError || !strings.Contains(res.Content[0].Text, "Liegestütze") {
		t.Fatalf("list_saved_programs: %s", res.Content[0].Text)
	}
	res = callTool("unselect_program", jsonRaw(map[string]any{"account": "c@b.com", "key": "pushups30"}))
	if !res.IsError {
		t.Fatal("edited program removed without consent")
	}
	res = callTool("unselect_program", jsonRaw(map[string]any{"account": "c@b.com", "key": "pushups30", "discardEdits": true}))
	if res.IsError {
		t.Fatalf("unselect: %s", res.Content[0].Text)
	}
	res = callTool("set_settings", jsonRaw(map[string]any{"account": "c@b.com", "voiceMode": "min"}))
	if res.IsError || !strings.Contains(res.Content[0].Text, `"voiceMode": "min"`) {
		t.Fatalf("set_settings: %s", res.Content[0].Text)
	}
}

func TestPosesForFourFigures(t *testing.T) {
	d := setupMCP(t)
	png := []byte("\x89PNG\r\n\x1a\n0000")
	if err := d.PutPose("squat", "squat", "Squat", "m", png); err != nil {
		t.Fatal(err)
	}
	girl := append([]byte(nil), png...)
	girl = append(girl, 'g')
	if err := d.PutPose("squat", "squat", "Squat", "g", girl); err != nil {
		t.Fatal(err)
	}
	p, _ := d.GetPose("squat")
	if !p.HasM || p.HasF || p.HasB || !p.HasG {
		t.Fatalf("figures: %+v", p)
	}
	for sex, want := range map[string][]byte{"g": girl, "b": png, "f": png, "m": png} {
		if _, got, err := d.PoseImage("squat", sex); err != nil || !bytes.Equal(got, want) {
			t.Fatalf("figure %s: %v", sex, err)
		}
	}
}
