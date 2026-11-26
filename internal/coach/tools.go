package coach

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"morning30/internal/app"
	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/i18n"
)

// The coach's tools. Every tool works on the signed-in user's CURRENT program
// and their own exercise library; nothing here can reach another program,
// another account or the settings beyond the program's own workout settings.

var daysProp = map[string]any{
	"description": `Days of the 30-day cycle: "all", "today", a day number, a range like "4-30", or a list like [1,3,5]. Default "all".`,
}

var partProp = map[string]any{"type": "string", "enum": []string{"warmup", "main", "cooldown"}}

func fn(name, desc string, props map[string]any, required ...string) map[string]any {
	if props == nil {
		props = map[string]any{}
	}
	params := map[string]any{"type": "object", "properties": props}
	if len(required) > 0 {
		params["required"] = required
	}
	return map[string]any{"type": "function", "function": map[string]any{"name": name, "description": desc, "parameters": params}}
}

func toolDefs() []map[string]any {
	str := func(desc string) map[string]any { return map[string]any{"type": "string", "description": desc} }
	num := func(desc string) map[string]any { return map[string]any{"type": "integer", "description": desc} }
	return []map[string]any{
		fn("search_exercises", "Find exercises in the library (built-in and the user's own). Returns id, name, group, part of the day, unit, difficulty 1–5, jumps, equipment. Use it before adding or replacing.",
			map[string]any{
				"query":    str("Words of the name, in any language, or an id"),
				"group":    map[string]any{"type": "string", "enum": []string{"warmup", "core", "push", "back", "arms", "legs", "cardio", "stretch", "dumbbell", "pullup", "balance", "taichi", "neck", "kids", "custom"}},
				"part":     partProp,
				"no_jumps": map[string]any{"type": "boolean", "description": "Leave out exercises with jumps"},
			}),
		fn("similar_exercises", "Exercises that can replace one: easier, similar and harder.", map[string]any{"id": str("Exercise id")}, "id"),
		fn("show_program", "Show the current program again, day by day (after edits).", nil),
		fn("add_exercise", "Add an exercise to a part of the chosen days.", map[string]any{
			"exercise": str("Exercise id"), "part": partProp, "days": daysProp,
			"value": num("Reps, seconds or breaths (by the exercise's unit). Omit for a sensible default."),
		}, "exercise", "part"),
		fn("replace_exercise", "Replace an exercise with another on the chosen days; the dose carries over.", map[string]any{
			"from": str("Exercise id now in the program"), "to": str("Exercise id to put instead"), "days": daysProp,
		}, "from", "to"),
		fn("remove_exercise", "Remove an exercise from the chosen days (a part of a day is never left empty).", map[string]any{
			"exercise": str("Exercise id"), "part": partProp, "days": daysProp,
		}, "exercise"),
		fn("set_part", "Write a whole part (warm-up, main or cool-down) of the chosen days at once: the list replaces what is there, in this order. Use it to write a new program or rework many days: e.g. days 1-10, 11-20 and 21-30 with growing doses.", map[string]any{
			"part": partProp, "days": daysProp,
			"exercises": map[string]any{"type": "array", "description": "In order; value = reps, seconds or breaths by the exercise's unit",
				"items": map[string]any{"type": "object", "properties": map[string]any{"id": str("Exercise id"), "value": num("Dose")}, "required": []string{"id", "value"}}},
		}, "part", "exercises"),
		fn("set_dose", "Change how many reps / seconds / breaths an exercise has on the chosen days: value sets it, percent scales the current dose (80 = 20% less, 120 = 20% more).", map[string]any{
			"exercise": str("Exercise id"), "days": daysProp, "value": num("New dose"), "percent": num("Scale, 10–300"),
		}, "exercise"),
		fn("set_workout", "Change this program's workout settings: rounds of the main part 1–10, difficulty level −3..+3 (each step ±10% of every main dose), rests in seconds (warm-up and cool-down 0–40, main 5–40). Only pass what changes.", map[string]any{
			"rounds": num("1–10"), "level": num("−3..+3"), "rest_warmup": num("0–40"), "rest_main": num("5–40"), "rest_cooldown": num("0–40"),
		}),
		fn("rename_program", "Rename the current program (up to 40 characters).", map[string]any{"name": str("New name")}, "name"),
		fn("create_exercise", "Create a new exercise in the user's library when nothing in search_exercises fits. Write the texts in the user's language. picture_from is the id of an existing exercise whose picture looks closest. Returns the new id; then add it with add_exercise.", map[string]any{
			"name":         str("Short name"),
			"part":         partProp,
			"unit":         map[string]any{"type": "string", "enum": []string{"reps", "seconds", "breaths"}},
			"per_side":     map[string]any{"type": "boolean", "description": "Done on each side in turn"},
			"difficulty":   num("1–5"),
			"description":  str("One or two sentences: what it is and what it is for"),
			"steps":        map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "3–5 short steps of the technique"},
			"hint":         str("One short cue said during the exercise"),
			"warning":      str("Safety note, if any"),
			"picture_from": str("Existing exercise id to borrow the picture from"),
		}, "name", "part", "unit", "description", "steps", "picture_from"),
	}
}

