package mcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"morning30/internal/app"
	"morning30/internal/content"
	"morning30/internal/db"
)

// Operations of the local stdio MCP server (cmd/morning-mcp). There is no
// network API: inside the app, programs are edited by the built-in AI coach.

var accountSchema = map[string]any{"type": "string", "description": "User id, email, or unique display name"}

var daysSchema = map[string]any{
	"description": `Which days: "today" (the user's next day), "all" (default), a day number 1–30, or a list of day numbers`,
}

var partSchema = map[string]any{"type": "string", "enum": []string{"warmup", "main", "cooldown"}}

type op struct {
	Name    string
	Desc    string
	Props   map[string]any // input properties besides "account"
	Require []string       // required inputs besides "account"
	Account bool           // works on one account
	Run     func(acc db.Account, a args) (any, error)
}

// args is a tool call's input: a JSON object, plus path and query values on HTTP.
type args map[string]any

func (a args) str(k string) string {
	switch v := a[k].(type) {
	case string:
		return strings.TrimSpace(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case json.Number:
		return v.String()
	}
	return ""
}

func (a args) int(k string) int {
	switch v := a[k].(type) {
	case float64:
		return int(v)
	case string:
		n, _ := strconv.Atoi(strings.TrimSpace(v))
		return n
	}
	return 0
}

func (a args) bool(k string) bool {
	switch v := a[k].(type) {
	case bool:
		return v
	case string:
		b, _ := strconv.ParseBool(strings.TrimSpace(v))
		return b || strings.TrimSpace(v) == "1"
	case float64:
		return v != 0
	}
	return false
}

func (a args) raw(k string) json.RawMessage {
	v, ok := a[k]
	if !ok {
		return nil
	}
	if s, ok := v.(string); ok {
		return json.RawMessage(s)
	}
	b, _ := json.Marshal(v)
	return b
}

// days reads a day selection; missing means every day.
func (a args) days(userID int64) (content.Days, error) {
	switch v := a["days"].(type) {
	case nil:
		return content.AllDays(), nil
	case string:
		s := strings.ToLower(strings.TrimSpace(v))
		switch s {
		case "", "all":
			return content.AllDays(), nil
		case "today":
			return content.OneDay(app.CurrentDay(userID)), nil
		}
		var list []int
		for _, part := range strings.Split(s, ",") {
			n, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return content.Days{}, fmt.Errorf("days: %q is not a day", part)
			}
			list = append(list, n)
		}
		return content.Days{List: list}, nil
	case float64:
		return content.OneDay(int(v)), nil
	case []any:
		var list []int
		for _, x := range v {
			n, ok := x.(float64)
			if !ok {
				return content.Days{}, errors.New("days: list must hold day numbers")
			}
			list = append(list, int(n))
		}
		return content.Days{List: list}, nil
	}
	return content.Days{}, errors.New(`days: use "today", "all", a day number or a list`)
}

func (a args) part() (content.Slot, error) {
	p := content.Slot(a.str("part"))
	switch p {
	case content.Warmup, content.Main, content.Cooldown:
		return p, nil
	}
	return "", content.ErrBadPart
}

func langOr(a args, userID int64) string {
	if l := a.str("lang"); l != "" {
		return l
	}
	return app.LangOf(userID)
}

// readyView describes a catalog program; Program is set for a single one.
type readyView struct {
	Key      string            `json:"key"`
	Slug     string            `json:"slug,omitempty"`
	Category string            `json:"category"`
	Kids     bool              `json:"kids,omitempty"`
	Titles   map[string]string `json:"titles"`
	About    map[string]string `json:"about"`
	Tags     content.Tags      `json:"tags"`
	Minutes  [2]int            `json:"minutes"` // day 1 and peak day
	Days     int               `json:"days"`
	Program  *content.Resolved `json:"program,omitempty"`
}

func readyOf(p content.ReadyProgram, full bool) readyView {
	r := p.Build()
	first, peak := content.MinutesSpan(r)
	v := readyView{Key: p.Key, Slug: p.Slug, Category: p.Category, Kids: p.Kids, Titles: p.Title, About: p.About,
		Tags: content.ProgramTags(r), Minutes: [2]int{first, peak}, Days: content.TotalDays}
	if full {
		v.Program = &r
	}
	return v
}

func catalogList() []readyView {
	out := []readyView{}
	for _, p := range content.Catalog() {
		out = append(out, readyOf(p, false))
	}
	return out
}

func readyDetail(key string) (readyView, error) {
	p, ok := content.ReadyByKey(strings.TrimSpace(key))
	if !ok {
		return readyView{}, app.ErrUnknownCourse
	}
	return readyOf(p, true), nil
}

