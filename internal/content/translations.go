package content

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

// Translations added after the first seven languages live in i18n/, one set of
// files per language:
//
//	i18n/courses.<lang>.json   program key → {title, about}
//	i18n/details.<lang>.json   built-in exercise id → card and voice text
//	i18n/exercises.<lang>.json bundled exercise id → the document's per-language part
//
// They are merged into the same structures the original languages use, so the
// rest of the app needs no special case.

//go:embed i18n/*.json
var translationFS embed.FS

// translationFiles returns lang → raw file for one kind ("courses", ...).
func translationFiles(kind string) map[string][]byte {
	out := map[string][]byte{}
	names, err := fs.Glob(translationFS, "i18n/"+kind+".*.json")
	if err != nil {
		panic(err)
	}
	for _, name := range names {
		lang := strings.TrimSuffix(strings.TrimPrefix(path.Base(name), kind+"."), ".json")
		raw, err := translationFS.ReadFile(name)
		if err != nil {
			panic(err)
		}
		out[lang] = raw
	}
	return out
}

// detailDoc is the JSON form of Detail.
type detailDoc struct {
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

// loadDetailTranslations adds the card text of built-in exercises.
func loadDetailTranslations() {
	for lang, raw := range translationFiles("details") {
		var docs map[string]detailDoc
		if err := json.Unmarshal(raw, &docs); err != nil {
			panic(fmt.Sprintf("content: details.%s.json: %v", lang, err))
		}
		m := map[string]Detail{}
		for id, d := range docs {
			m[id] = Detail{Desc: d.Desc, HowTo: d.HowTo, Correct: d.Correct, Wrong: d.Wrong, Breathing: d.Breathing, Safety: d.Safety,
				Voice: Voice{Start: d.Voice.Start, Mid: d.Voice.Mid, Last: d.Voice.Last, Finish: d.Voice.Finish}}
		}
		detailsByLang[lang] = m
	}
}

// mergeBundledTranslations adds languages to the bundled exercise documents.
func mergeBundledTranslations() {
	for lang, raw := range translationFiles("exercises") {
		var locs map[string]json.RawMessage
		if err := json.Unmarshal(raw, &locs); err != nil {
			panic(fmt.Sprintf("content: exercises.%s.json: %v", lang, err))
		}
		for id, loc := range locs {
			doc, ok := bundledDocs[id]
			if !ok {
				panic(fmt.Sprintf("content: exercises.%s.json: unknown exercise %q", lang, id))
			}
			var fields map[string]json.RawMessage
			var content map[string]json.RawMessage
			if json.Unmarshal(doc, &fields) != nil || json.Unmarshal(fields["content"], &content) != nil {
				panic("content: bad bundled document " + id)
			}
			content[lang] = loc
			fields["content"], _ = json.Marshal(content)
			bundledDocs[id], _ = json.Marshal(fields)
		}
	}
}

// addCourseTranslations adds titles and descriptions of the ready programs,
// the morning workout included (key "demo").
func addCourseTranslations() {
	for lang, raw := range translationFiles("courses") {
		var text map[string]struct {
			Title string `json:"title"`
			About string `json:"about"`
		}
		if err := json.Unmarshal(raw, &text); err != nil {
			panic(fmt.Sprintf("content: courses.%s.json: %v", lang, err))
		}
		for key, t := range text {
			title, about := courseText(key)
			if title == nil {
				panic(fmt.Sprintf("content: courses.%s.json: unknown program %q", lang, key))
			}
			title[lang], about[lang] = t.Title, t.About
		}
	}
}

// courseText returns the title and description maps of a program.
func courseText(key string) (title, about map[string]string) {
	if key == DemoKey {
		return demoProgram.Title, demoProgram.About
	}
	for _, c := range series30 {
		if c.Key == key {
			return c.Title, c.About
		}
	}
	return nil, nil
}
