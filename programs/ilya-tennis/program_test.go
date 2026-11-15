package ilyaten_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"morning30/internal/catalog"
	"morning30/internal/content"
)

func TestSpokenMistakeCuesUseNegativeInstructions(t *testing.T) {
	raw, err := os.ReadFile("exercises.json")
	if err != nil {
		t.Fatal(err)
	}
	var docs []catalog.Doc
	if err := json.Unmarshal(raw, &docs); err != nil {
		t.Fatal(err)
	}
	for _, ex := range content.Library() {
		doc, ok := catalog.Builtin(ex.ID)
		if !ok {
			t.Fatalf("missing built-in %s", ex.ID)
		}
		docs = append(docs, doc)
	}
	// These are authoring conventions, not a runtime grammar converter. Keep
	// reviewing compound cues too: each prohibited action must stay negated.
	patterns := map[string]*regexp.Regexp{
		"ru": regexp.MustCompile(`^Не `),
		"en": regexp.MustCompile(`^(Don't |Don’t |Do not |Never )`),
		"tr": regexp.MustCompile(`(ma|me|mayın|meyin|masın|mesin)[.!]?$`),
		"de": regexp.MustCompile(`(?i)\b(nicht|kein\w*)\b`),
		"es": regexp.MustCompile(`^No `),
		"fr": regexp.MustCompile(`^(Ne |N['’])`),
		"it": regexp.MustCompile(`^Non `),
	}
	for _, doc := range docs {
		for lang, pattern := range patterns {
			cues := doc.Content[lang].Wrong
			if len(cues) == 0 {
				t.Errorf("%s %s has no mistake cues", doc.ID, lang)
			}
			for _, cue := range cues {
				// Built-ins without a translation currently fall back to English.
				fallback := !strings.HasPrefix(doc.ID, "IT") && patterns["en"].MatchString(cue)
				if !pattern.MatchString(cue) && !fallback {
					t.Errorf("%s %s must phrase the mistake as a negative instruction: %q", doc.ID, lang, cue)
				}
			}
		}
	}
}

func TestThirtyDayProgressionUsesStandardWorkoutRules(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	if p.FixedDifficulty || !p.SelfManagedSides || len(p.Days) != 30 {
		t.Fatal("expected ordinary difficulty controls and self-managed sides over 30 days")
	}
	for _, field := range []string{`"rest"`, `"sideRest"`, `"progression"`, `"effort"`} {
		if bytes.Contains(raw, []byte(field)) {
			t.Fatalf("program must use the existing workout system: %s", field)
		}
	}
	if p.WarmupRounds != 1 || len(p.Warmup) != 6 || len(p.Cooldown) != 6 {
		t.Fatal("every day needs six warm-up and six cool-down exercises once")
	}
	wantScreens := []int{
		22, 24, 20, 24, 22, 20,
		24, 26, 20, 26, 24, 22,
		26, 28, 22, 28, 26, 20,
		28, 30, 22, 30, 28, 22,
		30, 32, 22, 30, 22, 32,
	}
	rests := content.Rests{Warmup: 15, Main: 20, Cooldown: 10}
	faster := content.Rests{Warmup: 0, Main: 5, Cooldown: 0}
	phaseRest := func(slot content.Slot, r content.Rests) int {
		switch slot {
		case content.Warmup:
			return r.Warmup
		case content.Cooldown:
			return r.Cooldown
		default:
			return r.Main
		}
	}
	for di, day := range p.Days {
		if day.Rounds != 2 || len(day.Items) < 4 || len(day.Items) > 10 {
			t.Fatalf("day %d must have two complete rounds of 4–10 stations", di+1)
		}
		seen := map[string]bool{}
		for _, it := range day.Items {
			if seen[it.ID] {
				t.Fatalf("day %d repeats %s inside one round", di+1, it.ID)
			}
			seen[it.ID] = true
		}
		w := p.Workout(di+1, 0, rests)
		quick := p.Workout(di+1, 0, faster)
		if !w.SelfManagedSides {
			t.Fatal("both sides must share one exercise screen")
		}
		warmup, cooldown := p.WarmupFor(di), p.CooldownFor(di)
		if len(warmup) != 6 || len(cooldown) != 6 {
			t.Fatalf("day %d has incomplete preparation or stretching", di+1)
		}
		want := append([]content.RItem{}, warmup...)
		want = append(want, day.Items...)
		want = append(want, day.Items...)
		want = append(want, cooldown...)
		if len(w.Items) != len(want) || len(w.Items) != wantScreens[di] {
			t.Fatalf("day %d has missing phases or incorrect round expansion", di+1)
		}
		if w.EstSec > 31*60 {
			t.Errorf("day %d exceeds the 31-minute estimate at the current rest settings: %d seconds", di+1, w.EstSec)
		}
		delta := 0
		for n, it := range w.Items {
			round := 1
			if n >= len(warmup)+len(day.Items) && n < len(warmup)+2*len(day.Items) {
				round = 2
			}
			if it.ExerciseID != want[n].ID || it.Value != want[n].Value || it.Slot != want[n].Slot || it.Round != round {
				t.Fatalf("day %d item %d must retain prescribed order, dose and circuit", di+1, n)
			}
			expectedRest := 0
			if n+1 < len(w.Items) {
				expectedRest = phaseRest(w.Items[n+1].Slot, rests)
				delta += expectedRest - phaseRest(w.Items[n+1].Slot, faster)
			}
			if it.Rest != expectedRest {
				t.Fatalf("day %d item %d ignores rest preferences", di+1, n)
			}
		}
		if w.EstSec-quick.EstSec != delta {
			t.Fatalf("day %d estimated duration ignores rest preferences", di+1)
		}
		harder := p.Workout(di+1, 1, rests)
		for n, it := range harder.Items {
			expected := w.Items[n].Value
			if it.Slot == content.Main {
				expected = content.ScaleValue(it.Unit, expected, 1)
			}
			if it.Value != expected {
				t.Fatalf("day %d difficulty control ignored", di+1)
			}
		}
	}
}

