package content

import (
	"reflect"
	"testing"
)

func TestPrescriptionRoundTripAndWorkout(t *testing.T) {
	raw := `{"version":2,"fixedDifficulty":true,"warmupRounds":1,"warmup":[{"id":"warm","unit":"seconds","value":45,"slot":"warmup","rest":15}],"cooldown":[{"id":"cool","unit":"seconds","value":20,"slot":"cooldown"}],"days":[{"rounds":1,"items":[{"id":"slow","unit":"reps","value":6,"perSide":true,"slot":"main","round":1,"rest":90,"sideRest":10,"repSeconds":4},{"id":"slow","unit":"reps","value":6,"perSide":true,"slot":"main","round":1,"rest":0,"sideRest":10,"repSeconds":4}]},{"rounds":1,"warmup":[],"cooldown":[],"items":[{"id":"recover","unit":"seconds","value":60,"slot":"main","rest":90}]}]}`
	p, err := ParseResolved(raw)
	if err != nil {
		t.Fatal(err)
	}
	q, err := ParseResolved(p.Marshal())
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(p, q) {
		t.Fatal("export/import changed the prescription")
	}
	w := q.Workout(1, 3, Rests{Warmup: 5, Main: 18, Cooldown: 5})
	if !w.FixedDifficulty || len(w.Items) != 4 {
		t.Fatalf("workout: %+v", w)
	}
	for i, rest := range []int{15, 90, 0, 0} {
		if w.Items[i].Rest != rest {
			t.Errorf("item %d rest=%d want %d", i, w.Items[i].Rest, rest)
		}
	}
	if w.Items[1].Value != 6 || w.Items[1].SideRest != 10 {
		t.Fatalf("prescribed dose changed: %+v", w.Items[1])
	}
	if w.EstSec != 45+15+(6*4*2+10)*2+90+20 {
		t.Fatalf("duration=%d", w.EstSec)
	}
	recovery := q.Workout(2, 3, Rests{Warmup: 5, Main: 18, Cooldown: 5})
	if len(recovery.Items) != 1 || recovery.EstSec != 60 || recovery.Items[0].Rest != 0 {
		t.Fatalf("recovery inherited phases or final rest: %+v", recovery)
	}
}

func TestLegacyPhasesPreserveDayDifferences(t *testing.T) {
	p, err := ParseResolved(`{"days":[{"items":[{"id":"a","slot":"warmup","value":10},{"id":"m","slot":"main","value":5},{"id":"z","slot":"cooldown","value":20}]},{"items":[{"id":"b","slot":"warmup","value":15},{"id":"m","slot":"main","value":6}]},{"items":[{"id":"m","slot":"main","value":7}]}]}`)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Warmup) != 1 || p.Warmup[0].ID != "a" || len(p.Cooldown) != 1 {
		t.Fatal("legacy shared phases not migrated")
	}
	if p.Days[1].Warmup == nil || (*p.Days[1].Warmup)[0].ID != "b" || p.Days[1].Cooldown == nil || len(*p.Days[1].Cooldown) != 0 {
		t.Fatal("different legacy phases were lost")
	}
	if len(p.Workout(3, 0, Rests{}).Items) != 1 {
		t.Fatal("legacy recovery day inherited unrelated phases")
	}
}

func TestRestOverridesDoNotChangeLegacyDefaults(t *testing.T) {
	p := ResolveBuiltin()
	w := p.Workout(1, 0, Rests{Warmup: 5, Main: 18, Cooldown: 0})
	for i, it := range w.Items[:len(w.Items)-1] {
		if it.Rest != restAfter(w.Items[i+1].Slot, Rests{Warmup: 5, Main: 18, Cooldown: 0}) {
			t.Fatalf("legacy rest changed at %d", i)
		}
	}
	tooHigh, tooLow := 150, -1
	p = Resolved{Days: []RDay{{Rounds: 1, Items: []RItem{{ID: "a", Value: 1, Rest: &tooHigh}, {ID: "b", Value: 1, Rest: &tooLow}, {ID: "c", Value: 1}}}}}
	w = p.Workout(1, 0, Rests{Main: 18})
	if w.Items[0].Rest != 90 || w.Items[1].Rest != 0 {
		t.Fatal("invalid explicit rest was not bounded")
	}
}

func TestRoundCollapseDoesNotLosePrescriptions(t *testing.T) {
	a, b := 10, 20
	items := []RItem{{ID: "a", Value: 10, Round: 1, Rest: &a}, {ID: "a", Value: 10, Round: 2, Rest: &b}}
	got, rounds := collapseMainRounds(items)
	if rounds != 1 || len(got) != 2 {
		t.Fatal("different prescribed rounds were collapsed")
	}
}
