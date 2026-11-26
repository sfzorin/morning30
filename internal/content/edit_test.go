package content

import (
	"errors"
	"testing"
)

func count(items []RItem, id string) int {
	n := 0
	for _, it := range items {
		if it.ID == id {
			n++
		}
	}
	return n
}

func TestAddExerciseToMainOnChosenDays(t *testing.T) {
	r := ResolveDemo()
	before := r.Marshal()
	ex, _ := Get("C07")
	out, err := AddExercise(r, ex, Main, Days{List: []int{2, 5, 5}}, 0)
	if err != nil {
		t.Fatal(err)
	}
	if r.Marshal() != before {
		t.Fatal("input program was modified")
	}
	for i, d := range out.Days {
		want := 0
		if i == 1 || i == 4 {
			want = 1
		}
		if got := count(d.Items, "C07") - count(r.Days[i].Items, "C07"); got != want {
			t.Fatalf("day %d: added %d, want %d", i+1, got, want)
		}
	}
	last := out.Days[1].Items[len(out.Days[1].Items)-1]
	if last.Value != DefaultValue(Seconds) || !last.PerSide || last.Slot != Main {
		t.Fatalf("added item %+v", last)
	}
	if _, err := AddExercise(r, ex, Main, Days{List: []int{31}}, 5); !errors.Is(err, ErrNoDays) {
		t.Fatalf("day 31 accepted: %v", err)
	}
	if _, err := AddExercise(r, ex, "nope", AllDays(), 5); !errors.Is(err, ErrBadPart) {
		t.Fatalf("bad part accepted: %v", err)
	}
}

func TestAddToWarmupAllDaysEditsSharedPhaseOnly(t *testing.T) {
	r := ResolveDemo()
	ex, _ := Get("W08")
	out, err := AddExercise(r, ex, Warmup, AllDays(), 20)
	if err != nil {
		t.Fatal(err)
	}
	if count(out.Warmup, "W08") != 1 {
		t.Fatal("shared warm-up not extended")
	}
	for i := range out.Days {
		if out.Days[i].Warmup != nil {
			t.Fatalf("day %d got its own warm-up copy", i+1)
		}
		if count(out.WarmupFor(i), "W08") != 1 {
			t.Fatalf("day %d does not see the new warm-up exercise", i+1)
		}
	}
}

func TestAddToCooldownOneDayGivesThatDayItsOwnCopy(t *testing.T) {
	r := ResolveDemo()
	ex, _ := Get("CD04")
	out, err := AddExercise(r, ex, Cooldown, OneDay(3), 30)
	if err != nil {
		t.Fatal(err)
	}
	if count(out.CooldownFor(2), "CD04") != 1 || count(out.CooldownFor(3), "CD04") != 0 || count(out.Cooldown, "CD04") != 0 {
		t.Fatal("one-day cool-down edit leaked")
	}
	back, err := ParseResolved(out.Marshal())
	if err != nil || back.Marshal() != out.Marshal() {
		t.Fatalf("edited program does not round-trip: %v", err)
	}
}

func TestReplaceExerciseConvertsDoseAndKeepsPosition(t *testing.T) {
	r := ResolveDemo()  // C01 plank (seconds) closes most demo days
	to, _ := Get("C09") // shoulder taps (reps)
	out, n, err := ReplaceExercise(r, "C01", to, AllDays())
	if err != nil || n == 0 {
		t.Fatalf("replace: %d %v", n, err)
	}
	for i, d := range out.Days {
		if count(d.Items, "C01") != 0 {
			t.Fatalf("day %d still has C01", i+1)
		}
		for j, it := range d.Items {
			if r.Days[i].Items[j].ID == "C01" {
				if it.ID != "C09" || it.Unit != Reps || it.Value != ConvertValue(Seconds, Reps, r.Days[i].Items[j].Value) {
					t.Fatalf("day %d item %d: %+v", i+1, j, it)
				}
			}
		}
	}
	if _, _, err := ReplaceExercise(r, "C01", to, OneDay(0)); !errors.Is(err, ErrNoDays) {
		t.Fatalf("day 0 accepted: %v", err)
	}
	if _, _, err := ReplaceExercise(r, "ZZ9", to, AllDays()); !errors.Is(err, ErrNotInProgram) {
		t.Fatalf("missing exercise: %v", err)
	}
}

