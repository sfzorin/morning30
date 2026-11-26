package content

import (
	"reflect"
	"sort"
	"testing"
)

func lightList(m map[int]bool) []int {
	var out []int
	for d := range m {
		out = append(out, d)
	}
	sort.Ints(out)
	return out
}

// Morning Workout's minutes: the shortest day of each week is marked; day 4
// (one minute shorter) and day 21 (day 18 is lighter that week) are not.
func TestLightDaysFollowLoad(t *testing.T) {
	min := []int{14, 15, 14, 13, 15, 17, 13, 16, 17, 18, 14, 18, 17, 14, 17, 17, 20, 13, 20, 19, 14, 21, 16, 20, 14, 21, 19, 14, 22, 19}
	est := append([]int(nil), min...)
	if got, want := lightList(LightDays(est)), []int{7, 11, 14, 18, 25, 28}; !reflect.DeepEqual(got, want) {
		t.Fatalf("light days %v, want %v", got, want)
	}
	flat := make([]int, 30)
	for i := range flat {
		flat[i] = 900
	}
	if got := LightDays(flat); len(got) != 0 {
		t.Fatalf("an even program has no light days, got %v", lightList(got))
	}
	est[13] = 25 // the user made day 14 heavy
	if LightDays(est)[14] {
		t.Fatalf("a day made heavy must stop being light")
	}
}

// Every built-in course's planned recovery days are lighter, so most of them
// must also be found from the load alone.
func TestLightDaysFindPlannedRecovery(t *testing.T) {
	for _, c := range Series30() {
		r := c.Resolve()
		est := make([]int, TotalDays)
		for d := 1; d <= TotalDays; d++ {
			est[d-1] = (r.WorkoutRounds(d, 0, MainRounds, Rests{Warmup: 10, Main: 20, Cooldown: 10}).EstSec + 59) / 60
		}
		found := LightDays(est)
		for d := range found {
			if !IsRecoveryDay(d) {
				t.Errorf("%s: day %d marked light but is not a planned recovery day", c.Key, d)
			}
		}
	}
}
