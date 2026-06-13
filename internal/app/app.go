// Package app holds process-wide dependencies shared by the page/segment code.
package app

import (
	"time"

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

// UserWorkout builds a day's workout from the active program. The built-in path
// keeps full block-variant adaptation; a custom program applies substitutions only.
func UserWorkout(userID int64, day, rest int, a content.Adapt) content.Workout {
	if activeKey(userID) == "custom" {
		if j, err := DB.GetProgramJSON(userID); err == nil && j != "" {
			if r, err := content.ParseResolved(j); err == nil && len(r.Days) > 0 {
				return r.Workout(day, rest, a)
			}
		}
	}
	return content.BuildAdapted(day, rest, a)
}
