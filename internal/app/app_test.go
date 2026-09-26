package app

import (
	"path/filepath"
	"testing"

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
