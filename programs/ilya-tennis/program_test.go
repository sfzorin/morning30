package ilyaten_test

import (
	"bytes"
	"encoding/json"
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
	for _, field := range []string{`"rest"`, `"sideRest"`} {
		if bytes.Contains(raw, []byte(field)) {
			t.Fatalf("program must not own pauses: %s", field)
		}
	}
	if p.WarmupRounds != 1 || len(p.Warmup) != 5 || len(p.Cooldown) != 3 {
		t.Fatal("every day needs the same warm-up and cool-down once")
	}
	wantMain := []int{6, 5, 8, 6, 5, 8}
	pushups := []int{6, 8, 10, 12, 14}
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
		if day.Rounds != 2 || len(day.Items) != wantMain[di%6] {
			t.Fatalf("day %d must have two complete rounds", di+1)
		}
		seen := map[string]bool{}
		for n, it := range day.Items {
			if seen[it.ID] {
				t.Fatalf("day %d repeats %s inside one round", di+1, it.ID)
			}
			seen[it.ID] = true
			if it.ID == "IT09" && it.Value != pushups[di/6] {
				t.Fatalf("day %d push-ups=%d", di+1, it.Value)
			}
			if di >= 6 {
				before := p.Days[di-6].Items[n]
				if it.ID != before.ID || it.Value < before.Value {
					t.Fatalf("day %d breaks the planned progression", di+1)
				}
			}
		}
		if di%6 == 2 || di%6 == 5 {
			if !reflect.DeepEqual(day.Items, p.Days[2].Items) {
				t.Fatalf("day %d must remain light", di+1)
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

func TestEveryExerciseHasAnAlignedDemonstration(t *testing.T) {
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
	if len(docs) != 33 {
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
	for _, id := range []string{"IT02", "IT04", "IT05", "IT06", "IT16", "IT18", "IT29", "IT30"} {
		if ids[id].Unit != "reps" {
			t.Errorf("countable dynamic exercise %s must use repetitions", id)
		}
	}
	for _, id := range []string{"IT01", "IT13", "IT14", "IT15", "IT17", "IT23", "IT28", "IT31"} {
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