func TestPublishedDemonstrationsAreAligned(t *testing.T) {
	raw, err := os.ReadFile("exercises.json")
	if err != nil {
		t.Fatal(err)
	}
	var docs []catalog.Doc
	if err := json.Unmarshal(raw, &docs); err != nil {
		t.Fatal(err)
	}
	raw, err = os.ReadFile("images.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest []struct {
		ID   string `json:"id"`
		File string `json:"file"`
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatal(err)
	}
	images := map[string]string{}
	for _, entry := range manifest {
		images[entry.ID] = entry.File
	}
	for _, doc := range docs {
		a := doc.Animation
		if a == nil || !catalog.ValidAnimation(a) {
			t.Fatalf("missing/invalid demonstration: %s", doc.ID)
		}
		if doc.PerSide && a.Kind != "alternating" && a.OtherPose == "" {
			t.Errorf("%s must demonstrate both sides", doc.ID)
		}
		for _, pose := range []string{a.Pose, a.OtherPose} {
			if pose == "" {
				continue
			}
			path := images[pose]
			if !strings.HasPrefix(path, "animations/aligned/") {
				t.Fatalf("%s uses an unaligned/missing strip: %s", doc.ID, path)
			}
			file, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			cfg, err := png.DecodeConfig(file)
			file.Close()
			if err != nil {
				t.Fatal(err)
			}
			if cfg.Height != 640 || cfg.Width != a.Frames*cfg.Height {
				t.Fatalf("%s has unequal frames", doc.ID)
			}
		}
	}
}

