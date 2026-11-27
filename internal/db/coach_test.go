package db

import (
	"testing"
	"time"
)

func TestCoachHistoryIsBounded(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("talk@example.com", "hash", "ru", "Talk")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 60; i++ {
		undo := ""
		if i%2 == 1 {
			undo = `{"program":"x"}`
		}
		if _, err := d.AddCoachMessage(u.ID, CoachMessage{ProgramID: 7, Role: "coach", Text: "m", Undo: undo}); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := d.AddCoachMessage(u.ID, CoachMessage{ProgramID: 8, Role: "user", Text: "other program"}); err != nil {
		t.Fatal(err)
	}
	if err := d.PruneCoach(u.ID, 7, 50); err != nil {
		t.Fatal(err)
	}
	msgs, _ := d.CoachMessages(u.ID, 7, 100)
	if len(msgs) != 50 {
		t.Fatalf("%d messages kept, want 50", len(msgs))
	}
	undos := 0
	for _, m := range msgs {
		if m.Undo != "" {
			undos++
		}
	}
	if undos != 1 || msgs[len(msgs)-1].Undo == "" {
		t.Fatalf("%d undo copies kept, want only the newest", undos)
	}
	if other, _ := d.CoachMessages(u.ID, 8, 10); len(other) != 1 {
		t.Fatal("another program's thread was touched")
	}
	recent, _ := d.CoachRecent(u.ID, 7, time.Now().Add(-time.Hour), 10)
	if len(recent) != 10 {
		t.Fatalf("recent = %d, want 10", len(recent))
	}
	if old, _ := d.CoachRecent(u.ID, 7, time.Now().Add(time.Hour), 10); len(old) != 0 {
		t.Fatal("messages older than the window were returned")
	}
}
