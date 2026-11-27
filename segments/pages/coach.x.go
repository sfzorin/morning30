// Managed by GoX v0.1.36

//line coach.gox:1
package pages

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"morning30/assets"
	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/coach"
	"morning30/internal/db"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// The AI coach: a chat on the program page that edits the current program.

type coachPage struct {
	sess auth.Session
	path doors.Source[path.Path]
	busy *atomic.Bool // a reply is in progress on this page
}

type coachForm struct {
	Text string `form:"text"`
}

// coachLine is one message as the screen shows it.
type coachLine struct {
	ID      int64
	Mine    bool
	Text    string
	Changes []coach.Change
	CanUndo bool
	Undone  bool
}

// coachView is the whole screen state in one comparable-by-replacement value.
type coachView struct {
	Lines   []coachLine
	Pending string // the message being answered
	Err     string // i18n key or ready text of the last problem
}

const (
	coachMaxText = 600            // the longest message the coach accepts, in characters
	coachKeep    = 50             // messages kept per program's conversation
	coachMemory  = 24 * time.Hour // how far back the model sees the conversation
)

func coachProgramID(userID int64) int64 { return app.DB.CurrentSelected(userID) }

func loadCoach(userID int64) coachView {
	rows, err := app.DB.CoachMessages(userID, coachProgramID(userID), coachKeep)
	if err != nil {
		log.Printf("coach history %d: %v", userID, err)
	}
	var v coachView
	lastEdit := -1
	for i, m := range rows {
		line := coachLine{ID: m.ID, Mine: m.Role == "user", Text: m.Text, Undone: m.Undone}
		if m.Changes != "" {
			_ = json.Unmarshal([]byte(m.Changes), &line.Changes)
		}
		if m.Undo != "" && len(line.Changes) > 0 && !m.Undone {
			lastEdit = i
		}
		v.Lines = append(v.Lines, line)
	}
	// Only the latest edit can be undone: earlier ones would undo later ones too.
	if lastEdit >= 0 {
		v.Lines[lastEdit].CanUndo = true
		for _, l := range v.Lines[lastEdit+1:] {
			if !l.Mine && len(l.Changes) > 0 {
				v.Lines[lastEdit].CanUndo = false
			}
		}
	}
	return v
}

