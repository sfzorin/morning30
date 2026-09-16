package i18n

import (
	"morning30/internal/content"
	"testing"
)

// Every UI string must at least have an English value (the fallback column).
func TestUIStringsHaveEnglish(t *testing.T) {
	enCol := index(EN)
	for key, row := range ui {
		if row[enCol] == "" {
			t.Errorf("ui key %q has no English value", key)
		}
	}
}

// Every exercise in the pool must have a localized name in all 7 languages.
func TestExerciseNamesCoverPool(t *testing.T) {
	for id := range content.Pool {
		row, ok := exerciseNames[id]
		if !ok {
			t.Errorf("exercise %q has no translations", id)
			continue
		}
		for i, name := range row {
			if name == "" {
				t.Errorf("exercise %q missing translation for %s", id, Order[i])
			}
		}
	}
}

func TestParseAndFallback(t *testing.T) {
	if Parse("ru-RU") != RU {
		t.Errorf("Parse(ru-RU) != RU")
	}
	if Parse("xx") != Default {
		t.Errorf("Parse(xx) should fall back to default")
	}
	if FromAcceptLanguage("fr-FR,fr;q=0.9,en;q=0.8") != FR {
		t.Errorf("Accept-Language parse failed")
	}
	if T(RU, "nonexistent.key") != "nonexistent.key" {
		t.Errorf("unknown key should return the key")
	}
	if Month(RU, 6) == "" || Weekday(EN, 0) != "Mon" {
		t.Errorf("calendar localization failed")
	}
}

func TestRestJokes(t *testing.T) {
	if n := len(restJokes); n < 100 {
		t.Errorf("want at least 100 rest jokes, got %d", n)
	}
	seenRU, seenEN := map[string]bool{}, map[string]bool{}
	for i, row := range restJokes {
		if row[0] == "" || row[1] == "" {
			t.Errorf("joke %d missing ru or en", i)
		}
		if seenRU[row[0]] {
			t.Errorf("duplicate Russian joke: %q", row[0])
		}
		seenRU[row[0]] = true
		if seenEN[row[1]] {
			t.Errorf("duplicate English joke: %q", row[1])
		}
		seenEN[row[1]] = true
	}
	ru, en := RestJokes(RU), RestJokes(DE)
	if ru[0] == en[0] {
		t.Errorf("Russian jokes should not fall back to English")
	}
	if en[0] != RestJokes(EN)[0] {
		t.Errorf("non-Russian languages should use English jokes")
	}
	sample := RestJokesSample(RU, 12)
	if len(sample) != 12 {
		t.Errorf("sample size = %d, want 12", len(sample))
	}
	seen := map[string]bool{}
	for _, s := range sample {
		if seen[s] {
			t.Errorf("sample repeated %q", s)
		}
		seen[s] = true
	}
}
