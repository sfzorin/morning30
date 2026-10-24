package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestScheduledCycleRestartPreservesHistory(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("restart@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := d.sql.Exec(`UPDATE users SET position = 17 WHERE id = ?`, u.ID); err != nil {
		t.Fatal(err)
	}
	for _, day := range []string{"2026-10-22", "2026-10-23", "2026-10-24"} {
		if err := d.RecordActivity(u.ID, day); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := d.sql.Exec(`INSERT INTO workout_completions
		(user_id, event_id, day, position, completed_at, completed_on) VALUES (?, 'old-receipt', 17, 16, '2026-10-24T05:00:00Z', '2026-10-24')`, u.ID); err != nil {
		t.Fatal(err)
	}
	at := time.Date(2026, 10, 25, 0, 0, 0, 0, time.FixedZone("Europe/Istanbul", 3*60*60))
	ctx := context.Background()
	if err := d.ScheduleCycleRestart(ctx, u.ID, at); err != nil {
		t.Fatal(err)
	}
	var scheduled string
	if err := d.sql.QueryRow(`SELECT cycle_restart_on FROM users WHERE id = ?`, u.ID).Scan(&scheduled); err != nil {
		t.Fatal(err)
	}
	if scheduled != "2026-10-24T21:00:00Z" {
		t.Fatalf("stored restart = %q", scheduled)
	}
	if pos, changed, err := d.ApplyDueCycleRestart(ctx, u.ID, at.Add(-time.Second), 30); err != nil || changed || pos != 17 {
		t.Fatalf("before due = %d, %v, %v", pos, changed, err)
	}
	// A workout today may still advance the old cycle before the midnight restart.
	if err := d.AdvancePosition(u.ID); err != nil {
		t.Fatal(err)
	}
	if pos, changed, err := d.ApplyDueCycleRestart(ctx, u.ID, at, 30); err != nil || !changed || pos != 30 {
		t.Fatalf("at due = %d, %v, %v", pos, changed, err)
	}
	if pos, changed, err := d.ApplyDueCycleRestart(ctx, u.ID, at.Add(24*time.Hour), 30); err != nil || changed || pos != 30 {
		t.Fatalf("retry = %d, %v, %v", pos, changed, err)
	}
	if err := d.sql.QueryRow(`SELECT cycle_restart_on FROM users WHERE id = ?`, u.ID).Scan(&scheduled); err != nil || scheduled != "" {
		t.Fatalf("consumed schedule = %q, %v", scheduled, err)
	}
	if streak, err := d.Streak(u.ID, "2026-10-25"); err != nil || streak != 3 {
		t.Fatalf("preserved streak = %d, %v", streak, err)
	}
	var receipts int
	if err := d.sql.QueryRow(`SELECT COUNT(*) FROM workout_completions WHERE user_id = ? AND event_id = 'old-receipt'`, u.ID).Scan(&receipts); err != nil || receipts != 1 {
		t.Fatalf("preserved receipts = %d, %v", receipts, err)
	}
}

func TestCycleRestartAlwaysMovesToNextBoundary(t *testing.T) {
	for _, tc := range []struct{ position, total, want int }{
		{0, 30, 30}, {30, 30, 60}, {61, 30, 90}, {7, 6, 12},
	} {
		t.Run(fmt.Sprintf("%d_in_%d", tc.position, tc.total), func(t *testing.T) {
			d := openTest(t)
			u, err := d.CreateUser("boundary@example.com", "hash", "en", "Test")
			if err != nil {
				t.Fatal(err)
			}
			if _, err := d.sql.Exec(`UPDATE users SET position = ? WHERE id = ?`, tc.position, u.ID); err != nil {
				t.Fatal(err)
			}
			now := time.Date(2026, 10, 25, 0, 0, 0, 0, time.UTC)
			if err := d.ScheduleCycleRestart(context.Background(), u.ID, now); err != nil {
				t.Fatal(err)
			}
			if pos, changed, err := d.ApplyDueCycleRestart(context.Background(), u.ID, now, tc.total); err != nil || !changed || pos != tc.want {
				t.Fatalf("restart = %d, %v, %v; want %d", pos, changed, err, tc.want)
			}
		})
	}
}

func TestDueRestartRejectsStalePositionInCompletion(t *testing.T) {
	for _, position := range []int{0, 30, 37} {
		t.Run(fmt.Sprint(position), func(t *testing.T) {
			d := openTest(t)
			u, err := d.CreateUser("stale@example.com", "hash", "en", "Test")
			if err != nil {
				t.Fatal(err)
			}
			if _, err := d.sql.Exec(`UPDATE users SET position = ? WHERE id = ?`, position, u.ID); err != nil {
				t.Fatal(err)
			}
			ctx := context.Background()
			if err := d.ScheduleCycleRestart(ctx, u.ID, time.Now().Add(-time.Hour)); err != nil {
				t.Fatal(err)
			}
			event := WorkoutCompletion{EventID: "stale-workout", UserID: u.ID, Day: position%30 + 1, Position: position,
				CompletedAt: "2026-10-24T05:00:00Z", CompletedOn: "2026-10-24"}
			want := (position/30 + 1) * 30
			for range 2 {
				result, err := d.CompleteWorkout(ctx, event)
				if err != nil || !result.Saved || result.Position != want {
					t.Fatalf("stale completion = %+v, %v; want position %d", result, err, want)
				}
			}
			if has, err := d.HasActivity(u.ID, event.CompletedOn); err != nil || !has {
				t.Fatalf("stale workout activity must remain credited: %v, %v", has, err)
			}
		})
	}
}

func TestCompletionRetryAppliesRestartUsingServerTime(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("retry-restart@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	event := WorkoutCompletion{EventID: "saved-before-restart", UserID: u.ID, Day: 1, Position: 0,
		CompletedAt: time.Now().Add(24 * time.Hour).UTC().Format(time.RFC3339), CompletedOn: "2026-10-24"}
	if err := d.ScheduleCycleRestart(ctx, u.ID, time.Now().Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	// A future client timestamp cannot trigger an early restart.
	if result, err := d.CompleteWorkout(ctx, event); err != nil || result.Position != 1 {
		t.Fatalf("before scheduled server time = %+v, %v", result, err)
	}
	if err := d.ScheduleCycleRestart(ctx, u.ID, time.Now().Add(-time.Hour)); err != nil {
		t.Fatal(err)
	}
	// An already saved receipt must not bypass a restart that is now due.
	if result, err := d.CompleteWorkout(ctx, event); err != nil || result.Position != 30 || !result.Saved {
		t.Fatalf("retry after due = %+v, %v", result, err)
	}
}

func TestCycleRestartFailureIsAtomicAndRetryable(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("restart-rollback@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	now := time.Now()
	if err := d.ScheduleCycleRestart(ctx, u.ID, now); err != nil {
		t.Fatal(err)
	}
	if _, err := d.sql.Exec(`CREATE TRIGGER fail_restart BEFORE UPDATE OF position ON users
		BEGIN SELECT RAISE(ABORT, 'simulated write failure'); END`); err != nil {
		t.Fatal(err)
	}
	if _, changed, err := d.ApplyDueCycleRestart(ctx, u.ID, now, 30); err == nil || changed {
		t.Fatalf("failed restart acknowledged: changed=%v, err=%v", changed, err)
	}
	var position int
	var scheduled string
	if err := d.sql.QueryRow(`SELECT position, cycle_restart_on FROM users WHERE id = ?`, u.ID).Scan(&position, &scheduled); err != nil {
		t.Fatal(err)
	}
	if position != 0 || scheduled == "" {
		t.Fatalf("failed restart changed state: %d %q", position, scheduled)
	}
	if _, err := d.sql.Exec(`DROP TRIGGER fail_restart`); err != nil {
		t.Fatal(err)
	}
	if pos, changed, err := d.ApplyDueCycleRestart(ctx, u.ID, now, 30); err != nil || !changed || pos != 30 {
		t.Fatalf("restart retry = %d, %v, %v", pos, changed, err)
	}
}

func TestCycleRestartValidation(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("restart-invalid@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	now := time.Now()
	if err := d.ScheduleCycleRestart(ctx, u.ID, time.Time{}); err == nil {
		t.Fatal("accepted zero restart instant")
	}
	if err := d.ScheduleCycleRestart(ctx, u.ID+1, now); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("unknown account schedule = %v", err)
	}
	if _, _, err := d.ApplyDueCycleRestart(ctx, u.ID, now, 0); err == nil {
		t.Fatal("accepted empty cycle")
	}
	if err := d.ScheduleCycleRestart(ctx, u.ID, now); err != nil {
		t.Fatal(err)
	}
	if _, err := d.sql.Exec(`UPDATE users SET position = ? WHERE id = ?`, math.MaxInt, u.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := d.ApplyDueCycleRestart(ctx, u.ID, now, 30); err == nil {
		t.Fatal("accepted position overflow")
	}
	if _, err := d.sql.Exec(`UPDATE users SET position = 0, cycle_restart_on = 'invalid' WHERE id = ?`, u.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := d.ApplyDueCycleRestart(ctx, u.ID, now, 30); err == nil {
		t.Fatal("accepted malformed stored restart instant")
	}
}

func TestScheduledCycleRestartLookup(t *testing.T) {
	d := openTest(t)
	u, err := d.CreateUser("restart-lookup@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if at, err := d.ScheduledCycleRestart(ctx, u.ID); err != nil || !at.IsZero() {
		t.Fatalf("empty schedule = %v, %v", at, err)
	}
	at := time.Date(2026, 10, 25, 0, 0, 0, 123, time.FixedZone("Europe/Istanbul", 3*60*60))
	if err := d.ScheduleCycleRestart(ctx, u.ID, at); err != nil {
		t.Fatal(err)
	}
	if got, err := d.ScheduledCycleRestart(ctx, u.ID); err != nil || !got.Equal(at) {
		t.Fatalf("scheduled instant = %v, %v; want %v", got, err, at)
	}
	if _, err := d.ScheduledCycleRestart(ctx, u.ID+1); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("unknown account lookup = %v", err)
	}
	if _, err := d.sql.Exec(`UPDATE users SET cycle_restart_on = 'invalid' WHERE id = ?`, u.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := d.ScheduledCycleRestart(ctx, u.ID); err == nil {
		t.Fatal("accepted malformed scheduled instant")
	}
}
