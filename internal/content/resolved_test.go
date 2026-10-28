package content

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestLegacyPausesAreIgnoredAndTwoWholeRoundsAreCanonical(t *testing.T) {
	raw := `{"version":2,"fixedDifficulty":true,"warmupRounds":4,"warmup":[{"id":"warm","unit":"seconds","value":45,"slot":"main","rest":15}],"cooldown":[{"id":"cool","unit":"seconds","value":20,"slot":"main","rest":90}],"days":[{"rounds":1,"items":[{"id":"a","unit":"reps","value":6,"perSide":true,"slot":"cooldown","round":1,"rest":90,"sideRest":10,"repSeconds":4},{"id":"a","unit":"reps","value":6,"perSide":true,"slot":"warmup","round":2,"rest":0,"sideRest":20,"repSeconds":4},{"id":"b","unit":"seconds","value":10,"round":1,"rest":60},{"id":"b","unit":"seconds","value":10,"round":2,"rest":0}]},{"rounds":9,"warmup":[],"cooldown":[],"items":[{"id":"recover","unit":"seconds","value":60,"rest":90}]}]}`
	p, err := ParseResolved(raw)
	if err != nil {
		t.Fatal(err)
	}
	if p.WarmupRounds != 1 || p.Days[0].Rounds != 2 || len(p.Days[0].Items) != 2 {
		t.Fatalf("noncanonical rounds: %+v", p.Days[0])
	}
	serialized := p.Marshal()
	if strings.Contains(serialized, `"rest"`) || strings.Contains(serialized, `"sideRest"`) {
		t.Fatal("legacy pause overrides survived export")
	}
	q, err := ParseResolved(serialized)
	if err != nil || !reflect.DeepEqual(p, q) {
		t.Fatalf("round trip changed program: %v", err)
	}
	w := q.Workout(1, 3, Rests{Warmup: 5, Main: 18, Cooldown: 5})
	if !w.FixedDifficulty || len(w.Items) != 6 {
		t.Fatalf("workout: %+v", w)
	}
	for i, id := range []string{"warm", "a", "b", "a", "b", "cool"} {
		if w.Items[i].ExerciseID != id {
			t.Fatalf("sequence at %d = %s, want %s", i, w.Items[i].ExerciseID, id)
		}
	}
	for i, rest := range []int{18, 18, 18, 18, 5, 0} {
		if w.Items[i].Rest != rest {
			t.Errorf("item %d rest=%d want %d", i, w.Items[i].Rest, rest)
		}
	}
	if w.Items[1].Value != 6 || w.Items[1].Slot != Main || w.Items[2].Slot != Main {
		t.Fatal("dose or containing phase was changed")
	}
	if w.EstSec != 45+(6*4*2+10)*2+20+18*4+5 {
		t.Fatalf("duration=%d", w.EstSec)
	}
	recovery := q.Workout(2, 0, Rests{Main: 18, Cooldown: 5})
	if len(recovery.Items) != 4 || recovery.Items[0].ExerciseID != "warm" || recovery.Items[3].ExerciseID != "cool" {
		t.Fatalf("empty overrides disabled daily phases: %+v", recovery)
	}
}

func TestLegacyPhasesPreserveNonemptyVariantsAndInheritMissingPhases(t *testing.T) {
	p, err := ParseResolved(`{"days":[{"items":[{"id":"a","slot":"warmup","value":10},{"id":"m","slot":"main","value":5},{"id":"z","slot":"cooldown","value":20}]},{"items":[{"id":"b","slot":"warmup","value":15},{"id":"m","slot":"main","value":6}]},{"items":[{"id":"m","slot":"main","value":7}]}]}`)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Warmup) != 1 || p.Warmup[0].ID != "a" || len(p.Cooldown) != 1 {
		t.Fatal("legacy shared phases not migrated")
	}
	if p.Days[1].Warmup == nil || (*p.Days[1].Warmup)[0].ID != "b" || p.Days[1].Cooldown != nil {
		t.Fatal("nonempty variation or missing-phase inheritance lost")
	}
	for day := 1; day <= 3; day++ {
		w := p.Workout(day, 0, Rests{})
		if len(w.Items) != 4 || w.Items[0].Slot != Warmup || w.Items[3].Slot != Cooldown || len(mainOf(w)) != 2 {
			t.Fatalf("day %d lacks the mandatory phases/rounds: %+v", day, w)
		}
	}
}

