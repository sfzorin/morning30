package content

// Vlad / Level 2 is the second standard set: a fitter profile (jumps allowed,
// knees fine). It adds plyometric/cardio work (jumps, burpees, mountain
// climbers, lunges) on top of the shared strength/core/back library. Same
// structure as Sergey — 6 warm-up, main block × 2 rounds, 4 cool-down, 20→30
// screens — but with explosive elements introduced gradually.

// vladWarmupSeq is Vlad's 6-exercise warm-up (one round), the same every day.
// Low pogo jumps (J01) carry a lighter substitute (J02, marching high knees)
// for recovery days.
var vladWarmupSeq = []setDef{
	{"W07", 20}, {"W01", 30}, {"W08", 30}, {"W04", 10}, {"L10", 10}, {"J01", 20},
}

// vladCooldownSeq is Vlad's 4-exercise cool-down (one round), the same every day.
var vladCooldownSeq = []setDef{
	{"CD07", 40}, {"CD04", 30}, {"CD02", 30}, {"CD03", 30},
}

// vladDayBlocks is Vlad's explicit MAIN block per day (1..30): an ordered list
// of {exercise, per-round value}. Each block runs for 2 rounds. Per-side values
// are stored once (the player doubles them). The difficulty level scales these
// at runtime; warm-up/cool-down and breaths are never scaled.
var vladDayBlocks = map[int][]setDef{
	1:  {{"P01", 14}, {"L10", 15}, {"J03", 30}, {"B01", 35}, {"C01", 60}},
	2:  {{"L11", 8}, {"B06", 14}, {"J04", 30}, {"C04", 30}, {"P06", 6}, {"C07", 30}},
	3:  {{"P03", 10}, {"J05", 8}, {"B02", 14}, {"C13", 30}, {"P02", 8}, {"J06", 30}, {"C02", 20}},
	4:  {{"C12", 10}, {"L05", 10}, {"B03", 12}, {"L12", 12}, {"C01", 50}}, // recovery
	5:  {{"P01", 15}, {"J07", 30}, {"L11", 10}, {"B01", 40}, {"C05", 12}, {"P06", 7}, {"C04", 32}},
	6:  {{"J08", 6}, {"J05", 10}, {"B06", 16}, {"J04", 35}, {"P02", 8}, {"L04", 10}, {"C03", 35}, {"C08", 8}},
	7:  {{"C11", 12}, {"B01", 30}, {"L06", 14}, {"B02", 12}, {"C01", 55}}, // recovery
	8:  {{"P01", 16}, {"J03", 40}, {"B03", 12}, {"C05", 12}, {"L11", 10}, {"P02", 9}, {"C02", 22}},
	9:  {{"P04", 6}, {"J06", 35}, {"B06", 16}, {"C04", 35}, {"L10", 18}, {"C10", 8}, {"B04", 12}, {"C07", 35}},
	10: {{"J08", 7}, {"L11", 10}, {"C03", 38}, {"B01", 40}, {"P06", 7}, {"L05", 12}, {"C13", 35}, {"C01", 70}},
	11: {{"C12", 10}, {"B02", 12}, {"L03", 14}, {"L06", 14}, {"C01", 60}}, // recovery
	12: {{"P03", 11}, {"J05", 12}, {"B02", 16}, {"J04", 40}, {"P02", 9}, {"C05", 13}, {"B11", 12}, {"J06", 40}, {"C02", 25}},
	13: {{"J07", 40}, {"P01", 16}, {"L04", 11}, {"B06", 16}, {"C06", 25}, {"P06", 8}, {"L11", 12}, {"C04", 38}},
	14: {{"C11", 14}, {"B01", 35}, {"L05", 10}, {"B03", 12}, {"C01", 60}}, // recovery
	15: {{"P04", 8}, {"J05", 12}, {"B03", 14}, {"C03", 40}, {"C10", 9}, {"J06", 45}, {"C05", 14}, {"C08", 10}, {"C02", 26}},
	16: {{"J08", 8}, {"B06", 18}, {"L11", 12}, {"C04", 40}, {"P06", 8}, {"L04", 12}, {"C13", 40}, {"C01", 80}},
	17: {{"P01", 16}, {"J07", 45}, {"B01", 45}, {"J09", 6}, {"P02", 10}, {"C05", 14}, {"B11", 14}, {"L10", 18}, {"C02", 28}},
	18: {{"C12", 10}, {"B01", 35}, {"L03", 14}, {"L05", 12}, {"C01", 65}}, // recovery
	19: {{"P03", 12}, {"J05", 14}, {"B06", 18}, {"J04", 45}, {"P06", 9}, {"L11", 12}, {"C03", 45}, {"B02", 18}, {"C08", 11}, {"C02", 30}},
	20: {{"J08", 8}, {"J06", 45}, {"P04", 8}, {"C04", 42}, {"L04", 12}, {"B04", 14}, {"P02", 10}, {"C06", 30}, {"C01", 85}},
	21: {{"C11", 14}, {"B02", 12}, {"L03", 14}, {"L05", 12}, {"C01", 65}}, // recovery
	22: {{"P05", 6}, {"J07", 45}, {"J05", 14}, {"B06", 18}, {"J04", 45}, {"C05", 15}, {"P06", 9}, {"J06", 45}, {"B11", 14}, {"C01", 90}},
	23: {{"P01", 14}, {"B01", 40}, {"L12", 15}, {"C07", 40}, {"B02", 16}, {"C03", 35}, {"C01", 75}}, // technique
	24: {{"J08", 9}, {"P03", 12}, {"J09", 8}, {"C04", 45}, {"P02", 10}, {"B03", 16}, {"L05", 14}, {"C13", 45}, {"J06", 45}, {"C02", 30}},
	25: {{"C12", 10}, {"B01", 35}, {"L03", 14}, {"B02", 12}, {"C01", 65}}, // recovery
	26: {{"P04", 9}, {"J05", 15}, {"B06", 20}, {"J04", 50}, {"P06", 10}, {"C05", 16}, {"L11", 14}, {"B04", 16}, {"C06", 35}, {"C01", 90}},
	27: {{"J07", 50}, {"P01", 16}, {"J06", 50}, {"L04", 13}, {"B02", 18}, {"P02", 10}, {"C03", 45}, {"C08", 12}, {"C02", 30}},
	28: {{"C11", 14}, {"B01", 35}, {"L06", 14}, {"B02", 12}, {"C01", 65}}, // recovery
	29: {{"J08", 10}, {"P04", 10}, {"J05", 16}, {"B06", 20}, {"J04", 50}, {"P06", 10}, {"J09", 8}, {"C05", 16}, {"B11", 16}, {"C02", 30}}, // peak
	30: {{"P01", 20}, {"J03", 60}, {"B06", 20}, {"J05", 15}, {"P02", 12}, {"C03", 50}, {"L11", 14}, {"B02", 20}, {"C05", 18}, {"C01", 90}}, // control
}

var vladSpec = programSpec{name: "Vlad", warmup: vladWarmupSeq, cooldown: vladCooldownSeq, days: vladDayBlocks}

// resolveVlad builds the Vlad / Level 2 standard program.
func resolveVlad() Resolved { return vladSpec.resolve() }
