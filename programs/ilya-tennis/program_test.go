package ilyaten_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"morning30/internal/catalog"
	"morning30/internal/content"
)

func TestPrescribedSixDayCycle(t *testing.T) {
	raw, err := os.ReadFile("program.json")
	if err != nil {
		t.Fatal(err)
	}
	p, err := content.ParseResolved(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	if !p.FixedDifficulty || len(p.Days) != 30 {
		t.Fatal("expected fixed 30-day program")
	}
	wantDuration := []int{1644, 1132, 731, 1612, 1132, 731}
	wantMain := []int{13, 11, 10, 12, 11, 10}
	for i, day := range p.Days {
		if day.Rounds != 1 || len(day.Items) != wantMain[i%6] {
			t.Fatalf("day %d: set order/count changed", i+1)
		}
		workout := p.Workout(i+1, 3, content.Rests{Warmup: 40, Main: 40, Cooldown: 40})
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
	if p.Days[0].Items[0].ID != "IT07" || p.Days[0].Items[1].ID != "IT07" || p.Days[0].Items[2].ID != "IT07" {
		t.Fatal("sets must be consecutive")
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
			if x.Name == "" || x.Hint == "" || x.Desc == "" || len(x.HowTo) == 0 || len(x.Correct) == 0 || len(x.Wrong) == 0 || x.Breathing == "" || x.Safety == "" || x.Voice.Start == "" || len(x.Voice.Mid) == 0 || x.Voice.Finish == "" {
				t.Errorf("incomplete %s %s", doc.ID, lang)
			}
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
