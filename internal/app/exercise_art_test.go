package app

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"morning30/assets"
	"morning30/internal/catalog"
	"morning30/internal/content"
)

const imageVersion = "?v=6105d6cc76af4003" // SHA-256 prefix of the fixture "image".

func useArtFiles(t *testing.T, files fstest.MapFS) {
	t.Helper()
	previous := exerciseArtwork
	exerciseArtwork = readExerciseArtwork(files)
	t.Cleanup(func() { exerciseArtwork = previous })
}

func TestExerciseArtworkUsesOnlyPublishedFigureAndPhases(t *testing.T) {
	root := exerciseArtDir + "/"
	file := &fstest.MapFile{Data: []byte("image")}
	files := fstest.MapFS{
		root + "motions.json":     {Data: []byte(`{"W01":{"frames":3,"kind":"sequence"},"P08":{"frames":2,"kind":"movement"}}`)},
		root + "W01/g/still.webp": file, root + "W01/g/sequence.webp": file,
		root + "W01/b/still.webp": file, root + "W01/m/sequence.webp": file,
		root + "P08/f/still.webp": file, root + "P08/f/movement.webp": file,
		root + "IT54/g/still.webp": file,
	}
	d := openAppDB(t)
	useArtFiles(t, files)
	for _, pid := range []string{"w01-1", "w01-2", "ilya-boy-it54-v1", "ilya-boy-it54-demo-v1"} {
		if err := d.PutPose(pid, "", "", "m", tinyPNG); err != nil {
			t.Fatal(err)
		}
	}
	builtin := catalog.New("")
	doc, _ := builtin.Doc("W01")
	doc.Poses = []string{"w01-1", "w01-2"}
	cat := catalog.NewLayered(map[string]string{"W01": doc.Compact()}, "")
	for _, tc := range []struct {
		id, figure, kind string
		frames           int
	}{
		{"W01", "g", "sequence", 3},
		{"W01", "b", "", 0}, // Completed still, pending motion.
		{"P08", "f", "movement", 2},
		{"IT54", "g", "", 0}, // The former boy animation must not replace this hold.
	} {
		base := "/static/" + root + tc.id + "/" + tc.figure + "/"
		if got := ExerciseFrames(cat, tc.id, tc.figure); !slices.Equal(got, []string{base + "still.webp" + imageVersion}) {
			t.Fatalf("%s/%s stills: %v", tc.id, tc.figure, got)
		}
		demo := ExerciseAnimation(cat, tc.id, tc.figure)
		if tc.frames == 0 {
			if demo != nil {
				t.Fatalf("unexpected demo: %+v", demo)
			}
		} else if demo == nil || demo.Src != base+tc.kind+".webp"+imageVersion || demo.Frames != tc.frames || demo.Kind != tc.kind {
			t.Fatalf("%s/%s demo: %+v", tc.id, tc.figure, demo)
		}
	}
	if got := ExerciseFrames(cat, "W01", "m"); len(got) != 0 {
		t.Fatalf("strip without a still brought back the mannequin images: %v", got)
	}
	delete(files, root+"motions.json")
	useArtFiles(t, files)
	if len(ExerciseFrames(cat, "W01", "g")) != 1 || ExerciseAnimation(cat, "W01", "g") != nil {
		t.Fatal("stills should work without the optional motion manifest")
	}
}

func TestSavedProgramExercisesUsePublishedArtwork(t *testing.T) {
	d := openAppDB(t)
	for _, tc := range []struct{ prefix, sex string }{
		{"sergey-man-", "m"},
		{"ilya-boy-", "b"},
	} {
		for _, suffix := range []string{"old-still", "old-animation"} {
			if err := d.PutPose(tc.prefix+suffix, "", "", tc.sex, tinyPNG); err != nil {
				t.Fatal(err)
			}
		}
		cat := catalog.New("")
		for _, exercise := range content.Library() {
			id := exercise.ID
			doc, _ := catalog.Builtin(id)
			doc.Content = map[string]catalog.Loc{"ru": {Name: "Название в личной программе"}}
			doc.SVG = "/static/ex/m/W01.png"
			doc.Poses = []string{tc.prefix + "old-still"}
			doc.Animation = &catalog.Animation{Pose: tc.prefix + "old-animation", Kind: "movement", Frames: 2}
			cat.Put(doc)
			for _, sex := range Figures {
				expected := exerciseArtwork[id][sex]
				if got := ExerciseFrames(cat, id, sex); len(got) != 1 ||
					!strings.HasPrefix(got[0], "/static/"+exerciseArtDir+"/") || got[0] != expected.still {
					t.Fatalf("%s%s/%s resolved to %v, want %q", tc.prefix, id, sex, got, expected.still)
				}
				if got := ExerciseAnimation(cat, id, sex); got != expected.demo {
					t.Fatalf("%s%s/%s animation: %+v, want %+v", tc.prefix, id, sex, got, expected.demo)
				}
			}
		}
	}
}

