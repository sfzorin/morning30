// Package fitness turns a short self-test (push-ups, squats in a minute, plank,
// pull-ups) into a fitness index from 0 to 100 and into a starting point for a
// program: which ready program, how hard, how many rounds.
//
// Each test is compared with a typical result for the person's age and sex:
// the typical result scores 50 points, twice that scores 100. The index is the
// average over the tests that were done. The reference values are rounded
// figures from common fitness-test tables; they set a starting load, they are
// not a medical assessment.
package fitness

import (
	"math"
	"strings"
)

// Concerns a person can mark on the "About you" screen.
const (
	Back     = "back"
	Knees    = "knees"
	Pressure = "pressure"
	Feet     = "feet"
)

// Concerns lists the known concerns in screen order.
var Concerns = []string{Back, Knees, Pressure, Feet}

// Profile is what the person tells about themselves.
type Profile struct {
	Child    bool     // the program is for a child
	Sex      string   // m | f | b | g (as the figure setting)
	Age      int      // years, 0 = not given
	Height   int      // centimetres, 0 = not given
	Weight   int      // kilograms, 0 = not given
	Concerns []string // subset of Concerns
}

// Has reports whether the profile marks concern c.
func (p Profile) Has(c string) bool {
	for _, x := range p.Concerns {
		if x == c {
			return true
		}
	}
	return false
}

// Female reports whether the norms for women and girls apply.
func (p Profile) Female() bool { return p.Sex == "f" || p.Sex == "g" }

// Results of one test. A negative value means the test was skipped.
type Results struct {
	PushUps  int `json:"pushUps"`
	Squats   int `json:"squats"`   // in one minute
	PlankSec int `json:"plankSec"` // seconds held
	PullUps  int `json:"pullUps"`
}

// Done reports whether at least one test has a result.
func (r Results) Done() bool {
	return r.PushUps >= 0 || r.Squats >= 0 || r.PlankSec >= 0 || r.PullUps >= 0
}

// Clamp keeps results in a sane range (a typo must not give 100 points).
func (r Results) Clamp() Results {
	c := func(v, max int) int {
		if v < 0 {
			return -1
		}
		if v > max {
			return max
		}
		return v
	}
	return Results{PushUps: c(r.PushUps, 200), Squats: c(r.Squats, 150), PlankSec: c(r.PlankSec, 1200), PullUps: c(r.PullUps, 100)}
}

// norm is the typical result in each test for one age band and sex.
type norm struct{ push, squat, plank, pull float64 }

// Adult bands: under 30, 30–39, 40–49, 50–59, 60 and over.
var (
	menNorms   = []norm{{25, 40, 90, 8}, {20, 36, 80, 6}, {16, 32, 70, 5}, {12, 28, 60, 3}, {10, 24, 50, 2}}
	womenNorms = []norm{{15, 35, 75, 2}, {13, 32, 70, 1}, {11, 28, 60, 1}, {9, 24, 50, 1}, {7, 20, 40, 1}}
	// Children: 6–9, 10–13, 14–17.
	boyNorms  = []norm{{8, 25, 30, 1}, {15, 32, 50, 3}, {22, 40, 70, 6}}
	girlNorms = []norm{{6, 23, 30, 1}, {10, 30, 45, 1}, {14, 36, 60, 2}}
)

func (p Profile) norm() norm {
	if p.Child {
		band := 1 // unknown age: the middle band
		switch {
		case p.Age > 0 && p.Age <= 9:
			band = 0
		case p.Age >= 14:
			band = 2
		}
		if p.Female() {
			return girlNorms[band]
		}
		return boyNorms[band]
	}
	band := 1 // unknown age: 30–39
	switch {
	case p.Age > 0 && p.Age < 30:
		band = 0
	case p.Age >= 60:
		band = 4
	case p.Age >= 50:
		band = 3
	case p.Age >= 40:
		band = 2
	}
	if p.Female() {
		return womenNorms[band]
	}
	return menNorms[band]
}

func points(v int, typical float64) (float64, bool) {
	if v < 0 {
		return 0, false
	}
	return math.Min(100, 50*float64(v)/typical), true
}

// Index scores the results from 0 to 100; ok is false when no test was done.
func Index(p Profile, r Results) (index int, ok bool) {
	n := p.norm()
	r = r.Clamp()
	sum, count := 0.0, 0
	for _, t := range []struct {
		v int
		n float64
	}{{r.PushUps, n.push}, {r.Squats, n.squat}, {r.PlankSec, n.plank}, {r.PullUps, n.pull}} {
		if pts, done := points(t.v, t.n); done {
			sum += pts
			count++
		}
	}
	if count == 0 {
		return 0, false
	}
	return int(math.Round(sum / float64(count))), true
}

// Level names for an index: "low", "fair", "good", "great".
func Level(index int) string {
	switch {
	case index >= 75:
		return "great"
	case index >= 55:
		return "good"
	case index >= 35:
		return "fair"
	}
	return "low"
}

// BMI returns the body-mass index, or 0 when height or weight is missing.
func (p Profile) BMI() float64 {
	if p.Height < 100 || p.Weight < 15 {
		return 0
	}
	h := float64(p.Height) / 100
	return float64(p.Weight) / (h * h)
}

// Plan is where the program starts.
type Plan struct {
	Program string // ready-program key
	Level   int    // difficulty −3..+3
	Rounds  int    // rounds of the main part
}

// Recommend picks the program and its starting load. hasIndex is false when
// the person skipped the test: the load then stays at the defaults.
func Recommend(p Profile, index int, hasIndex bool) Plan {
	plan := Plan{Program: program(p), Level: 0, Rounds: 2}
	if hasIndex {
		switch Level(index) {
		case "low":
			plan.Level = -2
		case "fair":
			plan.Level = 0
		case "good":
			plan.Level = 1
		case "great":
			plan.Level, plan.Rounds = 2, 3
		}
	}
	// Go easier where the body asks for it.
	if !p.Child && p.BMI() >= 32 {
		plan.Level--
	}
	if p.Has(Knees) {
		plan.Level--
	}
	if p.Has(Pressure) || (!p.Child && p.Age >= 60) {
		plan.Level = min(plan.Level, 0)
		plan.Rounds = min(plan.Rounds, 2)
	}
	plan.Level = max(-3, min(3, plan.Level))
	return plan
}

func program(p Profile) string {
	if p.Child {
		switch {
		case p.Has(Back):
			return "kidsposture"
		case p.Has(Feet):
			return "kidsfeet"
		}
		return "kids30"
	}
	switch {
	case p.Has(Back):
		return "backcare"
	case p.Has(Feet):
		return "feet"
	case p.Has(Pressure) || p.Age >= 50:
		return "age50"
	}
	return "demo"
}

// ParseConcerns reads a comma-separated list, keeping known concerns only.
func ParseConcerns(s string) []string {
	var out []string
	for _, c := range Concerns {
		for _, part := range strings.Split(s, ",") {
			if strings.TrimSpace(part) == c {
				out = append(out, c)
				break
			}
		}
	}
	return out
}