func TestEveryExerciseHasSevenCompleteLanguages(t *testing.T) {
	raw, err := os.ReadFile("exercises.json")
	if err != nil {
		t.Fatal(err)
	}
	var docs []catalog.Doc
	if err := json.Unmarshal(raw, &docs); err != nil {
		t.Fatal(err)
	}
	if len(docs) != 54 {
		t.Fatalf("got %d exercises", len(docs))
	}
	ids := map[string]catalog.Doc{}
	for _, doc := range docs {
		if !catalog.ValidDoc(doc) {
			t.Fatalf("invalid %s", doc.ID)
		}
		if _, exists := ids[doc.ID]; exists {
			t.Fatalf("duplicate %s", doc.ID)
		}
		ids[doc.ID] = doc
		for _, lang := range []string{"ru", "en", "tr", "de", "es", "fr", "it"} {
			x := doc.Content[lang]
			if doc.Unit == "reps" {
				for _, step := range x.HowTo {
					for _, signal := range []string{"сигнал", "signal", "işaret", "señal", "segnale"} {
						if strings.Contains(strings.ToLower(step), signal) {
							t.Errorf("%s %s tells a counted exercise to wait for a timer: %s", doc.ID, lang, step)
						}
					}
				}
			}
			if x.Name == "" || x.Hint == "" || x.Desc == "" || len(x.HowTo) == 0 || len(x.Correct) == 0 || len(x.Wrong) == 0 || x.Breathing == "" || x.Safety == "" || x.Voice.Start == "" || len(x.Voice.Mid) == 0 || x.Voice.Finish == "" {
				t.Errorf("incomplete %s %s", doc.ID, lang)
			}
		}
	}
	for n := 1; n <= 54; n++ {
		if _, exists := ids[fmt.Sprintf("IT%02d", n)]; !exists {
			t.Errorf("missing exercise IT%02d", n)
		}
	}
	for _, id := range []string{"IT02", "IT04", "IT05", "IT06", "IT08", "IT09", "IT16", "IT18", "IT29", "IT30", "IT38", "IT39", "IT40", "IT42", "IT44", "IT45", "IT46", "IT47", "IT48"} {
		if ids[id].Unit != "reps" {
			t.Errorf("countable dynamic exercise %s must use repetitions", id)
		}
	}
	for _, id := range []string{"IT01", "IT13", "IT14", "IT15", "IT17", "IT23", "IT28", "IT31", "IT34", "IT35", "IT36", "IT37", "IT41", "IT43", "IT49", "IT50", "IT51", "IT52", "IT53", "IT54"} {
		if ids[id].Unit != "seconds" {
			t.Errorf("continuous work or static hold %s must retain its duration", id)
		}
	}
	raw, _ = os.ReadFile("program.json")
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	for day := 1; day <= 30; day++ {
		for _, it := range p.Workout(day, 0, content.Rests{}).Items {
			doc, ok := ids[it.ExerciseID]
			if !ok || doc.Unit != string(it.Unit) || doc.PerSide != it.PerSide {
				t.Errorf("definition mismatch %s", it.ExerciseID)
			}
		}
	}
}

func TestDailyPhasesIncludePreparationAndStretching(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	type dose struct {
		id    string
		unit  content.Unit
		value int
		each  bool
	}
	warmupA := []dose{
		{"IT37", content.Seconds, 40, false},
		{"IT02", content.Reps, 12, false},
		{"IT04", content.Reps, 12, false},
		{"IT08", content.Reps, 4, true},
		{"IT19", content.Reps, 6, true},
		{"IT06", content.Reps, 8, false},
	}
	warmupB := []dose{
		{"IT37", content.Seconds, 40, false},
		{"IT02", content.Reps, 12, false},
		{"IT04", content.Reps, 12, false},
		{"IT05", content.Reps, 8, false},
		{"IT18", content.Reps, 16, false},
		{"IT06", content.Reps, 8, false},
	}
	cooldown := []dose{
		{"IT13", content.Seconds, 20, true},
		{"IT51", content.Seconds, 20, true},
		{"IT50", content.Seconds, 20, true},
		{"IT52", content.Seconds, 20, true},
		{"IT54", content.Seconds, 15, true},
		{"IT53", content.Seconds, 30, false},
	}
	check := func(day int, phase string, got []content.RItem, want []dose) {
		t.Helper()
		if len(got) != len(want) {
			t.Fatalf("day %d %s: got %d exercises", day, phase, len(got))
		}
		for i, it := range got {
			w := want[i]
			if it.ID != w.id || it.Unit != w.unit || it.Value != w.value || it.PerSide != w.each {
				t.Fatalf("day %d %s item %d: wrong exercise or dose", day, phase, i+1)
			}
		}
	}
	for day := range p.Days {
		warmup := warmupA
		if p.Days[day].Warmup != nil {
			warmup = append([]dose{}, warmupB...)
			if isRecoveryDay(day + 1) {
				warmup[0].value = 30
			}
			for _, it := range p.Days[day].Items {
				if it.ID == "IT18" {
					warmup[4] = dose{"IT19", content.Reps, 4, true}
				}
			}
		}
		check(day+1, "warm-up", p.WarmupFor(day), warmup)
		check(day+1, "cool-down", p.CooldownFor(day), cooldown)
	}
}

