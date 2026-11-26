package db

import "testing"

func TestFitnessProfileAndTests(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("fit@example.com", "hash", "ru", "Fit")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok, err := d.FitnessProfile(u.ID); err != nil || ok {
		t.Fatalf("new user has a profile: %v %v", ok, err)
	}
	if err := d.SaveFitnessProfile(u.ID, FitnessProfile{Age: 41, Height: 180, Weight: 82, Concerns: "back"}); err != nil {
		t.Fatal(err)
	}
	if err := d.SaveFitnessProfile(u.ID, FitnessProfile{Child: true, Age: 9, Concerns: "feet"}); err != nil {
		t.Fatal(err)
	}
	p, ok, err := d.FitnessProfile(u.ID)
	if err != nil || !ok || !p.Child || p.Age != 9 || p.Height != 0 || p.Concerns != "feet" {
		t.Fatalf("profile = %+v %v %v", p, ok, err)
	}
	for _, score := range []int{40, 57} {
		if _, err := d.AddFitnessTest(u.ID, FitnessTest{PushUps: 12, Squats: 34, PlankSec: 72, PullUps: -1, Score: score}); err != nil {
			t.Fatal(err)
		}
	}
	tests, err := d.FitnessTests(u.ID, 2)
	if err != nil || len(tests) != 2 || tests[0].Score != 57 || tests[1].Score != 40 || tests[0].PullUps != -1 {
		t.Fatalf("tests = %+v %v", tests, err)
	}
	if tests[0].TakenAt.IsZero() {
		t.Fatal("no time")
	}
}
