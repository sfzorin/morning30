package coach

import (
	"fmt"
	"sort"

	"morning30/internal/app"
	"morning30/internal/content"
	"morning30/internal/i18n"
)

// Change is one line of the list under a reply. It is computed from the
// program before and after the reply, never taken from the model's words.
type Change struct {
	Kind string `json:"kind"` // replace | add | remove | dose | new | rounds | level | rest_warmup | rest_main | rest_cooldown | rename
	Days []int  `json:"days,omitempty"`
	Part string `json:"part,omitempty"`
	From string `json:"from,omitempty"` // exercise name before (replace), old name (rename)
	To   string `json:"to,omitempty"`   // exercise name after / the exercise / new name
	Old  int    `json:"old,omitempty"`
	New  int    `json:"new,omitempty"`
	Pct  int    `json:"pct,omitempty"` // dose change in percent when doses differ by day
}

type changeKey struct {
	kind, part, from, to string
}

// Diff lists what changed in the current program since before.
func Diff(userID int64, lang string, before snapshot, created []string) []Change {
	l := i18n.Parse(lang)
	cat := app.UserCatalog(userID)
	name := func(id string) string { return cat.Name(l, id) }
	var out []Change
	for _, id := range created {
		out = append(out, Change{Kind: "new", To: name(id)})
	}
	old, err := content.ParseResolved(before.Program)
	now := app.UserProgram(userID)
	if err == nil && len(old.Days) == len(now.Days) {
		days := map[changeKey][]int{}
		doses := map[string][][2]int{} // exercise id → (old, new) per occurrence
		doseDays := map[string][]int{}
		var order []changeKey
		note := func(k changeKey, day int) {
			if _, ok := days[k]; !ok {
				order = append(order, k)
			}
			if d := days[k]; len(d) == 0 || d[len(d)-1] != day {
				days[k] = append(days[k], day)
			}
		}
		for i := range now.Days {
			for _, part := range []content.Slot{content.Warmup, content.Main, content.Cooldown} {
				a, b := partItems(old, i, part), partItems(now, i, part)
				removed, added := minus(a, b), minus(b, a)
				for len(removed) > 0 && len(added) > 0 {
					note(changeKey{"replace", "", removed[0].ID, added[0].ID}, i+1)
					removed, added = removed[1:], added[1:]
				}
				for _, it := range removed {
					note(changeKey{"remove", string(part), it.ID, ""}, i+1)
				}
				for _, it := range added {
					note(changeKey{"add", string(part), "", it.ID}, i+1)
				}
				for _, x := range a {
					for _, y := range b {
						if x.ID == y.ID && x.Value != y.Value {
							if _, ok := doses[x.ID]; !ok {
								order = append(order, changeKey{"dose", "", x.ID, ""})
							}
							doses[x.ID] = append(doses[x.ID], [2]int{x.Value, y.Value})
							if d := doseDays[x.ID]; len(d) == 0 || d[len(d)-1] != i+1 {
								doseDays[x.ID] = append(doseDays[x.ID], i+1)
							}
						}
					}
				}
			}
		}
		for _, k := range order {
			switch k.kind {
			case "replace":
				out = append(out, Change{Kind: "replace", Days: days[k], From: name(k.from), To: name(k.to)})
			case "remove":
				out = append(out, Change{Kind: "remove", Days: days[k], Part: k.part, To: name(k.from)})
			case "add":
				out = append(out, Change{Kind: "add", Days: days[k], Part: k.part, To: name(k.to)})
			case "dose":
				pairs := doses[k.from]
				c := Change{Kind: "dose", Days: doseDays[k.from], To: name(k.from)}
				same := true
				for _, p := range pairs {
					same = same && p == pairs[0]
				}
				if same {
					c.Old, c.New = pairs[0][0], pairs[0][1]
				} else {
					sumOld, sumNew := 0, 0
					for _, p := range pairs {
						sumOld += p[0]
						sumNew += p[1]
					}
					if sumOld > 0 {
						c.Pct = (sumNew*100+sumOld/2)/sumOld - 100
					}
				}
				out = append(out, c)
			}
		}
	}
	after := takeSnapshot(userID)
	for _, s := range []struct {
		kind     string
		old, new int
	}{
		{"rounds", before.Rounds, after.Rounds}, {"level", before.Level, after.Level},
		{"rest_warmup", before.RestWarmup, after.RestWarmup}, {"rest_main", before.RestMain, after.RestMain},
		{"rest_cooldown", before.RestCooldown, after.RestCooldown},
	} {
		if s.old != s.new {
			out = append(out, Change{Kind: s.kind, Old: s.old, New: s.new})
		}
	}
	if before.Name != after.Name {
		out = append(out, Change{Kind: "rename", From: before.Name, To: after.Name})
	}
	return out
}

func partItems(r content.Resolved, i int, part content.Slot) []content.RItem {
	switch part {
	case content.Warmup:
		return r.WarmupFor(i)
	case content.Cooldown:
		return r.CooldownFor(i)
	}
	return r.Days[i].Items
}

// minus returns the items of a whose ids are missing from b (as a multiset).
func minus(a, b []content.RItem) []content.RItem {
	left := map[string]int{}
	for _, it := range b {
		left[it.ID]++
	}
	var out []content.RItem
	for _, it := range a {
		if left[it.ID] > 0 {
			left[it.ID]--
			continue
		}
		out = append(out, it)
	}
	return out
}

// DayRanges writes day numbers as "1–3, 5, 7–30".
func DayRanges(days []int) string {
	d := append([]int(nil), days...)
	sort.Ints(d)
	s := ""
	for i := 0; i < len(d); {
		j := i
		for j+1 < len(d) && d[j+1] == d[j]+1 {
			j++
		}
		if s != "" {
			s += ", "
		}
		if i == j {
			s += fmt.Sprint(d[i])
		} else {
			s += fmt.Sprintf("%d–%d", d[i], d[j])
		}
		i = j + 1
	}
	return s
}
