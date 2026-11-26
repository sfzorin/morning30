package content

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
)

// Bundled exercises: complete exercise documents (definition plus names and
// technique in every language) kept as JSON in exercises/*.json. They join the
// built-in library at start-up, so programs, the library, the player and
// agents treat them like any other built-in exercise. Their pictures come from
// the shared pose catalog only; they have no legacy drawings.

//go:embed exercises/*.json
var bundleFS embed.FS

// bundledDef is the part of a bundled document this package needs.
type bundledDef struct {
	ID          string `json:"id"`
	Slot        Slot   `json:"slot"`
	Unit        Unit   `json:"unit"`
	PerSide     bool   `json:"perSide"`
	KneeLoad    string `json:"kneeLoad"`
	Difficulty  int    `json:"difficulty"`
	Replacement string `json:"replacement"`
	Jumps       bool   `json:"jumps"`
	Equipment   bool   `json:"equipment"`
}

var (
	bundledDocs = map[string]json.RawMessage{}
	bundledDefs = map[string]bundledDef{}
)

func init() {
	files, err := fs.Glob(bundleFS, "exercises/*.json")
	if err != nil {
		panic(err)
	}
	sort.Strings(files)
	for _, name := range files {
		raw, err := bundleFS.ReadFile(name)
		if err != nil {
			panic(err)
		}
		var docs []json.RawMessage
		if err := json.Unmarshal(raw, &docs); err != nil {
			panic(fmt.Sprintf("content: %s: %v", name, err))
		}
		for _, doc := range docs {
			var def bundledDef
			if err := json.Unmarshal(doc, &def); err != nil {
				panic(fmt.Sprintf("content: %s: %v", name, err))
			}
			if def.ID == "" || (def.Slot != Warmup && def.Slot != Main && def.Slot != Cooldown) ||
				(def.Unit != Reps && def.Unit != Seconds && def.Unit != Breaths) {
				panic(fmt.Sprintf("content: %s: invalid exercise %q", name, def.ID))
			}
			if _, dup := Pool[def.ID]; dup {
				panic(fmt.Sprintf("content: %s: duplicate exercise %q", name, def.ID))
			}
			ex := Exercise{ID: def.ID, Slot: def.Slot, Unit: def.Unit, PerSide: def.PerSide, KneeLoad: KneeLoad(def.KneeLoad),
				Difficulty: def.Difficulty, Replacement: def.Replacement, Bundled: true}
			if ex.KneeLoad == "" {
				ex.KneeLoad = KneeNone
			}
			library = append(library, ex)
			Pool[def.ID] = ex
			bundledDocs[def.ID] = doc
			bundledDefs[def.ID] = def
		}
	}
}

// BundledDoc returns the complete JSON document of a bundled exercise.
func BundledDoc(id string) (json.RawMessage, bool) {
	doc, ok := bundledDocs[id]
	return doc, ok
}

// IsBundled reports whether an exercise comes from a bundled document.
func IsBundled(id string) bool {
	_, ok := bundledDocs[id]
	return ok
}
