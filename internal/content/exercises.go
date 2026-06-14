// Package content defines the exercise library and the 30-day programs. The
// library is shared by every standard set; each set (Sergey, Vlad, …) selects
// from it per day. Floor/mat only, no equipment. The Sergey set is jump-free;
// the Vlad set adds plyometric/cardio movements (the J* group: jumps, burpees,
// mountain climbers) and lunges for a fitter profile.
package content

// Unit is how an exercise is measured.
type Unit string

const (
	Reps    Unit = "reps"    // counted repetitions
	Seconds Unit = "seconds" // held/timed
	Breaths Unit = "breaths" // counted breaths (paced by the user)
)

// Slot is where an exercise belongs in a session.
type Slot string

const (
	Warmup   Slot = "warmup"
	Main     Slot = "main"
	Cooldown Slot = "cooldown"
)

// KneeLoad flags exercises that load the knee, for the (phase-3) adaptation engine.
type KneeLoad string

const (
	KneeNone   KneeLoad = "none"
	KneeLow    KneeLoad = "low"
	KneeMedium KneeLoad = "medium"
)

// Exercise is one movement in the library. Magnitudes are NOT stored here — they
// come from the block tables per day. The ID doubles as the SVG asset name and
// the i18n key.
type Exercise struct {
	ID          string
	Slot        Slot
	Unit        Unit
	PerSide     bool
	KneeLoad    KneeLoad
	Difficulty  int    // 1–5
	Replacement string // substitute ID used by the adaptation engine
	Video       string // optional video/GIF URL (reserved; falls back to the SVG)
	Thumb       string // optional thumbnail URL (reserved)
}

// Pool is the library keyed by ID.
var Pool = func() map[string]Exercise {
	m := make(map[string]Exercise, len(library))
	for _, e := range library {
		m[e.ID] = e
	}
	return m
}()

// Get returns an exercise by ID.
func Get(id string) (Exercise, bool) {
	e, ok := Pool[id]
	return e, ok
}

// Library returns the built-in exercises in canonical (display) order.
func Library() []Exercise {
	out := make([]Exercise, len(library))
	copy(out, library)
	return out
}

