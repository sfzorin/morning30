// Package library exposes the imported Leap Fitness exercise library: a large,
// grouped, multi-language catalogue of demo exercises (each backed by a source
// YouTube video). It is browse-only reference content, separate from the curated
// 30-day programs. Data is embedded at build time (assets.LeapLibrary).
package library

import (
	"encoding/json"
	"sort"

	"morning30/assets"
)

// Localized is one card's text in a single language.
type Localized struct {
	ShortDescription string   `json:"short_description"`
	HowTo            []string `json:"how_to"`
	Correct          []string `json:"correct"`
	Mistakes         []string `json:"mistakes"`
}

// Source is the origin video for an exercise.
type Source struct {
	VideoTitle string `json:"video_title"`
	VideoURL   string `json:"video_url"`
	VideoID    string `json:"video_id"`
	Thumbnail  string `json:"thumbnail"`
}

// Card is one library exercise with names + cards per language.
type Card struct {
	ID           string               `json:"exercise_id"`
	Slug         string               `json:"slug"`
	Group        string               `json:"group"`
	GroupLabel   string               `json:"group_label"`
	Names        map[string]string    `json:"names"`
	Cards        map[string]Localized `json:"cards"`
	Source       Source               `json:"source"`
	ReviewStatus string               `json:"review_status"`
}

// Name returns the exercise name in language l, falling back to English.
func (c *Card) Name(l string) string {
	if n, ok := c.Names[l]; ok && n != "" {
		return n
	}
	return c.Names["en"]
}

// Text returns the card body in language l, falling back to English.
func (c *Card) Text(l string) Localized {
	if t, ok := c.Cards[l]; ok && t.ShortDescription != "" {
		return t
	}
	return c.Cards["en"]
}

// Group is one browse category with its exercises.
type Group struct {
	Key   string
	Label string
	Cards []*Card
}

// VideoRef links an existing app exercise to a matching demo video.
type VideoRef struct {
	ExerciseID string `json:"exercise_id"`
	MatchedFrom string `json:"matched_from"`
	VideoURL   string `json:"video_url"`
	VideoID    string `json:"video_id"`
	VideoTitle string `json:"video_title"`
}

// groupOrder is the display order for browse groups (broad → niche), matching
// the pipeline's grouping taxonomy in scripts/leap_common.py.
var groupOrder = []string{
	"plank", "push", "squat", "lunge", "glute", "core", "cardio", "back",
	"legs", "arms", "mobility", "yoga", "wall", "face", "dumbbell", "equipment", "other",
}

var (
	all      []*Card
	byID     = map[string]*Card{}
	groups   []Group
	videoMap = map[string]VideoRef{}
)

func init() {
	var cards []*Card
	if err := json.Unmarshal(assets.LeapLibrary, &cards); err != nil {
		panic("library: bad leap_library.json: " + err.Error())
	}
	all = cards
	for _, c := range cards {
		byID[c.ID] = c
	}

	rank := map[string]int{}
	for i, k := range groupOrder {
		rank[k] = i
	}
	byKey := map[string]*Group{}
	for _, c := range cards {
		g, ok := byKey[c.Group]
		if !ok {
			g = &Group{Key: c.Group, Label: c.GroupLabel}
			byKey[c.Group] = g
		}
		g.Cards = append(g.Cards, c)
	}
	for _, g := range byKey {
		sort.Slice(g.Cards, func(i, j int) bool { return g.Cards[i].Names["en"] < g.Cards[j].Names["en"] })
		groups = append(groups, *g)
	}
	sort.Slice(groups, func(i, j int) bool {
		ri, oki := rank[groups[i].Key]
		rj, okj := rank[groups[j].Key]
		if !oki {
			ri = len(groupOrder)
		}
		if !okj {
			rj = len(groupOrder)
		}
		if ri != rj {
			return ri < rj
		}
		return groups[i].Label < groups[j].Label
	})

	var refs []VideoRef
	if err := json.Unmarshal(assets.LeapVideoMap, &refs); err != nil {
		panic("library: bad leap_video_map.json: " + err.Error())
	}
	for _, r := range refs {
		videoMap[r.ExerciseID] = r
	}
}

// Groups returns the browse categories in display order.
func Groups() []Group { return groups }

// Count returns the total number of library exercises.
func Count() int { return len(all) }

// ByID returns the card with the given id.
func ByID(id string) (*Card, bool) {
	c, ok := byID[id]
	return c, ok
}

// Video returns a demo video matched to an existing app exercise id, if any.
func Video(existingID string) (VideoRef, bool) {
	v, ok := videoMap[existingID]
	return v, ok
}
