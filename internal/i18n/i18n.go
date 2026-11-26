// Package i18n holds UI strings, exercise names and calendar words for every
// supported language. Each language lives in its own file, locales/<code>.json;
// anything a language lacks falls back to English.
package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"morning30/internal/langs"
)

// Lang is one of our language codes (see package langs).
type Lang string

const (
	RU Lang = "ru"
	EN Lang = "en"
	TR Lang = "tr"
	DE Lang = "de"
	ES Lang = "es"
	FR Lang = "fr"
	IT Lang = "it"
)

// Default language when nothing else is known.
const Default = EN

// Meta describes a language for the picker.
type Meta struct {
	Code Lang
	Name string // endonym
	// BCP47 is the tag passed to the browser SpeechSynthesis / lang attribute.
	BCP47 string
	RTL   bool
}

// Languages is the ordered list shown in the UI; Order is the same as codes.
var (
	Languages []Meta
	Order     []Lang
)

// locale is the content of one locales/<code>.json file.
type locale struct {
	UI        map[string]string `json:"ui"`
	Exercises map[string]string `json:"exercises"`
	Months    []string          `json:"months"`
	Weekdays  []string          `json:"weekdays"`
}

//go:embed locales/*.json
var localeFS embed.FS

var locales = map[Lang]*locale{}

func init() {
	for _, m := range langs.All {
		l := Lang(m.Code)
		Languages = append(Languages, Meta{Code: l, Name: m.Name, BCP47: m.BCP47, RTL: m.RTL})
		Order = append(Order, l)
		loc := &locale{}
		if raw, err := localeFS.ReadFile("locales/" + m.Code + ".json"); err == nil {
			if err := json.Unmarshal(raw, loc); err != nil {
				panic(fmt.Sprintf("i18n: locales/%s.json: %v", m.Code, err))
			}
		}
		locales[l] = loc
	}
	if en := locales[EN]; len(en.UI) == 0 || len(en.Months) != 12 || len(en.Weekdays) != 7 {
		panic("i18n: the English locale is incomplete")
	}
}

func index(l Lang) int {
	for i, x := range Order {
		if x == l {
			return i
		}
	}
	return -1
}

// Valid reports whether code is a supported language.
func Valid(code string) bool {
	return index(Lang(code)) >= 0
}

// aliases maps browser and legacy tags to our codes.
var aliases = map[string]Lang{
	"no": "nb", "nn": "nb", "iw": "he", "in": "id", "pt-br": "pt", "pt-pt": "pt",
	"zh-hant": "zh-hant", "zh-tw": "zh-hant", "zh-hk": "zh-hant", "zh-mo": "zh-hant",
	"zh-hans": "zh", "zh-cn": "zh", "zh-sg": "zh", "zh-my": "zh",
}

// match maps one language tag (e.g. "zh-Hant-TW", "pt_BR", "fr") to a
// supported language.
func match(tag string) (Lang, bool) {
	tag = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(tag)), "_", "-")
	if tag == "" {
		return "", false
	}
	if index(Lang(tag)) >= 0 {
		return Lang(tag), true
	}
	if strings.HasPrefix(tag, "zh-") {
		for _, part := range strings.Split(tag, "-")[1:] {
			if l, ok := aliases["zh-"+part]; ok {
				return l, true
			}
		}
		return "zh", true
	}
	if l, ok := aliases[tag]; ok {
		return l, true
	}
	base := tag
	if i := strings.Index(base, "-"); i >= 0 {
		base = base[:i]
	}
	if l, ok := aliases[base]; ok {
		return l, true
	}
	if index(Lang(base)) >= 0 {
		return Lang(base), true
	}
	return "", false
}

// Parse normalizes a code to a supported Lang, falling back to Default.
func Parse(code string) Lang {
	if l, ok := match(code); ok {
		return l
	}
	return Default
}

// FromAcceptLanguage picks the best supported language from an Accept-Language header.
func FromAcceptLanguage(header string) Lang {
	for _, part := range strings.Split(header, ",") {
		tag := part
		if i := strings.Index(tag, ";"); i >= 0 {
			tag = tag[:i]
		}
		if l, ok := match(tag); ok {
			return l
		}
	}
	return Default
}

// BCP47 returns the speech/lang tag for a language.
func BCP47(l Lang) string {
	for _, m := range Languages {
		if m.Code == l {
			return m.BCP47
		}
	}
	return "en-US"
}

// RTL reports whether a language is written right to left.
func RTL(l Lang) bool {
	for _, m := range Languages {
		if m.Code == l {
			return m.RTL
		}
	}
	return false
}

// Dir returns the HTML dir attribute value for a language.
func Dir(l Lang) string {
	if RTL(l) {
		return "rtl"
	}
	return "ltr"
}

func lookup(m func(*locale) map[string]string, l Lang, key string) (string, bool) {
	if loc := locales[l]; loc != nil {
		if s := m(loc)[key]; s != "" {
			return s, true
		}
	}
	if s := m(locales[EN])[key]; s != "" {
		return s, true
	}
	return "", false
}

// T looks up a UI string by key, falling back to English then the key itself.
func T(l Lang, key string) string {
	if s, ok := lookup(func(x *locale) map[string]string { return x.UI }, l, key); ok {
		return s
	}
	return key
}

// Tf looks up a UI string with %s placeholders and fills them in.
func Tf(l Lang, key string, args ...any) string {
	return fmt.Sprintf(T(l, key), args...)
}

// Exercise returns the localized name for an exercise ID.
func Exercise(l Lang, id string) string {
	if s, ok := lookup(func(x *locale) map[string]string { return x.Exercises }, l, id); ok {
		return s
	}
	return id
}

// Cue returns the localized spoken cue text for a key (used by the voice engine).
func Cue(l Lang, key string) string {
	return T(l, "cue."+key)
}
