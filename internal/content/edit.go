package content

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Program edits used by the library ("add to program"), the day screen
// ("replace", "remove") and agents. Each returns a changed copy; the input
// program is never modified.

// Days selects which days of the cycle an edit touches.
type Days struct {
	All  bool  `json:"all,omitempty"`
	List []int `json:"list,omitempty"` // 1-based day numbers
}

// AllDays and OneDay are the two common selections.
func AllDays() Days       { return Days{All: true} }
func OneDay(day int) Days { return Days{List: []int{day}} }

var (
	ErrNoDays       = errors.New("no valid days selected")
	ErrBadPart      = errors.New("part must be warmup, main or cooldown")
	ErrNotInProgram = errors.New("exercise is not in the selected days")
	ErrEmptyPhase   = errors.New("a part of the day cannot be left empty")
)

// indexes returns zero-based day indexes for a selection.
func (d Days) indexes(total int) ([]int, error) {
	if d.All {
		out := make([]int, total)
		for i := range out {
			out[i] = i
		}
		return out, nil
	}
	seen := map[int]bool{}
	var out []int
	for _, day := range d.List {
		if day < 1 || day > total {
			return nil, fmt.Errorf("%w: day %d", ErrNoDays, day)
		}
		if !seen[day] {
			seen[day] = true
			out = append(out, day-1)
		}
	}
	if len(out) == 0 {
		return nil, ErrNoDays
	}
	return out, nil
}

// DefaultValue is the dose a newly added exercise starts with.
func DefaultValue(u Unit) int {
	switch u {
	case Seconds:
		return 30
	case Breaths:
		return 5
	}
	return 10
}

func clone(r Resolved) Resolved {
	var out Resolved
	b, _ := json.Marshal(r)
	_ = json.Unmarshal(b, &out)
	return out
}

func validPart(p Slot) bool { return p == Warmup || p == Main || p == Cooldown }

// phase returns a pointer to the list an edit should change for one day and
// part. Warm-up and cool-down are shared; editing only some days gives those
// days their own copy first.
func (r *Resolved) phase(idx int, part Slot, all bool) *[]RItem {
	switch part {
	case Main:
		return &r.Days[idx].Items
	case Warmup:
		if all && r.Days[idx].Warmup == nil {
			return nil // the shared list is edited once by the caller
		}
		if r.Days[idx].Warmup == nil || len(*r.Days[idx].Warmup) == 0 {
			own := append([]RItem(nil), r.WarmupFor(idx)...)
			r.Days[idx].Warmup = &own
		}
		return r.Days[idx].Warmup
	default:
		if all && r.Days[idx].Cooldown == nil {
			return nil
		}
		if r.Days[idx].Cooldown == nil || len(*r.Days[idx].Cooldown) == 0 {
			own := append([]RItem(nil), r.CooldownFor(idx)...)
			r.Days[idx].Cooldown = &own
		}
		return r.Days[idx].Cooldown
	}
}

// lists returns every list touched by an edit of part on the selected days.
func (r *Resolved) lists(part Slot, days Days) ([]*[]RItem, error) {
	idx, err := days.indexes(len(r.Days))
	if err != nil {
		return nil, err
	}
	var out []*[]RItem
	if days.All && part != Main {
		if part == Warmup {
			if len(r.Warmup) == 0 {
				r.Warmup = append([]RItem(nil), r.WarmupFor(-1)...)
			}
			out = append(out, &r.Warmup)
		} else {
			if len(r.Cooldown) == 0 {
				r.Cooldown = append([]RItem(nil), r.CooldownFor(-1)...)
			}
			out = append(out, &r.Cooldown)
		}
	}
	for _, i := range idx {
		if l := r.phase(i, part, days.All); l != nil {
			out = append(out, l)
		}
	}
	return out, nil
}

// AddExercise appends an exercise to a part of the selected days.
// value <= 0 uses DefaultValue for the exercise's unit.
func AddExercise(r Resolved, ex Exercise, part Slot, days Days, value int) (Resolved, error) {
	if !validPart(part) {
		return r, ErrBadPart
	}
	if value <= 0 {
		value = DefaultValue(ex.Unit)
	}
	out := clone(r)
	lists, err := out.lists(part, days)
	if err != nil {
		return r, err
	}
	for _, l := range lists {
		*l = append(*l, RItem{ID: ex.ID, Unit: ex.Unit, Value: value, PerSide: ex.PerSide, Slot: part, Round: 1})
	}
	return out, nil
}

// ReplaceExercise swaps every occurrence of from for to on the selected days,
// in every part. The dose carries over, converted when the unit differs.
// It returns how many occurrences changed.
func ReplaceExercise(r Resolved, from string, to Exercise, days Days) (Resolved, int, error) {
	out := clone(r)
	n := 0
	for _, part := range []Slot{Warmup, Main, Cooldown} {
		if !out.contains(part, days, from) {
			continue
		}
		lists, err := out.lists(part, days)
		if err != nil {
			return r, 0, err
		}
		for _, l := range lists {
			for i, it := range *l {
				if it.ID != from {
					continue
				}
				value := it.Value
				if it.Unit != to.Unit {
					value = ConvertValue(it.Unit, to.Unit, it.Value)
				}
				(*l)[i] = RItem{ID: to.ID, Unit: to.Unit, Value: value, PerSide: to.PerSide, Slot: it.Slot, Round: it.Round}
				n++
			}
		}
	}
	if n == 0 {
		if _, err := days.indexes(len(r.Days)); err != nil {
			return r, 0, err
		}
		return r, 0, ErrNotInProgram
	}
	return out, n, nil
}

// RemoveExercise takes an exercise out of the selected days. part "" means
// every part. A part may not be left empty. It returns how many were removed.
func RemoveExercise(r Resolved, id string, part Slot, days Days) (Resolved, int, error) {
	if part != "" && !validPart(part) {
		return r, 0, ErrBadPart
	}
	parts := []Slot{Warmup, Main, Cooldown}
	if part != "" {
		parts = []Slot{part}
	}
	out := clone(r)
	n := 0
	for _, p := range parts {
		if !out.contains(p, days, id) {
			continue
		}
		lists, err := out.lists(p, days)
		if err != nil {
			return r, 0, err
		}
		for _, l := range lists {
			kept := (*l)[:0]
			for _, it := range *l {
				if it.ID == id {
					n++
					continue
				}
				kept = append(kept, it)
			}
			if len(kept) == 0 {
				return r, 0, ErrEmptyPhase
			}
			*l = kept
		}
	}
	if n == 0 {
		if _, err := days.indexes(len(r.Days)); err != nil {
			return r, 0, err
		}
		return r, 0, ErrNotInProgram
	}
	return out, n, nil
}

// contains reports whether a part of any selected day holds the exercise, so
// edits never create per-day copies of a phase they do not change.
func (r Resolved) contains(part Slot, days Days, id string) bool {
	idx, err := days.indexes(len(r.Days))
	if err != nil {
		return false
	}
	for _, i := range idx {
		var items []RItem
		switch part {
		case Main:
			items = r.Days[i].Items
		case Warmup:
			items = r.WarmupFor(i)
		default:
			items = r.CooldownFor(i)
		}
		for _, it := range items {
			if it.ID == id {
				return true
			}
		}
	}
	return false
}