//line coach.gox:90
func (c coachPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line coach.gox:92
		c.busy = new(atomic.Bool)
		l := i18n.Lang(c.sess.Lang)
		view := doors.NewSourceEqual(loadCoach(c.sess.UserID), func(coachView, coachView) bool { return false })
		enabled := coach.FromEnv().Enabled()
		examples := []string{i18n.T(l, "coach.ex1"), i18n.T(l, "coach.ex2"), i18n.T(l, "coach.ex3")}

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line coach.gox:98
			__e = __c.Any(i18n.T(l, "coach.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line coach.gox:99
			__e = __c.Set("class", "screen coach"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line coach.gox:100
				__e = __c.Set("class", "lib-head"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:101
				__e = (doors.ALink{Model: path.Path{Page: path.Program}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line coach.gox:101
						__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line coach.gox:101
						__e = __c.Set("aria-label", i18n.T(l, "workout.prev")); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line coach.gox:101
						__e = __c.Any(icon("back", false)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h1"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line coach.gox:103
						__e = __c.Any(i18n.T(l, "coach.title")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line coach.gox:104
						__e = __c.Set("class", "coach-sub"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line coach.gox:104
						__e = __c.Any(app.ProgramTitle(c.sess.UserID, c.sess.Lang)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:107
			__e = __c.Any(view.Bind(func(v coachView) gox.Elem {
			return gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("div"); if __e != nil { return }
				{
//line coach.gox:108
					__e = __c.Set("class", "coach-thread"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line coach.gox:109
					if len(v.Lines) == 0 && v.Pending == "" {
						__e = __c.Init("p"); if __e != nil { return }
						{
//line coach.gox:110
							__e = __c.Set("class", "coach-lead"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line coach.gox:110
							__e = __c.Any(i18n.T(l, "coach.lead")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line coach.gox:111
						if enabled {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line coach.gox:112
								__e = __c.Set("class", "coach-examples"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line coach.gox:113
								for _, ex := range examples {
//line coach.gox:114
									__e = (doors.AClick{On: c.sendText(view, ex)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
										__e = __c.Init("button"); if __e != nil { return }
										{
//line coach.gox:114
											__e = __c.Set("class", "chip-btn"); if __e != nil { return }
//line coach.gox:114
											__e = __c.Set("type", "button"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line coach.gox:114
											__e = __c.Any(ex); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									return })); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
//line coach.gox:119
					for _, line := range v.Lines {
//line coach.gox:120
						if line.Mine {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line coach.gox:121
								__e = __c.Set("class", "bubble me"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line coach.gox:121
								__e = __c.Any(line.Text); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						} else  {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line coach.gox:123
								__e = __c.Set("class", "bubble ai"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("p"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line coach.gox:124
									__e = __c.Any(line.Text); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line coach.gox:125
								if len(line.Changes) > 0 {
									__e = __c.Init("ul"); if __e != nil { return }
									{
//line coach.gox:126
										__e = __c.Set("class", coachChangesClass(line.Undone)); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line coach.gox:127
										for _, ch := range line.Changes {
											__e = __c.Init("li"); if __e != nil { return }
											{
												__e = __c.Submit(); if __e != nil { return }
//line coach.gox:128
												__e = __c.Any(coachChange(l, ch)); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
									}
									__e = __c.Close(); if __e != nil { return }
								}
//line coach.gox:132
								if line.CanUndo {
//line coach.gox:133
									__e = (doors.AClick{On: c.undo(view, line.ID)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
										__e = __c.Init("button"); if __e != nil { return }
										{
//line coach.gox:133
											__e = __c.Set("class", "coach-undo"); if __e != nil { return }
//line coach.gox:133
											__e = __c.Set("type", "button"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line coach.gox:133
											__e = __c.Any(i18n.T(l, "coach.undo")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									return })); if __e != nil { return }
//line coach.gox:134
								} else if line.Undone {
									__e = __c.Init("p"); if __e != nil { return }
									{
//line coach.gox:135
										__e = __c.Set("class", "coach-undone"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line coach.gox:135
										__e = __c.Any(i18n.T(l, "coach.undone")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
//line coach.gox:140
					if v.Pending != "" {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line coach.gox:141
							__e = __c.Set("class", "bubble me"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line coach.gox:141
							__e = __c.Any(v.Pending); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line coach.gox:142
							__e = __c.Set("class", "bubble ai typing"); if __e != nil { return }
//line coach.gox:142
							__e = __c.Set("role", "status"); if __e != nil { return }
//line coach.gox:142
							__e = __c.Set("aria-label", i18n.T(l, "coach.thinking")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
//line coach.gox:144
					if v.Err != "" {
						__e = __c.Init("p"); if __e != nil { return }
						{
//line coach.gox:145
							__e = __c.Set("class", "err"); if __e != nil { return }
//line coach.gox:145
							__e = __c.Set("role", "alert"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line coach.gox:145
							__e = __c.Any(coachErr(l, v.Err)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Init("script"); if __e != nil { return }
					{
//line coach.gox:147
						__e = __c.Set("src", assets.CoachScroll); if __e != nil { return }
						__e = __c.Set("inline", true); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Raw(""); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })
//line coach.gox:149
		})); if __e != nil { return }
//line coach.gox:150
			if enabled {
//line coach.gox:151
				__e = (doors.ASubmit[coachForm]{On: c.submit(view)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("form"); if __e != nil { return }
					{
//line coach.gox:151
						__e = __c.Set("class", "coach-input"); if __e != nil { return }
//line coach.gox:151
						__e = __c.Set("autocomplete", "off"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.InitVoid("input"); if __e != nil { return }
						{
//line coach.gox:152
							__e = __c.Set("class", "text-input"); if __e != nil { return }
//line coach.gox:152
							__e = __c.Set("name", "text"); if __e != nil { return }
//line coach.gox:152
							__e = __c.Set("type", "text"); if __e != nil { return }
//line coach.gox:152
							__e = __c.Set("maxlength", coachMaxText); if __e != nil { return }
//line coach.gox:152
							__e = __c.Set("placeholder", i18n.T(l, "coach.placeholder")); if __e != nil { return }
//line coach.gox:152
							__e = __c.Set("aria-label", i18n.T(l, "coach.placeholder")); if __e != nil { return }
						}
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("button"); if __e != nil { return }
						{
//line coach.gox:153
							__e = __c.Set("class", "coach-send"); if __e != nil { return }
//line coach.gox:153
							__e = __c.Set("type", "submit"); if __e != nil { return }
//line coach.gox:153
							__e = __c.Set("aria-label", i18n.T(l, "coach.send")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line coach.gox:153
							__e = __c.Any(icon("up", false)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("script"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Raw("// The back arrow returns to where the coach was opened from (home or\n\t\t\t\t// the program page); opened directly, it goes to the program page.\n\t\t\t\tif (!window.__coachBack) {\n\t\t\t\t\twindow.__coachBack = true\n\t\t\t\t\twindow.addEventListener(\"click\", (e) => {\n\t\t\t\t\t\tif (e.target.closest && e.target.closest(\".coach .lib-back\") && history.length > 1) {\n\t\t\t\t\t\t\te.preventDefault()\n\t\t\t\t\t\t\te.stopImmediatePropagation()\n\t\t\t\t\t\t\thistory.back()\n\t\t\t\t\t\t}\n\t\t\t\t\t}, true)\n\t\t\t\t}\n\t\t\t\t$on(\"coach-sent\", () => {\n\t\t\t\t\tconst field = document.querySelector(\".coach-input input[name=text]\")\n\t\t\t\t\tif (field) field.value = \"\"\n\t\t\t\t})"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			} else  {
				__e = __c.Init("p"); if __e != nil { return }
				{
//line coach.gox:174
					__e = __c.Set("class", "coach-lead"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line coach.gox:174
					__e = __c.Any(i18n.T(l, "coach.off")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line coach.gox:177
}

func coachChangesClass(undone bool) string {
	if undone {
		return "coach-changes undone"
	}
	return "coach-changes"
}

func coachErr(l i18n.Lang, key string) string {
	if key == "coach.limit" {
		return i18n.Tf(l, key, coach.DailyLimit)
	}
	return i18n.T(l, key)
}

// coachDays writes the days of a change: "All days", "Day 5", "Days 4–30".
func coachDays(l i18n.Lang, days []int) string {
	switch {
	case len(days) == 0:
		return ""
	case len(days) >= 30:
		return i18n.T(l, "coach.all_days")
	case len(days) == 1:
		return i18n.Tf(l, "fmt.day", days[0])
	}
	return i18n.Tf(l, "coach.days", coach.DayRanges(days))
}

func coachPart(l i18n.Lang, part string) string {
	switch part {
	case "warmup":
		return i18n.T(l, "workout.warmup")
	case "cooldown":
		return i18n.T(l, "workout.cooldown")
	case "main":
		return i18n.T(l, "workout.main")
	}
	return ""
}

//line coach.gox:218
func coachChange(l i18n.Lang, ch coach.Change) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line coach.gox:220
		days := coachDays(l, ch.Days)
		prefix := days
		if part := coachPart(l, ch.Part); part != "" {
			prefix = days + ", " + strings.ToLower(part)
		}

//line coach.gox:226
		if ch.Kind == "replace" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:227
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:227
				__e = __c.Any(prefix + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:227
			__e = __c.Any(ch.From); if __e != nil { return }
//line coach.gox:227
			__e = __c.Any(icon("to", false)); if __e != nil { return }
//line coach.gox:227
			__e = __c.Any(ch.To); if __e != nil { return }
//line coach.gox:228
		} else if ch.Kind == "add" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:229
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:229
				__e = __c.Any(prefix + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:229
			__e = __c.Any(i18n.Tf(l, "coach.added", ch.To)); if __e != nil { return }
//line coach.gox:230
		} else if ch.Kind == "remove" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:231
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:231
				__e = __c.Any(prefix + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:231
			__e = __c.Any(i18n.Tf(l, "coach.removed", ch.To)); if __e != nil { return }
//line coach.gox:232
		} else if ch.Kind == "dose" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:233
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:233
				__e = __c.Any(days + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:233
			__e = __c.Any(ch.To); if __e != nil { return }
//line coach.gox:234
			if ch.Pct != 0 {
//line coach.gox:235
				__e = __c.Many(" ", fmt.Sprintf("%+d%%", ch.Pct)); if __e != nil { return }
			} else  {
//line coach.gox:237
				__e = __c.Many(" ", ch.Old); if __e != nil { return }
//line coach.gox:237
				__e = __c.Any(icon("to", false)); if __e != nil { return }
//line coach.gox:237
				__e = __c.Any(ch.New); if __e != nil { return }
			}
//line coach.gox:239
		} else if ch.Kind == "new" {
//line coach.gox:240
			__e = __c.Any(i18n.Tf(l, "coach.new", ch.To)); if __e != nil { return }
//line coach.gox:241
		} else if ch.Kind == "rename" {
//line coach.gox:242
			__e = __c.Any(i18n.Tf(l, "coach.renamed", ch.To)); if __e != nil { return }
//line coach.gox:243
		} else if ch.Kind == "rounds" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:244
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:244
				__e = __c.Any(i18n.T(l, "settings.rounds") + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:244
			__e = __c.Any(ch.Old); if __e != nil { return }
//line coach.gox:244
			__e = __c.Any(icon("to", false)); if __e != nil { return }
//line coach.gox:244
			__e = __c.Any(ch.New); if __e != nil { return }
//line coach.gox:245
		} else if ch.Kind == "level" {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:246
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:246
				__e = __c.Any(i18n.T(l, "settings.difficulty") + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:246
			__e = __c.Any(levelLabel(l, ch.Old)); if __e != nil { return }
//line coach.gox:246
			__e = __c.Any(icon("to", false)); if __e != nil { return }
//line coach.gox:246
			__e = __c.Any(levelLabel(l, ch.New)); if __e != nil { return }
		} else  {
			__e = __c.Init("span"); if __e != nil { return }
			{
//line coach.gox:248
				__e = __c.Set("class", "muted"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line coach.gox:248
				__e = __c.Any(i18n.T(l, "settings."+ch.Kind) + ": "); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line coach.gox:248
			__e = __c.Any(i18n.Tf(l, "fmt.sec", ch.Old)); if __e != nil { return }
//line coach.gox:248
			__e = __c.Any(icon("to", false)); if __e != nil { return }
//line coach.gox:248
			__e = __c.Any(i18n.Tf(l, "fmt.sec", ch.New)); if __e != nil { return }
		}
	return })
//line coach.gox:250
}

func (c coachPage) submit(view doors.Source[coachView]) func(context.Context, doors.RequestForm[coachForm]) bool {
	return func(ctx context.Context, r doors.RequestForm[coachForm]) bool {
		c.send(ctx, view, r.Data().Text)
		r.After(doors.ActionEmit{Name: "coach-sent"}) // empties the field (see the script by the form)
		return false
	}
}

func (c coachPage) sendText(view doors.Source[coachView], text string) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, _ doors.RequestPointer) bool {
		c.send(ctx, view, text)
		return false
	}
}

// send shows the message at once and lets the coach answer in the background.
func (c coachPage) send(ctx context.Context, view doors.Source[coachView], text string) {
	text = strings.TrimSpace(text)
	if r := []rune(text); len(r) > coachMaxText {
		text = string(r[:coachMaxText])
	}
	v := loadCoach(c.sess.UserID)
	if text == "" {
		return
	}
	if !c.busy.CompareAndSwap(false, true) {
		return // one message at a time
	}
	if coach.LimitReached(c.sess.UserID) {
		c.busy.Store(false)
		v.Err = "coach.limit"
		view.Update(ctx, v)
		return
	}
	v.Pending = text
	view.Update(ctx, v)
	uid, lang := c.sess.UserID, c.sess.Lang
	bg := doors.InstanceContext(ctx)
	go func() {
		defer c.busy.Store(false)
		before := coachProgramID(uid)
		// Only today's conversation goes to the model; the program itself is
		// always sent in full.
		history, _ := app.DB.CoachRecent(uid, before, time.Now().Add(-coachMemory), 10)
		askCtx, cancel := context.WithTimeout(bg, 2*time.Minute)
		res, snap, err := coach.Ask(askCtx, coach.FromEnv(), uid, lang, history, text)
		cancel()
		if err != nil {
			log.Printf("coach %d: %v", uid, err)
			next := loadCoach(uid)
			next.Err = "coach.error"
			if err == coach.ErrDisabled {
				next.Err = "coach.off"
			}
			<-view.XUpdate(bg, next)
			return
		}
		program := coachProgramID(uid)
		reply := res.Reply
		if reply == "" {
			reply = i18n.T(i18n.Lang(lang), "coach.done")
		}
		changes, undo := "", ""
		if len(res.Changes) > 0 {
			b, _ := json.Marshal(res.Changes)
			changes, undo = string(b), coach.SnapshotJSON(snap)
		}
		if _, err := app.DB.AddCoachMessage(uid, db.CoachMessage{ProgramID: program, Role: "user", Text: text}); err != nil {
			log.Printf("coach save %d: %v", uid, err)
		}
		if _, err := app.DB.AddCoachMessage(uid, db.CoachMessage{ProgramID: program, Role: "coach", Text: reply, Changes: changes, Undo: undo}); err != nil {
			log.Printf("coach save %d: %v", uid, err)
		}
		if err := app.DB.PruneCoach(uid, program, coachKeep); err != nil {
			log.Printf("coach prune %d: %v", uid, err)
		}
		<-view.XUpdate(bg, loadCoach(uid))
	}()
}

// undo brings the program back to how it was before a reply.
func (c coachPage) undo(view doors.Source[coachView], id int64) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, _ doors.RequestPointer) bool {
		m, err := app.DB.CoachMessage(c.sess.UserID, id)
		if err == nil && !m.Undone && m.Undo != "" {
			if err = coach.Undo(c.sess.UserID, m.Undo); err == nil {
				err = app.DB.MarkCoachUndone(c.sess.UserID, id)
			}
		}
		v := loadCoach(c.sess.UserID)
		if err != nil {
			log.Printf("coach undo %d: %v", c.sess.UserID, err)
			v.Err = "coach.error"
		}
		view.Update(ctx, v)
		return false
	}
}
