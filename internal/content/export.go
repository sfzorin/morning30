package content

// Snapshot types describe the whole program in a serializable shape so it can
// be exported to JSON (download) and, later, replaced by a custom uploaded one.

type SnapExercise struct {
	ID          string   `json:"id"`
	Unit        Unit     `json:"unit"`
	PerSide     bool     `json:"perSide"`
	KneeLoad    KneeLoad `json:"kneeLoad"`
	Difficulty  int      `json:"difficulty"`
	Replacement string   `json:"replacement,omitempty"`
	SVG         string   `json:"svg"`
}

type SnapSet struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type SnapBlockSet struct {
	ID   string `json:"id"`
	Vals [3]int `json:"vals"` // [−, base, +]
}

type SnapSeq struct {
	Rounds int       `json:"rounds"`
	Sets   []SnapSet `json:"sets"`
}

type SnapDay struct {
	Day      int    `json:"day"`
	Block    string `json:"block,omitempty"`
	Variant  int    `json:"variant"`
	Recovery bool   `json:"recovery,omitempty"`
	Control  bool   `json:"control,omitempty"`
}

// Snapshot is the full, serializable program (structure only; localized text is
// added by the program exporter which can read the i18n tables).
type Snapshot struct {
	Version         int                       `json:"version"`
	TotalDays       int                       `json:"totalDays"`
	MainRounds      int                       `json:"mainRounds"`
	Warmup          SnapSeq                   `json:"warmup"`
	Cooldown        SnapSeq                   `json:"cooldown"`
	Recovery        SnapSeq                   `json:"recovery"`
	Control         SnapSeq                   `json:"control"`
	ControlPlankSec int                       `json:"controlPlankSec"`
	Blocks          map[string][]SnapBlockSet `json:"blocks"`
	Calendar        []SnapDay                 `json:"calendar"`
	Exercises       []SnapExercise            `json:"exercises"`
}

func seq(rounds int, defs []setDef) SnapSeq {
	s := SnapSeq{Rounds: rounds, Sets: make([]SnapSet, 0, len(defs))}
	for _, d := range defs {
		s.Sets = append(s.Sets, SnapSet{ID: d.ID, Value: d.Val})
	}
	return s
}

// TakeSnapshot returns the built-in program as a serializable Snapshot.
func TakeSnapshot() Snapshot {
	s := Snapshot{
		Version:         1,
		TotalDays:       TotalDays,
		MainRounds:      2,
		Warmup:          seq(2, warmupSeq),
		Cooldown:        seq(1, cooldownSeq),
		Recovery:        seq(1, recoverySeq),
		Control:         seq(2, controlSeq),
		ControlPlankSec: controlPlankSec,
		Blocks:          map[string][]SnapBlockSet{},
	}
	for code, sets := range blocks {
		bs := make([]SnapBlockSet, 0, len(sets))
		for _, x := range sets {
			bs = append(bs, SnapBlockSet{ID: x.ID, Vals: x.Vals})
		}
		s.Blocks[code] = bs
	}
	for day := 1; day <= TotalDays; day++ {
		p := calendar[day]
		s.Calendar = append(s.Calendar, SnapDay{
			Day: day, Block: p.Code, Variant: p.Variant, Recovery: p.Recovery, Control: p.Control,
		})
	}
	for _, e := range library {
		s.Exercises = append(s.Exercises, SnapExercise{
			ID: e.ID, Unit: e.Unit, PerSide: e.PerSide, KneeLoad: e.KneeLoad,
			Difficulty: e.Difficulty, Replacement: e.Replacement,
			SVG: "/static/ex/" + e.ID + ".svg",
		})
	}
	return s
}