func isRecoveryDay(day int) bool {
	switch day {
	case 3, 6, 9, 12, 15, 18, 21, 24, 27, 29:
		return true
	}
	return false
}

func TestDailyStrengthVariesAndProgressesWithRecovery(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	hardPushups := map[int]int{2: 6, 4: 6, 8: 8, 10: 8, 14: 10, 16: 10, 20: 12, 22: 12, 26: 14, 28: 14, 30: 14}
	highSpeed := map[string]bool{"IT34": true, "IT35": true, "IT36": true, "IT37": true, "IT43": true, "IT49": true, "IT07": true, "IT24": true}
	sequences := map[string]int{}
	var blockSeconds [5]int
	var blockStations [5]int
	for di, day := range p.Days {
		ids := make([]string, 0, len(day.Items))
		pushups, planks, pushupReps := 0, 0, 0
		for n, it := range day.Items {
			ids = append(ids, it.ID)
			if (it.ID == "IT07" || it.ID == "IT24") && n != 0 {
				t.Errorf("day %d must put landing practice before fatigue", di+1)
			}
			if isRecoveryDay(di+1) && highSpeed[it.ID] {
				t.Errorf("recovery day %d includes high-speed work %s", di+1, it.ID)
			}
			switch it.ID {
			case "IT09":
				pushups++
				pushupReps = it.Value
				if it.Unit != content.Reps || it.PerSide {
					t.Errorf("day %d must count push-ups", di+1)
				}
			case "IT41":
				planks++
				if it.Unit != content.Seconds || !it.PerSide {
					t.Errorf("day %d side plank must hold both sides", di+1)
				}
			case "IT39", "IT40", "IT48":
				planks++
				if it.Unit != content.Reps || it.PerSide {
					t.Errorf("day %d moving plank must count alternating actions total", di+1)
				}
			}
		}
		if pushups != 1 || planks != 1 {
			t.Errorf("day %d needs one push-up and one plank station per round; got %d and %d", di+1, pushups, planks)
		}
		if want, hard := hardPushups[di+1]; hard && pushupReps != want {
			t.Errorf("hard day %d push-ups: got %d, want %d", di+1, pushupReps, want)
		}
		if isRecoveryDay(di + 1) {
			hardTarget := 6 + 2*(di/6)
			if pushupReps < 4 || pushupReps > 6 || pushupReps >= hardTarget {
				t.Errorf("recovery day %d must keep fewer push-ups than the hard days: %d", di+1, pushupReps)
			}
			if len(day.Items) > 5 {
				t.Errorf("recovery day %d has too many stations: %d", di+1, len(day.Items))
			}
		}
		sequence := strings.Join(ids, ",")
		if previous, exists := sequences[sequence]; exists {
			t.Errorf("day %d repeats the exercise sequence of day %d", di+1, previous)
		}
		sequences[sequence] = di + 1
		blockStations[di/6] += len(day.Items)
		blockSeconds[di/6] += p.Workout(di+1, 0, content.Rests{Warmup: 15, Main: 20, Cooldown: 10}).EstSec
	}
	for block := 1; block < 5; block++ {
		if blockStations[block] <= blockStations[block-1] || blockSeconds[block] <= blockSeconds[block-1] {
			t.Errorf("block %d must increase average session volume and duration", block+1)
		}
	}
}

func TestReactionDrillKeepsAcceleratingCues(t *testing.T) {
	raw, err := os.ReadFile("exercises.json")
	if err != nil {
		t.Fatal(err)
	}
	var docs []map[string]any
	if err := json.Unmarshal(raw, &docs); err != nil {
		t.Fatal(err)
	}
	for _, doc := range docs {
		if doc["id"] == "IT36" {
			want := map[string]any{
				"directions": []any{"left", "right"},
				"cadence": map[string]any{
					"startSeconds": float64(3),
					"endSeconds":   float64(1),
				},
			}
			if !reflect.DeepEqual(doc["reaction"], want) {
				t.Fatal("reaction drill needs app-driven left/right cues accelerating from 3 to 1 second")
			}
		} else if _, hasReaction := doc["reaction"]; hasReaction {
			t.Fatalf("unexpected reaction drill: %s", doc["id"])
		}
	}
}
