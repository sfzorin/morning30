package content

// The ready-program catalog: everything a user can pick from the "Programs"
// screen. It is the morning workout (the public demo) plus the goal-named
// "… in 30 days" courses. Picking one copies it into the user's selected
// programs; the catalog itself never changes per user.

// Category keys group ready programs on the Programs screen.
const (
	CatStart       = "start"       // first steps
	CatStrength    = "strength"    // plank, abs, push-ups, legs
	CatWeightLoss  = "weightloss"  // cardio circuits
	CatStretch     = "stretch"     // flexibility
	CatBackBalance = "backbalance" // posture, back, balance
	CatKids        = "kids"        // children
	CatSport       = "sport"       // supplementary training for athletes
)

// Categories lists the category keys in display order.
func Categories() []string {
	return []string{CatStart, CatStrength, CatWeightLoss, CatStretch, CatBackBalance, CatKids, CatSport}
}

// DemoKey is the catalog key of the morning workout (the public demo).
const DemoKey = "demo"

// ReadyProgram is one program of the catalog.
type ReadyProgram struct {
	Key      string            // stable id stored with a selected program
	Slug     string            // public URL name ("" when it has no public page)
	Category string            // one of the Cat* keys
	Kids     bool              // a children's program: shown with a child figure
	Title    map[string]string // language code → title
	About    map[string]string // language code → one-sentence description
	Build    func() Resolved
}

// TitleIn returns the title in language l, falling back to English.
func (p ReadyProgram) TitleIn(l string) string {
	if t := p.Title[l]; t != "" {
		return t
	}
	return p.Title["en"]
}

// AboutIn returns the description in language l, falling back to English.
func (p ReadyProgram) AboutIn(l string) string {
	if t := p.About[l]; t != "" {
		return t
	}
	return p.About["en"]
}

// Tags are facts read from the program itself, so they can never drift from
// what the user actually gets.
type Tags struct {
	Jumps     bool `json:"jumps"`     // has at least one jumping exercise
	Equipment bool `json:"equipment"` // needs anything beyond a mat
}

// Tags reports the program's jump and equipment needs.
func (p ReadyProgram) Tags() Tags { return ProgramTags(p.Build()) }

// ProgramTags reads the tags of any program, ready or edited.
func ProgramTags(r Resolved) Tags {
	var t Tags
	visit := func(items []RItem) {
		for _, it := range items {
			if IsJump(it.ID) {
				t.Jumps = true
			}
			if NeedsEquipment(it.ID) {
				t.Equipment = true
			}
		}
	}
	visit(r.Warmup)
	visit(r.Cooldown)
	for i, d := range r.Days {
		visit(d.Items)
		visit(r.WarmupFor(i))
		visit(r.CooldownFor(i))
	}
	return t
}

// jumpIDs are the built-in exercises that leave the floor.
var jumpIDs = map[string]bool{"J01": true, "J03": true, "J05": true, "J06": true, "J07": true, "J08": true, "J09": true}

// IsJump reports whether a built-in exercise involves jumping.
func IsJump(id string) bool { return jumpIDs[id] || bundledDefs[id].Jumps }

// NeedsEquipment reports whether a built-in exercise needs equipment: the
// light dumbbells of the D group, or a bundled exercise that says so.
func NeedsEquipment(id string) bool {
	return (len(id) == 3 && id[0] == 'D') || bundledDefs[id].Equipment
}

// MinutesSpan is a program's session length on day 1 and on the peak day 29,
// with the rest settings of a new account.
func MinutesSpan(r Resolved) (first, peak int) {
	rests := Rests{Warmup: 10, Main: 20, Cooldown: 10}
	min := func(sec int) int {
		if m := (sec + 59) / 60; m > 1 {
			return m
		}
		return 1
	}
	return min(r.Workout(1, 0, rests).EstSec), min(r.Workout(TotalDays-1, 0, rests).EstSec)
}

var demoProgram = ReadyProgram{
	Key:      DemoKey,
	Category: CatStart,
	Title: map[string]string{
		"ru": "Утренняя зарядка", "en": "Morning Workout", "tr": "Sabah Egzersizi", "de": "Morgengymnastik",
		"es": "Rutina matutina", "fr": "Gym du matin", "it": "Ginnastica mattutina",
	},
	About: map[string]string{
		"ru": "Короткая зарядка на всё тело без прыжков: мягкий старт дня и плавный рост нагрузки.",
		"en": "A short whole-body routine with no jumps: a gentle start to the day and a steady build-up.",
		"tr": "Zıplamasız kısa tüm vücut rutini: güne yumuşak bir başlangıç ve kademeli artış.",
		"de": "Kurzes Ganzkörper-Workout ohne Sprünge: sanfter Start in den Tag und stetige Steigerung.",
		"es": "Rutina corta de cuerpo completo sin saltos: un inicio suave del día y progresión constante.",
		"fr": "Une courte routine pour tout le corps, sans sauts : un réveil en douceur et une progression régulière.",
		"it": "Breve routine per tutto il corpo senza salti: un inizio dolce della giornata e una crescita graduale.",
	},
	Build: ResolveDemo,
}

// courseCategory places each "… in 30 days" course in a catalog section.
var courseCategory = map[string]string{
	"plank30": CatStrength, "abs30": CatStrength, "pushups30": CatStrength, "legs30": CatStrength,
	"slim30": CatWeightLoss, "posture30": CatBackBalance,
}

// Catalog returns every ready program in display order: the morning workout
// first, then the courses.
func Catalog() []ReadyProgram {
	out := []ReadyProgram{demoProgram}
	for _, c := range series30 {
		c := c
		cat := c.Category
		if cat == "" {
			cat = courseCategory[c.Key]
		}
		if cat == "" {
			cat = CatStrength
		}
		out = append(out, ReadyProgram{
			Key: c.Key, Slug: c.Slug, Category: cat, Kids: c.Kids, Title: c.Title, About: c.About,
			Build: c.Resolve,
		})
	}
	return out
}

// ReadyByKey returns a catalog program by its key.
func ReadyByKey(key string) (ReadyProgram, bool) {
	for _, p := range Catalog() {
		if p.Key == key {
			return p, true
		}
	}
	return ReadyProgram{}, false
}

// PopularKeys are the programs offered on a newcomer's home screen.
func PopularKeys() []string { return []string{DemoKey, "plank30", "abs30", "slim30"} }

// MatchReady returns the key of the ready program a program is an untouched
// copy of (the name is ignored), or "" when it matches none.
func MatchReady(r Resolved) string {
	body := withoutName(r)
	for _, p := range Catalog() {
		if withoutName(p.Build()) == body {
			return p.Key
		}
	}
	return ""
}

// SameAsReady reports whether a program is an untouched copy of the given
// ready program (the name is ignored).
func SameAsReady(r Resolved, key string) bool {
	p, ok := ReadyByKey(key)
	return ok && withoutName(p.Build()) == withoutName(r)
}

func withoutName(r Resolved) string {
	r.Name = ""
	return r.Marshal()
}
