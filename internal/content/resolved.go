package content

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Resolved is the editable program. Shared phases can be overridden for an
// individual day; an explicit empty day phase disables the shared sequence.
type Resolved struct {
	Version          int     `json:"version"`
	Name             string  `json:"name"`
	TotalDays        int     `json:"totalDays"`
	WarmupRounds     int     `json:"warmupRounds"`
	Warmup           []RItem `json:"warmup"`
	Days             []RDay  `json:"days"`
	Cooldown         []RItem `json:"cooldown"`
	FixedDifficulty  bool    `json:"fixedDifficulty,omitempty"`
	SelfManagedSides bool    `json:"selfManagedSides,omitempty"`
}

// RDay is one day's main sequence, repeated Rounds times (1–10).
// Nil phase overrides inherit the shared phase; a pointer to [] skips it.
type RDay struct {
	Day      int      `json:"day"`
	Rounds   int      `json:"rounds"`
	Items    []RItem  `json:"items"`
	Warmup   *[]RItem `json:"warmup,omitempty"`
	Cooldown *[]RItem `json:"cooldown,omitempty"`
}

// RItem is one exercise occurrence. Explicit rests belong after this occurrence,
// including phase transitions; omitted rests retain the user's phase settings.
type RItem struct {
	ID         string `json:"id"`
	Unit       Unit   `json:"unit"`
	Value      int    `json:"value"`
	PerSide    bool   `json:"perSide"`
	Slot       Slot   `json:"slot"`
	Round      int    `json:"round"`
	Rest       *int   `json:"rest,omitempty"`
	SideRest   int    `json:"sideRest,omitempty"`
	RepSeconds int    `json:"repSeconds,omitempty"`
}

// WarmupFor and CooldownFor return the effective phases for a zero-based day.
func (r Resolved) WarmupFor(index int) []RItem {
	if index >= 0 && index < len(r.Days) && r.Days[index].Warmup != nil {
		return *r.Days[index].Warmup
	}
	return r.Warmup
}

func (r Resolved) CooldownFor(index int) []RItem {
	if index >= 0 && index < len(r.Days) && r.Days[index].Cooldown != nil {
		return *r.Days[index].Cooldown
	}
	return r.Cooldown
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
		// Legacy full-day lists: retain differing phases as per-day overrides.
		legacy := false
		for _, d := range r.Days {
			for _, it := range d.Items {
				legacy = legacy || it.Slot == Warmup || it.Slot == Cooldown
			}
		}
		if legacy {
			for di := range r.Days {
				d := &r.Days[di]
				wu, cd, main := []RItem{}, []RItem{}, []RItem{}
				for _, it := range d.Items {
					switch it.Slot {
					case Warmup:
						if it.Round <= 1 {
							wu = append(wu, it)
						}
					case Cooldown:
						cd = append(cd, it)
					default:
						main = append(main, it)
					}
				}
				if di == 0 {
					r.Warmup, r.Cooldown = wu, cd
				}
				if d.Warmup == nil && !reflect.DeepEqual(wu, r.Warmup) {
					d.Warmup = &wu
				}
				if d.Cooldown == nil && !reflect.DeepEqual(cd, r.Cooldown) {
					d.Cooldown = &cd
				}
				d.Items = main
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
			left, right := byRound[r][i], base[i]
			left.Round, right.Round = 1, 1
			if !reflect.DeepEqual(left, right) {
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
		if scale && !r.FixedDifficulty {
			val = ScaleValue(unit, val, level)
		}
		sideRest := clampRest(ri.SideRest, 0, 90)
		if r.SelfManagedSides {
			sideRest = 0
		}
		items = append(items, Item{
			ExerciseID: ri.ID, Unit: unit, Value: val,
			Slot: ri.Slot, PerSide: perSide, Round: round,
			restOverride: ri.Rest, SideRest: sideRest, RepSeconds: ri.RepSeconds,
		})
	}

	for rnd := 1; rnd <= wuRounds; rnd++ {
		for _, ri := range r.WarmupFor(day - 1) {
			emit(ri, rnd, false)
		}
	}
	for rnd := 1; rnd <= mainRounds; rnd++ {
		for _, ri := range r.Days[day-1].Items {
			emit(ri, rnd, true)
		}
	}
	for _, ri := range r.CooldownFor(day - 1) {
		emit(ri, 1, false)
	}

	est := assignRestsAndEst(items, rests)
	return Workout{Day: day, Items: items, EstSec: est, FixedDifficulty: r.FixedDifficulty, SelfManagedSides: r.SelfManagedSides}
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
