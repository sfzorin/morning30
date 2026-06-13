// Package program serializes the workout program to JSON (for download, and
// later for custom user-uploaded programs).
package program

import (
	"encoding/json"

	"danicc/internal/content"
	"danicc/internal/i18n"
)

// VoiceLoc is the per-language voice cues for an exercise.
type VoiceLoc struct {
	Start  string   `json:"start,omitempty"`
	Mid    []string `json:"mid,omitempty"`
	Last   string   `json:"last,omitempty"`
	Finish string   `json:"finish,omitempty"`
}

// ExLoc is the localized text for one exercise in one language.
type ExLoc struct {
	Name      string   `json:"name"`
	Hint      string   `json:"hint,omitempty"`
	Warning   string   `json:"warning,omitempty"`
	Desc      string   `json:"desc,omitempty"`
	HowTo     []string `json:"howTo,omitempty"`
	Correct   []string `json:"correct,omitempty"`
	Wrong     []string `json:"wrong,omitempty"`
	Breathing string   `json:"breathing,omitempty"`
	Safety    string   `json:"safety,omitempty"`
	Voice     VoiceLoc `json:"voice,omitempty"`
}

// ProgramJSON is the full, self-contained program: structure + localized text.
type ProgramJSON struct {
	content.Snapshot
	Languages []string                    `json:"languages"`
	Content   map[string]map[string]ExLoc `json:"content"`
}

// ExportJSON returns the built-in program as pretty-printed JSON.
func ExportJSON() []byte {
	snap := content.TakeSnapshot()
	out := ProgramJSON{
		Snapshot:  snap,
		Languages: []string{"ru", "en", "tr", "de", "es", "fr", "it"},
		Content:   map[string]map[string]ExLoc{},
	}
	for _, ex := range snap.Exercises {
		byLang := map[string]ExLoc{}
		detail, hasDetail := content.GetDetail(ex.ID)
		for _, m := range i18n.Languages {
			l := m.Code
			loc := ExLoc{
				Name:    i18n.Exercise(l, ex.ID),
				Hint:    i18n.Hint(l, ex.ID),
				Warning: i18n.Warning(l, ex.ID),
			}
			// Rich content + voice currently exist in Russian; exported as-is.
			if l == i18n.RU && hasDetail {
				loc.Desc = detail.Desc
				loc.HowTo = detail.HowTo
				loc.Correct = detail.Correct
				loc.Wrong = detail.Wrong
				loc.Breathing = detail.Breathing
				loc.Safety = detail.Safety
				loc.Voice = VoiceLoc{
					Start:  detail.Voice.Start,
					Mid:    detail.Voice.Mid,
					Last:   detail.Voice.Last,
					Finish: detail.Voice.Finish,
				}
			}
			byLang[string(l)] = loc
		}
		out.Content[ex.ID] = byLang
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return []byte("{}")
	}
	return b
}
