// Managed by GoX v0.1.36

//line settings.gox:1
package pages

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"danicc/assets"
	"danicc/internal/app"
	"danicc/internal/auth"
	"danicc/internal/catalog"
	"danicc/internal/content"
	"danicc/internal/i18n"
	"danicc/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// settingsForm is the decoded settings submission.
type settingsForm struct {
	Name      string `form:"name"`
	Rest      int    `form:"rest"`
	VoiceMode string `form:"voice_mode"`
}

// editState is the editor's working copy: the selected day plus the program as
// JSON. It is a single comparable value so one Source/Bind drives the whole
// editor (day switch + every edit) — nested Binds don't patch reliably.
type editState struct {
	Day  int
	JSON string
}

// settingsPage lets the user change rest time, language, voice and reset progress.
type settingsPage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

//line settings.gox:43
func (s settingsPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:45
		l := i18n.Lang(s.sess.Lang)
		saved := doors.NewSource(false)
		rest := s.sess.Rest
		if rest < 10 {
			rest = 10
		}
		if rest > 40 {
			rest = 40
		}
		restVal := doors.NewSource(rest)
		streak, _ := app.DB.Streak(s.sess.UserID, app.Today())
		active, _ := app.DB.ActiveProgram(s.sess.UserID)
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

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:69
			__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line settings.gox:70
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:71
			__e = __c.Any(header{sess: s.sess, auth: s.auth, path: s.path, streak: streak}); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line settings.gox:72
				__e = __c.Set("class", "settings"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h1"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line settings.gox:73
					__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line settings.gox:75
				vm := s.sess.VoiceMode
				if vm == "" {
					vm = "normal"
				}

				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:80
					__e = __c.Set("class", "settings-card"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:82
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:83
							__e = __c.Any(i18n.T(l, "settings.language")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:84
						__e = __c.Any(langSwitch{sess: s.sess, auth: s.auth, persist: true}); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:87
					__e = (doors.ASubmit[settingsForm]{On: s.save(saved)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("form"); if __e != nil { return }
						{
//line settings.gox:87
							__e = __c.Set("class", "settings-form"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:88
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:89
									__e = __c.Any(i18n.T(l, "auth.name")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:90
									__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:90
									__e = __c.Set("name", "name"); if __e != nil { return }
//line settings.gox:90
									__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:90
									__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:90
									__e = __c.Set("value", s.sess.Name); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:92
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:93
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:94
										__e = __c.Any(i18n.T(l, "settings.rest")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:95
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:96
										__e = __c.Any(restVal.Bind(func(v int) gox.Elem {
									return gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
//line settings.gox:97
										__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
									return })
//line settings.gox:98
								})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:102
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:103
									__e = __c.Set("name", "rest"); if __e != nil { return }
//line settings.gox:104
									__e = __c.Set("min", "10"); if __e != nil { return }
//line settings.gox:105
									__e = __c.Set("max", "40"); if __e != nil { return }
//line settings.gox:106
									__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:107
									__e = __c.Set("value", rest); if __e != nil { return }
//line settings.gox:108
									__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								if n := r.Event().Number; n != nil {
									restVal.Update(ctx, int(*n))
								}
								return false
							}}); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:115
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:116
									__e = __c.Any(i18n.T(l, "settings.voice")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:117
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:117
									__e = __c.Set("name", "voice_mode"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:118
										__e = __c.Set("value", "off"); if __e != nil { return }
//line settings.gox:118
										__e = __c.Set("selected", func() any { return vm == "off" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:118
										__e = __c.Any(i18n.T(l, "voice.off")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:119
										__e = __c.Set("value", "min"); if __e != nil { return }
//line settings.gox:119
										__e = __c.Set("selected", func() any { return vm == "min" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:119
										__e = __c.Any(i18n.T(l, "voice.min")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:120
										__e = __c.Set("value", "normal"); if __e != nil { return }
//line settings.gox:120
										__e = __c.Set("selected", func() any { return vm == "normal" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:120
										__e = __c.Any(i18n.T(l, "voice.normal")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:121
										__e = __c.Set("value", "detailed"); if __e != nil { return }
//line settings.gox:121
										__e = __c.Set("selected", func() any { return vm == "detailed" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:121
										__e = __c.Any(i18n.T(l, "voice.detailed")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:124
								__e = __c.Set("class", "srow save-row"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:125
									__e = __c.Set("class", "btn primary"); if __e != nil { return }
//line settings.gox:125
									__e = __c.Set("type", "submit"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:125
									__e = __c.Any(i18n.T(l, "settings.save")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:126
								__e = __c.Any(saved.Bind(func(ok bool) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:127
								if ok {
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:128
										__e = __c.Set("class", "saved-msg"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:128
										__e = __c.Any(i18n.T(l, "settings.saved")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							return })
//line settings.gox:130
						})); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
					__e = __c.Init("details"); if __e != nil { return }
					{
//line settings.gox:135
						__e = __c.Set("class", "cycle-details"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("summary"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:136
							__e = __c.Any(i18n.T(l, "settings.cycle")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:139
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:140
								__e = __c.Any(i18n.T(l, "settings.program")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:141
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:141
								__e = __c.Modify(doors.AChange{On: s.selectProgram}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:142
									__e = __c.Set("value", "builtin"); if __e != nil { return }
//line settings.gox:142
									__e = __c.Set("selected", func() any { return active != "custom" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Text("dani.cc"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:143
								if hasCustom {
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:144
										__e = __c.Set("value", "custom"); if __e != nil { return }
//line settings.gox:144
										__e = __c.Set("selected", func() any { return active == "custom" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:144
										__e = __c.Any(i18n.T(l, "settings.my_version")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:149
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:150
								__e = __c.Any(i18n.T(l, "home.day")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:151
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:151
								__e = __c.Modify(doors.AChange{On: s.setDay(editSrc)}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:152
								for d := 1; d <= len(prog.Days); d++ {
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:153
										__e = __c.Set("value", d); if __e != nil { return }
//line settings.gox:153
										__e = __c.Set("selected", func() any { return d == 1 }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:153
										__e = __c.Any(d); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:159
						__e = __c.Any(editSrc.Bind(func(st editState) gox.Elem {
						return gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
//line settings.gox:161
							p, _ := content.ParseResolved(st.JSON)
							idx := clampDay(st.Day, len(p.Days)) - 1
							if idx < 0 {
								idx = 0
							}

//line settings.gox:167
							if len(p.Days) == 0 {
							} else  {
//line settings.gox:169
								rd := p.Days[idx]
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:170
									__e = __c.Set("class", "cyc-head"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:170
									__e = __c.Many(rd.Label, " · ", len(rd.Items), " ", i18n.T(l, "home.exercises")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("ol"); if __e != nil { return }
								{
//line settings.gox:171
									__e = __c.Set("class", "cyc-list edit"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:172
									for ii := range rd.Items {
//line settings.gox:173
										it := rd.Items[ii]
										__e = __c.Init("li"); if __e != nil { return }
										{
//line settings.gox:174
											__e = __c.Set("class", "cyc-row"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
											__e = __c.Init("select"); if __e != nil { return }
											{
//line settings.gox:175
												__e = __c.Set("class", "cyc-select"); if __e != nil { return }
//line settings.gox:175
												__e = __c.Modify(doors.AChange{On: s.editReplace(editSrc, cat, ii)}); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:176
												__e = __c.Any(s.exOptions(entries, l, it.ID)); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("span"); if __e != nil { return }
											{
//line settings.gox:178
												__e = __c.Set("class", "cyc-edit"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
												__e = __c.InitVoid("input"); if __e != nil { return }
												{
//line settings.gox:179
													__e = __c.Set("class", "cyc-val"); if __e != nil { return }
//line settings.gox:179
													__e = __c.Set("type", "number"); if __e != nil { return }
//line settings.gox:179
													__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line settings.gox:179
													__e = __c.Set("min", "1"); if __e != nil { return }
//line settings.gox:179
													__e = __c.Set("max", "600"); if __e != nil { return }
//line settings.gox:179
													__e = __c.Set("value", it.Value); if __e != nil { return }
//line settings.gox:179
													__e = __c.Modify(doors.AChange{On: s.editSetValue(editSrc, ii)}); if __e != nil { return }
												}
												__e = __c.Submit(); if __e != nil { return }
												__e = __c.Init("span"); if __e != nil { return }
												{
//line settings.gox:180
													__e = __c.Set("class", "cyc-unit"); if __e != nil { return }
													__e = __c.Submit(); if __e != nil { return }
//line settings.gox:180
													__e = __c.Any(unitWord(l, it.Unit, it.PerSide)); if __e != nil { return }
												}
												__e = __c.Close(); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("span"); if __e != nil { return }
											{
//line settings.gox:182
												__e = __c.Set("class", "cyc-ops"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:183
												__e = (doors.AClick{On: s.editMove(editSrc, ii, -1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
													ctx := __c.Context(); _ = ctx
													__e = __c.Init("button"); if __e != nil { return }
													{
//line settings.gox:183
														__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:183
														__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:183
														__e = __c.Set("aria-label", i18n.T(l, "settings.move_up")); if __e != nil { return }
														__e = __c.Submit(); if __e != nil { return }
														__e = __c.Text("↑"); if __e != nil { return }
													}
													__e = __c.Close(); if __e != nil { return }
												return })); if __e != nil { return }
//line settings.gox:184
												__e = (doors.AClick{On: s.editMove(editSrc, ii, 1)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
													ctx := __c.Context(); _ = ctx
													__e = __c.Init("button"); if __e != nil { return }
													{
//line settings.gox:184
														__e = __c.Set("class", "cyc-op"); if __e != nil { return }
//line settings.gox:184
														__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:184
														__e = __c.Set("aria-label", i18n.T(l, "settings.move_down")); if __e != nil { return }
														__e = __c.Submit(); if __e != nil { return }
														__e = __c.Text("↓"); if __e != nil { return }
													}
													__e = __c.Close(); if __e != nil { return }
												return })); if __e != nil { return }
//line settings.gox:185
												__e = (doors.AClick{On: s.editDelete(editSrc, ii)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
													ctx := __c.Context(); _ = ctx
													__e = __c.Init("button"); if __e != nil { return }
													{
//line settings.gox:185
														__e = __c.Set("class", "cyc-op danger"); if __e != nil { return }
//line settings.gox:185
														__e = __c.Set("type", "button"); if __e != nil { return }
//line settings.gox:185
														__e = __c.Set("aria-label", i18n.T(l, "settings.remove")); if __e != nil { return }
														__e = __c.Submit(); if __e != nil { return }
														__e = __c.Text("✕"); if __e != nil { return }
													}
													__e = __c.Close(); if __e != nil { return }
												return })); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:190
								__e = (doors.AClick{On: s.editAdd(editSrc, cat)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:190
										__e = __c.Set("class", "btn small cyc-add"); if __e != nil { return }
//line settings.gox:190
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Text("＋ "); if __e != nil { return }
//line settings.gox:190
										__e = __c.Any(i18n.T(l, "settings.add_exercise")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
						return })
//line settings.gox:192
					})); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:194
							__e = __c.Set("class", "srow cyc-save"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line settings.gox:195
								__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:195
								__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:195
								__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:195
								__e = __c.Set("value", prog.Name); if __e != nil { return }
//line settings.gox:195
								__e = __c.Set("placeholder", i18n.T(l, "settings.version_name")); if __e != nil { return }
//line settings.gox:196
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								nameSrc.Update(ctx, r.Event().Value)
								return false
							}}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:200
							__e = (doors.AClick{On: s.saveEditor(editSrc, nameSrc)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:200
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:200
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:200
									__e = __c.Any(i18n.T(l, "settings.save_version")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:203
						if hasCustom {
//line settings.gox:204
							__e = (doors.AClick{On: s.deleteProgram}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:204
									__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:204
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:204
									__e = __c.Any(i18n.T(l, "settings.delete_version")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("details"); if __e != nil { return }
					{
//line settings.gox:209
						__e = __c.Set("class", "cycle-details"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("summary"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:210
							__e = __c.Any(i18n.T(l, "settings.exercise_json")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:211
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:212
								__e = __c.Any(i18n.T(l, "settings.exercise")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:213
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:213
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
							selSrc.Update(ctx, r.Event().Value)
							return false
						}}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:217
								__e = __c.Any(s.exOptions(entries, l, firstID)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:220
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:221
								__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line settings.gox:222
							__e = __c.Any(selSrc.Bind(func(id string) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:224
								doc, _ := cat.Doc(id)
								b := doc.Marshal()

								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:227
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:227
									__e = __c.Modify(doors.ResourceBytes(b)); if __e != nil { return }
//line settings.gox:227
									__e = __c.Set("name", id + ".json"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:227
									__e = __c.Many(id, ".json"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })
//line settings.gox:228
						})); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:230
						__e = (doors.ARawSubmit{On: s.importExercise}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("form"); if __e != nil { return }
							{
//line settings.gox:230
								__e = __c.Set("class", "ex-import-form"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("textarea"); if __e != nil { return }
								{
//line settings.gox:231
									__e = __c.Set("class", "ex-import"); if __e != nil { return }
//line settings.gox:231
									__e = __c.Set("name", "exjson"); if __e != nil { return }
//line settings.gox:231
									__e = __c.Set("rows", "3"); if __e != nil { return }
//line settings.gox:231
									__e = __c.Set("placeholder", i18n.T(l, "settings.paste_json")); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:232
									__e = __c.Set("class", "srow ex-import-row"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:233
										__e = __c.Set("class", "ex-file"); if __e != nil { return }
//line settings.gox:233
										__e = __c.Set("type", "file"); if __e != nil { return }
//line settings.gox:233
										__e = __c.Set("accept", ".json,application/json"); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:234
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:234
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:234
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
//line settings.gox:237
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
//line settings.gox:240
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:241
							__e = __c.Any(i18n.T(l, "settings.account")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:242
						if s.sess.IsGuest {
//line settings.gox:243
							__e = (doors.ALink{Model: path.Path{Page: path.Register}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:243
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:243
									__e = __c.Any(i18n.T(l, "auth.register")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						} else  {
							__e = __c.Init("span"); if __e != nil { return }
							{
//line settings.gox:245
								__e = __c.Set("class", "email"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:245
								__e = __c.Any(s.sess.Email); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:249
						__e = __c.Set("class", "srow save-row"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:250
						__e = (doors.AClick{On: s.reset}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line settings.gox:250
								__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:250
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:250
								__e = __c.Any(i18n.T(l, "settings.reset")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:255
}

// exOptions renders the exercise dropdown options grouped by slot.
//line settings.gox:258
func (s settingsPage) exOptions(entries []catalog.Entry, l i18n.Lang, cur string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:259
		__e = __c.Any(s.exGroup(entries, l, cur, "warmup", "workout.warmup")); if __e != nil { return }
//line settings.gox:260
		__e = __c.Any(s.exGroup(entries, l, cur, "main", "workout.main")); if __e != nil { return }
//line settings.gox:261
		__e = __c.Any(s.exGroup(entries, l, cur, "cooldown", "workout.cooldown")); if __e != nil { return }
	return })
//line settings.gox:262
}

//line settings.gox:264
func (s settingsPage) exGroup(entries []catalog.Entry, l i18n.Lang, cur, slot, labelKey string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("optgroup"); if __e != nil { return }
		{
//line settings.gox:265
			__e = __c.Set("label", i18n.T(l, labelKey)); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:266
			for _, e := range entries {
//line settings.gox:267
				if e.Slot == slot {
					__e = __c.Init("option"); if __e != nil { return }
					{
//line settings.gox:268
						__e = __c.Set("value", e.ID); if __e != nil { return }
//line settings.gox:268
						__e = __c.Set("selected", func() any { return e.ID == cur }()); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:268
						__e = __c.Any(e.Name); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line settings.gox:272
}

// save persists rest + voice and updates the reactive session.
func (s settingsPage) save(saved doors.Source[bool]) func(context.Context, doors.RequestForm[settingsForm]) bool {
	return func(ctx context.Context, r doors.RequestForm[settingsForm]) bool {
		f := r.Data()
		rest := f.Rest
		if rest < 10 {
			rest = 10
		}
		if rest > 40 {
			rest = 40
		}
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
		_ = app.DB.UpdateSettings(s.sess.UserID, name, s.sess.Lang, rest, mode)
		s.auth.Mutate(ctx, func(cur auth.Session) auth.Session {
			cur.Name = name
			cur.Rest = rest
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

// mutateProg parses the working-copy JSON, applies fn to the selected day's
// item list, and stores it back (triggering a re-render).
func mutateProg(ctx context.Context, editSrc doors.Source[editState], fn func(items []content.RItem) []content.RItem) {
	st := editSrc.Get()
	p, err := content.ParseResolved(st.JSON)
	if err != nil {
		return
	}
	d := clampDay(st.Day, len(p.Days))
	if d < 1 {
		return
	}
	p.Days[d-1].Items = fn(p.Days[d-1].Items)
	st.JSON = p.Marshal()
	editSrc.Update(ctx, st)
}

// editReplace swaps the exercise at index ii for the one picked in the dropdown,
// resetting unit/side/slot from the chosen exercise and converting the value.
func (s settingsPage) editReplace(editSrc doors.Source[editState], cat catalog.Catalog, ii int) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
	return func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
		newID := r.Event().Value
		def, ok := cat.Def(newID)
		if !ok {
			return false
		}
		mutateProg(ctx, editSrc, func(its []content.RItem) []content.RItem {
			if ii < 0 || ii >= len(its) {
				return its
			}
			old := its[ii]
			its[ii] = content.RItem{
				ID: newID, Unit: def.Unit, PerSide: def.PerSide, Slot: def.Slot,
				Value: content.ConvertValue(old.Unit, def.Unit, old.Value), Round: old.Round,
			}
			return its
		})
		return false
	}
}

// editSetValue updates the value (reps/seconds) at index ii.
func (s settingsPage) editSetValue(editSrc doors.Source[editState], ii int) func(context.Context, doors.RequestEvent[doors.ChangeEvent]) bool {
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
		mutateProg(ctx, editSrc, func(its []content.RItem) []content.RItem {
			if ii >= 0 && ii < len(its) {
				its[ii].Value = v
			}
			return its
		})
		return false
	}
}

// editMove shifts the item at index ii by delta (-1 up / +1 down).
func (s settingsPage) editMove(editSrc doors.Source[editState], ii, delta int) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		mutateProg(ctx, editSrc, func(its []content.RItem) []content.RItem {
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
func (s settingsPage) editDelete(editSrc doors.Source[editState], ii int) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		mutateProg(ctx, editSrc, func(its []content.RItem) []content.RItem {
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

// editAdd appends a default exercise (push-ups) to the selected day.
func (s settingsPage) editAdd(editSrc doors.Source[editState], cat catalog.Catalog) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		it := defaultItem(cat)
		mutateProg(ctx, editSrc, func(its []content.RItem) []content.RItem {
			return append(append([]content.RItem{}, its...), it)
		})
		return false
	}
}

// defaultItem is the exercise added by the "add" button.
func defaultItem(cat catalog.Catalog) content.RItem {
	id := "M01"
	unit, perSide, slot := content.Reps, false, content.Main
	if def, ok := cat.Def(id); ok {
		unit, perSide, slot = def.Unit, def.PerSide, def.Slot
	}
	val := 10
	if unit == content.Seconds {
		val = 30
	}
	return content.RItem{ID: id, Unit: unit, Value: val, PerSide: perSide, Slot: slot, Round: 1}
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
			name = "my dani.cc"
		}
		if len(name) > 40 {
			name = name[:40]
		}
		p.Name = name
		_ = app.DB.SetProgramJSON(s.sess.UserID, p.Marshal())
		_ = app.DB.SetActiveProgram(s.sess.UserID, "custom")
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
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
	if !validDoc(d) {
		return false
	}
	_ = app.SaveCustomExercise(s.sess.UserID, d)
	_ = r.After(doors.ActionLocationReload{})
	return false
}

// validDoc sanity-checks an imported exercise document.
func validDoc(d catalog.Doc) bool {
	if strings.TrimSpace(d.ID) == "" {
		return false
	}
	switch content.Slot(d.Slot) {
	case content.Warmup, content.Main, content.Cooldown:
	default:
		return false
	}
	switch content.Unit(d.Unit) {
	case content.Reps, content.Seconds, content.Breaths:
	default:
		return false
	}
	return true
}

// selectProgram switches the active program (builtin | custom) and reloads.
func (s settingsPage) selectProgram(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
	app.DB.SetActiveProgram(s.sess.UserID, r.Event().Value)
	_ = r.After(doors.ActionLocationReload{})
	return false
}

// deleteProgram removes the custom program and reverts to the built-in (which
// can never be deleted), then reloads.
func (s settingsPage) deleteProgram(ctx context.Context, r doors.RequestPointer) bool {
	_ = app.DB.SetProgramJSON(s.sess.UserID, "")
	_ = app.DB.SetActiveProgram(s.sess.UserID, "builtin")
	_ = r.After(doors.ActionLocationReload{})
	return false
}

// reset clears all progress and returns to the home screen.
func (s settingsPage) reset(ctx context.Context, r doors.RequestPointer) bool {
	_ = app.DB.ResetProgress(s.sess.UserID)
	s.path.Update(ctx, path.Path{Page: path.Home})
	return false
}
