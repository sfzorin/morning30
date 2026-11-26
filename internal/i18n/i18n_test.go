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
		for key, got := range loc.UI {
			if _, ok := en.UI[key]; ok {
				continue
			}
			base, ok := pluralVariant(key)
			if _, known := en.UI[base]; !ok || !known {
				t.Errorf("%s: ui %q is not in English", l, key)
				continue
			}
			// A plural form keeps the number or, like Arabic "two", spells it out.
			if p := placeholders(got); p != "" && p != placeholders(en.UI[base]) {
				t.Errorf("%s: plural %q has placeholders %q", l, key, p)
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

func TestPluralForms(t *testing.T) {
	for _, c := range []struct {
		lang Lang
		key  string
		n    int
		want string
	}{
		{"ru", "fmt.exercises", 1, "1 упражнение"}, {"ru", "fmt.exercises", 3, "3 упражнения"},
		{"ru", "fmt.exercises", 12, "12 упражнений"}, {"ru", "fmt.exercises", 21, "21 упражнение"},
		{"ru", "fmt.rounds", 5, "5 кругов"}, {"ru", "fmt.rounds", 2, "2 круга"},
		{"pl", "fmt.exercises", 1, "1 ćwiczenie"}, {"pl", "fmt.exercises", 22, "22 ćwiczenia"},
		{"pl", "fmt.exercises", 25, "25 ćwiczeń"}, {"pl", "fmt.exercises", 13, "13 ćwiczeń"},
		{"pl", "fmt.days", 1, "1 dzień"}, {"pl", "fmt.rounds", 3, "3 rundy"},
		{"cs", "fmt.exercises", 1, "1 cvik"}, {"cs", "fmt.exercises", 4, "4 cviky"},
		{"cs", "fmt.exercises", 25, "25 cviků"}, {"cs", "fmt.days", 30, "30 dní"},
		{"ar", "fmt.exercises", 1, "تمرين واحد"}, {"ar", "fmt.exercises", 2, "تمرينان"},
		{"ar", "fmt.exercises", 5, "5 تمارين"}, {"ar", "fmt.exercises", 25, "25 تمرينًا"},
		{"ar", "fmt.days", 30, "30 يومًا"}, {"ar", "fmt.rounds", 10, "10 جولات"},
		{"en", "fmt.exercises", 25, "25 exercises"},
	} {
		if got := Tn(c.lang, c.key, c.n); got != c.want {
			t.Errorf("Tn(%s, %s, %d) = %q, want %q", c.lang, c.key, c.n, got, c.want)
		}
	}
}