// args is a decoded tool call.
type args map[string]any

func (a args) str(k string) string {
	if v, ok := a[k].(string); ok {
		return strings.TrimSpace(v)
	}
	return ""
}

func (a args) num(k string) (int, bool) {
	switch v := a[k].(type) {
	case float64:
		return int(v), true
	case string:
		n, err := strconv.Atoi(strings.TrimSpace(v))
		return n, err == nil
	}
	return 0, false
}

func (a args) flag(k string) bool {
	v, _ := a[k].(bool)
	return v
}

// days reads "all", "today", 5, "4-30", "1,3,5" or [1,3,5].
func (a args) days(userID int64) (content.Days, error) {
	var list []int
	add := func(s string) error {
		s = strings.TrimSpace(s)
		if from, to, ok := strings.Cut(s, "-"); ok {
			f, err1 := strconv.Atoi(strings.TrimSpace(from))
			t, err2 := strconv.Atoi(strings.TrimSpace(to))
			if err1 != nil || err2 != nil || f > t {
				return fmt.Errorf("days: %q is not a range", s)
			}
			for d := f; d <= t; d++ {
				list = append(list, d)
			}
			return nil
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("days: %q is not a day", s)
		}
		list = append(list, n)
		return nil
	}
	switch v := a["days"].(type) {
	case nil:
		return content.AllDays(), nil
	case float64:
		return content.OneDay(int(v)), nil
	case string:
		s := strings.ToLower(strings.TrimSpace(v))
		switch s {
		case "", "all":
			return content.AllDays(), nil
		case "today":
			return content.OneDay(app.CurrentDay(userID)), nil
		}
		for _, part := range strings.Split(s, ",") {
			if err := add(part); err != nil {
				return content.Days{}, err
			}
		}
	case []any:
		for _, x := range v {
			switch n := x.(type) {
			case float64:
				list = append(list, int(n))
			case string:
				if err := add(n); err != nil {
					return content.Days{}, err
				}
			}
		}
	default:
		return content.Days{}, errors.New(`days: use "all", "today", a day, a range or a list`)
	}
	if len(list) == content.TotalDays {
		return content.AllDays(), nil
	}
	return content.Days{List: list}, nil
}

// run executes one tool call and returns the text the model sees.
func (s *session) run(name, rawArgs string) string {
	var a args
	if strings.TrimSpace(rawArgs) == "" {
		rawArgs = "{}"
	}
	if err := json.Unmarshal([]byte(rawArgs), &a); err != nil {
		return "error: arguments are not valid JSON"
	}
	out, err := s.call(name, a)
	if err != nil {
		return "error: " + err.Error()
	}
	return out
}

