package app

import (
	"bytes"
	"io/fs"
	"log"
	"path"
	"strings"

	"morning30/internal/i18n"
)

// poseKinds is the taxonomy support for each rendered frame (1, then 2).
// Slugs stay ASCII because the catalog kind column is a slug.
var poseKinds = map[string][2]string{
	"W01": {"stand", "stand"}, "W02": {"stand", "stand"}, "W03": {"stand", "stand"},
	"W04": {"stand", "hinge"}, "W05": {"stand", "squat"}, "W06": {"plank", "plank"},
	"W09": {"stand", "stand"},
	"C01": {"plank", "plank"}, "C02": {"plank", "plank"}, "C03": {"supine", "supine"},
	"C04": {"supine", "supine"}, "C05": {"supine", "supine"}, "C06": {"supine", "supine"},
	"C07": {"side", "side"}, "C08": {"side", "side"}, "C09": {"plank", "plank"},
	"C10": {"plank", "plank"}, "C11": {"supine", "supine"}, "C12": {"supine", "supine"},
	"C13": {"supine", "supine"},
	"P01": {"plank", "plank"}, "P02": {"plank", "plank"}, "P03": {"plank", "plank"},
	"P04": {"plank", "plank"}, "P05": {"plank", "plank"}, "P06": {"plank", "plank"},
	"P08": {"plank", "plank"},
	"B01": {"prone", "prone"}, "B02": {"prone", "prone"}, "B03": {"prone", "prone"},
	"B04": {"prone", "prone"}, "B05": {"prone", "prone"}, "B06": {"prone", "prone"},
	"B10": {"prone", "prone"},
	"A01": {"stand", "stand"}, "A02": {"stand", "stand"}, "A03": {"stand", "stand"},
	"A04": {"stand", "stand"},
	"L01": {"stand", "squat"}, "L02": {"stand", "squat"}, "L03": {"stand", "hinge"},
	"L04": {"supine", "supine"}, "L05": {"supine", "supine"}, "L06": {"supine", "supine"},
	"L07": {"stand", "stand"}, "L10": {"stand", "squat"}, "L11": {"stand", "squat"},
	"L12": {"stand", "squat"},
	"J02": {"stand", "stand"}, "J03": {"stand", "stand"}, "J05": {"squat", "air"},
	"J06": {"stand", "stand"}, "J07": {"stand", "stand"},
	"D01": {"stand", "stand"}, "D02": {"hinge", "hinge"}, "D03": {"stand", "squat"},
	"D04": {"stand", "stand"}, "D05": {"stand", "stand"},
	"CD01": {"supine", "supine"}, "CD02": {"supine", "supine"}, "CD03": {"supine", "supine"},
	"CD04": {"prone", "prone"}, "CD07": {"kneel", "kneel"},
}

// SeedPoseStills loads {ID}-1.png and {ID}-2.png into the shared pose catalog
// and points each exercise at that pair when it does not already have one.
// Activity and program position are not written.
func SeedPoseStills(fsys fs.FS) error {
	if DB == nil {
		return nil
	}
	frames := map[string][2][]byte{}
	err := fs.WalkDir(fsys, ".", func(name string, d fs.DirEntry, err error) error {
		if err != nil || d == nil || d.IsDir() {
			return err
		}
		base := path.Base(name)
		ex, n, ok := splitStill(base)
		if !ok {
			return nil
		}
		b, err := fs.ReadFile(fsys, name)
		if err != nil {
			return err
		}
		pair := frames[ex]
		pair[n-1] = b
		frames[ex] = pair
		return nil
	})
	if err != nil {
		return err
	}
	cat := UserCatalog(0)
	n := 0
	for ex, pair := range frames {
		if len(pair[0]) == 0 || len(pair[1]) == 0 {
			continue
		}
		kinds := poseKinds[ex]
		name := i18n.Exercise(i18n.RU, ex)
		for i := 0; i < 2; i++ {
			id := strings.ToLower(ex) + "-" + string(rune('1'+i))
			if sameStill(id, pair[i]) {
				continue
			}
			label := name
			if label == "" {
				label = ex
			}
			label += " " + string(rune('1'+i))
			if err := DB.PutPose(id, kinds[i], label, "m", pair[i]); err != nil {
				return err
			}
			n++
		}
		if len(cat.PoseIDs(ex)) > 0 {
			continue
		}
		if err := SetLibraryPoses(ex, []string{strings.ToLower(ex) + "-1", strings.ToLower(ex) + "-2"}); err != nil {
			log.Printf("pose assign %s: %v", ex, err)
		}
	}
	if n > 0 {
		log.Printf("pose stills updated: %d", n)
	}
	return nil
}

func splitStill(base string) (ex string, frame int, ok bool) {
	base = strings.TrimSuffix(base, ".png")
	if len(base) < 3 || base[len(base)-2] != '-' {
		return "", 0, false
	}
	n := base[len(base)-1]
	if n != '1' && n != '2' {
		return "", 0, false
	}
	ex = strings.ToUpper(base[:len(base)-2])
	if ex == "" {
		return "", 0, false
	}
	return ex, int(n - '0'), true
}

func sameStill(id string, next []byte) bool {
	_, cur, err := DB.PoseImage(id, "m")
	return err == nil && bytes.Equal(cur, next)
}
