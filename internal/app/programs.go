package app

import (
	"errors"
	"strings"

	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
)

// Selected programs: the user's own shortlist. Each keeps its own copy and its
// own cycle day; switching never loses edits or progress. The serie (streak)
// and the activity calendar stay common to all programs.

var (
	// ErrUnknownCourse is returned when a ready-program key matches nothing.
	ErrUnknownCourse = errors.New("unknown program")
	// ErrEmptyProgramName is returned when a program is renamed to nothing.
	ErrEmptyProgramName = errors.New("program name is empty")
	// ErrHasEdits is returned when removing an edited program without consent.
	ErrHasEdits = errors.New("the program has your edits; confirm to discard them")
	// ErrUnknownExercise is returned when an exercise id matches nothing.
	ErrUnknownExercise = errors.New("unknown exercise")
)

// SelectedView is one selected program as the screens and agents see it.
type SelectedView struct {
	ID      int64  `json:"id"`
	Source  string `json:"source"` // catalog key, "" for the user's own
	Name    string `json:"name"`
	Day     int    `json:"day"`     // the next day of its own cycle, 1..30
	Current bool   `json:"current"` // the program the user runs now
	Edited  bool   `json:"edited"`  // removing it would lose the user's work
}

// EnsureSelected brings an account into the selected-programs model. It is
// idempotent and never changes the program, its day or the streak:
//   - a personal program in the active slot becomes a selected program;
//   - programs kept by the earlier "saved programs" list become selected programs.
func EnsureSelected(userID int64) {
	if DB == nil {
		return
	}
	SnapshotLegacy(userID)
	if DB.CurrentSelected(userID) == 0 {
		if raw, err := DB.GetProgramJSON(userID); err == nil && raw != "" {
			name, source := "My program", ""
			if r, err := content.ParseResolved(raw); err == nil {
				if strings.TrimSpace(r.Name) != "" {
					name = r.Name
				}
				source = content.MatchReady(r)
			}
			_, _ = DB.AdoptActive(userID, source, name, raw)
		}
	}
	_, _ = DB.TakeSavedPrograms(userID, func(raw string) string {
		r, err := content.ParseResolved(raw)
		if err != nil {
			return ""
		}
		return content.MatchReady(r)
	})
}

// SelectedPrograms lists the user's selected programs, oldest first.
func SelectedPrograms(userID int64) []SelectedView {
	EnsureSelected(userID)
	rows, err := DB.SelectedPrograms(userID)
	if err != nil {
		return nil
	}
	current := DB.CurrentSelected(userID)
	out := make([]SelectedView, 0, len(rows))
	for _, row := range rows {
		out = append(out, SelectedView{
			ID: row.ID, Source: row.Source, Name: row.Name, Day: row.Position%content.TotalDays + 1,
			Current: row.ID == current, Edited: isEdited(row),
		})
	}
	return out
}

func isEdited(row db.SelectedProgram) bool {
	if row.Source == "" {
		return true
	}
	r, err := content.ParseResolved(row.Program)
	if err != nil || !content.SameAsReady(r, row.Source) {
		return true
	}
	// A name of the user's own is their work too.
	p, _ := content.ReadyByKey(row.Source)
	for _, title := range p.Title {
		if title == row.Name {
			return false
		}
	}
	return true
}

func findSelected(userID, id int64) (db.SelectedProgram, bool) {
	rows, err := DB.SelectedPrograms(userID)
	if err != nil {
		return db.SelectedProgram{}, false
	}
	for _, row := range rows {
		if row.ID == id {
			return row, true
		}
	}
	return db.SelectedProgram{}, false
}

// adoptPlainDemo gives the plain morning workout a row before the user leaves
// it, so its cycle day is kept. A newcomer who never trained gets no row.
func adoptPlainDemo(userID int64, lang string) error {
	if DB.CurrentSelected(userID) != 0 {
		return nil
	}
	if pos, _ := DB.Position(userID); pos == 0 && !HasCustomProgram(userID) {
		return nil
	}
	demo, _ := content.ReadyByKey(content.DemoKey)
	r := demo.Build()
	r.Name = demo.TitleIn(lang)
	_, err := DB.AdoptActive(userID, content.DemoKey, r.Name, r.Marshal())
	return err
}

// SelectReady switches to a ready program and adds it to the selected ones.
// When the user already has it, their own version (and its day) comes back.
func SelectReady(userID int64, key, lang string) error {
	p, ok := content.ReadyByKey(strings.TrimSpace(key))
	if !ok {
		return ErrUnknownCourse
	}
	EnsureSelected(userID)
	// Keep the plain morning workout's day first, so picking the morning
	// workout itself finds that row instead of adding a second one.
	if err := adoptPlainDemo(userID, lang); err != nil {
		return err
	}
	id, err := starReady(userID, p, lang)
	if err != nil {
		return err
	}
	if err := DB.SwitchSelected(userID, id); err != nil {
		return err
	}
	adjustFigure(userID, p.Kids)
	return nil
}

