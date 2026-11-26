// Managed by GoX v0.1.36

//line settings.gox:1
package pages

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"morning30/assets"
	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

type newExForm struct {
	ID      string `form:"ex_id"`
	Name    string `form:"ex_name"`
	Slot    string `form:"ex_slot"`
	Unit    string `form:"ex_unit"`
	PerSide string `form:"ex_perside"`
	Figure  string `form:"ex_figure"`
	PoseA   string `form:"ex_pose_a"`
	PoseB   string `form:"ex_pose_b"`
}

// editState is the editor's working copy: the selected day plus the program as
// JSON. It is a single comparable value so one Source/Bind drives the whole
// editor (day switch + every edit) — nested Binds don't patch reliably.
type editState struct {
	Day  int
	JSON string
	Open string // which row's content preview is expanded ("target:index"), "" = none
}

// settingsPage lets the user change rest time, language, voice and reset progress.
type settingsPage struct {
	sess   auth.Session
	auth   doors.Source[auth.Session]
	path   doors.Source[path.Path]
	editor bool // the "program and exercise editor" page instead of settings
}

type posePickForm struct {
	ID    string `form:"pose_ex"`
	PoseA string `form:"pose_a"`
	PoseB string `form:"pose_b"`
}

type mcpTokState struct {
	Has   bool
	Fresh string // plaintext, only right after issue/reissue
}