// currentProgram is what program-changing operations return.
func currentProgram(userID int64) (programView, error) {
	acc, err := app.DB.AccountByID(userID)
	if err != nil {
		return programView{}, err
	}
	return programView{Account: acc, Personal: app.ActiveProgramKey(userID) == "custom", Program: app.UserProgram(userID),
		Day: app.CurrentDay(userID), Selected: app.SelectedPrograms(userID)}, nil
}

func then(userID int64, err error) (any, error) {
	if err != nil {
		return nil, err
	}
	return currentProgram(userID)
}

// selectArgs picks by selected-program id or by catalog key.
func selectProgram(acc db.Account, a args) (any, error) {
	if id := int64(a.int("id")); id != 0 {
		return then(acc.ID, app.SelectByID(acc.ID, id))
	}
	key := a.str("key")
	if key == "" {
		key = a.str("course")
	}
	if key == "" {
		return nil, errors.New("give key (a catalog program) or id (a selected program)")
	}
	return then(acc.ID, app.SelectReady(acc.ID, key, langOr(a, acc.ID)))
}

func programRaw(a args) (string, error) {
	return programJSON(a.raw("program"))
}

var ops = []op{
	{
		Name: "list_programs",
		Desc: "List every ready program of the catalog: key, category (start, strength, weightloss, stretch, backbalance, kids, sport), titles and one-line descriptions in every supported language, tags (jumps, equipment) and minutes per session on day 1 and on the peak day.",
		Run:  func(db.Account, args) (any, error) { return catalogList(), nil },
	},
	{
		Name: "get_ready_program",
		Desc:  "Get one ready program with its full 30-day plan (what the program preview shows).",
		Props: map[string]any{"key": map[string]any{"type": "string", "description": "Key from list_programs"}}, Require: []string{"key"},
		Run: func(_ db.Account, a args) (any, error) {
			key := a.str("key")
			if key == "" {
				key = a.str("course")
			}
			return readyDetail(key)
		},
	},
	{
		Name: "list_selected", Account: true,
		Desc: "List the account's selected programs: id, name, source catalog key (empty for the user's own), its own next day, whether it is current, and whether it holds the user's edits.",
		Run:  func(acc db.Account, _ args) (any, error) { return app.SelectedPrograms(acc.ID), nil },
	},
	{
		Name: "select_program", Account: true,
		Desc: "Switch the account to a program: a catalog key (it is added to the selected programs; if the user already has it, their own version and day come back) or the id of a selected program. Each program keeps its own day; the streak is common. Nothing is lost.",
		Props: map[string]any{
			"key":  map[string]any{"type": "string", "description": "Catalog key from list_programs"},
			"id":   map[string]any{"type": "integer", "description": "Selected program id from list_selected"},
			"lang": map[string]any{"type": "string", "description": "Language of the name of a newly added program (default: the account's)"},
		},
		Run: selectProgram,
	},
	{
		Name: "star_program", Account: true,
		Desc:  "Add a catalog program to the selected programs without switching to it (the star in the catalog).",
		Props: map[string]any{"key": map[string]any{"type": "string"}, "lang": map[string]any{"type": "string"}}, Require: []string{"key"},
		Run: func(acc db.Account, a args) (any, error) {
			if _, err := app.StarReady(acc.ID, a.str("key"), langOr(a, acc.ID)); err != nil {
				return nil, err
			}
			return app.SelectedPrograms(acc.ID), nil
		},
	},
	{
		Name: "unselect_program", Account: true,
		Desc: "Remove a program from the selected programs, by id or catalog key. A program holding the user's edits needs discardEdits=true, because the edits cannot be brought back. Removing the current program switches to another selected one, or to the morning workout.",
		Props: map[string]any{
			"id":           map[string]any{"type": "integer"},
			"key":          map[string]any{"type": "string"},
			"discardEdits": map[string]any{"type": "boolean"},
		},
		Run: func(acc db.Account, a args) (any, error) {
			var err error
			if id := int64(a.int("id")); id != 0 {
				err = app.Unselect(acc.ID, id, a.bool("discardEdits"))
			} else {
				err = app.UnstarReady(acc.ID, a.str("key"), a.bool("discardEdits"))
			}
			return then(acc.ID, err)
		},
	},
	{
		Name: "rename_program", Account: true,
		Desc:  "Rename the current program (up to 40 characters). Exercises, day and streak are unchanged.",
		Props: map[string]any{"name": map[string]any{"type": "string"}}, Require: []string{"name"},
		Run: func(acc db.Account, a args) (any, error) {
			return then(acc.ID, app.RenameProgram(acc.ID, a.str("name")))
		},
	},
	{
		Name: "set_program", Account: true,
		Desc:  "Save a new version of the CURRENT program from a full 30-day program JSON (get_program, change, set_program). The day and streak stay. To add a different program instead, use import_program.",
		Props: map[string]any{"program": map[string]any{"description": "Resolved program object or JSON string"}}, Require: []string{"program"},
		Run: func(acc db.Account, a args) (any, error) {
			raw, err := programRaw(a)
			if err != nil {
				return nil, err
			}
			return then(acc.ID, app.SetUserProgram(acc.ID, raw))
		},
	},
	{
		Name: "import_program", Account: true,
		Desc:  "Add a program from JSON as a new selected program and switch to it at day 1. Every program the user had stays selected.",
		Props: map[string]any{"program": map[string]any{"description": "Resolved program object or JSON string"}}, Require: []string{"program"},
		Run: func(acc db.Account, a args) (any, error) {
			raw, err := programRaw(a)
			if err != nil {
				return nil, err
			}
			return then(acc.ID, app.ImportProgram(acc.ID, raw))
		},
	},
	{
		Name: "restore_demo", Account: true,
		Desc: "Switch to the morning workout. Every selected program, with its edits and day, stays.",
		Run:  func(acc db.Account, _ args) (any, error) { return then(acc.ID, app.RestoreDemo(acc.ID)) },
	},
	{
		Name: "get_day", Account: true,
		Desc:  "The exercises of a day of the current program with the doses the user will do (warm-up, the main part listed once, cool-down). Without day: the user's next day.",
		Props: map[string]any{"day": map[string]any{"type": "integer", "description": "1–30"}, "lang": map[string]any{"type": "string"}},
		Run: func(acc db.Account, a args) (any, error) {
			return app.Day(acc.ID, langOr(a, acc.ID), a.int("day")), nil
		},
	},
	{
		Name: "add_exercise", Account: true,
		Desc: "Add an exercise to a part of the current program on the chosen days (the library's \"+\"). value is reps, seconds or breaths for the exercise's unit; omitted uses a sensible default.",
		Props: map[string]any{
			"exercise": map[string]any{"type": "string", "description": "Exercise id from list_exercises"},
			"part":     partSchema, "days": daysSchema, "value": map[string]any{"type": "integer"},
		},
		Require: []string{"exercise", "part"},
		Run: func(acc db.Account, a args) (any, error) {
			part, err := a.part()
			if err != nil {
				return nil, err
			}
			days, err := a.days(acc.ID)
			if err != nil {
				return nil, err
			}
			return then(acc.ID, app.AddToCurrent(acc.ID, a.str("exercise"), part, days, a.int("value")))
		},
	},
	{
		Name: "replace_exercise", Account: true,
		Desc: "Replace an exercise with another on the chosen days of the current program; the dose carries over (converted if the unit differs). See similar_exercises for easier, similar and harder options.",
		Props: map[string]any{
			"from": map[string]any{"type": "string"}, "to": map[string]any{"type": "string"}, "days": daysSchema,
		},
		Require: []string{"from", "to"},
		Run: func(acc db.Account, a args) (any, error) {
			days, err := a.days(acc.ID)
			if err != nil {
				return nil, err
			}
			_, err = app.ReplaceInCurrent(acc.ID, a.str("from"), a.str("to"), days)
			return then(acc.ID, err)
		},
	},
	{
		Name: "remove_exercise", Account: true,
		Desc: "Remove an exercise from the chosen days of the current program (part optional: every part). A part of a day is never left empty.",
		Props: map[string]any{
			"exercise": map[string]any{"type": "string"}, "part": partSchema, "days": daysSchema,
		},
		Require: []string{"exercise"},
		Run: func(acc db.Account, a args) (any, error) {
			var part content.Slot
			if a.str("part") != "" {
				p, err := a.part()
				if err != nil {
					return nil, err
				}
				part = p
			}
			days, err := a.days(acc.ID)
			if err != nil {
				return nil, err
			}
			_, err = app.RemoveFromCurrent(acc.ID, a.str("exercise"), part, days)
			return then(acc.ID, err)
		},
	},
	{
		Name: "list_exercises", Account: true,
		Desc: "The exercise library this account can use (built-in plus custom): id, name, group (warmup, core, push, back, arms, legs, cardio, stretch, dumbbell, custom), part of the day, unit, per side, difficulty 1–5, jumps, equipment.",
		Props: map[string]any{
			"group": map[string]any{"type": "string"}, "query": map[string]any{"type": "string", "description": "Part of a name or id"},
			"noJumps": map[string]any{"type": "boolean"}, "lang": map[string]any{"type": "string"},
		},
		Run: func(acc db.Account, a args) (any, error) {
			list := app.Library(acc.ID, langOr(a, acc.ID), a.str("group"), a.str("query"), a.bool("noJumps"))
			if list == nil {
				list = []app.ExerciseInfo{}
			}
			return list, nil
		},
	},
	{
		Name: "similar_exercises", Account: true,
		Desc:  "Exercises that can replace one: same group and part of the day, split into easier, same and harder.",
		Props: map[string]any{"id": map[string]any{"type": "string"}, "lang": map[string]any{"type": "string"}}, Require: []string{"id"},
		Run: func(acc db.Account, a args) (any, error) {
			return app.SimilarExercises(acc.ID, langOr(a, acc.ID), a.str("id"))
		},
	},
	{
		Name: "get_settings", Account: true,
		Desc: "The account's settings: name, lang, the three rests in seconds, voiceMode, figure (m man, f woman, b boy, g girl), difficulty level −3..+3 and main rounds 1..10.",
		Run:  func(acc db.Account, _ args) (any, error) { return app.GetSettings(acc.ID) },
	},
	{
		Name: "set_settings", Account: true,
		Desc: "Change only the given settings. Rests: warm-up and cool-down 0–40 s, main 5–40 s. voiceMode off|min|normal|detailed. figure m|f|b|g. level −3..+3. rounds 1..10 (how many times the main part repeats, for every program).",
		Props: map[string]any{
			"name": map[string]any{"type": "string"}, "lang": map[string]any{"type": "string"},
			"restWarmup": map[string]any{"type": "integer"}, "restMain": map[string]any{"type": "integer"}, "restCooldown": map[string]any{"type": "integer"},
			"voiceMode": map[string]any{"type": "string"}, "figure": map[string]any{"type": "string"}, "level": map[string]any{"type": "integer"},
			"rounds": map[string]any{"type": "integer"},
		},
		Run: func(acc db.Account, a args) (any, error) {
			var p app.SettingsPatch
			if err := json.Unmarshal(a.raw("_self"), &p); err != nil {
				return nil, err
			}
			return app.UpdateSettings(acc.ID, p)
		},
	},
	// Earlier names, kept so existing agents keep working.
	{
		Name: "use_course", Account: true, Desc: "Older name of select_program with key (course).",
		Props: map[string]any{"course": map[string]any{"type": "string"}, "lang": map[string]any{"type": "string"}}, Require: []string{"course"},
		Run: selectProgram,
	},
	{
		Name: "list_courses", Desc: "Older name of list_programs.",
		Run: func(db.Account, args) (any, error) { return catalogList(), nil },
	},
	{
		Name: "get_course", Desc: "Older name of get_ready_program (course = key).",
		Props: map[string]any{"course": map[string]any{"type": "string"}}, Require: []string{"course"},
		Run: func(_ db.Account, a args) (any, error) { return readyDetail(a.str("course")) },
	},
	{
		Name: "list_saved_programs", Account: true, Desc: "Older name of list_selected.",
		Run: func(acc db.Account, _ args) (any, error) { return app.SelectedPrograms(acc.ID), nil },
	},
	{
		Name: "restore_saved_program", Account: true, Desc: "Older name of select_program with id.",
		Props: map[string]any{"id": map[string]any{"type": "integer"}}, Require: []string{"id"},
		Run: selectProgram,
	},
}

