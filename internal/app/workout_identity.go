package app

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"morning30/internal/content"
)

// CurrentPosition applies an explicitly scheduled new cycle without changing
// the separate activity calendar. Old completion receipts retain old positions.
func CurrentPosition(userID int64) int {
	position, _, err := DB.ApplyDueCycleRestart(context.Background(), userID, time.Now(), content.TotalDays)
	if err != nil {
		log.Printf("read workout position: %v", err)
		position, _ = DB.Position(userID)
	}
	return position
}

// CycleRestartAt lets a page left open overnight revalidate a scheduled restart.
func CycleRestartAt(userID int64) string {
	at, err := DB.ScheduledCycleRestart(context.Background(), userID)
	if err != nil {
		log.Printf("read cycle restart: %v", err)
		return ""
	}
	if at.IsZero() {
		return ""
	}
	return at.UTC().Format(time.RFC3339Nano)
}

// WorkoutProgramKey invalidates saved player state when the prescribed program
// or rest settings change. Language, illustrations and saved difficulty do not
// affect it: a paused workout keeps its own chosen difficulty.
func WorkoutProgramKey(userID int64) string {
	warmup, main, cooldown := DB.Rests(userID)
	// Timing rules are part of the checkpoint contract even if program JSON
	// is unchanged. A combined timer saved before side-change preparation
	// cannot safely be interpreted as one segment of the new timed exercise.
	data := fmt.Sprintf("%s\n%d,%d,%d\nside-change:%d", UserProgram(userID).Marshal(), warmup, main, cooldown, content.SideChangeSeconds)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}
