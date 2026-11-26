package content

import (
	"encoding/json"

	"morning30/internal/langs"
	"regexp"
	"strings"
	"testing"
)

// bundledLoc mirrors the per-language content of a bundled document.
type bundledLoc struct {
	Name      string   `json:"name"`
	Hint      string   `json:"hint"`
	Warning   string   `json:"warning"`
	Desc      string   `json:"desc"`
	HowTo     []string `json:"howTo"`
	Correct   []string `json:"correct"`
	Wrong     []string `json:"wrong"`
	Breathing string   `json:"breathing"`
	Safety    string   `json:"safety"`
	Voice     struct {
		Start  string   `json:"start"`
		Mid    []string `json:"mid"`
		Last   string   `json:"last"`
		Finish string   `json:"finish"`
	} `json:"voice"`
}

var bundleLangs = langs.Codes()

// Mistakes are spoken as instructions, so every language phrases them as a
// prohibition (the same convention the existing exercises follow).
var negative = map[string]*regexp.Regexp{
	"ru":      regexp.MustCompile(`^Не `),
	"en":      regexp.MustCompile(`^(Don't |Don’t |Do not |Never )`),
	"tr":      regexp.MustCompile(`(ma|me|mayın|meyin|masın|mesin)[.!]?$`),
	"de":      regexp.MustCompile(`(?i)\b(nicht|kein\w*)\b`),
	"es":      regexp.MustCompile(`^No `),
	"fr":      regexp.MustCompile(`^(Ne |N['’])`),
	"it":      regexp.MustCompile(`^Non `),
	"pt":      regexp.MustCompile(`^Não `),
	"pl":      regexp.MustCompile(`^Nie `),
	"nl":      regexp.MustCompile(`(?i)\b(niet|geen|nooit)\b`),
	"sv":      regexp.MustCompile(`(?i)\b(inte|inga|ingen|aldrig)\b`),
	"da":      regexp.MustCompile(`(?i)\b(ikke|ingen|aldrig)\b`),
	"nb":      regexp.MustCompile(`(?i)\b(ikke|ingen|aldri)\b`),
	"fi":      regexp.MustCompile(`^(Älä|Älkää)\s`),
	"cs":      regexp.MustCompile(`^Ne`),
	"ro":      regexp.MustCompile(`^Nu[ -]`),
	"hu":      regexp.MustCompile(`^Ne `),
	"el":      regexp.MustCompile(`^(Μην|Μη) `),
	"ja":      regexp.MustCompile(`(ないで|ないように|ません|しない|ないこと)(ください|ね)?[。．.！]?$`),
	"ko":      regexp.MustCompile(`(마세요|마십시오|말아요)[.。]?$`),
	"zh":      regexp.MustCompile(`^(不要|别|勿|切勿|避免|不)`),
	"zh-hant": regexp.MustCompile(`^(不要|別|勿|切勿|避免|不)`),
	"id":      regexp.MustCompile(`^Jangan `),
	"vi":      regexp.MustCompile(`^(Đừng|Không) `),
	"th":      regexp.MustCompile(`^(อย่า|ห้าม)`),
	"ar":      regexp.MustCompile(`^لا `),
	"he":      regexp.MustCompile(`^(לא|אל) `),
	"fa":      regexp.MustCompile(`(^|\s)ن\S*[.؛]?$`),
}

func TestBundledExercisesAreComplete(t *testing.T) {
	if len(bundledDocs) == 0 {
		t.Fatal("no bundled exercises")
	}
	for id, raw := range bundledDocs {
		var doc struct {
			Content map[string]bundledLoc `json:"content"`
		}
		if err := json.Unmarshal(raw, &doc); err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		def := bundledDefs[id]
		if def.Difficulty < 1 || def.Difficulty > 5 {
			t.Errorf("%s: difficulty %d", id, def.Difficulty)
		}
		if _, ok := Get(id); !ok {
			t.Errorf("%s: not in the library", id)
		}
		for _, l := range bundleLangs {
			c, ok := doc.Content[l]
			if !ok {
				t.Errorf("%s: no %s content", id, l)
				continue
			}
			if strings.TrimSpace(c.Name) == "" || strings.TrimSpace(c.Hint) == "" || strings.TrimSpace(c.Desc) == "" ||
				strings.TrimSpace(c.Breathing) == "" || strings.TrimSpace(c.Safety) == "" {
				t.Errorf("%s %s: name, hint, desc, breathing and safety are required", id, l)
			}
			if len(c.HowTo) < 3 || len(c.Correct) < 2 || len(c.Wrong) < 2 {
				t.Errorf("%s %s: howTo %d, correct %d, wrong %d", id, l, len(c.HowTo), len(c.Correct), len(c.Wrong))
			}
			if c.Voice.Start == "" || c.Voice.Finish == "" {
				t.Errorf("%s %s: incomplete voice cues", id, l)
			}
			for _, w := range c.Wrong {
				if !negative[l].MatchString(strings.TrimSpace(w)) {
					t.Errorf("%s %s: mistake must be a prohibition: %q", id, l, w)
				}
			}
		}
	}
}
