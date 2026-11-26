package content

import (
	"fmt"

	"morning30/internal/langs"
	"strings"
	"testing"
)

var series30Langs = langs.Codes()

// The series must not leak into what accounts run today.
func TestSeries30IsSeparateFromExistingPrograms(t *testing.T) {
	if len(StandardPrograms) != 1 || StandardPrograms[0].Key != "demo" {
		t.Fatalf("standard programs changed: %+v", StandardPrograms)
	}
	seen := map[string]bool{}
	for _, c := range Series30() {
		if seen[c.Key] {
			t.Errorf("duplicate course key %s", c.Key)
		}
		seen[c.Key] = true
		if _, ok := StandardByKey(c.Key); ok {
			t.Errorf("%s is selectable as a standard program", c.Key)
		}
		if _, ok := LegacyNamed(c.Key); ok {
			t.Errorf("%s collides with a legacy program", c.Key)
		}
		if got, ok := CourseByKey(c.Key); !ok || got.Key != c.Key {
			t.Errorf("CourseByKey(%s) failed", c.Key)
		}
	}
	if len(seen) != 49 {
		t.Errorf("want 49 courses, got %d", len(seen))
	}
	if _, ok := CourseByKey("demo"); ok {
		t.Error("demo must not be a series course")
	}
}

// Every title carries the "30 days" promise in every supported language.
func TestSeries30Titles(t *testing.T) {
	for _, c := range Series30() {
		for _, l := range series30Langs {
			if title := c.Title[l]; title == "" || (strings.HasSuffix(c.Slug, "-in-30-days") && !strings.Contains(title, "30")) {
				t.Errorf("%s: bad %s title %q", c.Key, l, title)
			}
		}
		for _, l := range series30Langs {
			min := 25
			if l == "ja" || l == "ko" || l == "zh" || l == "zh-hant" {
				min = 10 // CJK says the same in far fewer characters
			}
			if len([]rune(c.About[l])) < min {
				t.Errorf("%s: missing %s description", c.Key, l)
			}
		}
		if got, ok := CourseBySlug(c.Slug); !ok || got.Key != c.Key || strings.ToLower(c.Slug) != c.Slug || strings.Contains(c.Slug, " ") {
			t.Errorf("%s: bad slug %q", c.Key, c.Slug)
		}
		if c.TitleIn("xx") != c.Title["en"] {
			t.Errorf("%s: unknown language must fall back to English", c.Key)
		}
		if c.Resolve().Name != c.Title["en"] {
			t.Errorf("%s: program name must be the English title", c.Key)
		}
	}
}

func TestSeries30Structure(t *testing.T) {
	for _, c := range Series30() {
		r := c.Resolve()
		if r.TotalDays != TotalDays || len(r.Days) != TotalDays {
			t.Fatalf("%s: %d days", c.Key, len(r.Days))
		}
		if len(r.Warmup) < 4 || len(r.Warmup) > 7 || len(r.Cooldown) != 4 {
			t.Errorf("%s: warmup=%d cooldown=%d", c.Key, len(r.Warmup), len(r.Cooldown))
		}
		check := func(where string, it RItem, slot Slot) {
			e, ok := Get(it.ID)
			if !ok {
				t.Errorf("%s %s: unknown exercise %s", c.Key, where, it.ID)
				return
			}
			// Warm-up and cool-down hold only their own kind of exercise; the
			// main part may use any (stretching programs stretch in the main part).
			if it.Slot != slot || (slot != Main && e.Slot != slot) {
				t.Errorf("%s %s: %s belongs to %s, used as %s", c.Key, where, it.ID, e.Slot, slot)
			}
			if it.Unit != e.Unit || it.PerSide != e.PerSide {
				t.Errorf("%s %s: %s unit/side mismatch", c.Key, where, it.ID)
			}
			if it.Value < 1 || (it.Unit == Seconds && (it.Value < 10 || it.Value%5 != 0)) {
				t.Errorf("%s %s: %s has dose %d", c.Key, where, it.ID, it.Value)
			}
		}
		for _, it := range r.Warmup {
			check("warm-up", it, Warmup)
		}
		for _, it := range r.Cooldown {
			check("cool-down", it, Cooldown)
		}
		for _, d := range r.Days {
			where := fmt.Sprintf("day %d", d.Day)
			if d.Rounds != MainRounds {
				t.Errorf("%s %s: %d rounds", c.Key, where, d.Rounds)
			}
			lo, hi := 4, 7
			if isRecoveryDay(d.Day) {
				lo, hi = 3, 4
			}
			if len(d.Items) < lo || len(d.Items) > hi {
				t.Errorf("%s %s: %d stations", c.Key, where, len(d.Items))
			}
			ids := map[string]bool{}
			for _, it := range d.Items {
				check(where, it, Main)
				if ids[it.ID] {
					t.Errorf("%s %s: %s repeats within the block", c.Key, where, it.ID)
				}
				ids[it.ID] = true
			}
		}
	}
}