func TestLegacyFullDayFilesWithEmptySharedFieldsStillMigrate(t *testing.T) {
	for _, shared := range []string{`"warmup":null,"cooldown":null`, `"warmup":[],"cooldown":[]`} {
		p, err := ParseResolved(`{"version":1,` + shared + `,"days":[{"items":[{"id":"w","unit":"seconds","value":10,"slot":"warmup"},{"id":"m","unit":"reps","value":6,"slot":"main"},{"id":"c","unit":"seconds","value":10,"slot":"cooldown"}]}]}`)
		if err != nil {
			t.Fatal(err)
		}
		w := p.Workout(1, 0, Rests{})
		if len(w.Items) != 4 || w.Items[0].ExerciseID != "w" || w.Items[3].ExerciseID != "c" || len(mainOf(w)) != 2 {
			t.Fatalf("legacy shared fields blocked migration: %+v", w)
		}
	}
}

func TestDirectProgramsObeyStructureAndDoNotMutateInput(t *testing.T) {
	empty := []RItem{}
	variant := []RItem{{ID: "variant", Unit: Seconds, Value: 5, Slot: Main, Round: 4}}
	p := Resolved{WarmupRounds: 6,
		Warmup:   []RItem{{ID: "warm", Unit: Seconds, Value: 5, Slot: Cooldown}},
		Cooldown: []RItem{{ID: "cool", Unit: Seconds, Value: 5, Slot: Main}},
		Days: []RDay{{Rounds: 10, Warmup: &empty, Cooldown: &variant, Items: []RItem{
			{ID: "a", Unit: Reps, Value: 6, Slot: Warmup, Round: 1},
			{ID: "b", Unit: Reps, Value: 8, Slot: Cooldown, Round: 1},
			{ID: "a", Unit: Reps, Value: 6, Slot: Main, Round: 2},
			{ID: "b", Unit: Reps, Value: 8, Slot: Main, Round: 2},
		}}}}
	before, _ := json.Marshal(p)
	w := p.Workout(1, 0, Rests{})
	if len(w.Items) != 6 {
		t.Fatalf("direct program repeated phases: %+v", w)
	}
	for i, want := range []struct {
		id    string
		slot  Slot
		round int
	}{
		{"warm", Warmup, 1}, {"a", Main, 1}, {"b", Main, 1}, {"a", Main, 2}, {"b", Main, 2}, {"variant", Cooldown, 1},
	} {
		it := w.Items[i]
		if it.ExerciseID != want.id || it.Slot != want.slot || it.Round != want.round {
			t.Fatalf("item %d: %+v", i, it)
		}
	}
	q, err := ParseResolved(p.Marshal())
	if err != nil || q.WarmupRounds != 1 || q.Days[0].Rounds != 2 || len(q.Days[0].Items) != 2 {
		t.Fatalf("direct export is not canonical: %v %+v", err, q.Days[0])
	}
	after, _ := json.Marshal(p)
	if string(before) != string(after) {
		t.Fatal("playback/export mutated the caller's program")
	}
}

