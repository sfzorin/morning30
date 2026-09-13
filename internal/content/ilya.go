package content

// Ilya is the third standard set: a home tennis-prep routine for an 8-year-old
// who trains alone and races yesterday, not a partner. Same skeleton as
// Sergey/Vlad — 6 warm-up, main block × 2 rounds, 4 cool-down, 20→30 screens —
// with kid volumes, no burpees/jump-lunges/pike/RKC, and five 0.5 kg dumbbell
// drills (D01–D05). Dumbbells stay out of the warm-up, cool-down and all jumps.

// ilyaWarmupSeq is Ilya's 6-exercise warm-up (one round), the same every day.
var ilyaWarmupSeq = []setDef{
	{"W09", 15}, {"W01", 15}, {"W03", 6}, {"W04", 8}, {"W05", 8}, {"J02", 20},
}

// ilyaCooldownSeq is Ilya's 4-exercise cool-down (one round), the same every day.
var ilyaCooldownSeq = []setDef{
	{"CD01", 20}, {"CD02", 20}, {"CD03", 20}, {"CD07", 25},
}

// ilyaDayBlocks is Ilya's MAIN block per day (1..30): {exercise, per-round value}.
// Each block runs for 2 rounds. Per-side values are stored once. The difficulty
// level scales these at runtime; warm-up/cool-down and breaths are never scaled.
// Recovery days 4/7/11/14/18/21/25/28 stay at 5 exercises; day 23 is technique,
// 29 is peak, 30 is control (slightly easier than 29).
var ilyaDayBlocks = map[int][]setDef{
	1:  {{"P08", 6}, {"L10", 8}, {"J03", 20}, {"D01", 8}, {"C01", 20}},
	2:  {{"L11", 6}, {"B03", 8}, {"J07", 20}, {"C13", 20}, {"D02", 8}, {"C07", 10}},
	3:  {{"P08", 6}, {"J06", 15}, {"D03", 8}, {"C12", 6}, {"L10", 8}, {"J03", 20}, {"C01", 22}},
	4:  {{"C11", 8}, {"B01", 20}, {"L06", 8}, {"D01", 8}, {"C01", 20}}, // recovery
	5:  {{"L10", 10}, {"J03", 25}, {"C09", 6}, {"D04", 10}, {"L11", 6}, {"C13", 20}, {"C07", 12}},
	6:  {{"P08", 7}, {"J06", 20}, {"B06", 8}, {"D05", 12}, {"C09", 6}, {"L05", 6}, {"C13", 22}, {"C01", 25}},
	7:  {{"C11", 10}, {"B01", 25}, {"L06", 10}, {"D01", 8}, {"C01", 22}}, // recovery
	8:  {{"L10", 10}, {"J07", 25}, {"D02", 8}, {"C05", 8}, {"L11", 6}, {"C07", 12}, {"C01", 25}},
	9:  {{"P08", 8}, {"J06", 20}, {"C13", 24}, {"D03", 8}, {"B04", 8}, {"C09", 8}, {"L07", 12}, {"C01", 25}},
	10: {{"J03", 25}, {"L11", 8}, {"D05", 12}, {"C07", 12}, {"L05", 8}, {"C13", 22}, {"C01", 28}},
	11: {{"C12", 6}, {"B02", 8}, {"L03", 10}, {"D01", 8}, {"C01", 22}}, // recovery
	12: {{"P08", 8}, {"J05", 6}, {"D02", 8}, {"J06", 20}, {"L10", 10}, {"C13", 24}, {"C09", 8}, {"L07", 12}, {"C01", 28}},
	13: {{"L11", 8}, {"J07", 25}, {"D04", 12}, {"C08", 6}, {"L10", 10}, {"B10", 8}, {"C13", 24}, {"C01", 30}},
	14: {{"C11", 10}, {"B01", 25}, {"L06", 10}, {"D01", 8}, {"C01", 25}}, // recovery
	15: {{"P01", 5}, {"J06", 22}, {"C12", 8}, {"D03", 10}, {"B03", 10}, {"C07", 14}, {"J03", 25}, {"C01", 30}},
	16: {{"L10", 12}, {"J05", 8}, {"B06", 10}, {"C13", 25}, {"P08", 8}, {"D05", 15}, {"C09", 8}, {"L07", 14}, {"C01", 32}},
	17: {{"J07", 25}, {"B01", 30}, {"L11", 8}, {"J06", 22}, {"D02", 10}, {"C07", 14}, {"C01", 30}},
	18: {{"C12", 6}, {"B01", 25}, {"L03", 10}, {"D01", 8}, {"C01", 25}}, // recovery
	19: {{"P08", 8}, {"J05", 8}, {"D04", 12}, {"J06", 25}, {"L10", 12}, {"C13", 26}, {"C09", 8}, {"L11", 8}, {"C01", 32}},
	20: {{"P01", 6}, {"J03", 30}, {"C13", 26}, {"L11", 8}, {"D02", 10}, {"C07", 15}, {"J06", 25}, {"L07", 14}, {"C09", 8}, {"C01", 35}},
	21: {{"C11", 10}, {"B02", 10}, {"L03", 12}, {"D01", 8}, {"C01", 25}}, // recovery
	22: {{"P08", 8}, {"J05", 8}, {"D03", 10}, {"J07", 25}, {"L10", 12}, {"C13", 28}, {"C08", 6}, {"L11", 8}, {"C01", 35}},
	23: {{"L12", 8}, {"D01", 8}, {"C12", 8}, {"P08", 6}, {"D02", 8}, {"C07", 12}, {"C01", 25}}, // technique
	24: {{"P01", 6}, {"J06", 25}, {"B06", 12}, {"J05", 8}, {"L11", 8}, {"C13", 28}, {"D04", 15}, {"L07", 14}, {"C07", 15}, {"C01", 35}},
	25: {{"C12", 8}, {"B01", 25}, {"L06", 10}, {"D01", 8}, {"C01", 25}}, // recovery
	26: {{"L10", 12}, {"J07", 30}, {"D05", 15}, {"J06", 25}, {"P08", 8}, {"C13", 28}, {"L11", 8}, {"C09", 10}, {"L07", 14}, {"C01", 35}},
	27: {{"P01", 8}, {"J03", 30}, {"B06", 12}, {"D03", 10}, {"C07", 15}, {"J05", 8}, {"C13", 28}, {"C01", 35}},
	28: {{"C11", 12}, {"B01", 30}, {"L06", 12}, {"D01", 10}, {"C01", 28}}, // recovery
	29: {{"P01", 8}, {"J06", 28}, {"C13", 30}, {"D03", 10}, {"J05", 8}, {"D04", 15}, {"C09", 10}, {"L11", 8}, {"C07", 16}, {"C01", 40}}, // peak
	30: {{"P08", 8}, {"J03", 30}, {"D02", 10}, {"L10", 12}, {"C13", 28}, {"L11", 8}, {"D05", 15}, {"J06", 25}, {"L07", 14}, {"C01", 35}}, // control
}

var ilyaSpec = programSpec{name: "Ilya", warmup: ilyaWarmupSeq, cooldown: ilyaCooldownSeq, days: ilyaDayBlocks}

// resolveIlya builds the Ilya kid-tennis standard program.
func resolveIlya() Resolved { return ilyaSpec.resolve() }