// Within one session type a station never gets easier, stations only join,
// and recovery days are lighter than the training days around them.
func TestSeries30Progression(t *testing.T) {
	rests := Rests{Warmup: 5, Main: 15, Cooldown: 0}
	for _, c := range Series30() {
		for name, session := range map[string][]dose{"a": c.plan.a, "b": c.plan.b, "c": c.plan.c, "recovery": c.plan.recovery} {
			for _, s := range session {
				if s.To < s.From || s.Since >= peakDay {
					t.Errorf("%s %s: %s has a bad range", c.Key, name, s.ID)
				}
			}
			prev := map[string]int{}
			count := 0
			for day := 1; day < TotalDays; day++ {
				block := dosesOn(session, day)
				if len(block) < count {
					t.Errorf("%s %s: stations drop on day %d", c.Key, name, day)
				}
				count = len(block)
				for _, s := range block {
					if s.Val < prev[s.ID] {
						t.Errorf("%s %s: %s drops to %d on day %d", c.Key, name, s.ID, s.Val, day)
					}
					prev[s.ID] = s.Val
				}
			}
			for _, s := range session {
				if prev[s.ID] != s.To {
					t.Errorf("%s %s: %s peaks at %d, want %d", c.Key, name, s.ID, prev[s.ID], s.To)
				}
			}
		}
		r := c.Resolve()
		est := func(day int) int { return r.Workout(day, 0, rests).EstSec }
		for day := 1; day <= TotalDays; day++ {
			if e := est(day); e < 5*60 || e > 25*60 {
				t.Errorf("%s day %d: estimated %d s", c.Key, day, e)
			}
			if isRecoveryDay(day) && (est(day) >= est(day-1) || est(day) >= est(day+1)) {
				t.Errorf("%s day %d: recovery is not lighter than its neighbours", c.Key, day)
			}
		}
		if est(peakDay) <= est(1) {
			t.Errorf("%s: day 29 (%d s) is not longer than day 1 (%d s)", c.Key, est(peakDay), est(1))
		}
	}
}

// A course is an ordinary program file: it survives export → import unchanged,
// so it can be installed through the existing JSON import or MCP paths.
func TestSeries30RoundTripsAsProgramJSON(t *testing.T) {
	for _, c := range Series30() {
		r := c.Resolve()
		back, err := ParseResolved(r.Marshal())
		if err != nil {
			t.Fatalf("%s: %v", c.Key, err)
		}
		if back.Marshal() != r.Marshal() {
			t.Errorf("%s: program changes on import", c.Key)
		}
		for day := 1; day <= TotalDays; day++ {
			w := back.Workout(day, 0, Rests{Warmup: 5, Main: 15})
			if want := len(r.WarmupFor(day-1)) + 2*len(r.Days[day-1].Items) + len(r.CooldownFor(day-1)); len(w.Items) != want {
				t.Errorf("%s day %d: %d screens, want %d", c.Key, day, len(w.Items), want)
			}
		}
	}
}

// Run with -v to review the whole series day by day.
func TestSeries30Table(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	rests := Rests{Warmup: 5, Main: 15, Cooldown: 0}
	for _, c := range Series30() {
		r := c.Resolve()
		t.Logf("== %s · %s", c.Key, c.Title["ru"])
		for _, d := range r.Days {
			parts := make([]string, len(d.Items))
			for i, it := range d.Items {
				parts[i] = fmt.Sprintf("%s %d", it.ID, it.Value)
			}
			mark := " "
			if isRecoveryDay(d.Day) {
				mark = "R"
			}
			t.Logf("%2d %s %4.1f min  %s", d.Day, mark, float64(r.Workout(d.Day, 0, rests).EstSec)/60, strings.Join(parts, ", "))
		}
	}
}
