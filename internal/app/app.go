// Package app holds process-wide dependencies shared by the page/segment code.
package app

import (
	"time"

	"danicc/internal/catalog"
	"danicc/internal/content"
	"danicc/internal/db"
)

// DB is the application database, set once at startup in main.
var DB *db.DB

// Today returns the server-local calendar date as YYYY-MM-DD.
func Today() string {
	return time.Now().Format("2006-01-02")
}

// ProgramByKey returns the program for "custom" (the stored one, if any) or the
// built-in program (flattened). The built-in is always available.
func ProgramByKey(userID int64, key string) content.Resolved {
	if key == "custom" {
		if j, err := DB.GetProgramJSON(userID); err == nil && j != "" {
			if r, err := content.ParseResolved(j); err == nil && len(r.Days) > 0 {
				return r
			}
		}
	}
	return content.ResolveBuiltin()
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
	c := UserCatalog(userID)
	c.Put(d)
	return DB.SetCustomExercises(userID, c.Marshal())
}

// HasCustomProgram reports whether the user saved a custom program.
func HasCustomProgram(userID int64) bool {
	j, err := DB.GetProgramJSON(userID)
	return err == nil && j != ""
}

func activeKey(userID int64) string {
	k, _ := DB.ActiveProgram(userID)
	return k
}

// UserProgram returns the user's active program (built-in or custom).
func UserProgram(userID int64) content.Resolved {
	return ProgramByKey(userID, activeKey(userID))
}

// UserWorkout builds a day's workout from the active program (built-in or custom),
// scaled by the user's universal difficulty level.
func UserWorkout(userID int64, day, rest int) content.Workout {
	return UserProgram(userID).Workout(day, rest, DB.GetLevel(userID))
}