// StarReady adds a ready program to the selected ones without switching to it
// and returns its id. A program already selected is returned as is.
func StarReady(userID int64, key, lang string) (int64, error) {
	p, ok := content.ReadyByKey(strings.TrimSpace(key))
	if !ok {
		return 0, ErrUnknownCourse
	}
	EnsureSelected(userID)
	return starReady(userID, p, lang)
}

func starReady(userID int64, p content.ReadyProgram, lang string) (int64, error) {
	rows, err := DB.SelectedPrograms(userID)
	if err != nil {
		return 0, err
	}
	for _, row := range rows {
		if row.Source == p.Key {
			return row.ID, nil
		}
	}
	r := p.Build()
	r.Name = p.TitleIn(lang)
	return DB.AddSelected(userID, p.Key, r.Name, r.Marshal(), 0)
}

// SelectByID switches to one of the user's selected programs.
func SelectByID(userID, id int64) error {
	EnsureSelected(userID)
	row, ok := findSelected(userID, id)
	if !ok {
		return db.ErrSelectedNotFound
	}
	if err := adoptPlainDemo(userID, LangOf(userID)); err != nil {
		return err
	}
	if err := DB.SwitchSelected(userID, id); err != nil {
		return err
	}
	if p, ok := content.ReadyByKey(row.Source); ok {
		adjustFigure(userID, p.Kids)
	}
	return nil
}

// Unselect removes a program from the selected ones. A program with the
// user's edits needs discardEdits, because they cannot be brought back.
// Removing the current program switches to the most recently added other
// one, or to the plain morning workout when none is left.
func Unselect(userID, id int64, discardEdits bool) error {
	EnsureSelected(userID)
	row, ok := findSelected(userID, id)
	if !ok {
		return db.ErrSelectedNotFound
	}
	if isEdited(row) && !discardEdits {
		return ErrHasEdits
	}
	if DB.CurrentSelected(userID) == id {
		rows, _ := DB.SelectedPrograms(userID)
		var next int64
		for _, other := range rows {
			if other.ID != id {
				next = other.ID
			}
		}
		var err error
		if next != 0 {
			err = DB.SwitchSelected(userID, next)
		} else {
			err = DB.ClearActive(userID, content.DefaultStandardKey())
		}
		if err != nil {
			return err
		}
	}
	return DB.DeleteSelected(userID, id)
}

// UnstarReady removes a ready program from the selected ones (the catalog's
// star). It needs discardEdits when the user changed that program.
func UnstarReady(userID int64, key string, discardEdits bool) error {
	EnsureSelected(userID)
	rows, err := DB.SelectedPrograms(userID)
	if err != nil {
		return err
	}
	for _, row := range rows {
		if row.Source == key {
			return Unselect(userID, row.ID, discardEdits)
		}
	}
	return db.ErrSelectedNotFound
}

// ImportProgram adds a program from JSON as a new selected program and
// switches to it at day 1. Nothing the user had is lost.
func ImportProgram(userID int64, raw string) error {
	r, err := content.ParseResolved(strings.TrimSpace(raw))
	if err != nil || len(r.Days) == 0 {
		return ErrInvalidProgram
	}
	if strings.TrimSpace(r.Name) == "" {
		r.Name = "My program"
	}
	r.Name = CleanProgramName(r.Name)
	EnsureSelected(userID)
	if err := adoptPlainDemo(userID, LangOf(userID)); err != nil {
		return err
	}
	id, err := DB.AddSelected(userID, content.MatchReady(r), r.Name, r.Marshal(), 0)
	if err != nil {
		return err
	}
	return DB.SwitchSelected(userID, id)
}

// SetUserProgram saves a new version of the current program from JSON (the
// editor's "save", an agent's update). The cycle day and streak stay.
func SetUserProgram(userID int64, raw string) error {
	r, err := content.ParseResolved(strings.TrimSpace(raw))
	if err != nil || len(r.Days) == 0 {
		return ErrInvalidProgram
	}
	return SaveCurrent(userID, r)
}

// SaveCurrent stores a changed current program. The plain morning workout
// becomes a selected program the first time it is changed.
func SaveCurrent(userID int64, r content.Resolved) error {
	if strings.TrimSpace(r.Name) == "" {
		r.Name = "My program"
	}
	r.Name = CleanProgramName(r.Name)
	EnsureSelected(userID)
	if DB.CurrentSelected(userID) == 0 {
		if _, err := DB.AdoptActive(userID, content.DefaultStandardKey(), r.Name, r.Marshal()); err != nil {
			return err
		}
	}
	return DB.SaveActive(userID, r.Name, r.Marshal())
}

