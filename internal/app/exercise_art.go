package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"
	"slices"
	"strings"

	"morning30/assets"
	"morning30/internal/catalog"
	"morning30/internal/content"
)

const exerciseArtDir = "exercise-art/v1"

// Only complete stills are advertised. Motion can be added in a later build.
type figureArt struct {
	still string
	demo  *ExerciseDemo
}

type exerciseArt struct {
	figures   map[string]figureArt
	poses     []string
	animation *catalog.Animation
	name      string
	howTo     []string
}

var exerciseArtwork = readExerciseArtwork(assets.Static())

func readExerciseArtwork(files fs.FS) map[string]exerciseArt {
	index := make(map[string]exerciseArt)
	dirs, err := fs.ReadDir(files, exerciseArtDir)
	if errors.Is(err, fs.ErrNotExist) {
		return index
	}
	if err != nil {
		log.Printf("exercise artwork: %v", err)
		return index
	}
	motions := readArtMap[catalog.Animation](files, "motions.json")
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		id := dir.Name()
		doc, ok := catalog.Builtin(id)
		if !ok {
			continue
		}
		art := artForBuiltin(doc)
		motion := motions[id]
		// The manifest only describes the order and number of square frames.
		hasMotion := catalog.ValidFrameSequence(motion.Frames, motion.Kind)
		for _, sex := range Figures {
			base := exerciseArtDir + "/" + id + "/" + sex + "/"
			still := artFileURL(files, base+"still.webp")
			if still == "" {
				continue
			}
			figure := figureArt{still: still}
			if hasMotion {
				if strip := artFileURL(files, base+motion.Kind+".webp"); strip != "" {
					figure.demo = &ExerciseDemo{Src: strip, Frames: motion.Frames, Kind: motion.Kind}
				}
			}
			art.figures[sex] = figure
		}
		index[id] = art
	}
	addSharedArt(index, motions, readArtMap[string](files, "shared-movements.json"))
	return index
}

func readArtMap[T any](files fs.FS, name string) map[string]T {
	raw, err := fs.ReadFile(files, exerciseArtDir+"/"+name)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	var entries map[string]T
	if err == nil {
		err = json.Unmarshal(raw, &entries)
	}
	if err != nil {
		log.Printf("exercise artwork %s: %v", name, err)
		return nil
	}
	return entries
}

func artForBuiltin(doc catalog.Doc) exerciseArt {
	en := doc.Content["en"]
	return exerciseArt{figures: map[string]figureArt{}, poses: doc.Poses, animation: doc.Animation, name: en.Name, howTo: en.HowTo}
}

// Share complete figure variants, keeping the target's pose assignments and technique.
// The table names direct sources only, so its map iteration order is immaterial.
func addSharedArt(index map[string]exerciseArt, motions map[string]catalog.Animation, aliases map[string]string) {
	if motions == nil {
		return // Without the manifest we cannot tell whether a missing strip is intentional.
	}
	for target, source := range aliases {
		sourceArt, ok := index[source]
		if !ok || aliases[source] != "" {
			continue
		}
		motion, dynamic := motions[target]
		sourceMotion, sourceDynamic := motions[source]
		if dynamic != sourceDynamic || motion != sourceMotion || (dynamic && !catalog.ValidFrameSequence(motion.Frames, motion.Kind)) {
			continue
		}
		targetArt, ok := index[target]
		if !ok {
			doc, ok := catalog.Builtin(target)
			if !ok {
				continue
			}
			targetArt = artForBuiltin(doc)
		}
		for sex, figure := range sourceArt.figures {
			if _, own := targetArt.figures[sex]; !own && (!dynamic || figure.demo != nil) {
				targetArt.figures[sex] = figure
			}
		}
		index[target] = targetArt
	}
}

// The PWA caches static URLs. Correcting a source file must refresh every exercise
// displaying it, including those that share the same illustration.
func artFileURL(files fs.FS, name string) string {
	f, err := files.Open(name)
	if err != nil {
		return ""
	}
	defer f.Close()
	hash := sha256.New()
	count, err := io.Copy(hash, f)
	if err != nil || count == 0 {
		return ""
	}
	return "/static/" + name + "?v=" + hex.EncodeToString(hash.Sum(nil)[:8])
}

// A personal exercise can reuse a built-in ID for a different movement.
func artForExercise(cat catalog.Catalog, doc catalog.Doc, sex string) (exerciseArt, figureArt, bool) {
	art := exerciseArtwork[doc.ID]
	figure, ok := art.figures[sex]
	if content.FigureKey(doc.SVG, doc.ID) != doc.ID {
		return art, figure, false
	}
	if ok && cat.IsCustom(doc.ID) {
		en := doc.Content["en"]
		ok = en.Name == art.name && slices.Equal(en.HowTo, art.howTo)
	}
	return art, figure, ok
}

func (art exerciseArt) hasAssignedStills(id string, poses []string) bool {
	if len(poses) == 0 || slices.Equal(poses, art.poses) {
		return false
	}
	// SeedPoseStills assigns this exact pair to the original built-ins.
	key := strings.ToLower(id)
	return len(poses) != 2 || poses[0] != key+"-1" || poses[1] != key+"-2"
}

func (art exerciseArt) hasAssignedAnimation(a *catalog.Animation) bool {
	return a != nil && (art.animation == nil || *a != *art.animation)
}
