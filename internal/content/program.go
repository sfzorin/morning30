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

// setDef is a fixed-value entry (warm-up, cool-down, recovery, control).
type setDef struct {
	ID  string
	Val int
}

// blockSet carries the three variant magnitudes [−, base, +].
type blockSet struct {
	ID   string
	Vals [3]int
}

// ---- Fixed sequences ----

var warmupSeq = []setDef{
	{"W01", 5}, {"W02", 5}, {"W03", 5}, {"W04", 5}, {"W05", 10},
	{"W06", 10}, {"W07", 5}, {"W08", 5}, {"W09", 5}, {"W10", 5},
}

var cooldownSeq = []setDef{
	{"CD01", 5}, {"CD02", 25}, {"CD03", 25}, {"CD04", 25}, {"CD05", 25}, {"CD06", 25},
}

// Recovery (R) block — 1 round, no long plank / scissors / flutter / narrow
// push-ups / mini-squat / RKC, per the spec.
var recoverySeq = []setDef{
	{"C04", 6}, {"G01", 12}, {"B01", 10}, {"S01", 15}, {"C05", 20},
}

// Day-30 control block — 2 rounds, then a 90 s plank challenge.
var controlSeq = []setDef{
	{"M01", 16}, {"C01", 30}, {"C03", 30}, {"C04", 10},
	{"B01", 15}, {"B02", 15}, {"L01", 10}, {"A01", 12},
}

const controlPlankSec = 90

// ---- Variant blocks (A–E), 2 rounds each ----

var blocks = map[string][]blockSet{
	"A": {
		{"M01", [3]int{10, 11, 12}},
		{"C01", [3]int{20, 22, 25}},
		{"G01", [3]int{14, 14, 15}},
		{"B01", [3]int{12, 12, 13}},
		{"A01", [3]int{8, 9, 10}},
		{"S01", [3]int{20, 22, 25}},
		{"C02", [3]int{45, 50, 55}}, // plank — last
	},
	"B": {
		{"M01", [3]int{12, 13, 14}},
		{"C03", [3]int{25, 27, 30}},
		{"C04", [3]int{8, 9, 10}},
		{"B02", [3]int{12, 13, 14}},
		{"M02", [3]int{6, 7, 8}},
		{"A01", [3]int{10, 11, 12}},
		{"C05", [3]int{30, 32, 35}},
		{"C02", [3]int{55, 60, 65}}, // plank — last
	},
	"C": {
		{"M03", [3]int{10, 11, 12}},
		{"C01", [3]int{25, 28, 30}},
		{"C06", [3]int{10, 11, 12}},
		{"B03", [3]int{20, 22, 25}},
		{"B04", [3]int{12, 13, 14}},
		{"G02", [3]int{16, 17, 18}},
		{"L01", [3]int{8, 9, 10}},
		{"C05", [3]int{35, 38, 40}},
		{"C02", [3]int{60, 65, 70}}, // plank — last
	},
	"D": {
		{"M01", [3]int{14, 15, 16}},
		{"M02", [3]int{8, 9, 10}},
		{"C03", [3]int{30, 33, 35}},
		{"C07", [3]int{20, 23, 25}},
		{"B02", [3]int{14, 15, 16}},
		{"B05", [3]int{10, 11, 12}},
		{"A01", [3]int{12, 13, 14}},
		{"A02", [3]int{15, 18, 20}},
		{"C08", [3]int{25, 28, 30}}, // RKC plank — last
	},
	"E": {
		{"M01", [3]int{16, 16, 16}},
		{"M02", [3]int{10, 10, 12}},
		{"C01", [3]int{35, 38, 40}},
		{"C03", [3]int{35, 38, 40}},
		{"C04", [3]int{10, 11, 12}},
		{"B01", [3]int{15, 16, 18}},
		{"B02", [3]int{15, 16, 18}},
		{"L01", [3]int{10, 12, 12}},
		{"A01", [3]int{12, 14, 15}},
		{"C02", [3]int{75, 80, 90}}, // plank — last
	},
}