// EditCurrent applies a change to the current program and saves it.
func EditCurrent(userID int64, change func(content.Resolved) (content.Resolved, error)) error {
	r := UserProgram(userID)
	if !HasCustomProgram(userID) {
		r.Name = ProgramTitle(userID, LangOf(userID))
	}
	out, err := change(r)
	if err != nil {
		return err
	}
	return SaveCurrent(userID, out)
}

// exerciseDef resolves an exercise the user can use (built-in or custom).
func exerciseDef(userID int64, id string) (content.Exercise, error) {
	def, ok := UserCatalog(userID).Def(strings.TrimSpace(id))
	if !ok {
		return content.Exercise{}, ErrUnknownExercise
	}
	return def, nil
}

// AddToCurrent adds an exercise to a part of the current program.
func AddToCurrent(userID int64, exerciseID string, part content.Slot, days content.Days, value int) error {
	def, err := exerciseDef(userID, exerciseID)
	if err != nil {
		return err
	}
	return EditCurrent(userID, func(r content.Resolved) (content.Resolved, error) {
		return content.AddExercise(r, def, part, days, value)
	})
}

// ReplaceInCurrent swaps one exercise for another in the current program.
func ReplaceInCurrent(userID int64, from, to string, days content.Days) (int, error) {
	def, err := exerciseDef(userID, to)
	if err != nil {
		return 0, err
	}
	n := 0
	err = EditCurrent(userID, func(r content.Resolved) (content.Resolved, error) {
		out, changed, err := content.ReplaceExercise(r, strings.TrimSpace(from), def, days)
		n = changed
		return out, err
	})
	return n, err
}

// RemoveFromCurrent takes an exercise out of the current program.
func RemoveFromCurrent(userID int64, id string, part content.Slot, days content.Days) (int, error) {
	n := 0
	err := EditCurrent(userID, func(r content.Resolved) (content.Resolved, error) {
		out, removed, err := content.RemoveExercise(r, strings.TrimSpace(id), part, days)
		n = removed
		return out, err
	})
	return n, err
}

// CurrentDay is the next day of the current program's cycle (1..30).
func CurrentDay(userID int64) int { return CurrentPosition(userID)%content.TotalDays + 1 }

// CurrentSource is the catalog key the current program came from ("" for the
// user's own; the morning workout for a newcomer).
func CurrentSource(userID int64) string {
	id := DB.CurrentSelected(userID)
	if id == 0 {
		if HasCustomProgram(userID) {
			return ""
		}
		return content.DemoKey
	}
	if row, ok := findSelected(userID, id); ok {
		return row.Source
	}
	return ""
}

// ProgramNameMax is the longest program name, in characters.
const ProgramNameMax = 40

// CleanProgramName trims a user-typed program name and clips it on whole
// characters. An empty result means "no name given".
func CleanProgramName(name string) string {
	name = strings.Join(strings.Fields(name), " ")
	if runes := []rune(name); len(runes) > ProgramNameMax {
		name = strings.TrimSpace(string(runes[:ProgramNameMax]))
	}
	return name
}

// ProgramTitle is the name of the current program: its own name, or the
// morning workout's title for a newcomer.
func ProgramTitle(userID int64, lang string) string {
	if HasCustomProgram(userID) {
		if name := strings.TrimSpace(UserProgram(userID).Name); name != "" {
			return name
		}
	}
	demo, _ := content.ReadyByKey(content.DemoKey)
	return demo.TitleIn(lang)
}

// RenameProgram gives the current program a name of the user's choosing.
// Exercises, doses, cycle day and streak are untouched.
func RenameProgram(userID int64, name string) error {
	name = CleanProgramName(name)
	if name == "" {
		return ErrEmptyProgramName
	}
	r := UserProgram(userID)
	r.Name = name
	return SaveCurrent(userID, r)
}

// UseCourse switches to a ready program (kept for callers of the course API).
func UseCourse(userID int64, key, lang string) error { return SelectReady(userID, key, lang) }

// RestoreDemo switches to the morning workout. Every selected program stays.
func RestoreDemo(userID int64) error {
	return SelectReady(userID, content.DemoKey, LangOf(userID))
}

// LangOf returns the user's interface language.
func LangOf(userID int64) string {
	if u, err := DB.UserByID(userID); err == nil && u.Lang != "" {
		return u.Lang
	}
	return string(i18n.EN)
}

// SetDoseInCurrent changes an exercise's dose on the chosen days of the
// current program: value sets it, percent scales it.
func SetDoseInCurrent(userID int64, id string, days content.Days, value, percent int) (int, error) {
	n := 0
	err := EditCurrent(userID, func(r content.Resolved) (content.Resolved, error) {
		out, changed, err := content.SetDose(r, strings.TrimSpace(id), days, value, percent)
		n = changed
		return out, err
	})
	return n, err
}