//line settings.gox:61
func (s settingsPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:63
		l := i18n.Lang(s.sess.Lang)
		saveStatus := doors.NewSource("")
		importStatus := doors.NewSource("")
		restW, restM, restC := app.DB.Rests(s.sess.UserID)
		rests := (content.Rests{Warmup: restW, Main: restM, Cooldown: restC}).Normalize()
		restW, restM, restC = rests.Warmup, rests.Main, rests.Cooldown
		restWVal := doors.NewSource(restW)
		restMVal := doors.NewSource(restM)
		restCVal := doors.NewSource(restC)
		cat := app.UserCatalog(s.sess.UserID)
		entries := cat.List(l)
		prog := app.UserProgram(s.sess.UserID) // the active program (built-in or custom)
		editSrc := doors.NewSource(editState{Day: 1, JSON: prog.Marshal()})
		nameSrc := doors.NewSource(prog.Name)
		firstID := "W01"
		if len(entries) > 0 {
			firstID = entries[0].ID
		}
		selSrc := doors.NewSource(firstID)
		levelSrc := doors.NewSource(s.sess.Level)
		curSex := app.DB.GetSex(s.sess.UserID)
		confirmReset := doors.NewSource(false)
		mcpSrc := doors.NewSource(mcpTokState{Has: app.DB.HasMCPToken(s.sess.UserID)})
		poseList, _ := app.DB.ListPoses()
		if poseList == nil {
			poseList = []db.Pose{}
		}

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:91
			if s.editor {
//line settings.gox:91
				__e = __c.Any(i18n.T(l, "settings.editor")); if __e != nil { return }
			} else  {
//line settings.gox:91
				__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line settings.gox:92
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line settings.gox:93
				__e = __c.Set("class", "settings"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line settings.gox:94
				if s.editor {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:95
						__e = __c.Set("class", "settings-head"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:96
						__e = (doors.ALink{Model: path.Path{Page: path.Settings}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("a"); if __e != nil { return }
							{
//line settings.gox:96
								__e = __c.Set("class", "back"); if __e != nil { return }
//line settings.gox:96
								__e = __c.Set("aria-label", i18n.T(l, "nav.settings")); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:96
								__e = __c.Any(icon("back", false)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
						__e = __c.Init("h1"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:97
							__e = __c.Any(i18n.T(l, "settings.editor")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line settings.gox:99
						__e = __c.Set("class", "editing"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:99
						__e = __c.Many(i18n.T(l, "settings.editing"), " "); if __e != nil { return }
						__e = __c.Init("b"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:99
							__e = __c.Any(app.ProgramTitle(s.sess.UserID, s.sess.Lang)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:100
						__e = __c.Set("class", "settings-card"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("details"); if __e != nil { return }
						{
//line settings.gox:102
							__e = __c.Set("class", "cycle-details"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("summary"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:103
								__e = __c.Any(i18n.T(l, "settings.cycle")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:105
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:106
									__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:107
								b := []byte(prog.Marshal())
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:108
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:108
									__e = __c.Modify(doors.ResourceBytes(b)); if __e != nil { return }
//line settings.gox:108
									__e = __c.Set("name", "workout.json"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Text("workout.json"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:110
							__e = (doors.ARawSubmit{On: s.importProgram(importStatus)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("form"); if __e != nil { return }
								{
//line settings.gox:110
									__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("textarea"); if __e != nil { return }
									{
//line settings.gox:111
										__e = __c.Set("class", "ex-import"); if __e != nil { return }
//line settings.gox:111
										__e = __c.Set("name", "progjson"); if __e != nil { return }
//line settings.gox:111
										__e = __c.Set("rows", "3"); if __e != nil { return }
//line settings.gox:111
										__e = __c.Set("placeholder", i18n.T(l, "settings.paste_program")); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:112
										__e = __c.Set("class", "srow ex-import-row"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.InitVoid("input"); if __e != nil { return }
										{
//line settings.gox:113
											__e = __c.Set("class", "ex-file"); if __e != nil { return }
//line settings.gox:113
											__e = __c.Set("type", "file"); if __e != nil { return }
//line settings.gox:113
											__e = __c.Set("accept", ".json,application/json"); if __e != nil { return }
										}
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("button"); if __e != nil { return }
										{
//line settings.gox:114
											__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:114
											__e = __c.Set("type", "submit"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:114
											__e = __c.Any(i18n.T(l, "settings.import")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
//line settings.gox:116
									__e = __c.Any(importStatus.Bind(func(key string) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:117
								if key != "" {
									__e = __c.Init("p"); if __e != nil { return }
									{
//line settings.gox:117
										__e = __c.Set("class", "err"); if __e != nil { return }
//line settings.gox:117
										__e = __c.Set("role", "alert"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:117
										__e = __c.Any(i18n.T(l, key)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							return })
//line settings.gox:118
						})); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:121
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:122
									__e = __c.Any(i18n.T(l, "home.day")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:123
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:123
									__e = __c.Modify(doors.AChange{On: s.setDay(editSrc)}); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:124
									for d := 1; d <= len(prog.Days); d++ {
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:125
											__e = __c.Set("value", d); if __e != nil { return }
//line settings.gox:125
											__e = __c.Set("selected", func() any { return d == 1 }()); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:125
											__e = __c.Any(d); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:130
							__e = __c.Any(editSrc.Bind(func(st editState) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:132
							p, _ := content.ParseResolved(st.JSON)
							idx := clampDay(st.Day, len(p.Days)) - 1
							if idx < 0 {
								idx = 0
							}
							warmTarget, coolTarget := "warmup", "cooldown"
							if len(p.Days) > 0 {
								if p.Days[idx].Warmup != nil { warmTarget = "day-warmup" }
								if p.Days[idx].Cooldown != nil { coolTarget = "day-cooldown" }
							}

							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:143
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:144
									__e = __c.Any(i18n.T(l, "settings.warmup_rounds")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:145
									__e = __c.Any(content.WarmupRounds); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:147
								__e = __c.Set("class", "cyc-head"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:147
								if warmTarget == "day-warmup" {
//line settings.gox:147
									__e = __c.Many(i18n.T(l, "workout.warmup"), " · ", i18n.T(l, "home.day"), " ", st.Day); if __e != nil { return }
								} else  {
//line settings.gox:147
									__e = __c.Any(i18n.T(l, "settings.warmup_series")); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:148
							__e = __c.Any(s.editList(editSrc, cat, entries, l, p.WarmupFor(idx), warmTarget, st.Open)); if __e != nil { return }
//line settings.gox:150
							if len(p.Days) > 0 {
//line settings.gox:151
								rd := p.Days[idx]
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:152
									__e = __c.Set("class", "srow"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:153
										__e = __c.Any(i18n.T(l, "settings.main_rounds")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:154
										__e = __c.Any(content.MainRounds); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:156
									__e = __c.Set("class", "cyc-head"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:156
									__e = __c.Many(i18n.T(l, "settings.main_block"), " · ", i18n.T(l, "home.day"), " ", st.Day, " · ", len(rd.Items), " ", i18n.T(l, "home.exercises")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:157
								__e = __c.Any(s.editList(editSrc, cat, entries, l, rd.Items, "day", st.Open)); if __e != nil { return }
							}
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:160
								__e = __c.Set("class", "cyc-head"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:160
								if coolTarget == "day-cooldown" {
//line settings.gox:160
									__e = __c.Many(i18n.T(l, "workout.cooldown"), " · ", i18n.T(l, "home.day"), " ", st.Day); if __e != nil { return }
								} else  {
//line settings.gox:160
									__e = __c.Any(i18n.T(l, "settings.cooldown_series")); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:161
							__e = __c.Any(s.editList(editSrc, cat, entries, l, p.CooldownFor(idx), coolTarget, st.Open)); if __e != nil { return }
						return })
//line settings.gox:162
					})); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:164
								__e = __c.Set("class", "srow cyc-save"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:165
								__e = (doors.AClick{On: s.saveEditor(editSrc, nameSrc)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:165
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:165
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:165
										__e = __c.Any(i18n.T(l, "settings.save_version")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("details"); if __e != nil { return }
						{
//line settings.gox:171
							__e = __c.Set("class", "cycle-details"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("summary"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:172
								__e = __c.Any(i18n.T(l, "settings.exercise_content")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:173
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:174
									__e = __c.Any(i18n.T(l, "settings.exercise")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:175
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:175
									__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
							selSrc.Update(ctx, r.Event().Value)
							return false
						}}); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:179
									__e = __c.Any(s.exOptions(entries, l, firstID)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:183
							__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:184
							__e = __c.Any(s.exDetail(l, cat, id)); if __e != nil { return }
						return })
//line settings.gox:185
					})); if __e != nil { return }
//line settings.gox:186
							__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:188
							cur := cat.PoseIDs(id)
							poseA, poseB := "", ""
							if len(cur) > 0 {
								poseA = cur[0]
							}
							if len(cur) > 1 {
								poseB = cur[1]
							}
							sex := app.DB.GetSex(s.sess.UserID)
							thumbs := app.ExerciseFrames(cat, id, sex)

//line settings.gox:199
							__e = (doors.ASubmit[posePickForm]{On: s.assignPoses}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("form"); if __e != nil { return }
								{
//line settings.gox:199
									__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:200
										__e = __c.Set("type", "hidden"); if __e != nil { return }
//line settings.gox:200
										__e = __c.Set("name", "pose_ex"); if __e != nil { return }
//line settings.gox:200
										__e = __c.Set("value", id); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:201
										__e = __c.Set("class", "pose-preview"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:202
										for _, u := range thumbs {
											__e = __c.InitVoid("img"); if __e != nil { return }
											{
//line settings.gox:203
												__e = __c.Set("class", "pose-thumb"); if __e != nil { return }
//line settings.gox:203
												__e = __c.Set("src", u); if __e != nil { return }
//line settings.gox:203
												__e = __c.Set("alt", ""); if __e != nil { return }
											}
											__e = __c.Submit(); if __e != nil { return }
										}
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:206
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:207
											__e = __c.Any(i18n.T(l, "settings.pose_a")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:208
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:208
											__e = __c.Set("name", "pose_a"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:209
											__e = __c.Any(s.poseOptions(poseList, poseA, i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:212
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:213
											__e = __c.Any(i18n.T(l, "settings.pose_b")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:214
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:214
											__e = __c.Set("name", "pose_b"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:215
											__e = __c.Any(s.poseOptions(poseList, poseB, i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:218
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:218
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:218
										__e = __c.Any(i18n.T(l, "settings.pose_save")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						return })
//line settings.gox:220
					})); if __e != nil { return }
//line settings.gox:221
							__e = (doors.ASubmit[newExForm]{On: s.createExercise}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("form"); if __e != nil { return }
								{
//line settings.gox:221
									__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:222
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:223
											__e = __c.Any(i18n.T(l, "settings.new_exercise")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.InitVoid("input"); if __e != nil { return }
										{
//line settings.gox:224
											__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:224
											__e = __c.Set("name", "ex_id"); if __e != nil { return }
//line settings.gox:224
											__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:224
											__e = __c.Set("maxlength", "12"); if __e != nil { return }
//line settings.gox:224
											__e = __c.Set("placeholder", "X01"); if __e != nil { return }
										}
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.InitVoid("input"); if __e != nil { return }
										{
//line settings.gox:225
											__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:225
											__e = __c.Set("name", "ex_name"); if __e != nil { return }
//line settings.gox:225
											__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:225
											__e = __c.Set("maxlength", "60"); if __e != nil { return }
//line settings.gox:225
											__e = __c.Set("placeholder", i18n.T(l, "settings.exercise")); if __e != nil { return }
										}
										__e = __c.Submit(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:227
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:228
											__e = __c.Any(i18n.T(l, "settings.pose_a")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:229
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:229
											__e = __c.Set("name", "ex_pose_a"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:230
											__e = __c.Any(s.poseOptions(poseList, "", i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:233
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:234
											__e = __c.Any(i18n.T(l, "settings.pose_b")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:235
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:235
											__e = __c.Set("name", "ex_pose_b"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:236
											__e = __c.Any(s.poseOptions(poseList, "", i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:239
										__e = __c.Set("class", "srow"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:240
											__e = __c.Any(i18n.T(l, "settings.figure")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:241
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:241
											__e = __c.Set("name", "ex_figure"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:242
											for _, e := range entries {
												__e = __c.Init("option"); if __e != nil { return }
												{
//line settings.gox:243
													__e = __c.Set("value", e.ID); if __e != nil { return }
													__e = __c.Submit(); if __e != nil { return }
//line settings.gox:243
													__e = __c.Any(e.Name); if __e != nil { return }
												}
												__e = __c.Close(); if __e != nil { return }
											}
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:247
										__e = __c.Set("class", "srow"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:248
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:248
											__e = __c.Set("name", "ex_slot"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:249
												__e = __c.Set("value", "warmup"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:249
												__e = __c.Any(i18n.T(l, "workout.warmup")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:250
												__e = __c.Set("value", "main"); if __e != nil { return }
//line settings.gox:250
												__e = __c.Set("selected", func() any { return true }()); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:250
												__e = __c.Any(i18n.T(l, "workout.main")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:251
												__e = __c.Set("value", "cooldown"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:251
												__e = __c.Any(i18n.T(l, "workout.cooldown")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:253
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:253
											__e = __c.Set("name", "ex_unit"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:254
												__e = __c.Set("value", "reps"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:254
												__e = __c.Any(i18n.T(l, "workout.x_reps")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:255
												__e = __c.Set("value", "seconds"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:255
												__e = __c.Any(i18n.T(l, "workout.sec")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:256
												__e = __c.Set("value", "breaths"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:256
												__e = __c.Any(i18n.T(l, "workout.breaths")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("label"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
											__e = __c.InitVoid("input"); if __e != nil { return }
											{
//line settings.gox:258
												__e = __c.Set("type", "checkbox"); if __e != nil { return }
//line settings.gox:258
												__e = __c.Set("name", "ex_perside"); if __e != nil { return }
//line settings.gox:258
												__e = __c.Set("value", "1"); if __e != nil { return }
											}
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:258
											__e = __c.Any(i18n.T(l, "workout.side")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("button"); if __e != nil { return }
										{
//line settings.gox:259
											__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:259
											__e = __c.Set("type", "submit"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:259
											__e = __c.Any(i18n.T(l, "settings.add_exercise")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:262
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:263
									__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:264
								__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:266
								doc, _ := cat.Doc(id)
								b := doc.Marshal()

								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:269
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:269
									__e = __c.Modify(doors.ResourceBytes(b)); if __e != nil { return }
//line settings.gox:269
									__e = __c.Set("name", id + ".json"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:269
									__e = __c.Many(id, ".json"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })
//line settings.gox:270
						})); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:272
							__e = (doors.ARawSubmit{On: s.importExercise}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("form"); if __e != nil { return }
								{
//line settings.gox:272
									__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("textarea"); if __e != nil { return }
									{
//line settings.gox:273
										__e = __c.Set("class", "ex-import"); if __e != nil { return }
//line settings.gox:273
										__e = __c.Set("name", "exjson"); if __e != nil { return }
//line settings.gox:273
										__e = __c.Set("rows", "3"); if __e != nil { return }
//line settings.gox:273
										__e = __c.Set("placeholder", i18n.T(l, "settings.paste_json")); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:274
										__e = __c.Set("class", "srow ex-import-row"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.InitVoid("input"); if __e != nil { return }
										{
//line settings.gox:275
											__e = __c.Set("class", "ex-file"); if __e != nil { return }
//line settings.gox:275
											__e = __c.Set("type", "file"); if __e != nil { return }
//line settings.gox:275
											__e = __c.Set("accept", ".json,application/json"); if __e != nil { return }
										}
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("button"); if __e != nil { return }
										{
//line settings.gox:276
											__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:276
											__e = __c.Set("type", "submit"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:276
											__e = __c.Any(i18n.T(l, "settings.import")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
							__e = __c.Init("script"); if __e != nil { return }
							{
//line settings.gox:279
								__e = __c.Set("src", assets.ExUpload); if __e != nil { return }
								__e = __c.Set("inline", true); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Raw(""); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:282
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:283
							__e = (doors.ALink{Model: path.Path{Page: path.Library}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:283
									__e = __c.Set("class", "row-link"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:283
										__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:283
										__e = __c.Set("class", "row-value"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:283
										__e = __c.Any(icon("next", false)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:285
						__e = __c.Any(mcpSrc.Bind(func(st mcpTokState) gox.Elem {
					return gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
//line settings.gox:287
						mcpBtn := i18n.T(l, "settings.mcp_create")
						if st.Has {
							mcpBtn = i18n.T(l, "settings.mcp_reissue")
						}

						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:292
							__e = __c.Set("class", "srow col"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:293
								__e = __c.Any(i18n.T(l, "settings.mcp")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:294
							if st.Fresh != "" {
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:295
									__e = __c.Set("class", "text-input compact mcp-token"); if __e != nil { return }
//line settings.gox:295
									__e = __c.Set("type", "text"); if __e != nil { return }
									__e = __c.Set("readonly", true); if __e != nil { return }
//line settings.gox:295
									__e = __c.Set("value", st.Fresh); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("p"); if __e != nil { return }
								{
//line settings.gox:296
									__e = __c.Set("class", "mcp-hint"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:296
									__e = __c.Any(i18n.T(l, "settings.mcp_copy")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:297
							} else if st.Has {
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:298
									__e = __c.Any(i18n.T(l, "settings.mcp_active")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							} else  {
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:300
									__e = __c.Any(i18n.T(l, "settings.mcp_none")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
//line settings.gox:302
							__e = (doors.AClick{On: s.issueMCP(mcpSrc)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:302
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:302
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:302
									__e = __c.Any(mcpBtn); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })
//line settings.gox:304
				})); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				} else  {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:308
						__e = __c.Set("class", "settings-head"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("h1"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:309
							__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:310
						__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("a"); if __e != nil { return }
							{
//line settings.gox:310
								__e = __c.Set("class", "quit"); if __e != nil { return }
//line settings.gox:310
								__e = __c.Set("aria-label", i18n.T(l, "nav.close")); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:310
								__e = __c.Any(icon("close", false)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:313
					vm := s.sess.VoiceMode
					if vm == "" {
						vm = "normal"
					}

					__e = __c.Init("h2"); if __e != nil { return }
					{
//line settings.gox:319
						__e = __c.Set("class", "sgroup"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:319
						__e = __c.Any(i18n.T(l, "settings.group_training")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:320
						__e = __c.Set("class", "settings-card"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:321
							__e = __c.Set("class", "srow col"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:322
								__e = __c.Set("class", "row-between"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:323
									__e = __c.Any(i18n.T(l, "settings.rest_warmup")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("output"); if __e != nil { return }
								{
//line settings.gox:324
									__e = __c.Set("class", "rest-out"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:324
									__e = __c.Any(restWVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:324
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:324
							})); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:326
								__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:326
								__e = __c.Set("min", "0"); if __e != nil { return }
//line settings.gox:326
								__e = __c.Set("max", content.RestMax); if __e != nil { return }
//line settings.gox:326
								__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:326
								__e = __c.Set("value", restW); if __e != nil { return }
//line settings.gox:326
								__e = __c.Set("aria-label", i18n.T(l, "settings.rest_warmup")); if __e != nil { return }
//line settings.gox:327
								__e = __c.Modify(doors.AChange{On: s.saveRest("warmup", restWVal, saveStatus)}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:329
							__e = __c.Set("class", "srow col"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:330
								__e = __c.Set("class", "row-between"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:331
									__e = __c.Any(i18n.T(l, "settings.rest_main")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("output"); if __e != nil { return }
								{
//line settings.gox:332
									__e = __c.Set("class", "rest-out"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:332
									__e = __c.Any(restMVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:332
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:332
							})); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:334
								__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:334
								__e = __c.Set("min", content.MainRestMin); if __e != nil { return }
//line settings.gox:334
								__e = __c.Set("max", content.RestMax); if __e != nil { return }
//line settings.gox:334
								__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:334
								__e = __c.Set("value", restM); if __e != nil { return }
//line settings.gox:334
								__e = __c.Set("aria-label", i18n.T(l, "settings.rest_main")); if __e != nil { return }
//line settings.gox:335
								__e = __c.Modify(doors.AChange{On: s.saveRest("main", restMVal, saveStatus)}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:337
							__e = __c.Set("class", "srow col"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:338
								__e = __c.Set("class", "row-between"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:339
									__e = __c.Any(i18n.T(l, "settings.rest_cooldown")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("output"); if __e != nil { return }
								{
//line settings.gox:340
									__e = __c.Set("class", "rest-out"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:340
									__e = __c.Any(restCVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:340
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:340
							})); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:342
								__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:342
								__e = __c.Set("min", "0"); if __e != nil { return }
//line settings.gox:342
								__e = __c.Set("max", content.RestMax); if __e != nil { return }
//line settings.gox:342
								__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:342
								__e = __c.Set("value", restC); if __e != nil { return }
//line settings.gox:342
								__e = __c.Set("aria-label", i18n.T(l, "settings.rest_cooldown")); if __e != nil { return }
//line settings.gox:343
								__e = __c.Modify(doors.AChange{On: s.saveRest("cooldown", restCVal, saveStatus)}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:345
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:346
								__e = __c.Any(i18n.T(l, "settings.voice")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:347
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:347
								__e = __c.Set("aria-label", i18n.T(l, "settings.voice")); if __e != nil { return }
//line settings.gox:347
								__e = __c.Modify(doors.AChange{On: s.saveSetting("voice", saveStatus)}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:348
									__e = __c.Set("value", "off"); if __e != nil { return }
//line settings.gox:348
									__e = __c.Set("selected", func() any { return vm == "off" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:348
									__e = __c.Any(i18n.T(l, "voice.off")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:349
									__e = __c.Set("value", "min"); if __e != nil { return }
//line settings.gox:349
									__e = __c.Set("selected", func() any { return vm == "min" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:349
									__e = __c.Any(i18n.T(l, "voice.min")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:350
									__e = __c.Set("value", "normal"); if __e != nil { return }
//line settings.gox:350
									__e = __c.Set("selected", func() any { return vm == "normal" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:350
									__e = __c.Any(i18n.T(l, "voice.normal")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:351
									__e = __c.Set("value", "detailed"); if __e != nil { return }
//line settings.gox:351
									__e = __c.Set("selected", func() any { return vm == "detailed" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:351
									__e = __c.Any(i18n.T(l, "voice.detailed")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:354
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:355
								__e = __c.Any(i18n.T(l, "settings.sex")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:356
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:356
								__e = __c.Set("aria-label", i18n.T(l, "settings.sex")); if __e != nil { return }
//line settings.gox:356
								__e = __c.Modify(doors.AChange{On: s.saveSetting("figure", saveStatus)}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:357
									__e = __c.Set("value", "m"); if __e != nil { return }
//line settings.gox:357
									__e = __c.Set("selected", func() any { return curSex == "m" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:357
									__e = __c.Any(i18n.T(l, "sex.male")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:358
									__e = __c.Set("value", "f"); if __e != nil { return }
//line settings.gox:358
									__e = __c.Set("selected", func() any { return curSex == "f" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:358
									__e = __c.Any(i18n.T(l, "sex.female")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:359
									__e = __c.Set("value", "b"); if __e != nil { return }
//line settings.gox:359
									__e = __c.Set("selected", func() any { return curSex == "b" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:359
									__e = __c.Any(i18n.T(l, "sex.boy")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:360
									__e = __c.Set("value", "g"); if __e != nil { return }
//line settings.gox:360
									__e = __c.Set("selected", func() any { return curSex == "g" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:360
									__e = __c.Any(i18n.T(l, "sex.girl")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:364
						if !prog.FixedDifficulty {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:365
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:366
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:367
										__e = __c.Any(i18n.T(l, "settings.difficulty")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:368
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:369
										__e = __c.Any(levelSrc.Bind(func(v int) gox.Elem {
									return gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
//line settings.gox:370
										__e = __c.Any(levelLabel(l, v)); if __e != nil { return }
									return })
//line settings.gox:371
								})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:374
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:374
									__e = __c.Set("min", "-3"); if __e != nil { return }
//line settings.gox:374
									__e = __c.Set("max", "3"); if __e != nil { return }
//line settings.gox:374
									__e = __c.Set("step", "1"); if __e != nil { return }
//line settings.gox:374
									__e = __c.Set("value", s.sess.Level); if __e != nil { return }
//line settings.gox:374
									__e = __c.Set("aria-label", i18n.T(l, "settings.difficulty")); if __e != nil { return }
//line settings.gox:374
									__e = __c.Modify(doors.AChange{On: s.setLevel(levelSrc)}); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("h2"); if __e != nil { return }
					{
//line settings.gox:379
						__e = __c.Set("class", "sgroup"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:379
						__e = __c.Any(i18n.T(l, "settings.group_program")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:380
						__e = __c.Set("class", "settings-card"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:381
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:382
							__e = (doors.ALink{Model: path.Path{Page: path.Programs}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:382
									__e = __c.Set("class", "row-link"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:382
										__e = __c.Any(i18n.T(l, "programs.title")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:382
										__e = __c.Set("class", "row-value"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:382
										__e = __c.Any(app.ProgramTitle(s.sess.UserID, s.sess.Lang)); if __e != nil { return }
//line settings.gox:382
										__e = __c.Any(icon("next", false)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:384
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:385
							__e = (doors.ALink{Model: path.Path{Page: path.Editor}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:385
									__e = __c.Set("class", "row-link"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:385
										__e = __c.Any(i18n.T(l, "settings.editor")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:385
										__e = __c.Set("class", "row-value"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:385
										__e = __c.Any(icon("next", false)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("h2"); if __e != nil { return }
					{
//line settings.gox:389
						__e = __c.Set("class", "sgroup"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:389
						__e = __c.Any(i18n.T(l, "settings.group_profile")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:390
						__e = __c.Set("class", "settings-card"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:391
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:392
								__e = __c.Any(i18n.T(l, "settings.language")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:393
							__e = __c.Any(langSwitch{sess: s.sess, auth: s.auth, persist: true}); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:395
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:396
								__e = __c.Any(i18n.T(l, "auth.name")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:397
								__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:397
								__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:397
								__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:397
								__e = __c.Set("value", s.sess.Name); if __e != nil { return }
//line settings.gox:397
								__e = __c.Set("aria-label", i18n.T(l, "auth.name")); if __e != nil { return }
//line settings.gox:397
								__e = __c.Modify(doors.AChange{On: s.saveSetting("name", saveStatus)}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:399
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:400
								__e = __c.Any(i18n.T(l, "settings.account")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:401
							if s.sess.IsGuest {
//line settings.gox:402
								__e = (doors.ALink{Model: path.Path{Page: path.Register}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("a"); if __e != nil { return }
									{
//line settings.gox:402
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:402
										__e = __c.Any(i18n.T(l, "auth.register")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							} else  {
								__e = __c.Init("span"); if __e != nil { return }
								{
//line settings.gox:404
									__e = __c.Set("class", "email"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:404
									__e = __c.Any(s.sess.Email); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:405
								__e = (doors.AClick{On: logout(s.sess, s.auth, s.path)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:405
										__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:405
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:405
										__e = __c.Any(i18n.T(l, "nav.logout")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:409
							__e = __c.Set("class", "srow save-row"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:410
							__e = __c.Any(confirmReset.Bind(func(c bool) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:411
							if !c {
//line settings.gox:412
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								confirmReset.Update(ctx, true)
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:415
										__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:415
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:415
										__e = __c.Any(i18n.T(l, "settings.reset")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							} else  {
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:417
									__e = __c.Set("class", "confirm-reset"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:418
										__e = __c.Set("class", "confirm-q"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:418
										__e = __c.Any(i18n.T(l, "settings.reset_confirm")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:419
										__e = __c.Set("class", "confirm-actions"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:420
										__e = (doors.AClick{On: s.reset}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line settings.gox:420
												__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:420
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:420
												__e = __c.Any(i18n.T(l, "settings.reset_yes")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										return })); if __e != nil { return }
//line settings.gox:421
										__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
										confirmReset.Update(ctx, false)
										return false
									}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line settings.gox:424
												__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:424
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:424
												__e = __c.Any(i18n.T(l, "settings.cancel")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										return })); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
						return })
//line settings.gox:428
					})); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:431
					__e = __c.Any(saveStatus.Bind(func(key string) gox.Elem {
					return gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
//line settings.gox:432
						if key != "" {
							__e = __c.Init("p"); if __e != nil { return }
							{
//line settings.gox:433
								__e = __c.Set("class", "err"); if __e != nil { return }
//line settings.gox:433
								__e = __c.Set("role", "alert"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:433
								__e = __c.Any(i18n.T(l, key)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					return })
//line settings.gox:435
				})); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:439
}

// exDetail renders the read-only content card for an exercise: description,
// how-to, correct, mistakes, breathing and safety, localized.
//line settings.gox:443
func (s settingsPage) exDetail(l i18n.Lang, cat catalog.Catalog, id string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:444
		d, _ := cat.Detail(l, id)
		__e = __c.Init("div"); if __e != nil { return }
		{
//line settings.gox:445
			__e = __c.Set("class", "ex-detail"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line settings.gox:446
				__e = __c.Set("class", "ex-detail-name"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line settings.gox:446
				__e = __c.Any(cat.Name(l, id)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line settings.gox:447
			if d.Desc != "" {
				__e = __c.Init("p"); if __e != nil { return }
				{
//line settings.gox:448
					__e = __c.Set("class", "ex-detail-desc"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line settings.gox:448
					__e = __c.Any(d.Desc); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:450
			if len(d.HowTo) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:451
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:452
						__e = __c.Any(i18n.T(l, "info.how")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ol"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:453
						for _, x := range d.HowTo {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:453
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:456
			if len(d.Correct) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:457
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:458
						__e = __c.Any(i18n.T(l, "info.correct")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ul"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:459
						for _, x := range d.Correct {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:459
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:462
			if len(d.Wrong) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:463
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:464
						__e = __c.Any(i18n.T(l, "info.wrong")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ul"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:465
						for _, x := range d.Wrong {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:465
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:468
			if d.Breathing != "" {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:469
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:470
						__e = __c.Any(i18n.T(l, "info.breathing")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:471
						__e = __c.Any(d.Breathing); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:474
			if d.Safety != "" {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:475
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("⚠️ "); if __e != nil { return }
//line settings.gox:476
						__e = __c.Any(i18n.T(l, "info.warning")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:477
						__e = __c.Any(d.Safety); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:481
}

// exOptions renders the exercise dropdown options grouped by slot.
//line settings.gox:484
func (s settingsPage) exOptions(entries []catalog.Entry, l i18n.Lang, cur string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:485
		__e = __c.Any(s.exGroup(entries, l, cur, "warmup", "workout.warmup")); if __e != nil { return }
//line settings.gox:486
		__e = __c.Any(s.exGroup(entries, l, cur, "main", "workout.main")); if __e != nil { return }
//line settings.gox:487
		__e = __c.Any(s.exGroup(entries, l, cur, "cooldown", "workout.cooldown")); if __e != nil { return }
	return })
//line settings.gox:488
}

//line settings.gox:490
func (s settingsPage) poseOptions(poses []db.Pose, cur, noneLabel string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("option"); if __e != nil { return }
		{
//line settings.gox:491
			__e = __c.Set("value", ""); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:491
			__e = __c.Any(noneLabel); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line settings.gox:492
		for _, p := range poses {
//line settings.gox:494
			lab := p.ID
			if p.Label != "" {
				lab = p.Label + " (" + p.ID + ")"
			}
			if p.Kind != "" {
				lab = p.Kind + " · " + lab
			}

			__e = __c.Init("option"); if __e != nil { return }
			{
//line settings.gox:502
				__e = __c.Set("value", p.ID); if __e != nil { return }
//line settings.gox:502
				__e = __c.Set("selected", func() any { return p.ID == cur }()); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line settings.gox:502
				__e = __c.Any(lab); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line settings.gox:504
}

//line settings.gox:506
func (s settingsPage) exGroup(entries []catalog.Entry, l i18n.Lang, cur, slot, labelKey string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("optgroup"); if __e != nil { return }
		{
//line settings.gox:507
			__e = __c.Set("label", i18n.T(l, labelKey)); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:508
			for _, e := range entries {
//line settings.gox:509
				if e.Slot == slot {
					__e = __c.Init("option"); if __e != nil { return }
					{
//line settings.gox:510
						__e = __c.Set("value", e.ID); if __e != nil { return }
//line settings.gox:510
						__e = __c.Set("selected", func() any { return e.ID == cur }()); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:510
						__e = __c.Any(e.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:514
}


// unitWord is the unit label shown next to an editable value.
func unitWord(l i18n.Lang, u content.Unit, perSide bool) string {
	var w string
	switch u {
	case content.Seconds:
		w = i18n.T(l, "workout.sec")
	case content.Breaths:
		w = i18n.T(l, "workout.breaths")
	default:
		w = i18n.T(l, "workout.x_reps")
	}
	if perSide {
		w += " /" + i18n.T(l, "workout.side")
	}
	return w
}

// clampDay clamps a day number to [1, n] (0 when n == 0).
func clampDay(d, n int) int {
	if n == 0 {
		return 0
	}
	if d < 1 {
		d = 1
	}
	if d > n {
		d = n
	}
	return d
}

// setDay switches the edited day, keeping the working JSON.
func (s settingsPage) setDay(editSrc doors.Source[editState]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		d, err := strconv.Atoi(r.Event().Value)
		if err != nil {
			return false
		}
		st := editSrc.Get()
		st.Day = d
		editSrc.Update(ctx, st)
		return false
	}
}

// slotForTarget maps an editor section to the slot its items belong to.
func slotForTarget(target string) content.Slot {
	switch target {
	case "warmup", "day-warmup":
		return content.Warmup
	case "cooldown", "day-cooldown":
		return content.Cooldown
	default:
		return content.Main
	}
}

// mutateList parses the working-copy JSON, applies fn to the target list — the
// shared warm-up / cool-down (whole series) or the selected day's main block —
// and stores it back (triggering a re-render).
func mutateList(ctx context.Context, editSrc doors.Source[editState], target string, fn func(items []content.RItem) []content.RItem) {
	st := editSrc.Get()
	p, err := content.ParseResolved(st.JSON)
	if err != nil {
		return
	}
	editProgramList(&p, st.Day, target, fn)
	st.JSON = p.Marshal()
	editSrc.Update(ctx, st)
}

// editProgramList keeps every phase nonempty while editing its exercise sequence.
func editProgramList(p *content.Resolved, day int, target string, fn func([]content.RItem) []content.RItem) {
	switch target {
	case "warmup":
		if items := fn(p.Warmup); len(items) > 0 { p.Warmup = items }
	case "cooldown":
		if items := fn(p.Cooldown); len(items) > 0 { p.Cooldown = items }
	default:
		d := clampDay(day, len(p.Days))
		if d < 1 { return }
		switch target {
		case "day-warmup":
			items := fn(p.WarmupFor(d-1))
			if len(items) > 0 { p.Days[d-1].Warmup = &items }
		case "day-cooldown":
			items := fn(p.CooldownFor(d-1))
			if len(items) > 0 { p.Days[d-1].Cooldown = &items }
		default:
			if items := fn(p.Days[d-1].Items); len(items) > 0 { p.Days[d-1].Items = items }
		}
	}
}

// editList renders one CRUD section (warm-up / day main / cool-down).
//line settings.gox:612
func (s settingsPage) editList(editSrc doors.Source[editState], cat catalog.Catalog, entries []catalog.Entry, l i18n.Lang, items []content.RItem, target, open string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ol"); if __e != nil { return }
		{
//line settings.gox:613
			__e = __c.Set("class", "cyc-list edit"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:614
			for ii := range items {
//line settings.gox:616
				it := items[ii]
				key := target + ":" + strconv.Itoa(ii)

				__e = __c.Init("li"); if __e != nil { return }
				{
//line settings.gox:619
					__e = __c.Set("class", "cyc-row"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line settings.gox:620
						__e = __c.Set("class", "cyc-num"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("select"); if __e != nil { return }
					{
//line settings.gox:621
						__e = __c.Set("class", "cyc-select"); if __e != nil { return }
//line settings.gox:621
						__e = __c.Modify(doors.AChange{On: s.editReplace(editSrc, cat, target, ii)}); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:622
						__e = __c.Any(s.exOptions(entries, l, it.ID)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.InitVoid("input"); if __e != nil { return }
					{
//line settings.gox:624
						__e = __c.Set("class", "cyc-val"); if __e != nil { return }
//line settings.gox:624
						__e = __c.Set("type", "number"); if __e != nil { return }
//line settings.gox:624
						__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line settings.gox:624
						__e = __c.Set("min", "1"); if __e != nil { return }
//line settings.gox:624
						__e = __c.Set("max", "600"); if __e != nil { return }
//line settings.gox:624
						__e = __c.Set("value", it.Value); if __e != nil { return }
//line settings.gox:624
						__e = __c.Modify(doors.AChange{On: s.editSetValue(editSrc, target, ii)}); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line settings.gox:625
						__e = __c.Set("class", "cyc-unit"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:625
						__e = __c.Any(unitWord(l, it.Unit, it.PerSide)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:626
					__e = (doors.AClick{On: s.editPeek(editSrc, target, ii)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:626
							__e = __c.Set("class", "cyc-op cyc-info"); if __e != nil { return }
//line settings.gox:626
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:626
							__e = __c.Set("aria-label", i18n.T(l, "settings.view")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("ⓘ"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:627
					__e = (doors.AClick{On: s.editMove(editSrc, target, ii, -1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:627
							__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:627
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:627
							__e = __c.Set("aria-label", i18n.T(l, "settings.move_up")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("↑"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:628
					__e = (doors.AClick{On: s.editMove(editSrc, target, ii, 1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:628
							__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:628
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:628
							__e = __c.Set("aria-label", i18n.T(l, "settings.move_down")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("↓"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:629
					__e = (doors.AClick{On: s.editDelete(editSrc, target, ii)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:629
							__e = __c.Set("class", "cyc-op danger"); if __e != nil { return }
//line settings.gox:629
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:629
							__e = __c.Set("disabled", len(items) <= 1); if __e != nil { return }
//line settings.gox:629
							__e = __c.Set("aria-label", i18n.T(l, "settings.remove")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("✕"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line settings.gox:631
				if open == key {
					__e = __c.Init("li"); if __e != nil { return }
					{
//line settings.gox:632
						__e = __c.Set("class", "cyc-peek"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:633
						__e = __c.Any(s.exDetail(l, cat, it.ID)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line settings.gox:638
		__e = (doors.AClick{On: s.editAdd(editSrc, cat, target)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("button"); if __e != nil { return }
			{
//line settings.gox:638
				__e = __c.Set("class", "btn small cyc-add"); if __e != nil { return }
//line settings.gox:638
				__e = __c.Set("type", "button"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("＋ "); if __e != nil { return }
//line settings.gox:638
				__e = __c.Any(i18n.T(l, "settings.add_exercise")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line settings.gox:639
}

// levelLabel formats the difficulty level as a percent (or the "base" word).
func levelLabel(l i18n.Lang, level int) string {
	if level == 0 {
		return i18n.T(l, "settings.level_base")
	}
	if level > 0 {
		return "+" + strconv.Itoa(level*10) + "%"
	}
	return strconv.Itoa(level*10) + "%"
}

// setLevel persists the universal difficulty level immediately and updates the session.
func (s settingsPage) setLevel(levelSrc doors.Source[int]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		n := r.Event().Number
		if n == nil {
			return false
		}
		lv := int(*n)
		if lv < -3 {
			lv = -3
		}
		if lv > 3 {
			lv = 3
		}
		_ = app.DB.SetLevel(s.sess.UserID, lv)
		s.auth.Mutate(ctx, func(cur auth.Session) auth.Session { cur.Level = lv; return cur })
		levelSrc.Update(ctx, lv)
		return false
	}
}

// editReplace swaps the exercise at index ii for the one picked in the dropdown,
// resetting unit/side/slot from the chosen exercise and converting the value.
func (s settingsPage) editReplace(editSrc doors.Source[editState], cat catalog.Catalog, target string, ii int) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		newID := r.Event().Value
		def, ok := cat.Def(newID)
		if !ok {
			return false
		}
		mutateList(ctx, editSrc, target, func(its []content.RItem) []content.RItem {
			if ii < 0 || ii >= len(its) {
				return its
			}
			old := its[ii]
			its[ii] = content.RItem{
				ID: newID, Unit: def.Unit, PerSide: def.PerSide, Slot: slotForTarget(target),
				Value: content.ConvertValue(old.Unit, def.Unit, old.Value), Round: old.Round,
			}
			return its
		})
		return false
	}
}

// editSetValue updates the value (reps/seconds) at index ii.
func (s settingsPage) editSetValue(editSrc doors.Source[editState], target string, ii int) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		n := r.Event().Number
		if n == nil {
			return false
		}
		v := int(*n)
		if v < 1 {
			v = 1
		}
		if v > 600 {
			v = 600
		}
		mutateList(ctx, editSrc, target, func(its []content.RItem) []content.RItem {
			if ii >= 0 && ii < len(its) {
				its[ii].Value = v
			}
			return its
		})
		return false
	}
}

// editMove shifts the item at index ii by delta (-1 up / +1 down).
func (s settingsPage) editMove(editSrc doors.Source[editState], target string, ii, delta int) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		mutateList(ctx, editSrc, target, func(its []content.RItem) []content.RItem {
			j := ii + delta
			if ii < 0 || ii >= len(its) || j < 0 || j >= len(its) {
				return its
			}
			its[ii], its[j] = its[j], its[ii]
			return its
		})
		return false
	}
}

// editDelete removes the item at index ii.
func (s settingsPage) editDelete(editSrc doors.Source[editState], target string, ii int) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		mutateList(ctx, editSrc, target, func(its []content.RItem) []content.RItem {
			if ii < 0 || ii >= len(its) {
				return its
			}
			nw := make([]content.RItem, 0, len(its)-1)
			nw = append(nw, its[:ii]...)
			nw = append(nw, its[ii+1:]...)
			return nw
		})
		return false
	}
}

// editPeek toggles the inline content preview for a row.
func (s settingsPage) editPeek(editSrc doors.Source[editState], target string, ii int) func(context.Context, doors.RequestPointer) bool {
	key := target + ":" + strconv.Itoa(ii)
	return func(ctx context.Context, r doors.RequestPointer) bool {
		st := editSrc.Get()
		if st.Open == key {
			st.Open = ""
		} else {
			st.Open = key
		}
		editSrc.Update(ctx, st)
		return false
	}
}

// editAdd appends a default exercise (by section) to the target list.
func (s settingsPage) editAdd(editSrc doors.Source[editState], cat catalog.Catalog, target string) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		it := defaultItem(cat, target)
		mutateList(ctx, editSrc, target, func(its []content.RItem) []content.RItem {
			return append(append([]content.RItem{}, its...), it)
		})
		return false
	}
}

// defaultItem is the exercise added by a section's "add" button.
func defaultItem(cat catalog.Catalog, target string) content.RItem {
	id := "M01"
	switch target {
	case "warmup", "day-warmup":
		id = "W03"
	case "cooldown", "day-cooldown":
		id = "CD02"
	}
	unit, perSide := content.Reps, false
	if def, ok := cat.Def(id); ok {
		unit, perSide = def.Unit, def.PerSide
	}
	val := 10
	if unit == content.Seconds {
		val = 30
	} else if unit == content.Breaths {
		val = 5
	}
	return content.RItem{ID: id, Unit: unit, Value: val, PerSide: perSide, Slot: slotForTarget(target), Round: 1}
}

// saveEditor stores the edited program as the user's custom version and reloads.
func (s settingsPage) saveEditor(editSrc doors.Source[editState], nameSrc doors.Source[string]) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		p, err := content.ParseResolved(editSrc.Get().JSON)
		if err != nil {
			return false
		}
		name := app.CleanProgramName(nameSrc.Get())
		if name == "" {
			name = "My program"
		}
		p.Name = name
		_ = app.SetUserProgram(s.sess.UserID, p.Marshal())
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
}

// importProgram validates a personal cycle before saving it. Rejected input
// stays in the form and never replaces the saved program or changes progress.
// An accepted import keeps the user's previous program in the saved list.
func (s settingsPage) importProgram(status doors.Source[string]) func(context.Context, doors.RequestRawForm) bool {
	return func(ctx context.Context, r doors.RequestRawForm) bool {
		status.Update(ctx, "")
		pf, err := r.ParseForm(1 << 20)
		if err != nil {
			status.Update(ctx, "settings.import_failed")
			return false
		}
		raw := strings.TrimSpace(pf.FormValue("progjson"))
		if raw == "" || app.ImportProgram(s.sess.UserID, raw) != nil {
			status.Update(ctx, "settings.import_failed")
			return false
		}
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
}

// createExercise adds a custom exercise that reuses a built-in figure.
func (s settingsPage) createExercise(ctx context.Context, r doors.RequestForm[newExForm]) bool {
	f := r.Data()
	id := strings.TrimSpace(f.ID)
	name := strings.TrimSpace(f.Name)
	if id == "" || name == "" {
		return false
	}
	fig := content.FigureKey(f.Figure, id)
	if _, ok := content.Get(fig); !ok {
		fig = "P08"
	}
	slot := f.Slot
	switch content.Slot(slot) {
	case content.Warmup, content.Main, content.Cooldown:
	default:
		slot = string(content.Main)
	}
	unit := f.Unit
	switch content.Unit(unit) {
	case content.Reps, content.Seconds, content.Breaths:
	default:
		unit = string(content.Reps)
	}
	lang := string(i18n.Lang(s.sess.Lang))
	d := catalog.Doc{
		ID: id, Slot: slot, Unit: unit, PerSide: f.PerSide == "1" || f.PerSide == "on",
		Difficulty: 2, SVG: fig, Poses: catalog.ClipPoses([]string{f.PoseA, f.PoseB}),
		Content: map[string]catalog.Loc{lang: {Name: name}, "en": {Name: name}},
	}
	if !catalog.ValidDoc(d) {
		return false
	}
	if app.SaveCustomExercise(s.sess.UserID, d) != nil {
		return false
	}
	_ = r.After(doors.ActionLocationReload{})
	return false
}

// assignPoses stores 0–2 catalog stills on the selected exercise.
func (s settingsPage) assignPoses(ctx context.Context, r doors.RequestForm[posePickForm]) bool {
	f := r.Data()
	if strings.TrimSpace(f.ID) == "" {
		return false
	}
	if app.AssignExercisePoses(s.sess.UserID, f.ID, []string{f.PoseA, f.PoseB}) != nil {
		return false
	}
	_ = r.After(doors.ActionLocationReload{})
	return false
}

// importExercise stores an uploaded/pasted per-exercise JSON into the user's
// custom library (it then appears in the dropdowns and is shown/spoken), reloads.
func (s settingsPage) importExercise(ctx context.Context, r doors.RequestRawForm) bool {
	pf, err := r.ParseForm(1 << 20)
	if err != nil {
		return false
	}
	raw := strings.TrimSpace(pf.FormValue("exjson"))
	if raw == "" {
		return false
	}
	var d catalog.Doc
	if json.Unmarshal([]byte(raw), &d) != nil {
		return false
	}
	if app.SaveCustomExercise(s.sess.UserID, d) != nil {
		return false
	}
	_ = r.After(doors.ActionLocationReload{})
	return false
}






// issueMCP creates or replaces this account's MCP bearer token and shows it once.
func (s settingsPage) issueMCP(src doors.Source[mcpTokState]) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		tok, err := app.DB.IssueMCPToken(s.sess.UserID)
		if err != nil || tok == "" {
			return false
		}
		src.Update(ctx, mcpTokState{Has: true, Fresh: tok})
		return false
	}
}

// reset clears all progress and returns to the home screen.
func (s settingsPage) reset(ctx context.Context, r doors.RequestPointer) bool {
	_ = app.DB.ResetProgress(s.sess.UserID)
	s.path.Update(ctx, path.Path{Page: path.Home})
	return false
}

// saveRest stores one of the three rests as soon as its slider is released.
func (s settingsPage) saveRest(which string, val doors.Source[int], status doors.Source[string]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		n := r.Event().Number
		if n == nil {
			return false
		}
		v := int(*n)
		var p app.SettingsPatch
		switch which {
		case "warmup":
			p.RestWarmup = &v
		case "main":
			p.RestMain = &v
		default:
			p.RestCooldown = &v
		}
		st, err := app.UpdateSettings(s.sess.UserID, p)
		if err != nil {
			status.Update(ctx, "settings.save_failed")
			return false
		}
		status.Update(ctx, "")
		val.Update(ctx, v)
		s.auth.Mutate(ctx, func(cur auth.Session) auth.Session {
			cur.Rest = st.RestMain
			return cur
		})
		return false
	}
}

// saveSetting stores the voice mode, the figure or the name as soon as it changes.
func (s settingsPage) saveSetting(which string, status doors.Source[string]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		v := r.Event().Value
		var p app.SettingsPatch
		switch which {
		case "voice":
			p.VoiceMode = &v
		case "figure":
			p.Figure = &v
		default:
			p.Name = &v
		}
		st, err := app.UpdateSettings(s.sess.UserID, p)
		if err != nil {
			status.Update(ctx, "settings.save_failed")
			return false
		}
		status.Update(ctx, "")
		s.auth.Mutate(ctx, func(cur auth.Session) auth.Session {
			cur.Name = st.Name
			cur.VoiceMode = st.VoiceMode
			cur.Voice = st.VoiceMode != "off"
			return cur
		})
		return false
	}
}
