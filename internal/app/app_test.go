package app

import (
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
)

func openAppDB(t *testing.T) *db.DB {
	t.Helper()
	d, err := db.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { d.Close() })
	DB = d
	return d
}

func TestSnapshotLegacyNamedKeepsPosition(t *testing.T) {
	d := openAppDB(t)
	u, err := d.CreateUser("ilya@example.com", "h", "ru", "Ilya")
	if err != nil {
		t.Fatal(err)
	}
	if err := d.SetActiveProgram(u.ID, "ilya"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 11; i++ {
		if err := d.AdvancePosition(u.ID); err != nil {
			t.Fatal(err)
		}
	}
	_ = d.RecordActivity(u.ID, "2026-09-01")

	SnapshotLegacy(u.ID)

	if k, _ := d.ActiveProgram(u.ID); k != "custom" {
		t.Fatalf("active_program = %q, want custom", k)
	}
	raw, _ := d.GetProgramJSON(u.ID)
	r, err := content.ParseResolved(raw)
	if err != nil {
		t.Fatal(err)
	}
	if r.Name != "Ilya" {
		t.Fatalf("program name = %q, want Ilya", r.Name)
	}
	if p, _ := d.Position(u.ID); p != 11 {
		t.Fatalf("position = %d, want 11", p)
	}
	if has, _ := d.HasActivity(u.ID, "2026-09-01"); !has {
		t.Fatal("activity was cleared")
	}
}

func TestSnapshotSergeyFromBuiltinRegistered(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("s@example.com", "h", "ru", "Sergey")
	_ = d.SetActiveProgram(u.ID, "builtin")
	for i := 0; i < 4; i++ {
		_ = d.AdvancePosition(u.ID)
	}
	SnapshotLegacy(u.ID)
	raw, _ := d.GetProgramJSON(u.ID)
	r, err := content.ParseResolved(raw)
	if err != nil {
		t.Fatal(err)
	}
	if r.Name != "Sergey" {
		t.Fatalf("program name = %q, want Sergey", r.Name)
	}
	if p, _ := d.Position(u.ID); p != 4 {
		t.Fatalf("position reset to %d", p)
	}
}

func TestSnapshotSkipsGuestAndDemo(t *testing.T) {
	d := openAppDB(t)
	g, _ := d.CreateGuest("guest:x", "en")
	_ = d.SetActiveProgram(g.ID, "builtin")
	SnapshotLegacy(g.ID)
	if HasCustomProgram(g.ID) {
		t.Fatal("guest on builtin was snapshotted")
	}

	u, _ := d.CreateUser("new@example.com", "h", "en", "New")
	SnapshotLegacy(u.ID)
	if HasCustomProgram(u.ID) {
		t.Fatal("new demo account was snapshotted")
	}
	if k := ActiveProgramKey(u.ID); k != "demo" {
		t.Fatalf("new account key = %q, want demo", k)
	}
}

func TestSnapshotDoesNotOverwritePersonalFile(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("own@example.com", "h", "en", "Own")
	demo := content.ResolveDemo()
	demo.Name = "Keep me"
	if err := SetUserProgram(u.ID, demo.Marshal()); err != nil {
		t.Fatal(err)
	}
	_ = d.SetActiveProgram(u.ID, "vlad")
	SnapshotLegacy(u.ID)
	raw, _ := d.GetProgramJSON(u.ID)
	r, _ := content.ParseResolved(raw)
	if r.Name != "Keep me" {
		t.Fatalf("personal file overwritten: %q", r.Name)
	}
}

func TestSetUserProgramAndRestoreDemo(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("p@example.com", "h", "en", "P")
	_ = d.AdvancePosition(u.ID)
	vlad := content.LegacyNamed
	r, ok := vlad("vlad")
	if !ok {
		t.Fatal("vlad missing")
	}
	if err := SetUserProgram(u.ID, r.Marshal()); err != nil {
		t.Fatal(err)
	}
	if UserProgram(u.ID).Name != "Vlad" {
		t.Fatalf("got %q", UserProgram(u.ID).Name)
	}
	if err := RestoreDemo(u.ID); err != nil {
		t.Fatal(err)
	}
	if k := ActiveProgramKey(u.ID); k != "demo" {
		t.Fatalf("key = %q", k)
	}
	if HasCustomProgram(u.ID) {
		t.Fatal("personal file still set")
	}
	if p, _ := d.Position(u.ID); p != 1 {
		t.Fatalf("position reset to %d", p)
	}
}

var tinyPNG = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
	0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4, 0x89, 0x00, 0x00, 0x00,
	0x0a, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0x63, 0x00, 0x01, 0x00, 0x00,
	0x05, 0x00, 0x01, 0x0d, 0x0a, 0x2d, 0xb4, 0x00, 0x00, 0x00, 0x00, 0x49,
	0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

func TestSeedBuiltinsKeepsPoseAssignment(t *testing.T) {
	d := openAppDB(t)
	SeedBuiltins()
	if err := d.PutPose("plank-high", "plank", "High", "m", tinyPNG); err != nil {
		t.Fatal(err)
	}
	if err := d.PutPose("plank-low", "plank", "Low", "m", tinyPNG); err != nil {
		t.Fatal(err)
	}
	if err := SetLibraryPoses("P08", []string{"plank-high", "plank-low"}); err != nil {
		t.Fatal(err)
	}
	SeedBuiltins()
	cat := UserCatalog(0)
	ids := cat.PoseIDs("P08")
	if len(ids) != 2 || ids[0] != "plank-high" || ids[1] != "plank-low" {
		t.Fatalf("poses after seed: %#v", ids)
	}
	frames := ExerciseFrames(cat, "P08", "m")
	if len(frames) != 2 || !strings.Contains(frames[0], "/poses/plank-high") {
		t.Fatalf("frames: %#v", frames)
	}
}

func TestExerciseFramesFallsBackToFigure(t *testing.T) {
	d := openAppDB(t)
	SeedBuiltins()
	_ = d
	frames := ExerciseFrames(UserCatalog(0), "P08", "f")
	if len(frames) != 1 || !strings.Contains(frames[0], "/static/ex/f/P08.png") {
		t.Fatalf("fallback: %#v", frames)
	}
}

func TestSeedPoseStillsKeepsProgressAndManualPoses(t *testing.T) {
	d := openAppDB(t)
	SeedBuiltins()
	u, _ := d.CreateUser("seed@b.com", "h", "en", "Seed")
	if err := d.RecordActivity(u.ID, "2026-09-01"); err != nil {
		t.Fatal(err)
	}
	fsys := fstest.MapFS{
		"body/out/P08-1.png": {Data: tinyPNG},
		"body/out/P08-2.png": {Data: append([]byte{}, tinyPNG...)},
	}
	if err := SeedPoseStills(fsys); err != nil {
		t.Fatal(err)
	}
	ids := UserCatalog(0).PoseIDs("P08")
	if len(ids) != 2 || ids[0] != "p08-1" || ids[1] != "p08-2" {
		t.Fatalf("assigned %#v", ids)
	}
	frames := ExerciseFrames(UserCatalog(0), "P08", "f")
	if len(frames) != 2 || !strings.Contains(frames[0], "/poses/p08-1") {
		t.Fatalf("frames %#v", frames)
	}
	if ok, err := d.HasActivity(u.ID, "2026-09-01"); err != nil || !ok {
		t.Fatalf("activity after seed: %v %v", ok, err)
	}
	if err := SetLibraryPoses("P08", []string{"p08-1"}); err != nil {
		t.Fatal(err)
	}
	if err := SeedPoseStills(fsys); err != nil {
		t.Fatal(err)
	}
	ids = UserCatalog(0).PoseIDs("P08")
	if len(ids) != 1 || ids[0] != "p08-1" {
		t.Fatalf("manual poses overwritten: %#v", ids)
	}
	if frames := ExerciseFrames(UserCatalog(0), "W01", "m"); len(frames) != 1 || !strings.Contains(frames[0], "/static/ex/") {
		t.Fatalf("unseeded figure: %#v", frames)
	}
}

func TestAssignExercisePosesCustomOverlay(t *testing.T) {
	d := openAppDB(t)
	SeedBuiltins()
	u, _ := d.CreateUser("p@b.com", "h", "en", "Pat")
	_ = d.PutPose("sit", "", "Sit", "m", tinyPNG)
	if err := SaveCustomExercise(u.ID, catalog.Doc{
		ID: "X01", Slot: "main", Unit: "reps", SVG: "P08",
		Content: map[string]catalog.Loc{"en": {Name: "Wall sit"}},
	}); err != nil {
		t.Fatal(err)
	}
	if err := AssignExercisePoses(u.ID, "X01", []string{"sit"}); err != nil {
		t.Fatal(err)
	}
	got, _ := UserCatalog(u.ID).Doc("X01")
	if len(got.Poses) != 1 || got.Poses[0] != "sit" {
		t.Fatalf("overlay poses: %#v", got.Poses)
	}
}
