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

// settingsForm is the decoded settings submission.
type settingsForm struct {
	Name         string `form:"name"`
	Rest         int    `form:"rest"`          // main set
	RestWarmup   int    `form:"rest_warmup"`
	RestCooldown int    `form:"rest_cooldown"`
	VoiceMode    string `form:"voice_mode"`
	Sex          string `form:"sex"`
}

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
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
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

//line settings.gox:70
func (s settingsPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:72
		l := i18n.Lang(s.sess.Lang)
		saved := doors.NewSource(false)
		restW, restM, restC := app.DB.Rests(s.sess.UserID)
		restWVal := doors.NewSource(restW)
		restMVal := doors.NewSource(restM)
		restCVal := doors.NewSource(restC)
		hasCustom := app.HasCustomProgram(s.sess.UserID)
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
//line settings.gox:98
			__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line settings.gox:99
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line settings.gox:100
				__e = __c.Set("class", "settings"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:101
					__e = __c.Set("class", "settings-head"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h1"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:102
						__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:103
					__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("a"); if __e != nil { return }
						{
//line settings.gox:103
							__e = __c.Set("class", "quit"); if __e != nil { return }
//line settings.gox:103
							__e = __c.Set("aria-label", i18n.T(l, "nav.close")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("✕"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line settings.gox:106
				vm := s.sess.VoiceMode
				if vm == "" {
					vm = "normal"
				}

				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:111
					__e = __c.Set("class", "settings-card"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:113
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:114
							__e = __c.Any(i18n.T(l, "settings.language")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:115
						__e = __c.Any(langSwitch{sess: s.sess, auth: s.auth, persist: true}); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:118
					__e = (doors.ASubmit[settingsForm]{On: s.save(saved)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("form"); if __e != nil { return }
						{
//line settings.gox:118
							__e = __c.Set("class", "settings-form"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:119
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:120
									__e = __c.Any(i18n.T(l, "auth.name")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:121
									__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:121
									__e = __c.Set("name", "name"); if __e != nil { return }
//line settings.gox:121
									__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:121
									__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:121
									__e = __c.Set("value", s.sess.Name); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:124
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:125
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:126
										__e = __c.Any(i18n.T(l, "settings.rest_warmup")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:127
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:127
										__e = __c.Any(restWVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:127
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:127
							})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:129
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:129
									__e = __c.Set("name", "rest_warmup"); if __e != nil { return }
//line settings.gox:129
									__e = __c.Set("min", "0"); if __e != nil { return }
//line settings.gox:129
									__e = __c.Set("max", "40"); if __e != nil { return }
//line settings.gox:129
									__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:129
									__e = __c.Set("value", restW); if __e != nil { return }
//line settings.gox:130
									__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								if n := r.Event().Number; n != nil { restWVal.Update(ctx, int(*n)) }
								return false
							}}); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:135
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:136
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:137
										__e = __c.Any(i18n.T(l, "settings.rest_main")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:138
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:138
										__e = __c.Any(restMVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:138
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:138
							})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:140
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:140
									__e = __c.Set("name", "rest"); if __e != nil { return }
//line settings.gox:140
									__e = __c.Set("min", "5"); if __e != nil { return }
//line settings.gox:140
									__e = __c.Set("max", "40"); if __e != nil { return }
//line settings.gox:140
									__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:140
									__e = __c.Set("value", restM); if __e != nil { return }
//line settings.gox:141
									__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								if n := r.Event().Number; n != nil { restMVal.Update(ctx, int(*n)) }
								return false
							}}); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:146
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:147
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:148
										__e = __c.Any(i18n.T(l, "settings.rest_cooldown")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:149
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:149
										__e = __c.Any(restCVal.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:149
									__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
								return })
//line settings.gox:149
							})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:151
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:151
									__e = __c.Set("name", "rest_cooldown"); if __e != nil { return }
//line settings.gox:151
									__e = __c.Set("min", "0"); if __e != nil { return }
//line settings.gox:151
									__e = __c.Set("max", "40"); if __e != nil { return }
//line settings.gox:151
									__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:151
									__e = __c.Set("value", restC); if __e != nil { return }
//line settings.gox:152
									__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								if n := r.Event().Number; n != nil { restCVal.Update(ctx, int(*n)) }
								return false
							}}); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:157
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:158
									__e = __c.Any(i18n.T(l, "settings.voice")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:159
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:159
									__e = __c.Set("name", "voice_mode"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:160
										__e = __c.Set("value", "off"); if __e != nil { return }
//line settings.gox:160
										__e = __c.Set("selected", func() any { return vm == "off" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:160
										__e = __c.Any(i18n.T(l, "voice.off")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:161
										__e = __c.Set("value", "min"); if __e != nil { return }
//line settings.gox:161
										__e = __c.Set("selected", func() any { return vm == "min" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:161
										__e = __c.Any(i18n.T(l, "voice.min")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:162
										__e = __c.Set("value", "normal"); if __e != nil { return }
//line settings.gox:162
										__e = __c.Set("selected", func() any { return vm == "normal" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:162
										__e = __c.Any(i18n.T(l, "voice.normal")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:163
										__e = __c.Set("value", "detailed"); if __e != nil { return }
//line settings.gox:163
										__e = __c.Set("selected", func() any { return vm == "detailed" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:163
										__e = __c.Any(i18n.T(l, "voice.detailed")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:166
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:167
									__e = __c.Any(i18n.T(l, "settings.sex")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:168
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:168
									__e = __c.Set("name", "sex"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:169
										__e = __c.Set("value", "m"); if __e != nil { return }
//line settings.gox:169
										__e = __c.Set("selected", func() any { return curSex != "f" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:169
										__e = __c.Any(i18n.T(l, "sex.male")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:170
										__e = __c.Set("value", "f"); if __e != nil { return }
//line settings.gox:170
										__e = __c.Set("selected", func() any { return curSex == "f" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:170
										__e = __c.Any(i18n.T(l, "sex.female")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:173
								__e = __c.Set("class", "srow save-row"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:174
									__e = __c.Set("class", "btn primary"); if __e != nil { return }
//line settings.gox:174
									__e = __c.Set("type", "submit"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:174
									__e = __c.Any(i18n.T(l, "settings.save")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:175
								__e = __c.Any(saved.Bind(func(ok bool) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:176
								if ok {
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:177
										__e = __c.Set("class", "saved-msg"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:177
										__e = __c.Any(i18n.T(l, "settings.saved")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							return })
//line settings.gox:179
						})); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:184
						__e = __c.Set("class", "srow col"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:185
							__e = __c.Set("class", "row-between"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:186
								__e = __c.Any(i18n.T(l, "settings.difficulty")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("output"); if __e != nil { return }
							{
//line settings.gox:187
								__e = __c.Set("class", "rest-out"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:188
								__e = __c.Any(levelSrc.Bind(func(v int) gox.Elem {
								return gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
//line settings.gox:189
									__e = __c.Any(levelLabel(l, v)); if __e != nil { return }
								return })
//line settings.gox:190
							})); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.InitVoid("input"); if __e != nil { return }
						{
//line settings.gox:193
							__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:193
							__e = __c.Set("min", "-3"); if __e != nil { return }
//line settings.gox:193
							__e = __c.Set("max", "3"); if __e != nil { return }
//line settings.gox:193
							__e = __c.Set("step", "1"); if __e != nil { return }
//line settings.gox:193
							__e = __c.Set("value", s.sess.Level); if __e != nil { return }
//line settings.gox:193
							__e = __c.Modify(doors.AChange{On: s.setLevel(levelSrc)}); if __e != nil { return }
						}
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("details"); if __e != nil { return }
					{
//line settings.gox:197
						__e = __c.Set("class", "cycle-details"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("summary"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:198
							__e = __c.Any(i18n.T(l, "settings.cycle")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:200
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:201
								__e = __c.Any(i18n.T(l, "settings.program")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:202
								if hasCustom && prog.Name != "" {
									__e = __c.Text(" prog.Name "); if __e != nil { return }
								} else  {
									__e = __c.Text(" i18n.T(l, \"program.demo\") "); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:205
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:206
								__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:207
							b := []byte(prog.Marshal())
							__e = __c.Init("a"); if __e != nil { return }
							{
//line settings.gox:208
								__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:208
								__e = __c.Modify(doors.ResourceBytes(b)); if __e != nil { return }
//line settings.gox:208
								__e = __c.Set("name", "workout.json"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Text("workout.json"); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:210
						__e = (doors.ARawSubmit{On: s.importProgram}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("form"); if __e != nil { return }
							{
//line settings.gox:210
								__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("textarea"); if __e != nil { return }
								{
//line settings.gox:211
									__e = __c.Set("class", "ex-import"); if __e != nil { return }
//line settings.gox:211
									__e = __c.Set("name", "progjson"); if __e != nil { return }
//line settings.gox:211
									__e = __c.Set("rows", "3"); if __e != nil { return }
//line settings.gox:211
									__e = __c.Set("placeholder", i18n.T(l, "settings.paste_program")); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:212
									__e = __c.Set("class", "srow ex-import-row"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:213
										__e = __c.Set("class", "ex-file"); if __e != nil { return }
//line settings.gox:213
										__e = __c.Set("type", "file"); if __e != nil { return }
//line settings.gox:213
										__e = __c.Set("accept", ".json,application/json"); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:214
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:214
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:214
										__e = __c.Any(i18n.T(l, "settings.import")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:218
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:219
								__e = __c.Any(i18n.T(l, "home.day")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:220
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:220
								__e = __c.Modify(doors.AChange{On: s.setDay(editSrc)}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:221
								for d := 1; d <= len(prog.Days); d++ {
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:222
										__e = __c.Set("value", d); if __e != nil { return }
//line settings.gox:222
										__e = __c.Set("selected", func() any { return d == 1 }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:222
										__e = __c.Any(d); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:227
						__e = __c.Any(editSrc.Bind(func(st editState) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:229
							p, _ := content.ParseResolved(st.JSON)
							idx := clampDay(st.Day, len(p.Days)) - 1
							if idx < 0 {
								idx = 0
							}
							mainRounds := 2
							if len(p.Days) > 0 {
								mainRounds = content.ClampRounds(p.Days[idx].Rounds)
							}

							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:239
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:240
									__e = __c.Any(i18n.T(l, "settings.warmup_rounds")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:241
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:241
									__e = __c.Modify(doors.AChange{On: s.setWarmupRounds(editSrc)}); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:242
									for n := 1; n <= 10; n++ {
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:243
											__e = __c.Set("value", n); if __e != nil { return }
//line settings.gox:243
											__e = __c.Set("selected", func() any { return n == p.WarmupRounds }()); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:243
											__e = __c.Any(n); if __e != nil { return }
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
								__e = __c.Set("class", "cyc-head"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:247
								__e = __c.Any(i18n.T(l, "settings.warmup_series")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:248
							__e = __c.Any(s.editList(editSrc, cat, entries, l, p.Warmup, "warmup", st.Open)); if __e != nil { return }
//line settings.gox:250
							if len(p.Days) > 0 {
//line settings.gox:251
								rd := p.Days[idx]
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:252
									__e = __c.Set("class", "srow"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:253
										__e = __c.Any(i18n.T(l, "settings.main_rounds")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:254
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:254
										__e = __c.Modify(doors.AChange{On: s.setMainRounds(editSrc)}); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:255
										for n := 1; n <= 10; n++ {
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:256
												__e = __c.Set("value", n); if __e != nil { return }
//line settings.gox:256
												__e = __c.Set("selected", func() any { return n == mainRounds }()); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:256
												__e = __c.Any(n); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:260
									__e = __c.Set("class", "cyc-head"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:260
									__e = __c.Many(i18n.T(l, "settings.main_block"), " · ", i18n.T(l, "home.day"), " ", st.Day, " · ", len(rd.Items), " ", i18n.T(l, "home.exercises")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:261
								__e = __c.Any(s.editList(editSrc, cat, entries, l, rd.Items, "day", st.Open)); if __e != nil { return }
							}
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:264
								__e = __c.Set("class", "cyc-head"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:264
								__e = __c.Any(i18n.T(l, "settings.cooldown_series")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:265
							__e = __c.Any(s.editList(editSrc, cat, entries, l, p.Cooldown, "cooldown", st.Open)); if __e != nil { return }
						return })
//line settings.gox:266
					})); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:268
							__e = __c.Set("class", "srow cyc-save"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:269
								__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:269
								__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:269
								__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:269
								__e = __c.Set("value", prog.Name); if __e != nil { return }
//line settings.gox:269
								__e = __c.Set("placeholder", i18n.T(l, "settings.version_name")); if __e != nil { return }
//line settings.gox:270
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								nameSrc.Update(ctx, r.Event().Value)
								return false
							}}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:274
							__e = (doors.AClick{On: s.saveEditor(editSrc, nameSrc)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:274
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:274
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:274
									__e = __c.Any(i18n.T(l, "settings.save_version")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:277
						if hasCustom {
//line settings.gox:278
							__e = (doors.AClick{On: s.deleteProgram}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:278
									__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:278
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:278
									__e = __c.Any(i18n.T(l, "settings.restore_demo")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("details"); if __e != nil { return }
					{
//line settings.gox:284
						__e = __c.Set("class", "cycle-details"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("summary"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:285
							__e = __c.Any(i18n.T(l, "settings.exercise_content")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:286
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:287
								__e = __c.Any(i18n.T(l, "settings.exercise")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:288
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:288
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
							selSrc.Update(ctx, r.Event().Value)
							return false
						}}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:292
								__e = __c.Any(s.exOptions(entries, l, firstID)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:296
						__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:297
							__e = __c.Any(s.exDetail(l, cat, id)); if __e != nil { return }
						return })
//line settings.gox:298
					})); if __e != nil { return }
//line settings.gox:299
						__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:301
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

//line settings.gox:312
							__e = (doors.ASubmit[posePickForm]{On: s.assignPoses}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("form"); if __e != nil { return }
								{
//line settings.gox:312
									__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:313
										__e = __c.Set("type", "hidden"); if __e != nil { return }
//line settings.gox:313
										__e = __c.Set("name", "pose_ex"); if __e != nil { return }
//line settings.gox:313
										__e = __c.Set("value", id); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:314
										__e = __c.Set("class", "pose-preview"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:315
										for _, u := range thumbs {
											__e = __c.InitVoid("img"); if __e != nil { return }
											{
//line settings.gox:316
												__e = __c.Set("class", "pose-thumb"); if __e != nil { return }
//line settings.gox:316
												__e = __c.Set("src", u); if __e != nil { return }
//line settings.gox:316
												__e = __c.Set("alt", ""); if __e != nil { return }
											}
											__e = __c.Submit(); if __e != nil { return }
										}
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:319
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:320
											__e = __c.Any(i18n.T(l, "settings.pose_a")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:321
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:321
											__e = __c.Set("name", "pose_a"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:322
											__e = __c.Any(s.poseOptions(poseList, poseA, i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:325
										__e = __c.Set("class", "srow col"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:326
											__e = __c.Any(i18n.T(l, "settings.pose_b")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("select"); if __e != nil { return }
										{
//line settings.gox:327
											__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:327
											__e = __c.Set("name", "pose_b"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:328
											__e = __c.Any(s.poseOptions(poseList, poseB, i18n.T(l, "settings.pose_none"))); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:331
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:331
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:331
										__e = __c.Any(i18n.T(l, "settings.pose_save")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						return })
//line settings.gox:333
					})); if __e != nil { return }
//line settings.gox:334
						__e = (doors.ASubmit[newExForm]{On: s.createExercise}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("form"); if __e != nil { return }
							{
//line settings.gox:334
								__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:335
									__e = __c.Set("class", "srow col"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:336
										__e = __c.Any(i18n.T(l, "settings.new_exercise")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:337
										__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:337
										__e = __c.Set("name", "ex_id"); if __e != nil { return }
//line settings.gox:337
										__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:337
										__e = __c.Set("maxlength", "12"); if __e != nil { return }
//line settings.gox:337
										__e = __c.Set("placeholder", "X01"); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:338
										__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:338
										__e = __c.Set("name", "ex_name"); if __e != nil { return }
//line settings.gox:338
										__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:338
										__e = __c.Set("maxlength", "60"); if __e != nil { return }
//line settings.gox:338
										__e = __c.Set("placeholder", i18n.T(l, "settings.exercise")); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:340
									__e = __c.Set("class", "srow col"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:341
										__e = __c.Any(i18n.T(l, "settings.pose_a")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:342
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:342
										__e = __c.Set("name", "ex_pose_a"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:343
										__e = __c.Any(s.poseOptions(poseList, "", i18n.T(l, "settings.pose_none"))); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:346
									__e = __c.Set("class", "srow col"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:347
										__e = __c.Any(i18n.T(l, "settings.pose_b")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:348
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:348
										__e = __c.Set("name", "ex_pose_b"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:349
										__e = __c.Any(s.poseOptions(poseList, "", i18n.T(l, "settings.pose_none"))); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:352
									__e = __c.Set("class", "srow"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:353
										__e = __c.Any(i18n.T(l, "settings.figure")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:354
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:354
										__e = __c.Set("name", "ex_figure"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:355
										for _, e := range entries {
											__e = __c.Init("option"); if __e != nil { return }
											{
//line settings.gox:356
												__e = __c.Set("value", e.ID); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:356
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
//line settings.gox:360
									__e = __c.Set("class", "srow"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:361
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:361
										__e = __c.Set("name", "ex_slot"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:362
											__e = __c.Set("value", "warmup"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:362
											__e = __c.Any(i18n.T(l, "workout.warmup")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:363
											__e = __c.Set("value", "main"); if __e != nil { return }
//line settings.gox:363
											__e = __c.Set("selected", func() any { return true }()); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:363
											__e = __c.Any(i18n.T(l, "workout.main")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:364
											__e = __c.Set("value", "cooldown"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:364
											__e = __c.Any(i18n.T(l, "workout.cooldown")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("select"); if __e != nil { return }
									{
//line settings.gox:366
										__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:366
										__e = __c.Set("name", "ex_unit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:367
											__e = __c.Set("value", "reps"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:367
											__e = __c.Any(i18n.T(l, "workout.x_reps")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:368
											__e = __c.Set("value", "seconds"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:368
											__e = __c.Any(i18n.T(l, "workout.sec")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("option"); if __e != nil { return }
										{
//line settings.gox:369
											__e = __c.Set("value", "breaths"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line settings.gox:369
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
//line settings.gox:371
											__e = __c.Set("type", "checkbox"); if __e != nil { return }
//line settings.gox:371
											__e = __c.Set("name", "ex_perside"); if __e != nil { return }
//line settings.gox:371
											__e = __c.Set("value", "1"); if __e != nil { return }
										}
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:371
										__e = __c.Any(i18n.T(l, "workout.side")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:372
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:372
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:372
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
//line settings.gox:375
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:376
								__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:377
							__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:379
								doc, _ := cat.Doc(id)
								b := doc.Marshal()

								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:382
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:382
									__e = __c.Modify(doors.ResourceBytes(b)); if __e != nil { return }
//line settings.gox:382
									__e = __c.Set("name", id + ".json"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:382
									__e = __c.Many(id, ".json"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })
//line settings.gox:383
						})); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:385
						__e = (doors.ARawSubmit{On: s.importExercise}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("form"); if __e != nil { return }
							{
//line settings.gox:385
								__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("textarea"); if __e != nil { return }
								{
//line settings.gox:386
									__e = __c.Set("class", "ex-import"); if __e != nil { return }
//line settings.gox:386
									__e = __c.Set("name", "exjson"); if __e != nil { return }
//line settings.gox:386
									__e = __c.Set("rows", "3"); if __e != nil { return }
//line settings.gox:386
									__e = __c.Set("placeholder", i18n.T(l, "settings.paste_json")); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:387
									__e = __c.Set("class", "srow ex-import-row"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:388
										__e = __c.Set("class", "ex-file"); if __e != nil { return }
//line settings.gox:388
										__e = __c.Set("type", "file"); if __e != nil { return }
//line settings.gox:388
										__e = __c.Set("accept", ".json,application/json"); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:389
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:389
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:389
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
//line settings.gox:392
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
//line settings.gox:395
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:396
							__e = __c.Any(i18n.T(l, "settings.account")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:397
						if s.sess.IsGuest {
//line settings.gox:398
							__e = (doors.ALink{Model: path.Path{Page: path.Register}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:398
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:398
									__e = __c.Any(i18n.T(l, "auth.register")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						} else  {
							__e = __c.Init("span"); if __e != nil { return }
							{
//line settings.gox:400
								__e = __c.Set("class", "email"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:400
								__e = __c.Any(s.sess.Email); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:404
					__e = __c.Any(mcpSrc.Bind(func(st mcpTokState) gox.Elem {
					return gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
//line settings.gox:406
						mcpBtn := i18n.T(l, "settings.mcp_create")
						if st.Has {
							mcpBtn = i18n.T(l, "settings.mcp_reissue")
						}

						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:411
							__e = __c.Set("class", "srow col"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:412
								__e = __c.Any(i18n.T(l, "settings.mcp")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:413
							if st.Fresh != "" {
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:414
									__e = __c.Set("class", "text-input compact mcp-token"); if __e != nil { return }
//line settings.gox:414
									__e = __c.Set("type", "text"); if __e != nil { return }
									__e = __c.Set("readonly", true); if __e != nil { return }
//line settings.gox:414
									__e = __c.Set("value", st.Fresh); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("p"); if __e != nil { return }
								{
//line settings.gox:415
									__e = __c.Set("class", "mcp-hint"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:415
									__e = __c.Any(i18n.T(l, "settings.mcp_copy")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:416
							} else if st.Has {
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:417
									__e = __c.Any(i18n.T(l, "settings.mcp_active")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							} else  {
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:419
									__e = __c.Any(i18n.T(l, "settings.mcp_none")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
//line settings.gox:421
							__e = (doors.AClick{On: s.issueMCP(mcpSrc)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:421
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:421
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:421
									__e = __c.Any(mcpBtn); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })
//line settings.gox:423
				})); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:425
						__e = __c.Set("class", "srow save-row"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:426
						__e = __c.Any(confirmReset.Bind(func(c bool) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:427
							if !c {
//line settings.gox:428
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								confirmReset.Update(ctx, true)
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:431
										__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:431
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:431
										__e = __c.Any(i18n.T(l, "settings.reset")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							} else  {
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:433
									__e = __c.Set("class", "confirm-reset"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:434
										__e = __c.Set("class", "confirm-q"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:434
										__e = __c.Any(i18n.T(l, "settings.reset_confirm")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("div"); if __e != nil { return }
									{
//line settings.gox:435
										__e = __c.Set("class", "confirm-actions"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:436
										__e = (doors.AClick{On: s.reset}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line settings.gox:436
												__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:436
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:436
												__e = __c.Any(i18n.T(l, "settings.reset_yes")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										return })); if __e != nil { return }
//line settings.gox:437
										__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
										confirmReset.Update(ctx, false)
										return false
									}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line settings.gox:440
												__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:440
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:440
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
//line settings.gox:444
					})); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:449
}

// exDetail renders the read-only content card for an exercise: description,
// how-to, correct, mistakes, breathing and safety, localized.
//line settings.gox:453
func (s settingsPage) exDetail(l i18n.Lang, cat catalog.Catalog, id string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:454
		d, _ := cat.Detail(l, id)
		__e = __c.Init("div"); if __e != nil { return }
		{
//line settings.gox:455
			__e = __c.Set("class", "ex-detail"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line settings.gox:456
				__e = __c.Set("class", "ex-detail-name"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line settings.gox:456
				__e = __c.Any(cat.Name(l, id)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line settings.gox:457
			if d.Desc != "" {
				__e = __c.Init("p"); if __e != nil { return }
				{
//line settings.gox:458
					__e = __c.Set("class", "ex-detail-desc"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line settings.gox:458
					__e = __c.Any(d.Desc); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:460
			if len(d.HowTo) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:461
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:462
						__e = __c.Any(i18n.T(l, "info.how")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ol"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:463
						for _, x := range d.HowTo {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:463
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:466
			if len(d.Correct) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:467
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:468
						__e = __c.Any(i18n.T(l, "info.correct")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ul"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:469
						for _, x := range d.Correct {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:469
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:472
			if len(d.Wrong) > 0 {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:473
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:474
						__e = __c.Any(i18n.T(l, "info.wrong")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("ul"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:475
						for _, x := range d.Wrong {
							__e = __c.Init("li"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:475
								__e = __c.Any(x); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:478
			if d.Breathing != "" {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:479
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:480
						__e = __c.Any(i18n.T(l, "info.breathing")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:481
						__e = __c.Any(d.Breathing); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
//line settings.gox:484
			if d.Safety != "" {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:485
					__e = __c.Set("class", "ex-detail-sec"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("h4"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("⚠️ "); if __e != nil { return }
//line settings.gox:486
						__e = __c.Any(i18n.T(l, "info.warning")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:487
						__e = __c.Any(d.Safety); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:491
}

// exOptions renders the exercise dropdown options grouped by slot.
//line settings.gox:494
func (s settingsPage) exOptions(entries []catalog.Entry, l i18n.Lang, cur string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:495
		__e = __c.Any(s.exGroup(entries, l, cur, "warmup", "workout.warmup")); if __e != nil { return }
//line settings.gox:496
		__e = __c.Any(s.exGroup(entries, l, cur, "main", "workout.main")); if __e != nil { return }
//line settings.gox:497
		__e = __c.Any(s.exGroup(entries, l, cur, "cooldown", "workout.cooldown")); if __e != nil { return }
	return })
//line settings.gox:498
}

//line settings.gox:500
func (s settingsPage) poseOptions(poses []db.Pose, cur, noneLabel string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("option"); if __e != nil { return }
		{
//line settings.gox:501
			__e = __c.Set("value", ""); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:501
			__e = __c.Any(noneLabel); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line settings.gox:502
		for _, p := range poses {
//line settings.gox:504
			lab := p.ID
			if p.Label != "" {
				lab = p.Label + " (" + p.ID + ")"
			}
			if p.Kind != "" {
				lab = p.Kind + " · " + lab
			}

			__e = __c.Init("option"); if __e != nil { return }
			{
//line settings.gox:512
				__e = __c.Set("value", p.ID); if __e != nil { return }
//line settings.gox:512
				__e = __c.Set("selected", func() any { return p.ID == cur }()); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line settings.gox:512
				__e = __c.Any(lab); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line settings.gox:514
}

//line settings.gox:516
func (s settingsPage) exGroup(entries []catalog.Entry, l i18n.Lang, cur, slot, labelKey string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("optgroup"); if __e != nil { return }
		{
//line settings.gox:517
			__e = __c.Set("label", i18n.T(l, labelKey)); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:518
			for _, e := range entries {
//line settings.gox:519
				if e.Slot == slot {
					__e = __c.Init("option"); if __e != nil { return }
					{
//line settings.gox:520
						__e = __c.Set("value", e.ID); if __e != nil { return }
//line settings.gox:520
						__e = __c.Set("selected", func() any { return e.ID == cur }()); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:520
						__e = __c.Any(e.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:524
}

// save persists rest + voice and updates the reactive session.
func (s settingsPage) save(saved doors.Source[bool]) func(context.Context, doors.RequestForm[settingsForm]) bool {
	return func(ctx context.Context, r doors.RequestForm[settingsForm]) bool {
		f := r.Data()
		restMain := clampRange(f.Rest, 5, 40)
		restWarm := clampRange(f.RestWarmup, 0, 40)
		restCool := clampRange(f.RestCooldown, 0, 40)
		name := strings.TrimSpace(f.Name)
		if len(name) > 40 {
			name = name[:40]
		}
		mode := f.VoiceMode
		switch mode {
		case "off", "min", "normal", "detailed":
		default:
			mode = "normal"
		}
		_ = app.DB.UpdateSettings(s.sess.UserID, name, s.sess.Lang, restWarm, restMain, restCool, mode)
		_ = app.DB.SetSex(s.sess.UserID, f.Sex)
		s.auth.Mutate(ctx, func(cur auth.Session) auth.Session {
			cur.Name = name
			cur.Rest = restMain
			cur.VoiceMode = mode
			cur.Voice = mode != "off"
			return cur
		})
		saved.Update(ctx, true)
		return false
	}
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

// clampRange clamps v to [lo, hi].
func clampRange(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
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
	case "warmup":
		return content.Warmup
	case "cooldown":
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
	switch target {
	case "warmup":
		p.Warmup = fn(p.Warmup)
	case "cooldown":
		p.Cooldown = fn(p.Cooldown)
	default:
		d := clampDay(st.Day, len(p.Days))
		if d < 1 {
			return
		}
		p.Days[d-1].Items = fn(p.Days[d-1].Items)
	}
	st.JSON = p.Marshal()
	editSrc.Update(ctx, st)
}

// editList renders one CRUD section (warm-up / day main / cool-down).
//line settings.gox:651
func (s settingsPage) editList(editSrc doors.Source[editState], cat catalog.Catalog, entries []catalog.Entry, l i18n.Lang, items []content.RItem, target, open string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("ol"); if __e != nil { return }
		{
//line settings.gox:652
			__e = __c.Set("class", "cyc-list edit"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:653
			for ii := range items {
//line settings.gox:655
				it := items[ii]
				key := target + ":" + strconv.Itoa(ii)

				__e = __c.Init("li"); if __e != nil { return }
				{
//line settings.gox:658
					__e = __c.Set("class", "cyc-row"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line settings.gox:659
						__e = __c.Set("class", "cyc-num"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("select"); if __e != nil { return }
					{
//line settings.gox:660
						__e = __c.Set("class", "cyc-select"); if __e != nil { return }
//line settings.gox:660
						__e = __c.Modify(doors.AChange{On: s.editReplace(editSrc, cat, target, ii)}); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:661
						__e = __c.Any(s.exOptions(entries, l, it.ID)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.InitVoid("input"); if __e != nil { return }
					{
//line settings.gox:663
						__e = __c.Set("class", "cyc-val"); if __e != nil { return }
//line settings.gox:663
						__e = __c.Set("type", "number"); if __e != nil { return }
//line settings.gox:663
						__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line settings.gox:663
						__e = __c.Set("min", "1"); if __e != nil { return }
//line settings.gox:663
						__e = __c.Set("max", "600"); if __e != nil { return }
//line settings.gox:663
						__e = __c.Set("value", it.Value); if __e != nil { return }
//line settings.gox:663
						__e = __c.Modify(doors.AChange{On: s.editSetValue(editSrc, target, ii)}); if __e != nil { return }
					}
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line settings.gox:664
						__e = __c.Set("class", "cyc-unit"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:664
						__e = __c.Any(unitWord(l, it.Unit, false)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:665
					__e = (doors.AClick{On: s.editPeek(editSrc, target, ii)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:665
							__e = __c.Set("class", "cyc-op cyc-info"); if __e != nil { return }
//line settings.gox:665
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:665
							__e = __c.Set("aria-label", i18n.T(l, "settings.view")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("ⓘ"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:666
					__e = (doors.AClick{On: s.editMove(editSrc, target, ii, -1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:666
							__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:666
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:666
							__e = __c.Set("aria-label", i18n.T(l, "settings.move_up")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("↑"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:667
					__e = (doors.AClick{On: s.editMove(editSrc, target, ii, 1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:667
							__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:667
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:667
							__e = __c.Set("aria-label", i18n.T(l, "settings.move_down")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("↓"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line settings.gox:668
					__e = (doors.AClick{On: s.editDelete(editSrc, target, ii)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line settings.gox:668
							__e = __c.Set("class", "cyc-op danger"); if __e != nil { return }
//line settings.gox:668
							__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:668
							__e = __c.Set("aria-label", i18n.T(l, "settings.remove")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("✕"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line settings.gox:670
				if open == key {
					__e = __c.Init("li"); if __e != nil { return }
					{
//line settings.gox:671
						__e = __c.Set("class", "cyc-peek"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:671
						__e = __c.Any(s.exDetail(l, cat, it.ID)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line settings.gox:675
		__e = (doors.AClick{On: s.editAdd(editSrc, cat, target)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("button"); if __e != nil { return }
			{
//line settings.gox:675
				__e = __c.Set("class", "btn small cyc-add"); if __e != nil { return }
//line settings.gox:675
				__e = __c.Set("type", "button"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("＋ "); if __e != nil { return }
//line settings.gox:675
				__e = __c.Any(i18n.T(l, "settings.add_exercise")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line settings.gox:676
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
	case "warmup":
		id = "W03"
	case "cooldown":
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
		name := strings.TrimSpace(nameSrc.Get())
		if name == "" {
			name = "My program"
		}
		if len(name) > 40 {
			name = name[:40]
		}
		p.Name = name
		_ = app.SetUserProgram(s.sess.UserID, p.Marshal())
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
}

func mutateProgram(ctx context.Context, editSrc doors.Source[editState], fn func(*content.Resolved)) {
	st := editSrc.Get()
	p, err := content.ParseResolved(st.JSON)
	if err != nil {
		return
	}
	fn(&p)
	p.WarmupRounds = content.ClampRounds(p.WarmupRounds)
	for i := range p.Days {
		p.Days[i].Rounds = content.ClampRounds(p.Days[i].Rounds)
	}
	st.JSON = p.Marshal()
	editSrc.Update(ctx, st)
}

// setWarmupRounds sets how many times the shared warm-up runs (1–10).
func (s settingsPage) setWarmupRounds(editSrc doors.Source[editState]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		n, err := strconv.Atoi(r.Event().Value)
		if err != nil {
			return false
		}
		mutateProgram(ctx, editSrc, func(p *content.Resolved) {
			p.WarmupRounds = content.ClampRounds(n)
		})
		return false
	}
}

// setMainRounds sets how many times this day's main block runs (1–10).
func (s settingsPage) setMainRounds(editSrc doors.Source[editState]) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		n, err := strconv.Atoi(r.Event().Value)
		if err != nil {
			return false
		}
		mutateProgram(ctx, editSrc, func(p *content.Resolved) {
			d := clampDay(editSrc.Get().Day, len(p.Days))
			if d < 1 {
				return
			}
			p.Days[d-1].Rounds = content.ClampRounds(n)
		})
		return false
	}
}

// importProgram stores an uploaded/pasted cycle JSON as this user's personal plan.
func (s settingsPage) importProgram(ctx context.Context, r doors.RequestRawForm) bool {
	pf, err := r.ParseForm(1 << 20)
	if err != nil {
		return false
	}
	raw := strings.TrimSpace(pf.FormValue("progjson"))
	if raw == "" {
		return false
	}
	if app.SetUserProgram(s.sess.UserID, raw) != nil {
		return false
	}
	_ = r.After(doors.ActionLocationReload{})
	return false
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

// deleteProgram clears the personal file and returns this account to the demo.
func (s settingsPage) deleteProgram(ctx context.Context, r doors.RequestPointer) bool {
	_ = app.RestoreDemo(s.sess.UserID)
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