func TestReplaceInSharedWarmupOnOneDay(t *testing.T) {
	r := ResolveDemo()
	to, _ := Get("W08")
	out, n, err := ReplaceExercise(r, "W09", to, OneDay(7))
	if err != nil || n != 1 {
		t.Fatalf("replace: %d %v", n, err)
	}
	if count(out.WarmupFor(6), "W08") != 1 || count(out.WarmupFor(6), "W09") != 0 {
		t.Fatal("day 7 warm-up not replaced")
	}
	if count(out.Warmup, "W09") != 1 || out.Days[5].Warmup != nil {
		t.Fatal("other days changed")
	}
}

func TestRemoveExerciseNeverEmptiesAPart(t *testing.T) {
	r := ResolveDemo()
	out, n, err := RemoveExercise(r, "C01", Main, AllDays())
	if err != nil || n == 0 {
		t.Fatalf("remove: %d %v", n, err)
	}
	for i, d := range out.Days {
		if count(d.Items, "C01") != 0 || len(d.Items) == 0 {
			t.Fatalf("day %d: %+v", i+1, d.Items)
		}
	}
	// A one-exercise cool-down cannot lose its only exercise.
	single := ResolveDemo()
	single.Cooldown = single.Cooldown[:1]
	if _, _, err := RemoveExercise(single, single.Cooldown[0].ID, Cooldown, AllDays()); !errors.Is(err, ErrEmptyPhase) {
		t.Fatalf("emptied cool-down: %v", err)
	}
}

func TestCatalogTagsAndMatching(t *testing.T) {
	seen := map[string]bool{}
	for _, p := range Catalog() {
		if seen[p.Key] {
			t.Fatalf("duplicate key %s", p.Key)
		}
		seen[p.Key] = true
		for _, l := range series30Langs {
			if p.Title[l] == "" || p.About[l] == "" {
				t.Errorf("%s: missing %s text", p.Key, l)
			}
		}
		valid := false
		for _, c := range Categories() {
			valid = valid || c == p.Category
		}
		if !valid {
			t.Errorf("%s: unknown category %q", p.Key, p.Category)
		}
		r := p.Build()
		r.Name = "Renamed"
		if MatchReady(r) != p.Key || !SameAsReady(r, p.Key) {
			t.Errorf("%s: renamed copy not matched", p.Key)
		}
	}
	if Catalog()[0].Key != DemoKey {
		t.Fatal("morning workout must come first")
	}
	for _, key := range PopularKeys() {
		if _, ok := ReadyByKey(key); !ok {
			t.Errorf("popular %s missing", key)
		}
	}
	jumps := map[string]bool{}
	for _, p := range Catalog() {
		jumps[p.Key] = p.Tags().Jumps
		if p.Tags().Equipment != (p.Key == "pullups30") {
			t.Errorf("%s equipment = %v", p.Key, p.Tags().Equipment)
		}
	}
	for key, want := range map[string]bool{"demo": false, "plank30": false, "abs30": false, "posture30": false, "slim30": true, "legs30": true,
		"stretch30": false, "taichi30": false, "kidsagility": true, "jumpsports": true, "age50": false} {
		if jumps[key] != want {
			t.Errorf("%s jumps = %v, want %v", key, jumps[key], want)
		}
	}
	edited := ResolveDemo()
	edited.Days[0].Items[0].Value++
	if MatchReady(edited) != "" || SameAsReady(edited, DemoKey) {
		t.Fatal("edited program matched a ready one")
	}
}