func opByName(name string) (op, bool) {
	for _, o := range ops {
		if o.Name == name {
			return o, true
		}
	}
	return op{}, false
}

func opTools() []map[string]any {
	out := make([]map[string]any, 0, len(ops))
	for _, o := range ops {
		props := map[string]any{}
		required := []string{}
		if o.Account {
			props["account"] = accountSchema
			required = append(required, "account")
		}
		for k, v := range o.Props {
			props[k] = v
		}
		required = append(required, o.Require...)
		schema := map[string]any{"type": "object", "properties": props}
		if len(required) > 0 {
			schema["required"] = required
		}
		out = append(out, map[string]any{"name": o.Name, "description": o.Desc, "inputSchema": schema})
	}
	return out
}

func parseArgs(raw []byte) (args, error) {
	a := args{}
	if len(strings.TrimSpace(string(raw))) == 0 {
		return a, nil
	}
	if err := json.Unmarshal(raw, &a); err != nil {
		return nil, err
	}
	a["_self"] = json.RawMessage(raw)
	return a, nil
}

// callStdio runs an operation for the stdio server, which addresses any account.
func (o op) callStdio(raw json.RawMessage) (any, error) {
	a, err := parseArgs(raw)
	if err != nil {
		return nil, err
	}
	var acc db.Account
	if o.Account {
		acc, err = app.DB.FindAccount(a.str("account"))
		if err != nil {
			return nil, err
		}
	}
	return o.Run(acc, a)
}