func TestArchivedPosesAreUnavailableInCustomExercisesAndOverHTTP(t *testing.T) {
	d := openAppDB(t)
	cat := catalog.New("")
	for _, tc := range []struct {
		id   string
		want int
	}{
		{"sergey-man-w01-v1", http.StatusNotFound},
		{"ilya-boy-it01-demo-v1", http.StatusNotFound},
		{"w01-1", http.StatusNotFound},
		{"custom-pose", http.StatusOK},
	} {
		if err := d.PutPose(tc.id, "", "", "m", tinyPNG); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/poses/"+tc.id, nil)
		req.SetPathValue("id", tc.id)
		res := httptest.NewRecorder()
		ServePose(res, req)
		if res.Code != tc.want {
			t.Fatalf("%s returned %d, want %d", tc.id, res.Code, tc.want)
		}
		cat.Put(catalog.Doc{ID: "X01", Poses: []string{tc.id}})
		if got := ExerciseFrames(cat, "X01", "m"); (len(got) == 1) != (tc.want == http.StatusOK) {
			t.Fatalf("unexpected custom exercise frames for %s: %v", tc.id, got)
		}
	}
	poses, err := ActiveExercisePoses()
	if err != nil || len(poses) != 1 || poses[0].ID != "custom-pose" {
		t.Fatalf("custom pose picker contains archived images: %v (%v)", poses, err)
	}
}

func TestSharedArtworkUsesCompleteSourcesAndOverridesOldAssignments(t *testing.T) {
	root := exerciseArtDir + "/"
	file := &fstest.MapFile{Data: []byte("image")}
	files := fstest.MapFS{
		root + "shared-movements.json": {Data: []byte(`{"IT09":"P01"}`)},
		root + "motions.json":          {Data: []byte(`{"P01":{"frames":2,"kind":"movement"},"IT09":{"frames":2,"kind":"movement"}}`)},
		root + "P01/f/still.webp":      file, root + "P01/f/movement.webp": file,
		root + "P01/b/still.webp": file, root + "P01/b/movement.webp": file,
		root + "P01/g/still.webp":  file, // A dynamic alias requires a complete source.
		root + "IT09/b/still.webp": file,
	}
	useArtFiles(t, files)
	doc, _ := catalog.Builtin("IT09")
	doc.Poses = []string{"old-assigned-pose"}
	cat := catalog.New("")
	cat.Put(doc) // A saved copy retains IT09's own name, technique and pose assignments.
	got := ExerciseFrames(cat, "IT09", "f")
	source := ExerciseFrames(cat, "P01", "f")
	if len(got) != 1 || !slices.Equal(got, source) || ExerciseAnimation(cat, "IT09", "f") == nil {
		t.Fatalf("alias with different built-in metadata: stills %v, source %v", got, source)
	}
	if own := ExerciseFrames(cat, "IT09", "b"); !slices.Equal(own, []string{"/static/" + root + "IT09/b/still.webp" + imageVersion}) || ExerciseAnimation(cat, "IT09", "b") != nil {
		t.Fatalf("own figure should take priority: %v", own)
	}
	if got := ExerciseFrames(cat, "IT09", "g"); len(got) != 0 {
		t.Fatalf("incomplete source advertised a fallback: %v", got)
	}
	files[root+"P01/f/still.webp"] = &fstest.MapFile{Data: []byte("corrected source illustration")}
	useArtFiles(t, files)
	refreshed := ExerciseFrames(cat, "IT09", "f")
	if len(refreshed) != 1 || slices.Equal(got, refreshed) || !slices.Equal(refreshed, ExerciseFrames(cat, "P01", "f")) {
		t.Fatalf("corrected source kept stale alias URL: old %v, new %v", got, refreshed)
	}
}

func TestPublishedSharedMovements(t *testing.T) {
	files := assets.Static()
	aliases := readArtMap[string](files, "shared-movements.json")
	motions := readArtMap[catalog.Animation](files, "motions.json")
	if len(aliases) == 0 || len(motions) == 0 {
		t.Fatal("expected both published artwork manifests")
	}
	for target, source := range aliases {
		_, hasSource := content.Get(source)
		_, hasTarget := content.Get(target)
		if !hasSource || !hasTarget || aliases[source] != "" || motions[source] != motions[target] {
			t.Errorf("incompatible published alias %s -> %s", target, source)
		}
		for sex, original := range exerciseArtwork[source] {
			_, dynamic := motions[source]
			if dynamic && original.demo == nil {
				continue
			}
			base := exerciseArtDir + "/" + target + "/" + sex + "/"
			if artFileURL(files, base+"still.webp") != "" {
				continue
			}
			if inherited := exerciseArtwork[target][sex]; inherited != original {
				t.Errorf("%s/%s missing completed source %s: %+v", target, sex, source, inherited)
			}
		}
	}
}
