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
	d := Doc{ID: "X01", Slot: "main", Unit: "reps", Poses: []string{"a", "b"}, Content: map[string]Loc{"en": {Name: "X"}}}
	var got Doc
	if err := json.Unmarshal([]byte(d.Compact()), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Poses) != 2 || got.Poses[0] != "a" {
		t.Fatalf("poses %#v", got.Poses)
	}
}
