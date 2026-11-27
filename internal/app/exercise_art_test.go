package app

import (
	"slices"
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
	if got := ExerciseFrames(cat, "W01", "m"); !slices.Equal(got, []string{"/poses/w01-1?sex=m", "/poses/w01-2?sex=m"}) {
		t.Fatalf("strip without a still hid the existing pictures: %v", got)
	}
	delete(files, root+"motions.json")
	useArtFiles(t, files)
	if len(ExerciseFrames(cat, "W01", "g")) != 1 || ExerciseAnimation(cat, "W01", "g") != nil {
		t.Fatal("stills should work without the optional motion manifest")
	}
}

func TestExerciseArtworkRespectsSavedAssignmentsAndCustomMovements(t *testing.T) {
	d := openAppDB(t)
	file := &fstest.MapFile{Data: []byte("image")}
	root := exerciseArtDir + "/P08/"
	useArtFiles(t, fstest.MapFS{root + "f/still.webp": file, root + "m/still.webp": file})
	for _, pid := range []string{"my-view", "my-motion"} {
		if err := d.PutPose(pid, "", "", "m", tinyPNG); err != nil {
			t.Fatal(err)
		}
	}
	doc, _ := catalog.Builtin("P08")
	doc.Poses = []string{"my-view"}
	doc.Animation = &catalog.Animation{Pose: "my-motion", Kind: "movement", Frames: 2}
	cat := catalog.New("")
	cat.Put(doc)
	if got := ExerciseFrames(cat, "P08", "m"); !slices.Equal(got, []string{"/poses/my-view?sex=m"}) {
		t.Fatalf("lost personal still: %v", got)
	}
	if got := ExerciseAnimation(cat, "P08", "m"); got == nil || got.Src != "/poses/my-motion?sex=m" {
		t.Fatalf("lost personal motion: %+v", got)
	}
	if got := ExerciseFrames(cat, "P08", "f"); !slices.Equal(got, []string{"/static/" + root + "f/still.webp" + imageVersion}) || ExerciseAnimation(cat, "P08", "f") != nil {
		t.Fatalf("different figure took precedence over available female artwork: %v", got)
	}
	doc.Content["en"] = catalog.Loc{Name: "An entirely different movement"}
	cat.Put(doc)
	if got := ExerciseFrames(cat, "P08", "f"); !slices.Equal(got, []string{"/poses/my-view?sex=f"}) {
		t.Fatalf("built-in artwork replaced a different custom exercise: %v", got)
	}
}

func TestSharedArtworkUsesCompleteSourcesAndPreservesTargetAssignments(t *testing.T) {
	root := exerciseArtDir + "/"
	file := &fstest.MapFile{Data: []byte("image")}
	files := fstest.MapFS{
		root + "shared-movements.json": {Data: []byte(`{"IT09":"P01"}`)},
		root + "motions.json":          {Data: []byte(`{"P01":{"frames":2,"kind":"movement"},"IT09":{"frames":2,"kind":"movement"}}`)},
		root + "P01/f/still.webp":      file, root + "P01/f/movement.webp": file,
		root + "P01/b/still.webp": file, root + "P01/b/movement.webp": file,
		root + "P01/g/still.webp":  file, // A missing strip must preserve IT09's legacy fallback.
		root + "IT09/b/still.webp": file,
	}
	d := openAppDB(t)
	useArtFiles(t, files)
	doc, _ := catalog.Builtin("IT09")
	for _, sex := range []string{"f", "g"} {
		if err := d.PutPose(doc.Poses[0], "", "", sex, tinyPNG); err != nil {
			t.Fatal(err)
		}
	}
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
	if got := ExerciseFrames(cat, "IT09", "g"); !slices.Equal(got, []string{poseURL(doc.Poses[0], "g")}) {
		t.Fatalf("incomplete source replaced the saved fallback: %v", got)
	}
	files[root+"P01/f/still.webp"] = &fstest.MapFile{Data: []byte("corrected source illustration")}
	useArtFiles(t, files)
	refreshed := ExerciseFrames(cat, "IT09", "f")
	if len(refreshed) != 1 || slices.Equal(got, refreshed) || !slices.Equal(refreshed, ExerciseFrames(cat, "P01", "f")) {
		t.Fatalf("corrected source kept stale alias URL: old %v, new %v", got, refreshed)
	}
	doc.Poses = []string{"personal"}
	cat.Put(doc)
	if err := d.PutPose("personal", "", "", "f", tinyPNG); err != nil {
		t.Fatal(err)
	}
	if got := ExerciseFrames(cat, "IT09", "f"); !slices.Equal(got, []string{poseURL("personal", "f")}) || ExerciseAnimation(cat, "IT09", "f") != nil {
		t.Fatalf("alias hid a personal assignment: %v", got)
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
		for sex, original := range exerciseArtwork[source].figures {
			_, dynamic := motions[source]
			if dynamic && original.demo == nil {
				continue
			}
			base := exerciseArtDir + "/" + target + "/" + sex + "/"
			if artFileURL(files, base+"still.webp") != "" {
				continue
			}
			if inherited := exerciseArtwork[target].figures[sex]; inherited != original {
				t.Errorf("%s/%s missing completed source %s: %+v", target, sex, source, inherited)
			}
		}
	}
}
