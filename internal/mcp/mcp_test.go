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

// A token issued in settings drives the whole course flow for its own account:
// list courses, start one, read it back, rename it, edit it, switch to another.
// Cycle position and other accounts are never touched.
func TestUserTokenRunsCourseFlowOverHTTP(t *testing.T) {
	d := setupMCP(t)
	a, _ := d.CreateUser("a@b.com", "h", "ru", "Ann")
	b, _ := d.CreateUser("b@b.com", "h", "en", "Bob")
	for i := 0; i < 3; i++ {
		_ = d.AdvancePosition(a.ID)
	}
	tok, err := d.IssueMCPToken(a.ID)
	if err != nil {
		t.Fatal(err)
	}
	h := Handler("")
	call := func(method, url, body string) (int, string) {
		req := httptest.NewRequest(method, url, strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+tok)
		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		return rr.Code, rr.Body.String()
	}
	view := func(body string) programView {
		var v programView
		if err := json.Unmarshal([]byte(body), &v); err != nil {
			t.Fatalf("bad program view: %v in %s", err, body)
		}
		return v
	}

	code, body := call("GET", "/api/courses", "")
	var courses []courseView
	if code != http.StatusOK || json.Unmarshal([]byte(body), &courses) != nil || len(courses) != len(content.Series30()) {
		t.Fatalf("courses: %d %s", code, body)
	}
	if courses[0].Key != "plank30" || courses[0].Titles["ru"] != "Планка за 30 дней" || courses[0].Program != nil {
		t.Fatalf("course list entry: %+v", courses[0])
	}
	code, body = call("GET", "/api/courses/plank30", "")
	var one courseView
	if code != http.StatusOK || json.Unmarshal([]byte(body), &one) != nil || one.Program == nil || len(one.Program.Days) != content.TotalDays {
		t.Fatalf("course detail: %d %s", code, body)
	}
	if code, _ = call("GET", "/api/courses/nope30", ""); code != http.StatusNotFound {
		t.Fatalf("unknown course detail: %d", code)
	}

	// Start a course: titled in the account's language, position kept.
	code, body = call("POST", "/api/accounts/me/course/plank30", "")
	v := view(body)
	if code != http.StatusOK || !v.Personal || v.Program.Name != "Планка за 30 дней" || !v.Account.HasProgram {
		t.Fatalf("use course: %d %s", code, body)
	}
	if code, _ = call("POST", "/api/accounts/me/course/nope30", ""); code != http.StatusNotFound {
		t.Fatalf("unknown course: %d", code)
	}
	if code, _ = call("POST", "/api/accounts/"+strconv.FormatInt(b.ID, 10)+"/course/abs30", ""); code != http.StatusForbidden {
		t.Fatalf("cross-account course: %d", code)
	}

	// Rename it, as plain text and as JSON.
	code, body = call("PUT", "/api/accounts/me/program/name", "Моя планка")
	if code != http.StatusOK || view(body).Program.Name != "Моя планка" {
		t.Fatalf("rename text: %d %s", code, body)
	}
	code, body = call("PUT", "/api/accounts/me/program/name", `{"name": "Планка по утрам"}`)
	if code != http.StatusOK || view(body).Program.Name != "Планка по утрам" {
		t.Fatalf("rename json: %d %s", code, body)
	}
	if code, _ = call("PUT", "/api/accounts/me/program/name", "   "); code != http.StatusBadRequest {
		t.Fatalf("empty rename: %d", code)
	}

	// The renamed course reads back as an ordinary program and can be edited
	// and uploaded again through the existing endpoint.
	code, body = call("GET", "/api/accounts/me/program", "")
	v = view(body)
	if code != http.StatusOK || v.Program.Name != "Планка по утрам" || len(v.Program.Days) != content.TotalDays {
		t.Fatalf("get program: %d %s", code, body)
	}
	v.Program.Days[0].Items[0].Value = 45
	code, body = call("PUT", "/api/accounts/me/program", v.Program.Marshal())
	v = view(body)
	if code != http.StatusOK || v.Program.Days[0].Items[0].Value != 45 || v.Program.Name != "Планка по утрам" {
		t.Fatalf("put edited course: %d %s", code, body)
	}

	// Switch to another course with an explicit title language.
	code, body = call("POST", "/api/accounts/me/course/abs30?lang=en", "")
	if code != http.StatusOK || view(body).Program.Name != "Abs in 30 Days" {
		t.Fatalf("switch course: %d %s", code, body)
	}

	if p, _ := d.Position(a.ID); p != 3 {
		t.Fatalf("position moved to %d", p)
	}
	if app.HasCustomProgram(b.ID) {
		t.Fatal("another account was changed")
	}
}

