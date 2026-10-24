package catalog

import (
	"encoding/json"
	"testing"
)

func TestClipPoses(t *testing.T) {
	got := ClipPoses([]string{" Plank-High ", "plank-high", "plank-low", "nope.png", ""})
	if len(got) != 2 || got[0] != "plank-high" || got[1] != "plank-low" {
		t.Fatalf("got %#v", got)
	}
	if len(ClipPoses(nil)) != 0 {
		t.Fatal("empty")
	}
}

func TestDocPosesRoundTrip(t *testing.T) {
	d := Doc{ID: "X01", Slot: "main", Unit: "reps", Poses: []string{"a", "b"}, Animation: &Animation{Pose: "motion", OtherPose: "other", Frames: 3, Kind: "movement"}, Content: map[string]Loc{"en": {Name: "X"}}}
	var got Doc
	if err := json.Unmarshal([]byte(d.Compact()), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Poses) != 2 || got.Poses[0] != "a" {
		t.Fatalf("poses %#v", got.Poses)
	}
	if got.Animation == nil || *got.Animation != *d.Animation {
		t.Fatalf("lost animation: %#v", got.Animation)
	}
}

func TestAnimationValidation(t *testing.T) {
	for _, a := range []*Animation{nil, {Pose: "motion-v1", Frames: 3, Kind: "alternating"}, {Pose: "a", OtherPose: "b", Frames: 3, Kind: "sequence"}} {
		if !ValidAnimation(a) {
			t.Fatalf("rejected valid animation: %#v", a)
		}
	}
	for _, a := range []*Animation{
		{Pose: "../a", Frames: 3, Kind: "movement"},
		{Pose: "a", Frames: 1, Kind: "movement"},
		{Pose: "a", Frames: 4, Kind: "movement"},
		{Pose: "a", Frames: 2, Kind: "alternating"},
		{Pose: "a", OtherPose: "../b", Frames: 3, Kind: "movement"},
		{Pose: "a", Frames: 3, Kind: "unknown"},
	} {
		d := Doc{ID: "X", Slot: "main", Unit: "reps", Animation: a}
		if ValidDoc(d) {
			t.Fatalf("accepted invalid animation: %#v", a)
		}
	}
}
