package i18n

import (
	"fmt"
	"strings"
)

// Plural forms. A counted template such as "fmt.exercises" ("%d exercises")
// may have variants per CLDR plural category: "fmt.exercises|one",
// "fmt.exercises|few", "fmt.exercises|many", and for Arabic also "|zero" and
// "|two". Tn picks the variant for the number and falls back to the plain key.
// A variant may leave the number out ("تمرينان" — two exercises).

// pluralCategory returns the CLDR plural category of a whole number.
func pluralCategory(l Lang, n int) string {
	if n < 0 {
		n = -n
	}
	mod10, mod100 := n%10, n%100
	switch l {
	case "ru":
		switch {
		case mod10 == 1 && mod100 != 11:
			return "one"
		case mod10 >= 2 && mod10 <= 4 && (mod100 < 12 || mod100 > 14):
			return "few"
		}
		return "many"
	case "pl":
		switch {
		case n == 1:
			return "one"
		case mod10 >= 2 && mod10 <= 4 && (mod100 < 12 || mod100 > 14):
			return "few"
		}
		return "many"
	case "cs":
		switch {
		case n == 1:
			return "one"
		case n >= 2 && n <= 4:
			return "few"
		}
		return "other"
	case "ar":
		switch {
		case n == 0:
			return "zero"
		case n == 1:
			return "one"
		case n == 2:
			return "two"
		case mod100 >= 3 && mod100 <= 10:
			return "few"
		case mod100 >= 11:
			return "many"
		}
		return "other"
	}
	if n == 1 {
		return "one"
	}
	return "other"
}

// Tn formats a counted template for n in language l.
func Tn(l Lang, key string, n int) string {
	tmpl := T(l, key)
	if loc := locales[l]; loc != nil {
		if s := loc.UI[key+"|"+pluralCategory(l, n)]; s != "" {
			tmpl = s
		}
	}
	if !strings.Contains(tmpl, "%") {
		return tmpl
	}
	return fmt.Sprintf(tmpl, n)
}

// pluralVariant reports whether key is a plural variant ("base|category") of
// an English key; the completeness test accepts such keys.
func pluralVariant(key string) (base string, ok bool) {
	i := strings.LastIndex(key, "|")
	if i < 0 {
		return "", false
	}
	switch key[i+1:] {
	case "zero", "one", "two", "few", "many", "other":
		return key[:i], true
	}
	return "", false
}