// The same course flow through the stdio tools.
func TestCourseToolsOverRPC(t *testing.T) {
	d := setupMCP(t)
	u, _ := d.CreateUser("c@b.com", "h", "de", "Cara")
	_ = d.AdvancePosition(u.ID)

	raw := handleRPC([]byte(`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`))
	for _, name := range []string{"list_courses", "get_course", "use_course", "rename_program"} {
		if !bytes.Contains(raw, []byte(`"`+name+`"`)) {
			t.Fatalf("tools/list lacks %s", name)
		}
	}
	res := callTool("list_courses", jsonRaw(map[string]any{}))
	if res.IsError || !strings.Contains(res.Content[0].Text, "slim30") {
		t.Fatalf("list_courses: %s", res.Content[0].Text)
	}
	res = callTool("get_course", jsonRaw(map[string]any{"course": "slim30"}))
	if res.IsError || !strings.Contains(res.Content[0].Text, `"totalDays": 30`) {
		t.Fatalf("get_course: %s", res.Content[0].Text)
	}
	res = callTool("use_course", jsonRaw(map[string]any{"account": "c@b.com", "course": "pushups30"}))
	if res.IsError || app.UserProgram(u.ID).Name != "Liegestütze in 30 Tagen" {
		t.Fatalf("use_course: %s", res.Content[0].Text)
	}
	res = callTool("rename_program", jsonRaw(map[string]any{"account": "c@b.com", "name": "Mein Plan"}))
	if res.IsError || app.UserProgram(u.ID).Name != "Mein Plan" {
		t.Fatalf("rename_program: %s", res.Content[0].Text)
	}
	if res = callTool("use_course", jsonRaw(map[string]any{"account": "c@b.com", "course": "nope30"})); !res.IsError {
		t.Fatal("unknown course accepted")
	}
	if res = callTool("rename_program", jsonRaw(map[string]any{"account": "c@b.com", "name": ""})); !res.IsError {
		t.Fatal("empty name accepted")
	}
	if app.UserProgram(u.ID).Name != "Mein Plan" {
		t.Fatal("failed calls changed the program")
	}
	if p, _ := d.Position(u.ID); p != 1 {
		t.Fatalf("position moved to %d", p)
	}
}

// An agent upload cannot wipe a user's own program either: it lands in the
// saved list, which the same token can read and restore from.
func TestAgentUploadKeepsOwnProgramAndCanRestoreIt(t *testing.T) {
	d := setupMCP(t)
	a, _ := d.CreateUser("a@b.com", "h", "en", "Ann")
	b, _ := d.CreateUser("b@b.com", "h", "en", "Bob")
	mine := content.ResolveDemo()
	mine.Name = "Ann's plan"
	mine.Days[0].Items[0].Value = 17
	if err := app.SetUserProgram(a.ID, mine.Marshal()); err != nil {
		t.Fatal(err)
	}
	want := app.UserProgram(a.ID).Marshal()
	tok, _ := d.IssueMCPToken(a.ID)
	tokB, _ := d.IssueMCPToken(b.ID)
	h := Handler("")
	call := func(token, method, url, body string) (int, string) {
		req := httptest.NewRequest(method, url, strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()
		h.ServeHTTP(rr, req)
		return rr.Code, rr.Body.String()
	}

	upload := content.ResolveDemo()
	upload.Name = "Agent plan"
	upload.Days[1].Items[0].Value = 31
	if code, body := call(tok, "PUT", "/api/accounts/me/program", upload.Marshal()); code != http.StatusOK {
		t.Fatalf("upload: %d %s", code, body)
	}
	code, body := call(tok, "GET", "/api/accounts/me/saved-programs", "")
	var saved []db.SavedProgram
	if code != http.StatusOK || json.Unmarshal([]byte(body), &saved) != nil || len(saved) != 1 || saved[0].Name != "Ann's plan" {
		t.Fatalf("saved list: %d %s", code, body)
	}
	id := strconv.FormatInt(saved[0].ID, 10)
	if code, _ = call(tokB, "POST", "/api/accounts/me/saved-programs/"+id+"/restore", ""); code != http.StatusNotFound {
		t.Fatalf("another account restored it: %d", code)
	}
	if code, body = call(tok, "POST", "/api/accounts/me/saved-programs/"+id+"/restore", ""); code != http.StatusOK {
		t.Fatalf("restore: %d %s", code, body)
	}
	if app.UserProgram(a.ID).Marshal() != want {
		t.Fatal("restored program differs from the original")
	}

	// The same through the stdio tools: the agent's plan is now the saved one.
	res := callTool("list_saved_programs", jsonRaw(map[string]any{"account": "a@b.com"}))
	if res.IsError || !strings.Contains(res.Content[0].Text, "Agent plan") {
		t.Fatalf("list_saved_programs: %s", res.Content[0].Text)
	}
	rows, _ := d.SavedPrograms(a.ID)
	res = callTool("restore_saved_program", jsonRaw(map[string]any{"account": "a@b.com", "id": rows[0].ID}))
	if res.IsError || app.UserProgram(a.ID).Name != "Agent plan" {
		t.Fatalf("restore_saved_program: %s", res.Content[0].Text)
	}
	if rows, _ = d.SavedPrograms(a.ID); len(rows) != 1 || rows[0].Name != "Ann's plan" {
		t.Fatalf("after swap: %+v", rows)
	}
	if app.HasCustomProgram(b.ID) {
		t.Fatal("another account was changed")
	}
}
