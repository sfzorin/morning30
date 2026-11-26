package app

import "morning30/internal/content"

// CycleLoad estimates how long each of the 30 days takes for this user (index
// 0 = day 1), with their level, rounds and rests, reading the settings once.
func CycleLoad(userID int64) []int {
	w, m, c := DB.Rests(userID)
	rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
	level, rounds := DB.GetLevel(userID), DB.GetRounds(userID)
	prog := UserProgram(userID)
	est := make([]int, content.TotalDays)
	for d := 1; d <= content.TotalDays; d++ {
		est[d-1] = prog.WorkoutRounds(d, level, rounds, rests).EstSec
	}
	return est
}
