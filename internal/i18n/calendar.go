package i18n

// calendarWord returns item i of a calendar list in language l, or English.
func calendarWord(l Lang, list func(*locale) []string, i int) string {
	if loc := locales[l]; loc != nil {
		if xs := list(loc); i < len(xs) && xs[i] != "" {
			return xs[i]
		}
	}
	if xs := list(locales[EN]); i < len(xs) {
		return xs[i]
	}
	return ""
}

// Month returns the localized name of month m (1-12).
func Month(l Lang, m int) string {
	if m < 1 || m > 12 {
		return ""
	}
	return calendarWord(l, func(x *locale) []string { return x.Months }, m-1)
}

// Weekday returns the localized short name for a Monday-first index (0-6).
func Weekday(l Lang, idx int) string {
	if idx < 0 || idx > 6 {
		return ""
	}
	return calendarWord(l, func(x *locale) []string { return x.Weekdays }, idx)
}
