package content

import (
	"encoding/json"
	"strings"
)

// Resolved is the editable, flat program: a shared warm-up and cool-down series
// (the same for every day) plus a per-day MAIN block. This is what the editor
// edits and what Workout builds from. Warm-up/cool-down are stored once so
// editing them applies to the whole series; only the main block is per-day.
// The cycle length is always 30 days. Warm-up and main run 1–10 rounds.
type Resolved struct {
	Version      int     `json:"version"`
	Name         string  `json:"name"`
	TotalDays    int     `json:"totalDays"`
	WarmupRounds int     `json:"warmupRounds"`
	Warmup       []RItem `json:"warmup"`   // shared, one round (repeated WarmupRounds times)
	Days         []RDay  `json:"days"`     // per-day MAIN items (one round, repeated Rounds times)
	Cooldown     []RItem `json:"cooldown"` // shared, one round
}

// RDay is one day's MAIN block: one sequence, repeated Rounds times (1–10).
type RDay struct {
	Day    int     `json:"day"`
	Rounds int     `json:"rounds"`
	Items  []RItem `json:"items"`
}

// RItem is one exercise occurrence with its value.
type RItem struct {
	ID      string `json:"id"`
	Unit    Unit   `json:"unit"`
	Value   int    `json:"value"`
	PerSide bool   `json:"perSide"`
	Slot    Slot   `json:"slot"`
	Round   int    `json:"round"`
}

// ResolveBuiltin flattens the legacy Sergey program (personal-file snapshot).
func ResolveBuiltin() Resolved { return sergeySpec.resolve() }

// ParseResolved parses a stored custom program JSON and normalizes it (defaults +
// migration of legacy full-day programs into shared warm-up/cool-down).
func ParseResolved(data string) (Resolved, error) {
	var r Resolved
	if err := json.Unmarshal([]byte(data), &r); err != nil {
		return r, err
	}
	r.normalize()
	return r, nil
}

// normalize fills defaults and migrates legacy programs (where warm-up/cool-down
// lived inside each day's Items) into the shared warm-up/cool-down model.
func (r *Resolved) normalize() {
	r.TotalDays = TotalDays
	r.WarmupRounds = ClampRounds(r.WarmupRounds)
	if len(r.Warmup) == 0 && len(r.Cooldown) == 0 && len(r.Days) > 0 {
		// Legacy: lift the shared warm-up (round 1) and cool-down out of day 1,
		// then strip warm-up/cool-down from every day, leaving main only.
		for _, it := range r.Days[0].Items {
			switch {
			case it.Slot == Warmup && it.Round <= 1:
				r.Warmup = append(r.Warmup, it)
			case it.Slot == Cooldown:
				r.Cooldown = append(r.Cooldown, it)
			}
		}
		if len(r.Warmup) > 0 || len(r.Cooldown) > 0 {
			for di := range r.Days {
				main := make([]RItem, 0, len(r.Days[di].Items))
				for _, it := range r.Days[di].Items {
					if it.Slot == Main {
						main = append(main, it)
					}
				}
				r.Days[di].Items = main
			}
		}
	}
	for di := range r.Days {
		r.Days[di].Day = di + 1
		r.Days[di].normalizeMain()
	}
	for len(r.Days) < TotalDays {
		r.Days = append(r.Days, RDay{Day: len(r.Days) + 1, Rounds: 2})
	}
	if len(r.Days) > TotalDays {
		r.Days = r.Days[:TotalDays]
	}
}

func (d *RDay) normalizeMain() {
	items, detected := collapseMainRounds(d.Items)
	if d.Rounds < 1 {
		d.Rounds = detected
	}
	d.Rounds = ClampRounds(d.Rounds)
	d.Items = items
}

// collapseMainRounds folds a legacy list that stored every round as duplicated
// items into one sequence. Asymmetric rounds are left expanded with rounds=1.
func collapseMainRounds(items []RItem) ([]RItem, int) {
	maxR := 1
	for _, it := range items {
		if it.Round > maxR {
			maxR = it.Round
		}
	}
	if maxR <= 1 {
		return items, 1
	}
	if maxR > 10 {
		maxR = 10
	}
	byRound := make([][]RItem, maxR)
	for _, it := range items {
		r := it.Round
		if r < 1 {
			r = 1
		}
		if r > maxR {
			continue
		}
		byRound[r-1] = append(byRound[r-1], it)
	}
	base := byRound[0]
	for r := 1; r < maxR; r++ {
		if len(byRound[r]) != len(base) {
			return items, 1
		}
		for i := range base {
			if byRound[r][i].ID != base[i].ID || byRound[r][i].Value != base[i].Value {
				return items, 1
			}
		}
	}
	out := make([]RItem, len(base))
	for i, it := range base {
		it.Round = 1
		out[i] = it
	}
	return out, maxR
}

// Marshal serializes the program to JSON.
func (r Resolved) Marshal() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// Workout builds a session for a day: the shared warm-up (× WarmupRounds), the
// day's main block repeated Rounds times and scaled by the difficulty level,
// then the shared cool-down. Warm-up/cool-down and breaths are never scaled.
func (r Resolved) Workout(day, level int, rests Rests) Workout {
	if len(r.Days) == 0 {
		return ResolveDemo().Workout(day, level, rests)
	}
	if day < 1 {
		day = 1
	}
	if day > len(r.Days) {
		day = len(r.Days)
	}
	wuRounds := ClampRounds(r.WarmupRounds)
	mainRounds := ClampRounds(r.Days[day-1].Rounds)

	var items []Item
	emit := func(ri RItem, round int, scale bool) {
		unit, perSide := ri.Unit, ri.PerSide
		if e, ok := Pool[ri.ID]; ok { // built-in def wins for unit/side
			unit, perSide = e.Unit, e.PerSide
		}
		val := ri.Value
		if scale {
			val = ScaleValue(unit, val, level)
		}
		items = append(items, Item{
			ExerciseID: ri.ID, Unit: unit, Value: val,
			Slot: ri.Slot, PerSide: perSide, Round: round,
		})
	}

	for rnd := 1; rnd <= wuRounds; rnd++ {
		for _, ri := range r.Warmup {
			emit(ri, rnd, false)
		}
	}
	for rnd := 1; rnd <= mainRounds; rnd++ {
		for _, ri := range r.Days[day-1].Items {
			emit(ri, rnd, true)
		}
	}
	for _, ri := range r.Cooldown {
		emit(ri, 1, false)
	}

	est := assignRestsAndEst(items, rests)
	return Workout{Day: day, Items: items, EstSec: est}
}

// FigureKey extracts a built-in illustration id from a stored SVG path or raw id.
func FigureKey(svg, fallback string) string {
	s := strings.TrimSpace(svg)
	if s == "" {
		return fallback
	}
	s = strings.TrimSuffix(s, ".png")
	if i := strings.LastIndex(s, "/"); i >= 0 {
		s = s[i+1:]
	}
	if s == "" || strings.Contains(s, " ") {
		return fallback
	}
	if _, ok := Get(s); !ok && fallback != "" {
		return fallback
	}
	return s
}
