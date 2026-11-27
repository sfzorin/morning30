package app

import (
	"path/filepath"
	"testing"

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

var tinyPNG = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
	0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4, 0x89, 0x00, 0x00, 0x00,
	0x0a, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0x63, 0x00, 0x01, 0x00, 0x00,
	0x05, 0x00, 0x01, 0x0d, 0x0a, 0x2d, 0xb4, 0x00, 0x00, 0x00, 0x00, 0x49,
	0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

func TestSeedBuiltinsKeepsPoseAssignment(t *testing.T) {
	useArtFiles(t, nil)
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

}

func TestExerciseAnimationSurvivesSeedingAndRequiresBothStrips(t *testing.T) {
	useArtFiles(t, nil)
	d := openAppDB(t)
	SeedBuiltins()
	doc, _ := UserCatalog(0).Doc("P08")
	doc.Animation = &catalog.Animation{Pose: "motion", OtherPose: "other", Frames: 3, Kind: "movement"}
	if err := d.SetExercise("P08", doc.Compact()); err != nil {
		t.Fatal(err)
	}
	SeedBuiltins()
	cat := UserCatalog(0)
	saved, _ := cat.Doc("P08")
	if saved.Animation == nil || *saved.Animation != *doc.Animation {
		t.Fatal("seeding discarded the animation")
	}
	custom := catalog.Doc{ID: "X01", Animation: saved.Animation, Poses: []string{"motion"}}
	cat.Put(custom)
	if err := d.PutPose("motion", "demo", "Movement", "m", tinyPNG); err != nil {
		t.Fatal(err)
	}
	if ExerciseAnimation(cat, "X01", "m") != nil {
		t.Fatal("published an incomplete demonstration")
	}
	if err := d.PutPose("other", "demo", "Other side", "m", tinyPNG); err != nil {
		t.Fatal(err)
	}
	a := ExerciseAnimation(cat, "X01", "m")
	if a == nil || a.Src != "/poses/motion?sex=m" || a.OtherSrc != "/poses/other?sex=m" || a.Frames != 3 {
		t.Fatalf("animation: %#v", a)
	}
	if len(ExerciseFrames(cat, "X01", "m")) == 0 {
		t.Fatal("animation removed still fallback")
	}
}

func TestExerciseFramesCanAwaitCustomIllustrations(t *testing.T) {
	useArtFiles(t, nil)
	openAppDB(t)
	cat := catalog.New(`{"X01":{"id":"X01","slot":"main","unit":"seconds","content":{"en":{"name":"Walking"}}}}`)
	if frames := ExerciseFrames(cat, "X01", "m"); len(frames) != 0 {
		t.Fatalf("unillustrated exercise must not show a broken or unrelated picture: %#v", frames)
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
