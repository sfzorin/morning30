package db

import (
	"testing"
	"time"
)

// Deleting an account removes the user and everything tied to it, and leaves
// other accounts alone.
func TestDeleteUserRemovesAllData(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("gone@example.com", "hash", "ru", "Gone")
	if err != nil {
		t.Fatal(err)
	}
	other, err := d.CreateUser("stays@example.com", "hash", "en", "Stays")
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []int64{u.ID, other.ID} {
		if err := d.RecordActivity(id, "2026-11-20"); err != nil {
			t.Fatal(err)
		}
		if _, err := d.AddSelected(id, "plank30", "Plank", `{"days":[]}`, 3); err != nil {
			t.Fatal(err)
		}
		if err := d.SaveFitnessProfile(id, FitnessProfile{Age: 40, Height: 180, Weight: 80, Concerns: "back"}); err != nil {
			t.Fatal(err)
		}
		if _, err := d.AddFitnessTest(id, FitnessTest{PushUps: 10, Squats: 30, PlankSec: 60, PullUps: -1, Score: 45}); err != nil {
			t.Fatal(err)
		}
	}
	if err := d.CreateSession("tok-gone", u.ID, time.Hour); err != nil {
		t.Fatal(err)
	}
	if err := d.DeleteUser(u.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := d.UserByID(u.ID); err == nil {
		t.Fatal("user still exists")
	}
	for table, want := range map[string]int{"activity": 1, "selected_programs": 1, "sessions": 0, "fitness_profiles": 1, "fitness_tests": 1} {
		var n int
		if err := d.sql.QueryRow(`SELECT COUNT(*) FROM ` + table).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != want {
			t.Errorf("%s: %d rows left, want %d (only the other account's)", table, n, want)
		}
	}
	if _, err := d.UserByID(other.ID); err != nil {
		t.Fatal("the other account was touched")
	}
}
