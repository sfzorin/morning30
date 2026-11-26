package content

import "sort"

// LightDays finds the noticeably easier days of a cycle from what each day
// actually takes (index 0 = day 1; any unit, e.g. the minutes shown), so an
// edited program is marked by its real load rather than by fixed day numbers.
// Days tied for the shortest of their week are all light.
//
// A day is light when it is the shortest of its week (three days either
// side) and at least 12% shorter than that week's typical (median) day.
func LightDays(est []int) map[int]bool {
	light := map[int]bool{}
	for i, sec := range est {
		lo, hi := i-3, i+4
		if lo < 0 {
			lo = 0
		}
		if hi > len(est) {
			hi = len(est)
		}
		week := append([]int(nil), est[lo:hi]...)
		sort.Ints(week)
		if sec > week[0] || len(week) < 2 {
			continue
		}
		n := len(week)
		median := float64(week[n/2])
		if n%2 == 0 {
			median = float64(week[n/2-1]+week[n/2]) / 2
		}
		if float64(sec) <= median*0.88 {
			light[i+1] = true
		}
	}
	return light
}
