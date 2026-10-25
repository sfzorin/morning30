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

func TestPrescribedSixDayCycle(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	if !p.FixedDifficulty || !p.SelfManagedSides || len(p.Days) != 30 {
		t.Fatal("expected fixed 30-day program with self-managed sides")
	}
	if bytes.Contains(raw, []byte(`"sideRest"`)) {
		t.Fatal("Ilya manages side changes without separate pauses")
	}
	wantDuration := []int{1580, 1238, 761, 1628, 1238, 761}
	wantMain := []int{6, 5, 8, 6, 5, 8}
	// Combined main-block dose across both rounds and both sides. Reordering
	// into circuits must not multiply the earlier two-set workload again.
	wantDose := []map[string]int{
		{"IT07": 12, "IT08": 24, "IT09": 12, "IT10": 16, "IT11": 32, "IT12": 20},
		{"IT15": 60, "IT16": 8, "IT17": 80, "IT18": 40, "IT19": 24},
		{"IT01": 120, "IT03": 16, "IT20": 8, "IT21": 12, "IT22": 8, "IT13": 80, "IT23": 80, "IT14": 60},
		{"IT24": 12, "IT25": 20, "IT09": 12, "IT26": 20, "IT27": 24, "IT28": 60},
		{"IT29": 12, "IT30": 20, "IT31": 60, "IT32": 12, "IT33": 24},
		{"IT01": 120, "IT03": 16, "IT20": 8, "IT21": 12, "IT22": 8, "IT13": 80, "IT23": 80, "IT14": 60},
	}
	for i, day := range p.Days {
		if day.Rounds != 2 || len(day.Items) != wantMain[i%6] {
			t.Fatalf("day %d: expected two complete circuits", i+1)
		}
		seen := map[string]bool{}
		for _, it := range day.Items {
			if seen[it.ID] {
				t.Fatalf("day %d repeats %s inside one round", i+1, it.ID)
			}
			seen[it.ID] = true
			if it.Rest == nil || *it.Rest <= 0 {
				t.Fatalf("day %d needs a prescribed pause after %s, including the round boundary", i+1, it.ID)
			}
		}
		workout := p.Workout(i+1, 3, content.Rests{Warmup: 40, Main: 40, Cooldown: 40})
		if !workout.SelfManagedSides {
			t.Fatal("player must receive self-managed sides")
		}
		dose := map[string]int{}
		main := []content.Item{}
		for _, it := range workout.Items {
			if it.SideRest != 0 {
				t.Fatalf("unexpected side pause in %s", it.ExerciseID)
			}
			if it.Slot != content.Main {
				continue
			}
			main = append(main, it)
			n := it.Value
			if it.PerSide {
				n *= 2
			}
			dose[it.ExerciseID] += n
		}
		if !reflect.DeepEqual(dose, wantDose[i%6]) {
			t.Fatalf("day %d main dose=%v want %v", i+1, dose, wantDose[i%6])
		}
		if len(main) != 2*len(day.Items) {
			t.Fatalf("day %d expanded main count", i+1)
		}
		for n, it := range main {
			if it.ExerciseID != day.Items[n%len(day.Items)].ID || it.Round != n/len(day.Items)+1 {
				t.Fatalf("day %d must run A→B→C, then repeat the round", i+1)
			}
		}
		if workout.EstSec != wantDuration[i%6] {
			t.Fatalf("day %d duration=%d want %d", i+1, workout.EstSec, wantDuration[i%6])
		}
		if i%3 == 2 && (len(p.WarmupFor(i)) != 0 || len(p.CooldownFor(i)) != 0) {
			t.Fatalf("recovery day %d inherited phases", i+1)
		}
		if i >= 6 && !reflect.DeepEqual(day.Items, p.Days[i%6].Items) {
			t.Fatalf("day %d adds unapproved progression", i+1)
		}
		base := p.Workout(i+1, 0, content.Rests{})
		if !reflect.DeepEqual(workout, base) {
			t.Fatalf("day %d affected by level/rest preferences", i+1)
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
