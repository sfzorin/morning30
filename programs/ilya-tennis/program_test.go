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
		t.Fatal("every day needs the same warm-up and cool-down once")
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
		if day.Rounds != 2 || len(day.Items) != 8 {
			t.Fatalf("day %d must have two complete rounds", di+1)
		}
		seen := map[string]bool{}
		for n, it := range day.Items {
			if seen[it.ID] {
				t.Fatalf("day %d repeats %s inside one round", di+1, it.ID)
			}
			seen[it.ID] = true
			if di >= 6 {
				before := p.Days[di-6].Items[n]
				if it.ID != before.ID || it.Value < before.Value {
					t.Fatalf("day %d breaks the planned progression", di+1)
				}
			}
		}
		w := p.Workout(di+1, 0, rests)
		quick := p.Workout(di+1, 0, faster)
		if !w.SelfManagedSides {
			t.Fatal("both sides must share one exercise screen")
		}
		want := append([]content.RItem{}, p.Warmup...)
		want = append(want, day.Items...)
		want = append(want, day.Items...)
		want = append(want, p.Cooldown...)
		if len(w.Items) != len(want) {
			t.Fatalf("day %d has missing phases or incorrect round expansion", di+1)
		}
		delta := 0
		for n, it := range w.Items {
			round := 1
			if n >= len(p.Warmup)+len(day.Items) && n < len(p.Warmup)+2*len(day.Items) {
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
		// New definitions are published before their images. Existing published
		// demonstrations remain required; any newly supplied strip is validated.
		if a == nil && doc.ID >= "IT37" && doc.ID <= "IT54" {
			continue
		}
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
	warmup := []dose{
		{"IT37", content.Seconds, 40, false},
		{"IT02", content.Reps, 12, false},
		{"IT04", content.Reps, 12, false},
		{"IT08", content.Reps, 4, true},
		{"IT19", content.Reps, 6, true},
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
		check(day+1, "warm-up", p.WarmupFor(day), warmup)
		check(day+1, "cool-down", p.CooldownFor(day), cooldown)
	}
}

func TestEveryDayHasTheApprovedStrengthAndMovementSequence(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	orders := [][]string{
		{"IT34", "IT09", "IT08", "IT37", "IT39", "IT38", "IT35", "IT36"},
		{"IT35", "IT09", "IT42", "IT37", "IT40", "IT27", "IT11", "IT36"},
		{"IT43", "IT09", "IT44", "IT34", "IT41", "IT45", "IT46", "IT36"},
		{"IT24", "IT09", "IT47", "IT35", "IT39", "IT27", "IT34", "IT36"},
		{"IT49", "IT09", "IT08", "IT37", "IT48", "IT27", "IT11", "IT36"},
		{"IT34", "IT09", "IT42", "IT35", "IT41", "IT45", "IT46", "IT36"},
	}
	// Five block targets per exercise, in the same order as each daily circuit.
	doses := [6][8][5]int{
		{{45, 45, 50, 50, 50}, {6, 8, 10, 12, 14}, {8, 8, 9, 9, 10}, {45, 45, 50, 50, 50}, {10, 10, 12, 12, 14}, {12, 12, 14, 14, 16}, {45, 45, 50, 50, 50}, {45, 45, 50, 50, 50}},
		{{40, 40, 40, 45, 45}, {5, 5, 6, 6, 7}, {6, 6, 7, 7, 8}, {40, 40, 40, 45, 45}, {8, 8, 10, 10, 12}, {8, 8, 9, 9, 10}, {12, 12, 12, 12, 12}, {40, 40, 40, 45, 45}},
		{{40, 40, 40, 45, 45}, {5, 5, 6, 6, 7}, {12, 12, 14, 14, 16}, {40, 40, 40, 45, 45}, {20, 20, 25, 25, 30}, {6, 6, 6, 6, 6}, {12, 12, 12, 12, 12}, {40, 40, 40, 45, 45}},
		{{6, 6, 6, 6, 6}, {6, 8, 10, 12, 14}, {8, 8, 9, 9, 10}, {45, 45, 50, 50, 50}, {12, 12, 14, 14, 16}, {10, 10, 11, 11, 12}, {45, 45, 50, 50, 50}, {45, 45, 50, 50, 50}},
		{{40, 40, 40, 45, 45}, {5, 5, 6, 6, 7}, {6, 6, 7, 7, 8}, {40, 40, 40, 45, 45}, {8, 8, 10, 10, 12}, {8, 8, 9, 9, 10}, {12, 12, 12, 12, 12}, {40, 40, 40, 45, 45}},
		{{40, 40, 40, 45, 45}, {5, 5, 6, 6, 7}, {6, 6, 7, 7, 8}, {40, 40, 40, 45, 45}, {20, 20, 25, 25, 30}, {6, 6, 6, 6, 6}, {12, 12, 12, 12, 12}, {40, 40, 40, 45, 45}},
	}
	for di, day := range p.Days {
		if len(day.Items) != 8 {
			t.Fatalf("day %d must have eight stations", di+1)
		}
		for n, it := range day.Items {
			if it.ID != orders[di%6][n] || it.Value != doses[di%6][n][di/6] {
				t.Errorf("day %d station %d: got %s ×%d, want %s ×%d", di+1, n+1, it.ID, it.Value, orders[di%6][n], doses[di%6][n][di/6])
			}
		}
		if day.Items[1].Unit != content.Reps || day.Items[1].PerSide {
			t.Fatalf("day %d must count push-ups", di+1)
		}
		plank := day.Items[4]
		if plank.ID == "IT41" {
			if plank.Unit != content.Seconds || !plank.PerSide {
				t.Fatalf("day %d side plank must hold both sides", di+1)
			}
		} else if plank.Unit != content.Reps || plank.PerSide {
			t.Fatalf("day %d moving plank must count alternating actions total", di+1)
		}
	}
	raw, err = os.ReadFile("exercises.json")
	if err != nil {
		t.Fatal(err)
	}
	var docs []map[string]any
	if err := json.Unmarshal(raw, &docs); err != nil {
		t.Fatal(err)
	}
	for _, doc := range docs {
		if doc["id"] == "IT36" {
			want := map[string]any{"directions": []any{"left", "right"}}
			if !reflect.DeepEqual(doc["reaction"], want) {
				t.Fatal("reaction drill needs app-driven left/right cues")
			}
		} else if _, hasReaction := doc["reaction"]; hasReaction {
			t.Fatalf("unexpected reaction drill: %s", doc["id"])
		}
	}
}
