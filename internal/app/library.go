package app

import (
	"sort"
	"strings"

	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
)

// Exercise library groups, in display order.
const (
	GroupWarmup   = "warmup"
	GroupCore     = "core"
	GroupPush     = "push"
	GroupBack     = "back"
	GroupArms     = "arms"
	GroupLegs     = "legs"
	GroupCardio   = "cardio"
	GroupStretch  = "stretch"
	GroupDumbbell = "dumbbell"
	GroupPullup   = "pullup"
	GroupBalance  = "balance"
	GroupTaiChi   = "taichi"
	GroupNeck     = "neck"
	GroupKids     = "kids"
	GroupCustom   = "custom"
)

// LibraryGroups lists the library groups in display order.
func LibraryGroups() []string {
	return []string{GroupWarmup, GroupCore, GroupPush, GroupPullup, GroupBack, GroupArms, GroupLegs, GroupCardio,
		GroupStretch, GroupNeck, GroupBalance, GroupTaiChi, GroupDumbbell, GroupKids, GroupCustom}
}

// ExerciseInfo is one library exercise as screens and agents see it.
type ExerciseInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Group      string `json:"group"`
	Slot       string `json:"slot"`
	Unit       string `json:"unit"`
	PerSide    bool   `json:"perSide"`
	Difficulty int    `json:"difficulty"`
	Jumps      bool   `json:"jumps"`
	Equipment  bool   `json:"equipment"`
	Custom     bool   `json:"custom"`
}

// ExerciseGroup places an exercise in a library group.
func ExerciseGroup(cat catalog.Catalog, id string) string {
	if cat.IsCustom(id) {
		return GroupCustom
	}
	switch {
	case strings.HasPrefix(id, "CD"), strings.HasPrefix(id, "ST"):
		return GroupStretch
	case strings.HasPrefix(id, "PU"):
		return GroupPullup
	case strings.HasPrefix(id, "BL"):
		return GroupBalance
	case strings.HasPrefix(id, "TC"):
		return GroupTaiChi
	case strings.HasPrefix(id, "NK"):
		return GroupNeck
	case strings.HasPrefix(id, "IT"):
		return GroupKids
	case strings.HasPrefix(id, "W"):
		return GroupWarmup
	case strings.HasPrefix(id, "C"):
		return GroupCore
	case strings.HasPrefix(id, "P"):
		return GroupPush
	case strings.HasPrefix(id, "B"):
		return GroupBack
	case strings.HasPrefix(id, "A"):
		return GroupArms
	case strings.HasPrefix(id, "L"):
		return GroupLegs
	case strings.HasPrefix(id, "J"):
		return GroupCardio
	case strings.HasPrefix(id, "D"):
		return GroupDumbbell
	}
	if def, ok := cat.Def(id); ok {
		switch def.Slot {
		case content.Warmup:
			return GroupWarmup
		case content.Cooldown:
			return GroupStretch
		}
	}
	return GroupCustom
}

func exerciseInfo(cat catalog.Catalog, l i18n.Lang, id string) (ExerciseInfo, bool) {
	def, ok := cat.Def(id)
	if !ok {
		return ExerciseInfo{}, false
	}
	return ExerciseInfo{
		ID: id, Name: cat.Name(l, id), Group: ExerciseGroup(cat, id), Slot: string(def.Slot), Unit: string(def.Unit),
		PerSide: def.PerSide, Difficulty: def.Difficulty, Jumps: content.IsJump(id),
		Equipment: content.NeedsEquipment(id), Custom: cat.IsCustom(id),
	}, true
}