func TestMissingPhasesAndMainUseGenericDefaults(t *testing.T) {
	for _, raw := range []string{`{}`, `{"version":2,"warmup":[],"cooldown":[],"days":[{"rounds":1,"items":[],"warmup":[],"cooldown":[]}]}`} {
		p, err := ParseResolved(raw)
		if err != nil {
			t.Fatal(err)
		}
		if len(p.Days) != TotalDays || len(p.Warmup) == 0 || len(p.Cooldown) == 0 {
			t.Fatal("incomplete normalized program")
		}
		for day := 1; day <= TotalDays; day++ {
			w := p.Workout(day, 0, Rests{})
			demo := ResolveDemo().Workout(day, 0, Rests{})
			if !reflect.DeepEqual(w.Items, demo.Items) {
				t.Fatalf("missing day %d did not use the generic demo", day)
			}
		}
	}
	w := (Resolved{}).Workout(1, 0, Rests{})
	if len(mainOf(w)) == 0 || w.Items[0].Slot != Warmup || w.Items[len(w.Items)-1].Slot != Cooldown {
		t.Fatal("direct empty struct bypassed defaults")
	}
}

func TestAsymmetricLegacyRoundImportIsExplicitlyRejected(t *testing.T) {
	for _, raw := range []string{
		`{"days":[{"items":[{"id":"a","unit":"reps","value":6,"round":1},{"id":"a","unit":"reps","value":8,"round":2}]}]}`,
		`{"days":[{"items":[{"id":"a","unit":"reps","value":6,"round":1},{"id":"b","unit":"reps","value":6,"round":2}]}]}`,
		`{"days":[{"items":[{"id":"a","unit":"reps","value":6,"round":1},{"id":"a","unit":"reps","value":6,"round":2},{"id":"b","unit":"reps","value":6,"round":2}]}]}`,
	} {
		if _, err := ParseResolved(raw); err == nil || !strings.Contains(err.Error(), "different exercises or doses") {
			t.Fatalf("asymmetric import should fail clearly: %v", err)
		}
	}
	// Direct structs cannot return an import error; use their first labelled
	// sequence twice, never multiply an entire expanded asymmetric day.
	p := Resolved{Days: []RDay{{Items: []RItem{{ID: "a", Unit: Reps, Value: 6, Round: 1}, {ID: "b", Unit: Reps, Value: 10, Round: 2}}}}}
	m := mainOf(p.Workout(1, 0, Rests{}))
	if len(m) != 2 || m[0].ExerciseID != "a" || m[1].ExerciseID != "a" || m[0].Value != 6 {
		t.Fatalf("direct legacy fallback doubled expanded work: %+v", m)
	}
}

func TestUserRestPreferencesOwnAllTransitionsAndEstimate(t *testing.T) {
	p := Resolved{
		Warmup:   []RItem{{ID: "w1", Unit: Seconds, Value: 10}, {ID: "w2", Unit: Seconds, Value: 10}},
		Cooldown: []RItem{{ID: "c1", Unit: Seconds, Value: 10}, {ID: "c2", Unit: Seconds, Value: 10}},
		Days:     []RDay{{Items: []RItem{{ID: "a", Unit: Seconds, Value: 10}, {ID: "b", Unit: Seconds, Value: 10}}}},
	}
	for _, r := range []Rests{{Warmup: 0, Main: 5, Cooldown: 0}, {Warmup: 10, Main: 20, Cooldown: 15}, {Warmup: 40, Main: 40, Cooldown: 40}, {Warmup: -20, Main: 90, Cooldown: 99}} {
		w := p.Workout(1, 0, r)
		n := r.Normalize()
		want := []int{n.Warmup, n.Main, n.Main, n.Main, n.Main, n.Cooldown, n.Cooldown, 0}
		est := 80
		for i, rest := range want {
			est += rest
			if w.Items[i].Rest != rest {
				t.Fatalf("settings %+v item %d rest=%d want %d", r, i, w.Items[i].Rest, rest)
			}
		}
		if w.EstSec != est {
			t.Fatalf("settings %+v estimate=%d want %d", r, w.EstSec, est)
		}
	}
}

func TestRestNormalizationMatchesSettingsRanges(t *testing.T) {
	for _, tc := range []struct{ in, want Rests }{
		{Rests{}, Rests{Main: 5}},
		{Rests{Warmup: -1, Main: -10, Cooldown: -5}, Rests{Main: 5}},
		{Rests{Warmup: 90, Main: 60, Cooldown: 45}, Rests{Warmup: 40, Main: 40, Cooldown: 40}},
		{Rests{Warmup: 10, Main: 20, Cooldown: 0}, Rests{Warmup: 10, Main: 20, Cooldown: 0}},
	} {
		if got := tc.in.Normalize(); got != tc.want {
			t.Fatalf("Normalize(%+v)=%+v want %+v", tc.in, got, tc.want)
		}
	}
}

