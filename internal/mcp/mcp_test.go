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
	if !bytes.Contains(raw, []byte("put_exercise")) || !bytes.Contains(raw, []byte("set_program")) {
		t.Fatalf("tools/list: %s", raw)
	}
}

func jsonRaw(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}
