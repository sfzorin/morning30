package content

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Resolved is an editable program. Each day has one warm-up, two complete
// main rounds and one cool-down. Nonempty day phases may replace shared phases.
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

// RDay stores one main sequence, repeated exactly twice. Empty phase overrides
// inherit shared phases. Rounds remains in JSON for older editors and exports.
type RDay struct {
	Day      int      `json:"day"`
	Rounds   int      `json:"rounds"`
	Items    []RItem  `json:"items"`
	Warmup   *[]RItem `json:"warmup,omitempty"`
	Cooldown *[]RItem `json:"cooldown,omitempty"`
}

// RItem prescribes an exercise and its dose. Pauses belong to user settings,
// never to an exercise occurrence. Legacy JSON rest/sideRest keys are ignored.
type RItem struct {
	ID         string `json:"id"`
	Unit       Unit   `json:"unit"`
	Value      int    `json:"value"`
	PerSide    bool   `json:"perSide"`
	Slot       Slot   `json:"slot"`
	Round      int    `json:"round"`
	RepSeconds int    `json:"repSeconds,omitempty"`
}

// WarmupFor and CooldownFor return nonempty phases for a zero-based day.
func (r Resolved) WarmupFor(index int) []RItem {
	if index >= 0 && index < len(r.Days) && r.Days[index].Warmup != nil && len(*r.Days[index].Warmup) > 0 {
		return *r.Days[index].Warmup
	}
	if len(r.Warmup) > 0 {
		return r.Warmup
	}
	return seqItems(demoWarmupSeq, Warmup)
}

func (r Resolved) CooldownFor(index int) []RItem {
	if index >= 0 && index < len(r.Days) && r.Days[index].Cooldown != nil && len(*r.Days[index].Cooldown) > 0 {
		return *r.Days[index].Cooldown
	}
	if len(r.Cooldown) > 0 {
		return r.Cooldown
	}
	return seqItems(demoCooldownSeq, Cooldown)
}

// ResolveBuiltin flattens the legacy Sergey program (personal-file snapshot).
func ResolveBuiltin() Resolved { return sergeySpec.resolve() }

// ParseResolved accepts old files while returning the same canonical contract
// used by new programs. Only legacy full-day files use item slots as containers.
func ParseResolved(data string) (Resolved, error) {
	var r Resolved
	if err := json.Unmarshal([]byte(data), &r); err != nil {
		return r, err
	}
	if r.Version < 2 && len(r.Warmup) == 0 && len(r.Cooldown) == 0 {
		r.migrateLegacyPhases()
	}
	if err := r.validateLegacyRounds(); err != nil {
		return r, err
	}
	r.normalize()
	return r, nil
}

// A single repeated sequence cannot represent different legacy rounds without
// dropping or multiplying work. Reject that import explicitly; identical old
// rounds (including ones differing only in discarded pause fields) are safe.
func (r Resolved) validateLegacyRounds() error {
	check := func(items []RItem, phase string) error {
		byRound := map[int][]RItem{}
		for _, it := range items {
			round := it.Round
			if round < 1 {
				round = 1
			}
			it.Round, it.Slot = 1, Main
			byRound[round] = append(byRound[round], it)
		}
		var first []RItem
		for _, items := range byRound {
			if first == nil {
				first = items
			} else if !reflect.DeepEqual(first, items) {
				return fmt.Errorf("%s: legacy rounds contain different exercises or doses; provide one exercise sequence per phase", phase)
			}
		}
		return nil
	}
	if err := check(r.Warmup, "warm-up"); err != nil {
		return err
	}
	if err := check(r.Cooldown, "cool-down"); err != nil {
		return err
	}
	for i, d := range r.Days {
		if err := check(d.Items, fmt.Sprintf("day %d main", i+1)); err != nil {
			return err
		}
		if d.Warmup != nil {
			if err := check(*d.Warmup, fmt.Sprintf("day %d warm-up", i+1)); err != nil {
				return err
			}
		}
		if d.Cooldown != nil {
			if err := check(*d.Cooldown, fmt.Sprintf("day %d cool-down", i+1)); err != nil {
				return err
			}
		}
	}
	return nil
}

