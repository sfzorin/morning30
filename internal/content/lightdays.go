package content

import "sort"

// LightDays finds the noticeably easier days of a cycle from what each day
// actually takes (index 0 = day 1; any unit, e.g. the minutes shown), so an
// edited program is marked by its real load rather than by fixed day numbers.
// Days tied for the shortest of their five are all light.
//
// A day is light when it is the shortest of the five days around it (two
// either side) and at least 12% shorter than their typical (median) day.
func LightDays(est []int) map[int]bool {
	light := map[int]bool{}
	for i, sec := range est {
		// Only days with two neighbours on each side are compared; the
		// cycle's first and last two days have nothing to stand out from.
		if i < 2 || i+2 >= len(est) {
			continue
		}
		week := append([]int(nil), est[i-2:i+3]...)
		sort.Ints(week)
		if sec > week[0] {
			continue
		}
		if float64(sec) <= float64(week[2])*0.88 { // week[2] is the median of five
			light[i+1] = true
		}
	}
	return light
}