// ---- 30-day calendar ----

type dayPlan struct {
	Code     string // block letter
	Variant  int    // 0=−, 1=base, 2=+
	Recovery bool
	Control  bool
}

// calendar[day] (1-indexed) maps each day to its block/variant. The variant per
// day still bakes the natural day-to-day progression into the base program; the
// runtime no longer shifts it (difficulty is the universal level instead).
var calendar = map[int]dayPlan{
	1: {Code: "A", Variant: 0}, 2: {Code: "A", Variant: 1}, 3: {Code: "A", Variant: 2},
	4: {Recovery: true},
	5: {Code: "B", Variant: 0}, 6: {Code: "B", Variant: 1},
	7: {Recovery: true},
	8: {Code: "B", Variant: 2}, 9: {Code: "C", Variant: 0},
	10: {Recovery: true},
	11: {Code: "C", Variant: 1}, 12: {Code: "C", Variant: 2},
	13: {Recovery: true},
	14: {Code: "B", Variant: 2}, 15: {Code: "C", Variant: 1}, 16: {Code: "D", Variant: 0},
	17: {Recovery: true},
	18: {Code: "D", Variant: 1}, 19: {Code: "C", Variant: 2},
	20: {Recovery: true},
	21: {Code: "D", Variant: 2}, 22: {Code: "D", Variant: 1},
	23: {Recovery: true},
	24: {Code: "E", Variant: 0}, 25: {Code: "D", Variant: 2},
	26: {Recovery: true},
	27: {Code: "E", Variant: 1},
	28: {Recovery: true},
	29: {Code: "E", Variant: 2},
	30: {Control: true},
}

func clampDay(day int) int {
	if day < 1 {
		return 1
	}
	if day > TotalDays {
		return TotalDays
	}
	return day
}

// restAfter applies the spec's rest rules.
func restAfter(e Exercise, lightRest int) int {
	switch e.Slot {
	case Warmup:
		return 5
	case Cooldown:
		return 0
	}
	switch e.ID {
	case "M01", "M02", "M03": // after push-ups
		return 35
	case "C02", "C05", "C08": // after a plank
		return 50
	}
	if lightRest < 10 {
		lightRest = 10
	}
	if lightRest > 40 {
		lightRest = 40
	}
	return lightRest
}

const restBetweenRounds = 75

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
const warmupRounds = 2

// seqItems converts a fixed sequence (warm-up / cool-down) to base RItems.
func seqItems(seq []setDef, slot Slot) []RItem {
	out := make([]RItem, 0, len(seq))
	for _, s := range seq {
		e := Pool[s.ID]
		out = append(out, RItem{ID: s.ID, Unit: e.Unit, Value: s.Val, PerSide: e.PerSide, Slot: slot, Round: 1})
	}
	return out
}

// mainItems builds the base MAIN items for a day from the calendar. The fixed
// per-day variant bakes the natural day-to-day progression into the base values;
// the runtime applies the universal level on top.
func mainItems(day int) []RItem {
	plan := calendar[clampDay(day)]
	var out []RItem
	add := func(id string, val, round int) {
		e := Pool[id]
		out = append(out, RItem{ID: id, Unit: e.Unit, Value: val, PerSide: e.PerSide, Slot: Main, Round: round})
	}
	switch {
	case plan.Control:
		for r := 1; r <= 2; r++ {
			for _, s := range controlSeq {
				add(s.ID, s.Val, r)
			}
		}
		add("C02", controlPlankSec, 3) // plank challenge
	case plan.Recovery:
		for _, s := range recoverySeq {
			add(s.ID, s.Val, 1)
		}
	default:
		for r := 1; r <= 2; r++ {
			for _, bs := range blocks[plan.Code] {
				add(bs.ID, bs.Vals[plan.Variant], r)
			}
		}
	}
	return out
}