var library = []Exercise{
	// ---- Warm-up (fixed, 6 exercises, 1 round) ----
	{ID: "W01", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},          // shoulder circles
	{ID: "W02", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},          // small arm circles
	{ID: "W03", Slot: Warmup, Unit: Reps, PerSide: true, Difficulty: 1, KneeLoad: KneeNone}, // standing twists
	{ID: "W04", Slot: Warmup, Unit: Reps, Difficulty: 1, KneeLoad: KneeNone},             // hip hinge no weight
	{ID: "W05", Slot: Warmup, Unit: Reps, Difficulty: 1, KneeLoad: KneeLow},              // slow shallow squat
	{ID: "W06", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},          // high plank weight shift
	{ID: "W07", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},          // standing breathing + reach (Vlad)
	{ID: "W08", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},          // standing torso rotations (Vlad)

	// ---- Abs / core ----
	{ID: "C01", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone},                      // forearm plank
	{ID: "C02", Slot: Main, Unit: Seconds, Difficulty: 4, KneeLoad: KneeNone, Replacement: "C01"},  // RKC plank
	{ID: "C03", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone, Replacement: "C12"},  // scissors
	{ID: "C04", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone, Replacement: "C12"},  // flutter kicks
	{ID: "C05", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeNone},                         // reverse crunch
	{ID: "C06", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone, Replacement: "C12"},  // hollow hold (easy)
	{ID: "C07", Slot: Main, Unit: Seconds, PerSide: true, Difficulty: 3, KneeLoad: KneeNone},       // side plank
	{ID: "C08", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 3, KneeLoad: KneeNone},          // side plank hip lift
	{ID: "C09", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeNone},                         // plank shoulder taps
	{ID: "C10", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeNone},                         // up-down plank
	{ID: "C11", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeNone},          // heel taps (per side)
	{ID: "C12", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeNone},          // dead bug advanced
	{ID: "C13", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone},                      // bicycle crunches (Vlad)

	// ---- Push-ups / triceps / shoulders ----
	{ID: "P01", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeNone},                         // classic push-up
	{ID: "P02", Slot: Main, Unit: Reps, Difficulty: 4, KneeLoad: KneeNone, Replacement: "P01"},     // narrow push-up
	{ID: "P03", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeNone, Replacement: "P01"},     // paused push-up
	{ID: "P04", Slot: Main, Unit: Reps, Difficulty: 4, KneeLoad: KneeNone, Replacement: "P03"},     // slow 3-1-3 push-up
	{ID: "P05", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 4, KneeLoad: KneeNone, Replacement: "P01"}, // staggered push-up (per side)
	{ID: "P06", Slot: Main, Unit: Reps, Difficulty: 4, KneeLoad: KneeNone, Replacement: "P03"},     // pike push-up
	{ID: "P08", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},                         // sphinx push-up

	// ---- Back / scapula ----
	{ID: "B01", Slot: Main, Unit: Seconds, Difficulty: 2, KneeLoad: KneeNone}, // prone cobra hold
	{ID: "B02", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // reverse snow angels
	{ID: "B03", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // prone W-raise
	{ID: "B04", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // prone Y-raise
	{ID: "B05", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // prone T-raise
	{ID: "B06", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // superman pull-down
	{ID: "B07", Slot: Main, Unit: Seconds, Difficulty: 2, KneeLoad: KneeNone}, // swimmers
	{ID: "B08", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeNone}, // reverse plank
	{ID: "B09", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // back extension pulses
	{ID: "B10", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeNone},    // cobra to W-pull
	{ID: "B11", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeNone}, // swimmers, counted per side (Vlad)

	// ---- Arms / biceps (self-resistance, no equipment) ----
	{ID: "A01", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeNone},    // self-resist curl
	{ID: "A02", Slot: Main, Unit: Seconds, PerSide: true, Difficulty: 2, KneeLoad: KneeNone}, // biceps isometric 90°
	{ID: "A03", Slot: Main, Unit: Seconds, Difficulty: 2, KneeLoad: KneeNone},                // hand-hook pull
	{ID: "A04", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeNone},    // slow negative curl

	// ---- Legs / glutes ----
	{ID: "L01", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeMedium, Replacement: "L03"}, // slow shallow squat
	{ID: "L02", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeMedium, Replacement: "L03"}, // paused squat
	{ID: "L03", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeLow},                        // hip hinge
	{ID: "L04", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeLow},         // single-leg glute bridge
	{ID: "L05", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeLow},         // glute bridge march
	{ID: "L06", Slot: Main, Unit: Reps, Difficulty: 1, KneeLoad: KneeLow},                        // glute bridge hold
	{ID: "L07", Slot: Main, Unit: Reps, Difficulty: 1, KneeLoad: KneeNone},                       // calf raises
	{ID: "L09", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 3, KneeLoad: KneeNone},        // side plank leg lift
	{ID: "L10", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeMedium, Replacement: "L01"},                // air squat (Vlad)
	{ID: "L11", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 2, KneeLoad: KneeMedium, Replacement: "L03"}, // reverse lunge (Vlad)
	{ID: "L12", Slot: Main, Unit: Reps, Difficulty: 2, KneeLoad: KneeMedium, Replacement: "L01"},                // slow air squat (Vlad)

	// ---- Cardio / plyometrics (Vlad set) ----
	{ID: "J01", Slot: Warmup, Unit: Seconds, Difficulty: 2, KneeLoad: KneeLow, Replacement: "J02"},              // low pogo jumps
	{ID: "J02", Slot: Warmup, Unit: Seconds, Difficulty: 1, KneeLoad: KneeLow},                                  // marching high knees (pogo substitute)
	{ID: "J03", Slot: Main, Unit: Seconds, Difficulty: 2, KneeLoad: KneeLow},                                    // jumping jacks
	{ID: "J04", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeLow},                                    // mountain climbers
	{ID: "J05", Slot: Main, Unit: Reps, Difficulty: 3, KneeLoad: KneeMedium, Replacement: "L10"},                // squat jumps
	{ID: "J06", Slot: Main, Unit: Seconds, Difficulty: 3, KneeLoad: KneeMedium, Replacement: "J03"},             // skater hops
	{ID: "J07", Slot: Main, Unit: Seconds, Difficulty: 2, KneeLoad: KneeLow, Replacement: "J02"},                // high knees
	{ID: "J08", Slot: Main, Unit: Reps, Difficulty: 4, KneeLoad: KneeMedium, Replacement: "J03"},                // burpees
	{ID: "J09", Slot: Main, Unit: Reps, PerSide: true, Difficulty: 4, KneeLoad: KneeMedium, Replacement: "L11"}, // jump lunges (per side)

	// ---- Cool-down (fixed, 4 exercises, 1 round) ----
	{ID: "CD01", Slot: Cooldown, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},                      // chest stretch
	{ID: "CD02", Slot: Cooldown, Unit: Seconds, PerSide: true, Difficulty: 1, KneeLoad: KneeNone},       // hamstring stretch
	{ID: "CD03", Slot: Cooldown, Unit: Seconds, PerSide: true, Difficulty: 1, KneeLoad: KneeNone},       // supine twist
	{ID: "CD04", Slot: Cooldown, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone, Replacement: "CD05"}, // sphinx pose
	{ID: "CD05", Slot: Cooldown, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},                      // lying breathing (sphinx replacement)
	{ID: "CD07", Slot: Cooldown, Unit: Seconds, Difficulty: 1, KneeLoad: KneeNone},                      // child's pose (Vlad)
}
