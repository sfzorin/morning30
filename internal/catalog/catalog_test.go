package catalog

import (
	"encoding/json"
	"reflect"
	"strings"
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

func TestReactionDocumentValidationAndRoundTrip(t *testing.T) {
	d := Doc{ID: "R01", Slot: "main", Unit: "seconds", Reaction: &Reaction{Directions: []string{"left", "right"}}}
	if !ValidDoc(d) {
		t.Fatal("valid timed reaction rejected")
	}
	var got Doc
	if err := json.Unmarshal(d.Marshal(), &got); err != nil || !reflect.DeepEqual(got.Reaction, d.Reaction) {
		t.Fatalf("reaction did not round trip: %#v, %v", got, err)
	}
	for _, r := range []*Reaction{{}, {Directions: []string{"left"}}, {Directions: []string{"left", "left"}}, {Directions: []string{"left", "up"}}, {Directions: []string{"left", "right", "left"}}} {
		d.Reaction = r
		if ValidDoc(d) {
			t.Fatalf("accepted invalid reaction: %#v", r)
		}
	}
	d.Reaction = &Reaction{Directions: []string{"right", "left"}}
	for _, unit := range []string{"reps", "breaths"} {
		d.Unit = unit
		if ValidDoc(d) {
			t.Fatalf("accepted reaction for %s", unit)
		}
	}
	d.Unit, d.PerSide = "seconds", true
	if ValidDoc(d) {
		t.Fatal("accepted reaction with a side-change timer")
	}
	for _, raw := range []string{
		`{"reaction":{"directions":["left","right"],"rest":60}}`,
		`{"reaction":{"directions":["left","right"],"interval":1}}`,
		`{"reaction":{"directions":["left",1]}}`,
	} {
		if err := json.Unmarshal([]byte(raw), &got); err == nil {
			t.Fatalf("accepted malformed reaction: %s", raw)
		}
	}
}

func TestCatalogReactionIsValidatedAndCopied(t *testing.T) {
	c := New(`{"R":{"id":"R","slot":"main","unit":"seconds","reaction":{"directions":["left","right"],"cadence":{"startSeconds":3,"endSeconds":1}}}}`)
	r := c.Reaction("R")
	if r == nil {
		t.Fatal("missing reaction")
	}
	r.Directions[0] = "changed"
	r.Cadence.StartSeconds, r.Cadence.EndSeconds = 10, 9
	if got := c.Reaction("R"); got.Directions[0] != "left" || *got.Cadence != (ReactionCadence{StartSeconds: 3, EndSeconds: 1}) {
		t.Fatal("payload modified stored reaction")
	}
	c.Put(Doc{ID: "invalid", Slot: "main", Unit: "reps", Reaction: &Reaction{Directions: []string{"left", "right"}}})
	if c.Reaction("invalid") != nil || c.Reaction("missing") != nil {
		t.Fatal("invalid or missing document produced reaction cues")
	}
}

func TestReactionCadenceRoundTripAndLegacyDefault(t *testing.T) {
	for _, cadence := range []*ReactionCadence{nil,
		{StartSeconds: 3, EndSeconds: 1}, {StartSeconds: 1, EndSeconds: 1},
		{StartSeconds: 10, EndSeconds: 10}, {StartSeconds: 10, EndSeconds: 1},
	} {
		d := Doc{ID: "R", Slot: "main", Unit: "seconds", Reaction: &Reaction{Directions: []string{"left", "right"}, Cadence: cadence}}
		if !ValidDoc(d) {
			t.Fatalf("rejected cadence: %#v", cadence)
		}
		raw := d.Marshal()
		if cadence == nil && strings.Contains(string(raw), `"cadence"`) {
			t.Fatal("legacy reaction must omit cadence, retaining the player's default")
		}
		var got Doc
		if err := json.Unmarshal(raw, &got); err != nil || !reflect.DeepEqual(got.Reaction, d.Reaction) {
			t.Fatalf("cadence did not round trip: %#v, %v", got.Reaction, err)
		}
		for _, unit := range []string{"reps", "breaths"} {
			got.Unit = unit
			if ValidDoc(got) {
				t.Fatalf("cadence enabled reaction on %s", unit)
			}
		}
		got.Unit, got.PerSide = "seconds", true
		if ValidDoc(got) {
			t.Fatal("cadence enabled a reaction side-change timer")
		}
	}
}

func TestReactionCadenceRejectsMalformedJSON(t *testing.T) {
	for _, raw := range []string{
		`null`, `[]`, `"3,1"`, `true`, `3`, `{}`,
		`{"startSeconds":3}`, `{"endSeconds":1}`,
		`{"startSeconds":null,"endSeconds":1}`,
		`{"startSeconds":3,"endSeconds":null}`,
		`{"startSeconds":true,"endSeconds":1}`,
		`{"startSeconds":"3","endSeconds":1}`,
		`{"startSeconds":3.0,"endSeconds":1}`,
		`{"startSeconds":3,"endSeconds":1.5}`,
		`{"startSeconds":0,"endSeconds":1}`,
		`{"startSeconds":11,"endSeconds":1}`,
		`{"startSeconds":3,"endSeconds":0}`,
		`{"startSeconds":3,"endSeconds":4}`,
		`{"startSeconds":3,"endSeconds":1,"rest":10}`,
		`{"StartSeconds":3,"endSeconds":1}`,
	} {
		t.Run(raw, func(t *testing.T) {
			var d Doc
			err := json.Unmarshal([]byte(`{"reaction":{"directions":["left","right"],"cadence":`+raw+`}}`), &d)
			if err == nil {
				t.Fatal("accepted malformed cadence")
			}
		})
	}
}

func TestReactionCadenceValidatesDirectStructs(t *testing.T) {
	for _, cadence := range []ReactionCadence{
		{}, {StartSeconds: -1, EndSeconds: 1}, {StartSeconds: 11, EndSeconds: 1},
		{StartSeconds: 3, EndSeconds: 0}, {StartSeconds: 3, EndSeconds: 4},
	} {
		d := Doc{ID: "R", Slot: "main", Unit: "seconds", Reaction: &Reaction{Directions: []string{"left", "right"}, Cadence: &cadence}}
		if ValidDoc(d) {
			t.Fatalf("accepted invalid cadence: %#v", cadence)
		}
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