// Library lists the exercises a user can use. group "" means every group;
// query matches the name or id in any case; noJumps hides jumping exercises.
func Library(userID int64, lang, group, query string, noJumps bool) []ExerciseInfo {
	cat := UserCatalog(userID)
	l := i18n.Lang(lang)
	query = strings.ToLower(strings.TrimSpace(query))
	var out []ExerciseInfo
	for _, e := range cat.List(l) {
		info, ok := exerciseInfo(cat, l, e.ID)
		if !ok || (group != "" && info.Group != group) || (noJumps && info.Jumps) {
			continue
		}
		if query != "" && !strings.Contains(strings.ToLower(info.Name), query) && !strings.Contains(strings.ToLower(info.ID), query) {
			continue
		}
		out = append(out, info)
	}
	rank := map[string]int{}
	for i, g := range LibraryGroups() {
		rank[g] = i
	}
	sort.SliceStable(out, func(i, j int) bool { return rank[out[i].Group] < rank[out[j].Group] })
	return out
}

// Similar are the exercises that can replace one: same group and part of the
// day, split by difficulty.
type Similar struct {
	Easier []ExerciseInfo `json:"easier"`
	Same   []ExerciseInfo `json:"same"`
	Harder []ExerciseInfo `json:"harder"`
}

// SimilarExercises finds replacements for an exercise.
func SimilarExercises(userID int64, lang, id string) (Similar, error) {
	cat := UserCatalog(userID)
	l := i18n.Lang(lang)
	base, ok := exerciseInfo(cat, l, strings.TrimSpace(id))
	if !ok {
		return Similar{}, ErrUnknownExercise
	}
	out := Similar{Easier: []ExerciseInfo{}, Same: []ExerciseInfo{}, Harder: []ExerciseInfo{}}
	for _, e := range Library(userID, lang, base.Group, "", false) {
		if e.ID == base.ID || e.Slot != base.Slot {
			continue
		}
		switch {
		case e.Difficulty < base.Difficulty:
			out.Easier = append(out.Easier, e)
		case e.Difficulty > base.Difficulty:
			out.Harder = append(out.Harder, e)
		default:
			out.Same = append(out.Same, e)
		}
	}
	return out, nil
}

// DayItem is one exercise of a day's plan, with the dose the user will do.
type DayItem struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Part    string `json:"part"` // warmup, main or cooldown
	Unit    string `json:"unit"`
	Value   int    `json:"value"`
	PerSide bool   `json:"perSide"`
}

// DayPlan is what a day of the current program holds.
type DayPlan struct {
	Program    string    `json:"program"`
	Day        int       `json:"day"`
	Exercises  int       `json:"exercises"` // screens in the workout, both rounds counted
	MainRounds int       `json:"mainRounds"`
	EstMinutes int       `json:"estMinutes"`
	Items      []DayItem `json:"items"` // the main part is listed once
}

// Day returns the plan of a day (1..30) of the current program; day 0 is the
// user's next day.
func Day(userID int64, lang string, day int) DayPlan {
	if day < 1 || day > content.TotalDays {
		day = CurrentDay(userID)
	}
	w := UserWorkout(userID, day)
	cat := UserCatalog(userID)
	l := i18n.Lang(lang)
	plan := DayPlan{Program: ProgramTitle(userID, lang), Day: day, Exercises: len(w.Items), MainRounds: DB.GetRounds(userID), EstMinutes: (w.EstSec + 59) / 60}
	for _, it := range w.Items {
		if it.Slot == content.Main && it.Round > 1 {
			continue
		}
		plan.Items = append(plan.Items, DayItem{
			ID: it.ExerciseID, Name: cat.Name(l, it.ExerciseID), Part: string(it.Slot),
			Unit: string(it.Unit), Value: it.Value, PerSide: it.PerSide,
		})
	}
	return plan
}

// Settings are the account preferences agents and the settings page change.
type Settings struct {
	Name         string `json:"name"`
	Lang         string `json:"lang"`
	RestWarmup   int    `json:"restWarmup"`
	RestMain     int    `json:"restMain"`
	RestCooldown int    `json:"restCooldown"`
	VoiceMode    string `json:"voiceMode"` // off, min, normal or detailed
	Figure       string `json:"figure"`    // m, f, b or g
	Level        int    `json:"level"`     // difficulty −3..+3
	Rounds       int    `json:"rounds"`    // main rounds 1..10
}

