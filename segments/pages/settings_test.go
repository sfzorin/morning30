package pages

import (
	"testing"

	"morning30/internal/content"
)

func TestEditorPreservesPrescribedProgram(t *testing.T) {
	p, err := content.ParseResolved(`{"fixedDifficulty":true,"warmup":[{"id":"w","value":45}],"cooldown":[{"id":"c","value":20}],"days":[{"rounds":1,"warmup":[],"cooldown":[],"items":[{"id":"a","value":6,"rest":0,"sideRest":10,"repSeconds":4}]}]}`)
	if err != nil {
		t.Fatal(err)
	}
	editProgramList(&p, 1, "day", func(items []content.RItem) []content.RItem { items[0].Value = 7; return items })
	editProgramList(&p, 1, "day-warmup", func(items []content.RItem) []content.RItem {
		return append(items, content.RItem{ID: "local", Value: 20})
	})
	editProgramList(&p, 1, "day-cooldown", func(items []content.RItem) []content.RItem { return nil })
	q, err := content.ParseResolved(p.Marshal())
	if err != nil {
		t.Fatal(err)
	}
	it := q.Days[0].Items[0]
	if !q.FixedDifficulty || it.Value != 7 || it.Rest == nil || *it.Rest != 0 || it.SideRest != 10 || it.RepSeconds != 4 {
		t.Fatalf("editor lost prescription: %+v", q)
	}
	if len(q.Warmup) != 1 || q.Warmup[0].ID != "w" || len(q.WarmupFor(0)) != 1 || q.WarmupFor(0)[0].ID != "local" {
		t.Fatal("day phase edits changed shared phases")
	}
	if q.Days[0].Cooldown == nil || len(q.CooldownFor(0)) != 0 {
		t.Fatal("empty day phase was lost during import")
	}
}
