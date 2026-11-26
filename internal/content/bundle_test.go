package content

import (
	"encoding/json"
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

var bundleLangs = []string{"ru", "en", "tr", "de", "es", "fr", "it"}

// Mistakes are spoken as instructions, so every language phrases them as a
// prohibition (the same convention the existing exercises follow).
var negative = map[string]*regexp.Regexp{
	"ru": regexp.MustCompile(`^Не `),
	"en": regexp.MustCompile(`^(Don't |Don’t |Do not |Never )`),
	"tr": regexp.MustCompile(`(ma|me|mayın|meyin|masın|mesin)[.!]?$`),
	"de": regexp.MustCompile(`(?i)\b(nicht|kein\w*)\b`),
	"es": regexp.MustCompile(`^No `),
	"fr": regexp.MustCompile(`^(Ne |N['’])`),
	"it": regexp.MustCompile(`^Non `),
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