// SettingsPatch changes only the fields that are set.
type SettingsPatch struct {
	Name         *string `json:"name,omitempty"`
	Lang         *string `json:"lang,omitempty"`
	RestWarmup   *int    `json:"restWarmup,omitempty"`
	RestMain     *int    `json:"restMain,omitempty"`
	RestCooldown *int    `json:"restCooldown,omitempty"`
	VoiceMode    *string `json:"voiceMode,omitempty"`
	Figure       *string `json:"figure,omitempty"`
	Level        *int    `json:"level,omitempty"`
	Rounds       *int    `json:"rounds,omitempty"`
}

// GetSettings reads the account preferences.
func GetSettings(userID int64) (Settings, error) {
	u, err := DB.UserByID(userID)
	if err != nil {
		return Settings{}, err
	}
	w, m, c := DB.Rests(userID)
	vm := u.VoiceMode
	if vm == "" {
		vm = "normal"
	}
	return Settings{Name: u.Name, Lang: u.Lang, RestWarmup: w, RestMain: m, RestCooldown: c,
		VoiceMode: vm, Figure: DB.GetSex(userID), Level: DB.GetLevel(userID), Rounds: DB.GetRounds(userID)}, nil
}

// ErrInvalidSettings is returned for an unsupported language, voice mode or figure.
var ErrInvalidSettings = errSettings{}

type errSettings struct{}

func (errSettings) Error() string {
	return "invalid settings: lang must be one of ru en tr de es fr it, voiceMode off|min|normal|detailed, figure m|f|b|g"
}

// UpdateSettings applies a patch; rests and level are clamped to their ranges.
func UpdateSettings(userID int64, p SettingsPatch) (Settings, error) {
	s, err := GetSettings(userID)
	if err != nil {
		return s, err
	}
	if p.Name != nil {
		s.Name = strings.TrimSpace(*p.Name)
		if r := []rune(s.Name); len(r) > 40 {
			s.Name = string(r[:40])
		}
	}
	if p.Lang != nil {
		if !i18n.Valid(*p.Lang) {
			return s, ErrInvalidSettings
		}
		s.Lang = string(i18n.Parse(*p.Lang))
	}
	if p.VoiceMode != nil {
		switch *p.VoiceMode {
		case "off", "min", "normal", "detailed":
			s.VoiceMode = *p.VoiceMode
		default:
			return s, ErrInvalidSettings
		}
	}
	if p.Figure != nil {
		if db.NormalizeSex(*p.Figure) != strings.ToLower(strings.TrimSpace(*p.Figure)) {
			return s, ErrInvalidSettings
		}
		s.Figure = db.NormalizeSex(*p.Figure)
	}
	if p.RestWarmup != nil {
		s.RestWarmup = *p.RestWarmup
	}
	if p.RestMain != nil {
		s.RestMain = *p.RestMain
	}
	if p.RestCooldown != nil {
		s.RestCooldown = *p.RestCooldown
	}
	rests := content.Rests{Warmup: s.RestWarmup, Main: s.RestMain, Cooldown: s.RestCooldown}.Normalize()
	if err := DB.UpdateSettings(userID, s.Name, s.Lang, rests.Warmup, rests.Main, rests.Cooldown, s.VoiceMode); err != nil {
		return s, err
	}
	if err := DB.SetSex(userID, s.Figure); err != nil {
		return s, err
	}
	if p.Level != nil {
		if err := DB.SetLevel(userID, *p.Level); err != nil {
			return s, err
		}
	}
	if p.Rounds != nil {
		if err := DB.SetRounds(userID, *p.Rounds); err != nil {
			return s, err
		}
	}
	return GetSettings(userID)
}
