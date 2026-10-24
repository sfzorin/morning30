package pages

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/db"
)

func TestStartResolvesCurrentCycleAfterScheduledRestart(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "workout.db"))
	if err != nil {
		t.Fatal(err)
	}
	previous := app.DB
	app.DB = database
	t.Cleanup(func() { app.DB = previous; database.Close() })
	u, err := database.CreateUser("restart@example.com", "hash", "ru", "Test")
	if err != nil {
		t.Fatal(err)
	}
	for range 39 {
		if err := database.AdvancePosition(u.ID); err != nil {
			t.Fatal(err)
		}
	}
	page := workoutPage{sess: auth.Session{UserID: u.ID}, day: 0}
	if got := page.clampDay(); got != 10 {
		t.Fatalf("start day before restart = %d", got)
	}
	if err := database.ScheduleCycleRestart(context.Background(), u.ID, time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	}
	if got := page.clampDay(); got != 1 {
		t.Fatalf("stale home must start new cycle day 1, got %d", got)
	}
	page.day = 7
	if got := page.clampDay(); got != 7 {
		t.Fatalf("manual calendar selection changed: %d", got)
	}
}