func TestSelfManagedSidesPreservesDoseAndRoundTrip(t *testing.T) {
	p, err := ParseResolved(`{"version":2,"selfManagedSides":true,"days":[{"rounds":2,"items":[{"id":"custom-reps","unit":"reps","value":6,"perSide":true,"slot":"main","rest":0,"sideRest":10},{"id":"custom-hold","unit":"seconds","value":20,"perSide":true,"slot":"main","rest":0,"sideRest":10}]}]}`)
	if err != nil {
		t.Fatal(err)
	}
	q, err := ParseResolved(p.Marshal())
	if err != nil || !reflect.DeepEqual(p, q) || !q.SelfManagedSides {
		t.Fatalf("self-managed sides lost on export/import: %v", err)
	}
	w := q.Workout(1, 0, Rests{Main: 20})
	m := mainOf(w)
	if !w.SelfManagedSides || len(m) != 4 {
		t.Fatalf("combined-side dose or rounds changed: %+v", w)
	}
	for i, it := range m {
		if !it.PerSide || it.Round != i/2+1 {
			t.Fatalf("item %d has wrong side/round: %+v", i, it)
		}
	}
	q.SelfManagedSides = false
	separate := q.Workout(1, 0, Rests{Main: 20})
	if separate.SelfManagedSides || separate.EstSec != w.EstSec || !reflect.DeepEqual(separate.Items, w.Items) {
		t.Fatal("side display preference changed doses or pauses")
	}
}

func TestTimedSideChangesAddPreparationWithoutConsumingWorkTime(t *testing.T) {
	p := Resolved{
		Warmup: []RItem{{ID: "warm-hold", Unit: Seconds, Value: 10, PerSide: true}},
		Days: []RDay{{Items: []RItem{
			{ID: "hold", Unit: Seconds, Value: 20, PerSide: true},
			{ID: "reps", Unit: Reps, Value: 6, PerSide: true, RepSeconds: 4},
			{ID: "breaths", Unit: Breaths, Value: 3, PerSide: true},
		}}},
		Cooldown: []RItem{{ID: "cool", Unit: Seconds, Value: 5}},
	}
	for _, selfManaged := range []bool{false, true} {
		p.SelfManagedSides = selfManaged
		w := p.Workout(1, 0, Rests{Main: 5})
		// 20s warmup + 3s switch; two rounds of 40s + 3s switch,
		// 48s counted reps and 24s breaths; 5s cooldown; six 5s rests.
		if w.EstSec != 288 {
			t.Fatalf("selfManaged=%v: duration=%d, want 288", selfManaged, w.EstSec)
		}
		if w.Items[0].Value != 10 || w.Items[1].Value != 20 || w.Items[4].Value != 20 {
			t.Fatal("preparation reduced the timed exercise dose")
		}
	}
}

func TestSideChangePreparationOnlyAppliesToTimedBilateralWork(t *testing.T) {
	for _, tc := range []struct {
		unit    Unit
		perSide bool
		want    int
	}{
		{Seconds, true, 23}, {Seconds, false, 10},
		{Reps, true, 60}, {Reps, false, 30},
		{Breaths, true, 80}, {Breaths, false, 40},
	} {
		items := []Item{{Unit: tc.unit, Value: 10, PerSide: tc.perSide, Slot: Main}}
		if got := assignRestsAndEst(items, Rests{Main: 40}); got != tc.want {
			t.Fatalf("unit=%s perSide=%v: duration=%d want %d", tc.unit, tc.perSide, got, tc.want)
		}
		if items[0].Rest != 0 {
			t.Fatal("side preparation introduced a final exercise rest")
		}
	}
}