func (s *session) call(name string, a args) (string, error) {
	uid, lang := s.userID, s.lang
	switch name {
	case "search_exercises":
		list := app.Library(uid, lang, a.str("group"), a.str("query"), a.flag("no_jumps"))
		part := a.str("part")
		var b strings.Builder
		n := 0
		for _, e := range list {
			if part != "" && e.Slot != part {
				continue
			}
			if n == 40 {
				b.WriteString("… more: narrow the search\n")
				break
			}
			b.WriteString(exerciseLine(e))
			n++
		}
		if n == 0 {
			return "nothing found; try other words or a group", nil
		}
		return b.String(), nil
	case "similar_exercises":
		sim, err := app.SimilarExercises(uid, lang, a.str("id"))
		if err != nil {
			return "", err
		}
		var b strings.Builder
		for _, g := range []struct {
			title string
			list  []app.ExerciseInfo
		}{{"easier", sim.Easier}, {"similar", sim.Same}, {"harder", sim.Harder}} {
			b.WriteString(g.title + ":\n")
			for i, e := range g.list {
				if i == 12 {
					break
				}
				b.WriteString("  " + exerciseLine(e))
			}
		}
		return b.String(), nil
	case "show_program":
		return Overview(uid, lang), nil
	case "add_exercise":
		days, err := a.days(uid)
		if err != nil {
			return "", err
		}
		value, _ := a.num("value")
		if err := app.AddToCurrent(uid, a.str("exercise"), content.Slot(a.str("part")), days, value); err != nil {
			return "", err
		}
		s.edited = true
		return "done", nil
	case "replace_exercise":
		days, err := a.days(uid)
		if err != nil {
			return "", err
		}
		n, err := app.ReplaceInCurrent(uid, a.str("from"), a.str("to"), days)
		if err != nil {
			return "", err
		}
		s.edited = true
		return fmt.Sprintf("done: %d replaced", n), nil
	case "remove_exercise":
		days, err := a.days(uid)
		if err != nil {
			return "", err
		}
		n, err := app.RemoveFromCurrent(uid, a.str("exercise"), content.Slot(a.str("part")), days)
		if err != nil {
			return "", err
		}
		s.edited = true
		return fmt.Sprintf("done: %d removed", n), nil
	case "set_part":
		days, err := a.days(uid)
		if err != nil {
			return "", err
		}
		raw, _ := a["exercises"].([]any)
		var list []app.PartItem
		for _, x := range raw {
			m, _ := x.(map[string]any)
			id, _ := m["id"].(string)
			v, _ := args(m).num("value")
			list = append(list, app.PartItem{ID: strings.TrimSpace(id), Value: v})
		}
		if len(list) == 0 {
			return "", errors.New("exercises must list at least one exercise")
		}
		if err := app.SetPartInCurrent(uid, content.Slot(a.str("part")), days, list); err != nil {
			return "", err
		}
		s.edited = true
		return fmt.Sprintf("done: %d exercises set", len(list)), nil
	case "set_dose":
		days, err := a.days(uid)
		if err != nil {
			return "", err
		}
		value, _ := a.num("value")
		percent, _ := a.num("percent")
		if percent != 0 && (percent < 10 || percent > 300) {
			return "", errors.New("percent must be 10–300")
		}
		n, err := app.SetDoseInCurrent(uid, a.str("exercise"), days, value, percent)
		if err != nil {
			return "", err
		}
		s.edited = true
		return fmt.Sprintf("done: %d changed", n), nil
	case "set_workout":
		var p app.SettingsPatch
		set := func(k string, dst **int) {
			if v, ok := a.num(k); ok {
				*dst = &v
			}
		}
		set("rounds", &p.Rounds)
		set("level", &p.Level)
		set("rest_warmup", &p.RestWarmup)
		set("rest_main", &p.RestMain)
		set("rest_cooldown", &p.RestCooldown)
		if _, err := app.UpdateSettings(uid, p); err != nil {
			return "", err
		}
		s.edited = true
		return "done", nil
	case "rename_program":
		if err := app.RenameProgram(uid, a.str("name")); err != nil {
			return "", err
		}
		s.edited = true
		return "done", nil
	case "create_exercise":
		id, err := createExercise(uid, lang, a)
		if err != nil {
			return "", err
		}
		s.created = append(s.created, id)
		s.edited = true
		return "created with id " + id + "; now add it with add_exercise", nil
	}
	return "", fmt.Errorf("unknown tool %q", name)
}

func exerciseLine(e app.ExerciseInfo) string {
	var flags []string
	if e.PerSide {
		flags = append(flags, "per side")
	}
	if e.Jumps {
		flags = append(flags, "jumps")
	}
	if e.Equipment {
		flags = append(flags, "equipment")
	}
	line := fmt.Sprintf("%s | %s | %s | %s | %s | difficulty %d", e.ID, e.Name, e.Group, e.Slot, e.Unit, e.Difficulty)
	if len(flags) > 0 {
		line += " | " + strings.Join(flags, ", ")
	}
	return line + "\n"
}

// createExercise writes a new exercise into the user's library.
func createExercise(userID int64, lang string, a args) (string, error) {
	name := a.str("name")
	if name == "" {
		return "", errors.New("name is required")
	}
	part, unit := a.str("part"), a.str("unit")
	cat := app.UserCatalog(userID)
	from, ok := cat.Doc(a.str("picture_from"))
	if !ok {
		return "", errors.New("picture_from must be an existing exercise id")
	}
	difficulty, _ := a.num("difficulty")
	difficulty = max(1, min(difficulty, 5))
	var steps []string
	if list, ok := a["steps"].([]any); ok {
		for _, x := range list {
			if s, ok := x.(string); ok && strings.TrimSpace(s) != "" {
				steps = append(steps, strings.TrimSpace(s))
			}
		}
	}
	id := "U" + strings.ToUpper(strconv.FormatInt(time.Now().UnixMilli(), 36))
	d := catalog.Doc{
		ID: id, Slot: part, Unit: unit, PerSide: a.flag("per_side"), Difficulty: difficulty,
		SVG: from.SVG, Poses: from.Poses, Animation: from.Animation,
		Content: map[string]catalog.Loc{string(i18n.Parse(lang)): {
			Name: name, Desc: a.str("description"), HowTo: steps, Hint: a.str("hint"), Warning: a.str("warning"),
		}},
	}
	if err := app.SaveCustomExercise(userID, d); err != nil {
		return "", err
	}
	return id, nil
}
