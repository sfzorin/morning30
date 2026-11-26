// Package langs is the single list of languages the app speaks. It has no
// dependencies, so both the UI strings (i18n) and the exercise content
// (content) can check their translations against it.
package langs

// Meta describes one language.
type Meta struct {
	Code  string // our code: used in URLs, the database and translation files
	Name  string // endonym shown in the language picker
	BCP47 string // tag for <html lang>, hreflang and speech synthesis
	RTL   bool   // written right to left
}

// All lists the languages in picker order. The first seven are the original
// ones; English is the fallback for anything missing.
var All = []Meta{
	{"ru", "Русский", "ru-RU", false},
	{"en", "English", "en-US", false},
	{"tr", "Türkçe", "tr-TR", false},
	{"de", "Deutsch", "de-DE", false},
	{"es", "Español", "es-ES", false},
	{"fr", "Français", "fr-FR", false},
	{"it", "Italiano", "it-IT", false},
	{"pt", "Português", "pt-BR", false},
	{"pl", "Polski", "pl-PL", false},
	{"nl", "Nederlands", "nl-NL", false},
	{"sv", "Svenska", "sv-SE", false},
	{"da", "Dansk", "da-DK", false},
	{"nb", "Norsk", "nb-NO", false},
	{"fi", "Suomi", "fi-FI", false},
	{"cs", "Čeština", "cs-CZ", false},
	{"ro", "Română", "ro-RO", false},
	{"hu", "Magyar", "hu-HU", false},
	{"el", "Ελληνικά", "el-GR", false},
	{"ja", "日本語", "ja-JP", false},
	{"ko", "한국어", "ko-KR", false},
	{"zh", "简体中文", "zh-CN", false},
	{"zh-hant", "繁體中文", "zh-TW", false},
	{"id", "Bahasa Indonesia", "id-ID", false},
	{"vi", "Tiếng Việt", "vi-VN", false},
	{"th", "ไทย", "th-TH", false},
	{"ar", "العربية", "ar", true},
	{"he", "עברית", "he-IL", true},
	{"fa", "فارسی", "fa-IR", true},
}

// Codes returns every language code in picker order.
func Codes() []string {
	out := make([]string, len(All))
	for i, m := range All {
		out[i] = m.Code
	}
	return out
}

// Get returns a language by its code.
func Get(code string) (Meta, bool) {
	for _, m := range All {
		if m.Code == code {
			return m, true
		}
	}
	return Meta{}, false
}
