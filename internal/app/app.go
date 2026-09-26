// Package app holds process-wide dependencies shared by the page/segment code.
package app

import (
	"errors"
	"strings"
	"time"

	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
)

var (
	// ErrInvalidProgram is returned when uploaded program JSON is not a usable cycle.
	ErrInvalidProgram = errors.New("invalid program")
	// ErrInvalidExercise is returned when a custom exercise document is unusable.
	ErrInvalidExercise = errors.New("invalid exercise")
)

// DB is the application database, set once at startup in main.
var DB *db.DB

// Today returns the server-local calendar date as YYYY-MM-DD.
func Today() string {
	return time.Now().Format("2006-01-02")
}

// ProgramByKey returns the program for "custom" (the stored one, if any) or a
// standard program by key, falling back to the default standard program.
func ProgramByKey(userID int64, key string) content.Resolved {
	if key == "custom" {
		if j, err := DB.GetProgramJSON(userID); err == nil && j != "" {
			if r, err := content.ParseResolved(j); err == nil && len(r.Days) > 0 {
				return r
			}
		}
	}
	if r, ok := content.StandardByKey(key); ok {
		return r
	}
	r, _ := content.StandardByKey(content.DefaultStandardKey())
	return r
}

// SeedBuiltins makes the global `exercises` table match the built-in library
// (authoritative upsert + prune), so it is the canonical base every user reads.
func SeedBuiltins() {
	ids := make([]string, 0, len(content.Library()))
	for _, ex := range content.Library() {
		if d, ok := catalog.Builtin(ex.ID); ok {
			_ = DB.SetExercise(ex.ID, d.Compact()) // upsert (code is canonical)
			ids = append(ids, ex.ID)
		}
	}
	_ = DB.PruneExercises(ids) // drop exercises removed from the library
}

// UserCatalog returns the user's exercise catalog: the global DB library overlaid
// with their custom exercises. Pages use it to resolve names/content/voice so
// custom exercises are selectable, shown and spoken.
func UserCatalog(userID int64) catalog.Catalog {
	base, _ := DB.AllExercises()
	j, _ := DB.GetCustomExercises(userID)
	return catalog.NewLayered(base, j)
}

// SaveCustomExercise adds or overrides a custom exercise in the user's library.
func SaveCustomExercise(userID int64, d catalog.Doc) error {
	d.ID = strings.TrimSpace(d.ID)
	if d.Slot == "" {
		d.Slot = string(content.Main)
	}
	if d.Unit == "" {
		d.Unit = string(content.Reps)
	}
	if !catalog.ValidDoc(d) {
		return ErrInvalidExercise
	}
	d.SVG = catalog.NormalizeFigure(d.SVG, d.ID)
	c := UserCatalog(userID)
	c.Put(d)
	return DB.SetCustomExercises(userID, c.Marshal())
}

// SnapshotAll copies leftover named built-in plans into personal files for every
// existing account. Idempotent; used at process start and by MCP.
func SnapshotAll() {
	if DB == nil {
		return
	}
	accs, err := DB.ListAccounts()
	if err != nil {
		return
	}
	for _, a := range accs {
		SnapshotLegacy(a.ID)
	}
}

// HasCustomProgram reports whether the user saved a custom program.
func HasCustomProgram(userID int64) bool {
	j, err := DB.GetProgramJSON(userID)
	return err == nil && j != ""
}

// SnapshotLegacy copies a previously named built-in plan (Sergey / Vlad / Ilya)
// into the user's personal file. Registered accounts still on the old default
// (`builtin`) get the Sergey cycle. Guests and new demo accounts are left on
// the live demo. Cycle position, streaks and activity are never touched.
func SnapshotLegacy(userID int64) {
	if DB == nil || HasCustomProgram(userID) {
		return
	}
	k, _ := DB.ActiveProgram(userID)
	u, err := DB.UserByID(userID)
	if err != nil {
		return
	}
	legacyKey := k
	switch k {
	case "ilya", "vlad", "sergey":
		// keep
	case "builtin", "":
		if u.IsGuest {
			return
		}
		legacyKey = "sergey"
	default:
		return
	}
	r, ok := content.LegacyNamed(legacyKey)
	if !ok {
		return
	}
	_ = DB.SetProgramJSON(userID, r.Marshal())
	_ = DB.SetActiveProgram(userID, "custom")
}

// SetUserProgram stores a personal cycle from JSON and makes it active.
// It does not reset cycle position, streaks, or activity.
func SetUserProgram(userID int64, raw string) error {
	r, err := content.ParseResolved(strings.TrimSpace(raw))
	if err != nil || len(r.Days) == 0 {
		return ErrInvalidProgram
	}
	if strings.TrimSpace(r.Name) == "" {
		r.Name = "My program"
	}
	if err := DB.SetProgramJSON(userID, r.Marshal()); err != nil {
		return err
	}
	return DB.SetActiveProgram(userID, "custom")
}

// RestoreDemo clears the personal file so the user runs the public demo.
// Progress is left as-is.
func RestoreDemo(userID int64) error {
	if err := DB.SetProgramJSON(userID, ""); err != nil {
		return err
	}
	return DB.SetActiveProgram(userID, content.DefaultStandardKey())
}

// ActiveProgramKey returns the user's selected program key: "custom" when a saved
// custom program is active, otherwise a valid standard program key (the default
// when the stored value is empty or a legacy/unknown key).
func ActiveProgramKey(userID int64) string {
	SnapshotLegacy(userID)
	k, _ := DB.ActiveProgram(userID)
	if k == "custom" && HasCustomProgram(userID) {
		return "custom"
	}
	if _, ok := content.StandardByKey(k); ok {
		return k
	}
	return content.DefaultStandardKey()
}

// UserProgram returns the user's active program (a standard set or custom).
func UserProgram(userID int64) content.Resolved {
	return ProgramByKey(userID, ActiveProgramKey(userID))
}

// UserWorkout builds a day's workout from the active program (built-in or custom),
// scaled by the user's universal difficulty level, with the user's per-phase rests.
func UserWorkout(userID int64, day int) content.Workout {
	w, m, c := DB.Rests(userID)
	rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
	return UserProgram(userID).Workout(day, DB.GetLevel(userID), rests)
}

// UserWorkoutBase builds the day's workout at level 0 (base, unscaled values) and
// returns the user's difficulty level separately, so the client can scale main
// exercises on the fly (the in-workout difficulty nudge).
func UserWorkoutBase(userID int64, day int) (content.Workout, int) {
	w, m, c := DB.Rests(userID)
	rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
	return UserProgram(userID).Workout(day, 0, rests), DB.GetLevel(userID)
}
