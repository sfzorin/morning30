package db

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func completionFor(userID int64) WorkoutCompletion {
	return WorkoutCompletion{
		EventID: "eea47236-50ec-4adb-9303-93fd090dc2d7", UserID: userID,
		Day: 1, Position: 0, CompletedAt: "2026-10-24T05:00:00Z", CompletedOn: "2026-10-24",
	}
}

func TestCompleteWorkoutDuplicateAndConflictingReceipt(t *testing.T) {
	d := openTest(t)
	u, _ := d.CreateUser("workout@example.com", "hash", "en", "Test")
	event := completionFor(u.ID)
	_ = d.RecordActivity(u.ID, "2026-10-23")
	for range 3 {
		result, err := d.CompleteWorkout(context.Background(), event)
		if err != nil || !result.Saved || result.EventID != event.EventID || result.Position != 1 || result.Streak != 2 {
			t.Fatalf("complete = %+v, %v", result, err)
		}
	}
	for _, change := range []func(*WorkoutCompletion){
		func(e *WorkoutCompletion) { e.Day = 2 },
		func(e *WorkoutCompletion) { e.Position = 1 },
		func(e *WorkoutCompletion) { e.CompletedAt = "2026-10-25T05:00:00Z" },
		func(e *WorkoutCompletion) { e.CompletedOn = "2026-10-25" },
	} {
		conflict := event
		change(&conflict)
		if _, err := d.CompleteWorkout(context.Background(), conflict); !errors.Is(err, ErrCompletionConflict) {
			t.Fatalf("conflicting receipt accepted: %v", err)
		}
	}
	if has, _ := d.HasActivity(u.ID, "2026-10-25"); has {
		t.Fatal("conflicting receipt changed activity")
	}
	other, _ := d.CreateUser("other@example.com", "hash", "en", "Other")
	event.UserID = other.ID
	if result, err := d.CompleteWorkout(context.Background(), event); err != nil || result.Position != 1 {
		t.Fatalf("event ID should be scoped to its account: %+v, %v", result, err)
	}
}

func TestCompleteWorkoutConcurrentDelivery(t *testing.T) {
	for _, distinct := range []bool{false, true} {
		t.Run(fmt.Sprintf("distinct_events_%v", distinct), func(t *testing.T) {
			d := openTest(t)
			u, _ := d.CreateUser("parallel@example.com", "hash", "en", "Test")
			const requests = 20
			var wg sync.WaitGroup
			for i := range requests {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					event := completionFor(u.ID)
					if distinct {
						event.EventID = fmt.Sprintf("%032d", i)
					}
					if result, err := d.CompleteWorkout(context.Background(), event); err != nil || result.Position != 1 {
						t.Errorf("parallel completion = %+v, %v", result, err)
					}
				}(i)
			}
			wg.Wait()
			var receipts int
			if err := d.sql.QueryRow(`SELECT COUNT(*) FROM workout_completions WHERE user_id = ?`, u.ID).Scan(&receipts); err != nil {
				t.Fatal(err)
			}
			want := 1
			if distinct {
				want = requests
			}
			if receipts != want {
				t.Fatalf("receipts = %d, want %d", receipts, want)
			}
		})
	}
}

func TestCompleteWorkoutPositionUsesCapturedCycle(t *testing.T) {
	for _, tc := range []struct {
		name                         string
		current, captured, day, want int
	}{
		{"later day in second cycle", 39, 39, 17, 47},
		{"earlier day replay", 39, 39, 5, 39},
		{"finish cycle", 59, 59, 30, 60},
		{"stale position", 39, 0, 30, 39},
		{"future position", 39, 60, 1, 39},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := openTest(t)
			u, _ := d.CreateUser("position@example.com", "hash", "en", "Test")
			_, _ = d.sql.Exec(`UPDATE users SET position = ? WHERE id = ?`, tc.current, u.ID)
			event := completionFor(u.ID)
			event.Position, event.Day = tc.captured, tc.day
			result, err := d.CompleteWorkout(context.Background(), event)
			if err != nil || result.Position != tc.want {
				t.Fatalf("complete = %+v, %v; want position %d", result, err, tc.want)
			}
			if has, _ := d.HasActivity(u.ID, event.CompletedOn); !has {
				t.Fatal("replay or stale completion should still credit its activity date")
			}
		})
	}
}

func TestCompleteWorkoutRollbackCanRetry(t *testing.T) {
	d := openTest(t)
	u, _ := d.CreateUser("rollback@example.com", "hash", "en", "Test")
	event := completionFor(u.ID)
	if _, err := d.sql.Exec(`CREATE TRIGGER fail_workout BEFORE UPDATE OF position ON users
		BEGIN SELECT RAISE(ABORT, 'simulated write failure'); END`); err != nil {
		t.Fatal(err)
	}
	if result, err := d.CompleteWorkout(context.Background(), event); err == nil || result.Saved {
		t.Fatalf("write failure acknowledged: %+v, %v", result, err)
	}
	if has, _ := d.HasActivity(u.ID, event.CompletedOn); has {
		t.Fatal("activity survived the failed transaction")
	}
	var receipts int
	_ = d.sql.QueryRow(`SELECT COUNT(*) FROM workout_completions`).Scan(&receipts)
	if receipts != 0 {
		t.Fatal("receipt survived the failed transaction")
	}
	if position, _ := d.Position(u.ID); position != 0 {
		t.Fatal("position changed despite rollback")
	}
	_, _ = d.sql.Exec(`DROP TRIGGER fail_workout`)
	if result, err := d.CompleteWorkout(context.Background(), event); err != nil || !result.Saved || result.Position != 1 {
		t.Fatalf("retry after rollback = %+v, %v", result, err)
	}
}

func TestWorkoutMigrationPreservesExistingUserData(t *testing.T) {
	d := openTest(t)
	u, _ := d.CreateUser("migration@example.com", "hash", "en", "Test")
	_, _ = d.sql.Exec(`UPDATE users SET position = 37, program_json = 'personal plan',
		active_program = 'custom', custom_exercises = 'personal exercises' WHERE id = ?`, u.ID)
	_ = d.RecordActivity(u.ID, "2026-10-23")
	// Recreate the schema state immediately before this migration.
	if _, err := d.sql.Exec(`DROP TABLE workout_completions`); err != nil {
		t.Fatal(err)
	}
	if _, err := d.sql.Exec(`ALTER TABLE users DROP COLUMN cycle_restart_on`); err != nil {
		t.Fatal(err)
	}
	if err := d.migrate(); err != nil {
		t.Fatal(err)
	}
	var position int
	var program, active, exercises, restart string
	if err := d.sql.QueryRow(`SELECT position, program_json, active_program, custom_exercises,
		cycle_restart_on FROM users WHERE id = ?`, u.ID).Scan(&position, &program, &active, &exercises, &restart); err != nil {
		t.Fatal(err)
	}
	if position != 37 || program != "personal plan" || active != "custom" || exercises != "personal exercises" || restart != "" {
		t.Fatalf("migration changed account: %d %q %q %q %q", position, program, active, exercises, restart)
	}
	if has, _ := d.HasActivity(u.ID, "2026-10-23"); !has {
		t.Fatal("migration removed activity")
	}
	event := completionFor(u.ID)
	event.Day, event.Position = 8, 37
	if result, err := d.CompleteWorkout(context.Background(), event); err != nil || result.Position != 38 {
		t.Fatalf("migrated receipt table is not usable: %+v, %v", result, err)
	}
}
