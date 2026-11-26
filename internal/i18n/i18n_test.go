package i18n

import (
	"regexp"
	"sort"
	"testing"

	"morning30/internal/content"
	"morning30/internal/langs"
)

var placeholder = regexp.MustCompile(`%[0-9.]*[sdvq]`)

func placeholders(s string) string {
	ps := placeholder.FindAllString(s, -1)
	sort.Strings(ps)
	out := ""
	for _, p := range ps {
		out += p
	}
	return out
}

// Every language has its file and every UI string English has, with the same
// placeholders, plus month and weekday names.
func TestEveryLanguageIsComplete(t *testing.T) {
	en := locales[EN]
	for _, m := range langs.All {
		l := Lang(m.Code)
		if _, err := localeFS.ReadFile("locales/" + m.Code + ".json"); err != nil {
			t.Errorf("%s: no locale file", l)
			continue
		}
		loc := locales[l]
		for key, want := range en.UI {
			got := loc.UI[key]
			if got == "" {
				t.Errorf("%s: missing ui %q", l, key)
				continue
			}
			if placeholders(got) != placeholders(want) {
				t.Errorf("%s: ui %q placeholders %q, English has %q", l, key, placeholders(got), placeholders(want))
			}
		}
		for key := range loc.UI {
			if _, ok := en.UI[key]; !ok {
				t.Errorf("%s: ui %q is not in English", l, key)
			}
		}
		if len(loc.Months) != 12 || len(loc.Weekdays) != 7 {
			t.Errorf("%s: %d months, %d weekdays", l, len(loc.Months), len(loc.Weekdays))
		}
	}
}

// Every built-in exercise has a name in every language.
func TestExerciseNamesCoverPool(t *testing.T) {
	for id := range content.Pool {
		if content.IsBundled(id) {
			continue // bundled documents carry their own names
		}
		for _, l := range Order {
			if locales[l].Exercises[id] == "" {
				t.Errorf("exercise %q missing name in %s", id, l)
			}
		}
	}
}

func TestParseAndFallback(t *testing.T) {
	for in, want := range map[string]Lang{
		"ru-RU": RU, "xx": Default, "pt-BR": "pt", "pt_PT": "pt", "zh-CN": "zh", "zh-Hans-CN": "zh", "zh": "zh",
		"zh-TW": "zh-hant", "zh-Hant": "zh-hant", "zh-HK": "zh-hant", "zh-hant": "zh-hant", "no": "nb", "nb-NO": "nb",
		"iw": "he", "in": "id", "ar-EG": "ar", "": Default,
	} {
		if got := Parse(in); got != want {
			t.Errorf("Parse(%q) = %q, want %q", in, got, want)
		}
	}
	if FromAcceptLanguage("fr-FR,fr;q=0.9,en;q=0.8") != FR {
		t.Errorf("Accept-Language parse failed")
	}
	if FromAcceptLanguage("hi-IN,zh-TW;q=0.9") != "zh-hant" {
		t.Errorf("Accept-Language must skip unsupported languages")
	}
	if T(RU, "nonexistent.key") != "nonexistent.key" {
		t.Errorf("unknown key should return the key")
	}
	if Month(RU, 6) == "" || Weekday(EN, 0) != "Mon" {
		t.Errorf("calendar localization failed")
	}
	if Dir("ar") != "rtl" || Dir("he") != "rtl" || Dir(RU) != "ltr" {
		t.Errorf("text direction is wrong")
	}
	if len(Languages) != len(langs.All) || len(Order) != len(langs.All) {
		t.Errorf("language lists differ")
	}
}
