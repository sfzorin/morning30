package content

import (
	"strings"
	"testing"
)

// base builds a day's workout from the built-in program at level 0.
func base(day int) Workout { return ResolveBuiltin().Workout(day, 18, 0) }

func mainOf(w Workout) []Item {
	var out []Item
	for _, it := range w.Items {
		if it.Slot == Main {
			out = append(out, it)
		}
	}
	return out
}

// wantMainCount is the per-day main-exercise count (per round), from the spec.
var wantMainCount = map[int]int{
	1: 5, 2: 6, 3: 7, 4: 5, 5: 7, 6: 8, 7: 5, 8: 7, 9: 8, 10: 7,
	11: 5, 12: 9, 13: 8, 14: 5, 15: 8, 16: 9, 17: 7, 18: 5, 19: 9, 20: 10,
	21: 5, 22: 9, 23: 7, 24: 10, 25: 5, 26: 10, 27: 8, 28: 5, 29: 10, 30: 10,
}

// Every day: 6 warm-up + (block × 2 rounds) main + 4 cool-down; screen totals
// follow the 20–30 progression.
func TestDayStructureAndCounts(t *testing.T) {
	for day := 1; day <= TotalDays; day++ {
		w := base(day)
		count := wantMainCount[day]
		wantTotal := 6 + count*2 + 4
		if len(w.Items) != wantTotal {
			t.Fatalf("day %d: %d items, want %d", day, len(w.Items), wantTotal)
		}
		for i := 0; i < 6; i++ {
			if w.Items[i].Slot != Warmup || !strings.HasPrefix(w.Items[i].ExerciseID, "W") {
				t.Fatalf("day %d: item %d not warm-up (%s)", day, i, w.Items[i].ExerciseID)
			}
		}
		for _, it := range w.Items[len(w.Items)-4:] {
			if it.Slot != Cooldown || !strings.HasPrefix(it.ExerciseID, "CD") {
				t.Fatalf("day %d: tail not cool-down (%s)", day, it.ExerciseID)
			}
		}
		if got := len(mainOf(w)); got != count*2 {
			t.Errorf("day %d: %d main items, want %d", day, got, count*2)
		}
		if w.Items[len(w.Items)-1].Rest != 0 {
			t.Errorf("day %d: last item has rest", day)
		}
		for _, it := range w.Items {
			if _, ok := Get(it.ExerciseID); !ok {
				t.Errorf("day %d: unknown exercise %s", day, it.ExerciseID)
			}
		}
	}
}

// Warm-up (6, one round) and cool-down (4) are shared and identical every day.
func TestSharedWarmupCooldown(t *testing.T) {
	r := ResolveBuiltin()
	if len(r.Warmup) != 6 || r.WarmupRounds != 1 || len(r.Cooldown) != 4 {
		t.Fatalf("warmup=%d rounds=%d cooldown=%d", len(r.Warmup), r.WarmupRounds, len(r.Cooldown))
	}
	for _, d := range r.Days {
		for _, it := range d.Items {
			if it.Slot != Main {
				t.Fatalf("day %d main has a non-main item %s", d.Day, it.ID)
			}
		}
	}
}

// Recovery days are lighter (5 exercises) and spaced (≤4 working days in a row).
func TestRecoveryCadence(t *testing.T) {
	run := 0
	for day := 1; day <= TotalDays; day++ {
		if isRecoveryDay(day) {
			if wantMainCount[day] != 5 {
				t.Errorf("recovery day %d has %d exercises, want 5", day, wantMainCount[day])
			}
			run = 0
			continue
		}
		run++
		if run > 4 {
			t.Errorf("more than 4 working days in a row by day %d", day)
		}
	}
}

func TestForbiddenExercisesAbsent(t *testing.T) {
	forbidden := []string{"jump", "burpee", "jack", "mountain", "lunge", "highknee", "stepup"}
	for id := range Pool {
		low := strings.ToLower(id)
		for _, f := range forbidden {
			if strings.Contains(low, f) {
				t.Errorf("forbidden exercise in pool: %s", id)
			}
		}
	}
}

// ScaleValue: linear ±10% per step on reps/seconds; breaths fixed; min clamps.
func TestScaleValue(t *testing.T) {
	cases := []struct {
		u              Unit
		val, lvl, want int
	}{
		{Reps, 12, 0, 12}, {Reps, 12, 1, 13}, {Reps, 12, 3, 16}, {Reps, 12, -3, 8},
		{Seconds, 60, 2, 72}, {Seconds, 60, -3, 42},
		{Breaths, 5, 3, 5}, {Reps, 1, -3, 1}, {Seconds, 5, -3, 5},
	}
	for _, c := range cases {
		if got := ScaleValue(c.u, c.val, c.lvl); got != c.want {
			t.Errorf("ScaleValue(%s,%d,%d) = %d, want %d", c.u, c.val, c.lvl, got, c.want)
		}
	}
}

// The level scales the main block but not warm-up/cool-down.
func TestLevelScalingInWorkout(t *testing.T) {
	r := ResolveBuiltin()
	wu := r.Workout(1, 18, 0).Items[0].Value
	for _, lvl := range []int{-3, 0, 3} {
		if v := r.Workout(1, 18, lvl).Items[0].Value; v != wu {
			t.Errorf("level %d changed the warm-up value", lvl)
		}
	}
	// Day 1 first main = push-ups = 12 → +1 = 13, +3 = 16.
	if got := mainOf(r.Workout(1, 18, 1))[0].Value; got != 13 {
		t.Errorf("level +1 push-ups = %d, want 13", got)
	}
	if got := mainOf(r.Workout(1, 18, 3))[0].Value; got != 16 {
		t.Errorf("level +3 push-ups = %d, want 16", got)
	}
}
