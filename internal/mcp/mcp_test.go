package mcp

import (
	"bytes"
	"encoding/json"
	"path/filepath"
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

func decode[T any](t *testing.T, body string) T {
	t.Helper()
	var v T
	if err := json.Unmarshal([]byte(body), &v); err != nil {
		t.Fatalf("decode %T: %v in %s", v, err, body)
	}
	return v
}

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
