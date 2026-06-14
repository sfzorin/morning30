package content

// TotalDays is the cycle length.
const TotalDays = 30

// Item is one step the player walks the user through.
type Item struct {
	ExerciseID string `json:"id"`
	Unit       Unit   `json:"unit"`
	Value      int    `json:"value"`
	Slot       Slot   `json:"slot"`
	PerSide    bool   `json:"perSide"`
	Round      int    `json:"round"`
	Rest       int    `json:"rest"` // seconds of rest AFTER this item (0 on the last)
}

// Workout is a full session for one day.
type Workout struct {
	Day    int    `json:"day"`
	Items  []Item `json:"items"`
	EstSec int    `json:"estSec"`
}

// ScaleValue applies the universal difficulty level to a value: ±10% per step
// (linear), level in [−3, +3]. Only reps/seconds scale; breaths stay fixed.
func ScaleValue(u Unit, value, level int) int {
	if level == 0 || u == Breaths {
		return value
	}
	out := (value*(10+level) + 5) / 10 // round half up (10+level is 7..13)
	if u == Reps && out < 1 {
		out = 1
	}
	if u == Seconds && out < 5 {
		out = 5
	}
	return out
}

// ConvertValue adapts a magnitude when swapping between units (used by manual
// replacement so e.g. 20 s does not become "20 reps"). ~2 s per rep.
func ConvertValue(from, to Unit, value int) int {
	if from == to {
		return value
	}
	switch {
	case from == Seconds && to == Reps:
		v := value / 2
		if v < 6 {
			v = 6
		}
		return v
	case from == Reps && to == Seconds:
		v := value * 2
		if v < 15 {
			v = 15
		}
		return v
	case to == Breaths:
		return 5
	}
	return value
}

// setDef is a fixed-value entry (warm-up / cool-down).
type setDef struct {
	ID  string
	Val int
}

// ---- Fixed daily sequences ----

// warmupSeq is the 6-exercise warm-up (one round), the same every day.
var warmupSeq = []setDef{
	{"W01", 20}, {"W02", 40}, {"W03", 8}, {"W04", 10}, {"W05", 8}, {"W06", 20},
}

// cooldownSeq is the 4-exercise cool-down (one round), the same every day.
var cooldownSeq = []setDef{
	{"CD01", 30}, {"CD02", 30}, {"CD03", 30}, {"CD04", 30},
}

// dayBlocks is the explicit main block for each day (1..30): an ordered list of
// {exercise, per-round value}. Each block runs for 2 rounds. Values are the base
// (level-0) magnitudes; the universal difficulty level scales them at runtime.
var dayBlocks = map[int][]setDef{
	1:  {{"P01", 12}, {"B01", 30}, {"C03", 30}, {"L01", 10}, {"C01", 60}},
	2:  {{"P01", 12}, {"B02", 12}, {"C04", 30}, {"L05", 8}, {"A01", 8}, {"C07", 25}},
	3:  {{"P03", 8}, {"B06", 12}, {"C03", 32}, {"L01", 12}, {"P02", 6}, {"C05", 10}, {"B03", 12}},
	4:  {{"C11", 10}, {"B01", 25}, {"L03", 12}, {"B02", 10}, {"C01", 50}}, // recovery
	5:  {{"P01", 13}, {"C04", 32}, {"A01", 9}, {"C08", 8}, {"L04", 8}, {"B04", 10}, {"C02", 20}},
	6:  {{"P04", 6}, {"B10", 10}, {"C03", 35}, {"L02", 8}, {"C10", 6}, {"C05", 11}, {"B05", 10}, {"C01", 65}},
	7:  {{"C11", 12}, {"B01", 30}, {"L06", 12}, {"B02", 10}, {"C01", 55}}, // recovery
	8:  {{"P01", 14}, {"B06", 14}, {"C04", 35}, {"L04", 9}, {"P02", 7}, {"C07", 30}, {"B03", 12}},
	9:  {{"P03", 9}, {"B01", 35}, {"C03", 36}, {"L01", 12}, {"A01", 10}, {"C05", 12}, {"B02", 14}, {"C02", 22}},
	10: {{"B01", 40}, {"L02", 9}, {"B03", 12}, {"L05", 10}, {"B06", 14}, {"C01", 70}, {"L07", 20}},
	11: {{"C12", 8}, {"B02", 10}, {"L03", 12}, {"L06", 12}, {"C01", 55}}, // recovery
	12: {{"P04", 7}, {"B06", 14}, {"C04", 36}, {"L04", 10}, {"P02", 8}, {"C08", 9}, {"B04", 12}, {"L02", 10}, {"C02", 23}},
	13: {{"P06", 5}, {"B02", 15}, {"C03", 38}, {"C05", 12}, {"A04", 6}, {"L05", 11}, {"C06", 22}, {"C01", 75}},
	14: {{"C11", 12}, {"B01", 30}, {"L06", 12}, {"B02", 12}, {"C01", 60}}, // recovery
	15: {{"P03", 10}, {"B10", 12}, {"C04", 38}, {"L01", 14}, {"A01", 11}, {"C05", 13}, {"B03", 13}, {"C02", 25}},
	16: {{"P04", 8}, {"B06", 16}, {"C03", 40}, {"L04", 11}, {"P02", 9}, {"C08", 10}, {"B05", 12}, {"C06", 25}, {"L07", 22}},
	17: {{"B01", 45}, {"L02", 10}, {"B02", 16}, {"L05", 12}, {"B04", 12}, {"C01", 80}, {"A03", 20}},
	18: {{"C12", 8}, {"B01", 30}, {"L03", 14}, {"L05", 10}, {"C01", 60}}, // recovery
	19: {{"P05", 5}, {"B06", 16}, {"C04", 40}, {"L02", 11}, {"P06", 6}, {"C05", 14}, {"B03", 14}, {"L04", 12}, {"C02", 27}},
	20: {{"P03", 11}, {"B01", 45}, {"C03", 42}, {"L01", 15}, {"P02", 9}, {"B04", 13}, {"C05", 14}, {"C08", 11}, {"A02", 22}, {"C01", 80}},
	21: {{"C11", 14}, {"B02", 12}, {"L03", 14}, {"L05", 10}, {"C01", 60}}, // recovery
	22: {{"P04", 8}, {"B10", 13}, {"C04", 42}, {"L04", 12}, {"P02", 10}, {"B05", 13}, {"C06", 28}, {"L02", 12}, {"C02", 28}},
	23: {{"P01", 14}, {"B03", 14}, {"C03", 35}, {"L05", 12}, {"A01", 12}, {"C07", 40}, {"C01", 80}}, // technique
	24: {{"P05", 5}, {"B06", 18}, {"C04", 44}, {"L02", 12}, {"P06", 7}, {"C05", 15}, {"B02", 18}, {"L04", 12}, {"C08", 12}, {"C02", 30}},
	25: {{"C12", 10}, {"B01", 35}, {"L03", 14}, {"B02", 12}, {"C01", 60}}, // recovery
	26: {{"P03", 12}, {"B10", 14}, {"C03", 45}, {"L01", 16}, {"P02", 10}, {"B04", 14}, {"C06", 30}, {"L05", 14}, {"A04", 8}, {"C01", 85}},
	27: {{"P01", 16}, {"B06", 16}, {"C04", 40}, {"L02", 12}, {"P08", 8}, {"C07", 45}, {"C05", 15}, {"B03", 15}},
	28: {{"C11", 14}, {"B01", 35}, {"L06", 14}, {"B02", 12}, {"C01", 65}}, // recovery
	29: {{"P04", 10}, {"B01", 50}, {"C03", 45}, {"L02", 14}, {"P06", 8}, {"C05", 16}, {"B02", 18}, {"L04", 14}, {"A02", 25}, {"C02", 30}}, // peak
	30: {{"P01", 16}, {"B06", 18}, {"C04", 45}, {"L01", 15}, {"P02", 10}, {"B03", 15}, {"C05", 15}, {"C08", 12}, {"L04", 12}, {"C01", 90}}, // control
}

