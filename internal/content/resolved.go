package content

import "encoding/json"

// Resolved is the program flattened to a per-day list of items. This is the
// shape stored in the DB for a user's custom program (viewable/editable), and
// what Workout builds from.
type Resolved struct {
	Version   int    `json:"version"`
	Name      string `json:"name"`
	TotalDays int    `json:"totalDays"`
	Days      []RDay `json:"days"`
}

// RDay is one day's ordered list of exercises.
type RDay struct {
	Day   int     `json:"day"`
	Label string  `json:"label"`
	Items []RItem `json:"items"`
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

// ResolveBuiltin flattens the built-in program into a Resolved (one entry per
// day, using each day's calendar block/variant).
func ResolveBuiltin() Resolved {
	r := Resolved{Version: 1, Name: "dani.cc", TotalDays: TotalDays}
	for day := 1; day <= TotalDays; day++ {
		w := Build(day, 0)
		rd := RDay{Day: day, Label: w.BlockCode}
		for _, it := range w.Items {
			rd.Items = append(rd.Items, RItem{
				ID: it.ExerciseID, Unit: it.Unit, Value: it.Value,
				PerSide: it.PerSide, Slot: it.Slot, Round: it.Round,
			})
		}
		r.Days = append(r.Days, rd)
	}
	return r
}

// ParseResolved parses a stored custom program JSON.
func ParseResolved(data string) (Resolved, error) {
	var r Resolved
	err := json.Unmarshal([]byte(data), &r)
	return r, err
}

// Marshal serializes the program to JSON.
func (r Resolved) Marshal() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// Workout builds a session for a day from this resolved program, applying the
// rest rules and pain substitutions. (Variant-level adaptation doesn't apply to
// a custom flat program — the user set fixed values.)
func (r Resolved) Workout(day, lightRest int, a Adapt) Workout {
	if len(r.Days) == 0 {
		return Build(day, lightRest)
	}
	if day < 1 {
		day = 1
	}
	if day > len(r.Days) {
		day = len(r.Days)
	}
	rd := r.Days[day-1]
	items := make([]Item, 0, len(rd.Items))
	for _, ri := range rd.Items {
		id, val := ri.ID, ri.Value
		if ri.Slot == Main {
			id, val = substitute(id, val, a)
		}
		unit, perSide := ri.Unit, ri.PerSide
		if e, ok := Pool[id]; ok {
			unit, perSide = e.Unit, e.PerSide
		}
		items = append(items, Item{
			ExerciseID: id, Unit: unit, Value: val,
			Slot: ri.Slot, PerSide: perSide, Round: ri.Round,
		})
	}
	est := assignRestsAndEst(items, lightRest)
	return Workout{Day: day, BlockCode: rd.Label, Items: items, EstSec: est}
}
