package fitness

import (
	"testing"

	"morning30/internal/content"
)

func TestIndexMatchesNorms(t *testing.T) {
	man40 := Profile{Sex: "m", Age: 41}
	// Exactly the typical results score 50; twice that caps at 100.
	if got, ok := Index(man40, Results{16, 32, 70, 5}); !ok || got != 50 {
		t.Fatalf("typical man 40-49 = %d %v, want 50", got, ok)
	}
	if got, _ := Index(man40, Results{40, 80, 200, 20}); got != 100 {
		t.Fatalf("strong man = %d, want 100", got)
	}
	// A skipped test does not count against the person.
	if got, _ := Index(man40, Results{16, 32, 70, -1}); got != 50 {
		t.Fatalf("skipped pull-ups = %d, want 50", got)
	}
	if _, ok := Index(man40, Results{-1, -1, -1, -1}); ok {
		t.Fatal("no tests must give no index")
	}
	// The same result means more for an older woman.
	young, _ := Index(Profile{Sex: "m", Age: 25}, Results{12, 30, 60, -1})
	older, _ := Index(Profile{Sex: "f", Age: 62}, Results{12, 30, 60, -1})
	if older <= young {
		t.Fatalf("older woman %d should score above young man %d", older, young)
	}
	// Children have their own norms.
	kid, _ := Index(Profile{Child: true, Sex: "b", Age: 8}, Results{8, 25, 30, 1})
	if kid != 50 {
		t.Fatalf("typical boy = %d, want 50", kid)
	}
}

func TestLevels(t *testing.T) {
	for index, want := range map[int]string{0: "low", 34: "low", 35: "fair", 54: "fair", 55: "good", 74: "good", 75: "great", 100: "great"} {
		if got := Level(index); got != want {
			t.Errorf("Level(%d) = %s, want %s", index, got, want)
		}
	}
}

func TestRecommend(t *testing.T) {
	cases := []struct {
		name  string
		p     Profile
		index int
		has   bool
		want  Plan
	}{
		{"fit adult", Profile{Sex: "m", Age: 35}, 80, true, Plan{"demo", 2, 3}},
		{"good with back", Profile{Sex: "m", Age: 41, Concerns: []string{Back}}, 57, true, Plan{"backcare", 1, 2}},
		{"beginner", Profile{Sex: "f", Age: 30}, 20, true, Plan{"demo", -2, 2}},
		{"no test", Profile{Sex: "f", Age: 30, Concerns: []string{Feet}}, 0, false, Plan{"feet", 0, 2}},
		{"over 60 stays gentle", Profile{Sex: "m", Age: 66}, 90, true, Plan{"age50", 0, 2}},
		{"pressure", Profile{Sex: "f", Age: 45, Concerns: []string{Pressure}}, 60, true, Plan{"age50", 0, 2}},
		{"knees and weight", Profile{Sex: "m", Age: 38, Height: 175, Weight: 105, Concerns: []string{Knees}}, 40, true, Plan{"demo", -2, 2}},
		{"child posture", Profile{Child: true, Sex: "g", Age: 11, Concerns: []string{Back}}, 50, true, Plan{"kidsposture", 0, 2}},
		{"child", Profile{Child: true, Sex: "b", Age: 7}, 80, true, Plan{"kids30", 2, 3}},
	}
	for _, c := range cases {
		if got := Recommend(c.p, c.index, c.has); got != c.want {
			t.Errorf("%s: %+v, want %+v", c.name, got, c.want)
		}
	}
}

// Every recommended program exists in the catalog.
func TestRecommendedProgramsExist(t *testing.T) {
	for _, child := range []bool{false, true} {
		for _, c := range append([]string{""}, Concerns...) {
			for _, age := range []int{8, 30, 55} {
				key := program(Profile{Child: child, Age: age, Concerns: []string{c}})
				if _, ok := content.ReadyByKey(key); !ok {
					t.Errorf("program %q is not in the catalog", key)
				}
			}
		}
	}
}

func TestParseConcerns(t *testing.T) {
	got := ParseConcerns("feet, back,unknown")
	if len(got) != 2 || got[0] != Back || got[1] != Feet {
		t.Fatalf("ParseConcerns = %v", got)
	}
}
