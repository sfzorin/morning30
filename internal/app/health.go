package app

import "morning30/internal/content"

// WorkoutKind is how the phone apps label a finished workout in Apple Health
// and Health Connect: "stretch" for stretching programs, else "strength".
func WorkoutKind(userID int64) string {
	if DB == nil {
		return "strength"
	}
	if row, ok := findSelected(userID, DB.CurrentSelected(userID)); ok {
		if p, ok := content.ReadyByKey(row.Source); ok && p.Category == content.CatStretch {
			return "stretch"
		}
	}
	return "strength"
}