// recoveryDays are the lighter days in the cycle.
var recoveryDays = map[int]bool{4: true, 7: true, 11: true, 14: true, 18: true, 21: true, 25: true, 28: true}

func clampDay(day int) int {
	if day < 1 {
		return 1
	}
	if day > TotalDays {
		return TotalDays
	}
	return day
}

// isRecoveryDay reports whether the given day is a lighter recovery day.
func isRecoveryDay(day int) bool { return recoveryDays[clampDay(day)] }

// restAfter is the rest after a main exercise (the user's 10..40 s setting; the
// spec suggests 10–20). Warm-up/cool-down items have a short/zero rest.
func restAfter(e Exercise, lightRest int) int {
	switch e.Slot {
	case Warmup:
		return 5
	case Cooldown:
		return 0
	}
	if lightRest < 10 {
		lightRest = 10
	}
	if lightRest > 40 {
		lightRest = 40
	}
	return lightRest
}

const restBetweenRounds = 60

// assignRestsAndEst fills each item's Rest (per the spec rules; longer between
// main rounds) and returns a rough total-duration estimate in seconds.
func assignRestsAndEst(items []Item, lightRest int) int {
	est := 0
	for i := range items {
		it := items[i]
		e := Pool[it.ExerciseID]
		if i < len(items)-1 {
			next := items[i+1]
			if it.Slot == Main && next.Slot == Main && next.Round != it.Round {
				items[i].Rest = restBetweenRounds
			} else {
				items[i].Rest = restAfter(e, lightRest)
			}
		}
		work := it.Value
		if it.Unit == Reps {
			work *= 3
		} else if it.Unit == Breaths {
			work *= 4
		}
		if it.PerSide {
			work *= 2
		}
		est += work + items[i].Rest
	}
	return est
}

// warmupRounds is how many times the shared warm-up series runs each session.
const warmupRounds = 1

// seqItems converts a fixed sequence (warm-up / cool-down) to base RItems.
func seqItems(seq []setDef, slot Slot) []RItem {
	out := make([]RItem, 0, len(seq))
	for _, s := range seq {
		e := Pool[s.ID]
		out = append(out, RItem{ID: s.ID, Unit: e.Unit, Value: s.Val, PerSide: e.PerSide, Slot: slot, Round: 1})
	}
	return out
}

// mainItems builds the base MAIN items for a day: the day's explicit block run
// for 2 rounds. The universal difficulty level scales these at runtime.
func mainItems(day int) []RItem {
	defs := dayBlocks[clampDay(day)]
	out := make([]RItem, 0, len(defs)*2)
	for r := 1; r <= 2; r++ {
		for _, s := range defs {
			e := Pool[s.ID]
			out = append(out, RItem{ID: s.ID, Unit: e.Unit, Value: s.Val, PerSide: e.PerSide, Slot: Main, Round: r})
		}
	}
	return out
}
