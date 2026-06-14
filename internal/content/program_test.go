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

func isPlank(id string) bool { return id == "C02" || id == "C05" || id == "C08" }

func TestStructureWarmupCooldown(t *testing.T) {
	for day := 1; day <= TotalDays; day++ {
		w := base(day)
		// Warm-up: 10 exercises × 2 rounds = first 20 items, all W*.
		if len(w.Items) < 26 {
			t.Fatalf("day %d: too few items (%d)", day, len(w.Items))
		}
		for i := 0; i < 20; i++ {
			if w.Items[i].Slot != Warmup || !strings.HasPrefix(w.Items[i].ExerciseID, "W") {
				t.Fatalf("day %d: item %d is not a warm-up (%s)", day, i, w.Items[i].ExerciseID)
			}
		}
		// Cool-down: last 6 items, all CD*, 1 round.
		cd := w.Items[len(w.Items)-6:]
		for _, it := range cd {
			if it.Slot != Cooldown || !strings.HasPrefix(it.ExerciseID, "CD") {
				t.Fatalf("day %d: expected cool-down at the end, got %s", day, it.ExerciseID)
			}
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

// Warm-up and cool-down are stored once and identical for every day.
func TestSharedWarmupCooldown(t *testing.T) {
	r := ResolveBuiltin()
	if len(r.Warmup) != 10 || r.WarmupRounds != 2 || len(r.Cooldown) != 6 {
		t.Fatalf("shared series: warmup=%d rounds=%d cooldown=%d", len(r.Warmup), r.WarmupRounds, len(r.Cooldown))
	}
	for _, d := range r.Days {
		for _, it := range d.Items {
			if it.Slot != Main {
				t.Fatalf("day %d main list has a non-main item %s (%s)", d.Day, it.ID, it.Slot)
			}
		}
	}
}

func TestPlankChallengeAndProgression(t *testing.T) {
	for day := 1; day <= TotalDays; day++ {
		p := calendar[day]
		if p.Recovery {
			continue
		}
		m := mainOf(base(day))
		last := m[len(m)-1]
		if !isPlank(last.ExerciseID) {
			t.Errorf("day %d: main does not end with a plank (got %s)", day, last.ExerciseID)
		}
		if last.ExerciseID == "C02" && last.Value < 45 {
			t.Errorf("day %d: plank %ds < 45s", day, last.Value)
		}
		if p.Control && last.Value != 90 {
			t.Errorf("day 30 control plank = %ds, want 90", last.Value)
		}
	}
}

func TestRecoveryDays(t *testing.T) {
	allowed := map[string]bool{"C04": true, "G01": true, "B01": true, "S01": true, "C05": true}
	rCount := 0
	for day := 1; day <= TotalDays; day++ {
		if !calendar[day].Recovery {
			continue
		}
		rCount++
		for _, it := range mainOf(base(day)) {
			if it.Round != 1 {
				t.Errorf("R day %d: main round %d, want single round", day, it.Round)
			}
			if !allowed[it.ExerciseID] {
				t.Errorf("R day %d: forbidden recovery exercise %s", day, it.ExerciseID)
			}
		}
	}
	if rCount < 7 {
		t.Errorf("only %d recovery days, expected several", rCount)
	}
}

func TestRecoveryCadence(t *testing.T) {
	run := 0
	for day := 1; day <= TotalDays; day++ {
		if calendar[day].Recovery {
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
		u             Unit
		val, lvl, want int
	}{
		{Reps, 11, 0, 11}, {Reps, 11, 3, 14}, {Reps, 11, -3, 8},
		{Seconds, 50, 2, 60}, {Seconds, 50, -3, 35},
		{Breaths, 5, 3, 5}, {Breaths, 5, -3, 5},
		{Reps, 1, -3, 1}, {Seconds, 5, -3, 5}, // clamps
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
	wuVal := r.Workout(1, 18, 0).Items[0].Value
	for _, lvl := range []int{-3, 0, 3} {
		if v := r.Workout(1, 18, lvl).Items[0].Value; v != wuVal {
			t.Errorf("level %d changed warm-up value (%d != %d)", lvl, v, wuVal)
		}
	}
	// Day 1 first push-up base = 10 reps → +1 = 11, +3 = 13, −3 = 7.
	if got := mainOf(r.Workout(1, 18, 1))[0].Value; got != 11 {
		t.Errorf("level +1 push-ups = %d, want 11", got)
	}
	if got := mainOf(r.Workout(1, 18, 3))[0].Value; got != 13 {
		t.Errorf("level +3 push-ups = %d, want 13", got)
	}
}

func TestRestRules(t *testing.T) {
	w := base(1) // block A
	for i, it := range w.Items {
		if i == len(w.Items)-1 {
			continue
		}
		switch {
		case it.ExerciseID == "M01" || it.ExerciseID == "M02" || it.ExerciseID == "M03":
			if it.Rest != 35 {
				t.Errorf("after push-ups rest = %d, want 35", it.Rest)
			}
		}
	}
}