// migrateLegacyPhases separates old full-day lists. A missing daily phase now
// inherits the shared sequence; nonempty variations retain their own content.
func (r *Resolved) migrateLegacyPhases() {
	for di := range r.Days {
		d := &r.Days[di]
		var wu, cd, main []RItem
		for _, it := range d.Items {
			switch it.Slot {
			case Warmup:
				wu = append(wu, it)
			case Cooldown:
				cd = append(cd, it)
			default:
				main = append(main, it)
			}
		}
		if di == 0 {
			r.Warmup, r.Cooldown = wu, cd
		}
		if d.Warmup == nil && len(wu) > 0 && !reflect.DeepEqual(wu, r.Warmup) {
			d.Warmup = &wu
		}
		if d.Cooldown == nil && len(cd) > 0 && !reflect.DeepEqual(cd, r.Cooldown) {
			d.Cooldown = &cd
		}
		d.Items = main
	}
}

// normalize builds owned slices so normalizing a value copy for export or
// playback never mutates the caller's program. Containers determine each slot.
func (r *Resolved) normalize() {
	r.TotalDays = TotalDays
	r.WarmupRounds = WarmupRounds
	r.Warmup = phaseItems(r.WarmupFor(-1), Warmup)
	r.Cooldown = phaseItems(r.CooldownFor(-1), Cooldown)
	r.Days = append([]RDay(nil), r.Days...)
	if len(r.Days) > TotalDays {
		r.Days = r.Days[:TotalDays]
	}
	for len(r.Days) < TotalDays {
		r.Days = append(r.Days, RDay{})
	}
	for di := range r.Days {
		d := &r.Days[di]
		d.Day = di + 1
		d.Rounds = MainRounds
		d.Items = phaseItems(d.Items, Main)
		if len(d.Items) == 0 {
			d.Items = mainItemsOnce(demoDayBlocks[d.Day])
		}
		d.Warmup = normalizeOverride(d.Warmup, Warmup)
		d.Cooldown = normalizeOverride(d.Cooldown, Cooldown)
	}
}

func normalizeOverride(items *[]RItem, slot Slot) *[]RItem {
	if items == nil || len(*items) == 0 {
		return nil
	}
	out := phaseItems(*items, slot)
	return &out
}

// phaseItems converts an old expanded list into one canonical round. Explicit
// round labels identify old copies; the first labelled round owns the sequence
// and doses, including when later rounds differ. Never repeat an already
// expanded day as though it were one round. Unlabelled repeated entries are
// intentional exercise occurrences and retain their order and dose.
func phaseItems(items []RItem, slot Slot) []RItem {
	firstRound := 0
	for _, it := range items {
		round := it.Round
		if round < 1 {
			round = 1
		}
		if firstRound == 0 || round < firstRound {
			firstRound = round
		}
	}
	out := make([]RItem, 0, len(items))
	for _, it := range items {
		round := it.Round
		if round < 1 {
			round = 1
		}
		if round != firstRound {
			continue
		}
		it.Slot, it.Round = slot, 1
		out = append(out, it)
	}
	return out
}

// Marshal exports the canonical program, including when called on a direct
// struct rather than a parsed file. Unsupported legacy pause keys are omitted.
func (r Resolved) Marshal() string {
	r.normalize()
	b, _ := json.Marshal(r)
	return string(b)
}

// Workout builds one warm-up, two complete main rounds, and one cool-down.
// Direct structs and imported programs obey the same contract. Dose scaling
// affects the main phase only; user rest preferences affect every transition.
func (r Resolved) Workout(day, level int, rests Rests) Workout {
	r.normalize()
	day = clampDay(day)
	var items []Item
	emit := func(ri RItem, slot Slot, round int) {
		unit, perSide := ri.Unit, ri.PerSide
		if e, ok := Pool[ri.ID]; ok {
			unit, perSide = e.Unit, e.PerSide
		}
		val := ri.Value
		if slot == Main && !r.FixedDifficulty {
			val = ScaleValue(unit, val, level)
		}
		items = append(items, Item{
			ExerciseID: ri.ID, Unit: unit, Value: val,
			Slot: slot, PerSide: perSide, Round: round, RepSeconds: ri.RepSeconds,
		})
	}
	for _, ri := range r.WarmupFor(day - 1) {
		emit(ri, Warmup, 1)
	}
	for round := 1; round <= MainRounds; round++ {
		for _, ri := range r.Days[day-1].Items {
			emit(ri, Main, round)
		}
	}
	for _, ri := range r.CooldownFor(day - 1) {
		emit(ri, Cooldown, 1)
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
