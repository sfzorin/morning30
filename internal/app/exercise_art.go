package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"

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

var exerciseArtwork = readExerciseArtwork(assets.Static())

type exerciseArtIndex map[string]map[string]figureArt

func readExerciseArtwork(files fs.FS) exerciseArtIndex {
	index := make(exerciseArtIndex)
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
		if _, ok := content.Get(id); !ok {
			continue
		}
		figures := make(map[string]figureArt)
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
			figures[sex] = figure
		}
		index[id] = figures
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

// Share complete figure variants between equivalent built-in exercises.
// The table names direct sources only, so its map iteration order is immaterial.
func addSharedArt(index exerciseArtIndex, motions map[string]catalog.Animation, aliases map[string]string) {
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
			if _, ok := content.Get(target); !ok {
				continue
			}
			targetArt = make(map[string]figureArt)
		}
		for sex, figure := range sourceArt {
			if _, own := targetArt[sex]; !own && (!dynamic || figure.demo != nil) {
				targetArt[sex] = figure
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
