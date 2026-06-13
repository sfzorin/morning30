// Managed by GoX v0.1.36

//line settings.gox:1
package pages

import (
	"context"
	"strconv"
	"strings"

	"danicc/internal/app"
	"danicc/internal/auth"
	"danicc/internal/content"
	"danicc/internal/i18n"
	"danicc/internal/program"
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

// settingsPage lets the user change rest time, language, voice and reset progress.
type settingsPage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

//line settings.gox:33
func (s settingsPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line settings.gox:35
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
		prog := app.UserProgram(s.sess.UserID) // the active program (built-in or custom)
		daySource := doors.NewSource(1)

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:51
			__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line settings.gox:52
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line settings.gox:53
			__e = __c.Any(header{sess: s.sess, auth: s.auth, path: s.path, streak: streak}); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line settings.gox:54
				__e = __c.Set("class", "settings"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h1"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line settings.gox:55
					__e = __c.Any(i18n.T(l, "settings.title")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line settings.gox:57
				vm := s.sess.VoiceMode
				if vm == "" {
					vm = "normal"
				}

				__e = __c.Init("div"); if __e != nil { return }
				{
//line settings.gox:62
					__e = __c.Set("class", "settings-card"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:64
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:65
							__e = __c.Any(i18n.T(l, "settings.language")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:66
						__e = __c.Any(langSwitch{sess: s.sess, auth: s.auth, persist: true}); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line settings.gox:69
					__e = (doors.ASubmit[settingsForm]{On: s.save(saved)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("form"); if __e != nil { return }
						{
//line settings.gox:69
							__e = __c.Set("class", "settings-form"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:70
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:71
									__e = __c.Any(i18n.T(l, "auth.name")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:72
									__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:72
									__e = __c.Set("name", "name"); if __e != nil { return }
//line settings.gox:72
									__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:72
									__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:72
									__e = __c.Set("value", s.sess.Name); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:74
								__e = __c.Set("class", "srow col"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:75
									__e = __c.Set("class", "row-between"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:76
										__e = __c.Any(i18n.T(l, "settings.rest")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("output"); if __e != nil { return }
									{
//line settings.gox:77
										__e = __c.Set("class", "rest-out"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:78
										__e = __c.Any(restVal.Bind(func(v int) gox.Elem {
									return gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
//line settings.gox:79
										__e = __c.Many(v, " ", i18n.T(l, "workout.sec")); if __e != nil { return }
									return })
//line settings.gox:80
								})); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.InitVoid("input"); if __e != nil { return }
								{
//line settings.gox:84
									__e = __c.Set("type", "range"); if __e != nil { return }
//line settings.gox:85
									__e = __c.Set("name", "rest"); if __e != nil { return }
//line settings.gox:86
									__e = __c.Set("min", "10"); if __e != nil { return }
//line settings.gox:87
									__e = __c.Set("max", "40"); if __e != nil { return }
//line settings.gox:88
									__e = __c.Set("step", "5"); if __e != nil { return }
//line settings.gox:89
									__e = __c.Set("value", rest); if __e != nil { return }
//line settings.gox:90
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
//line settings.gox:97
								__e = __c.Set("class", "srow"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:98
									__e = __c.Any(i18n.T(l, "settings.voice")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("select"); if __e != nil { return }
								{
//line settings.gox:99
									__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:99
									__e = __c.Set("name", "voice_mode"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:100
										__e = __c.Set("value", "off"); if __e != nil { return }
//line settings.gox:100
										__e = __c.Set("selected", func() any { return vm == "off" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:100
										__e = __c.Any(i18n.T(l, "voice.off")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:101
										__e = __c.Set("value", "min"); if __e != nil { return }
//line settings.gox:101
										__e = __c.Set("selected", func() any { return vm == "min" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:101
										__e = __c.Any(i18n.T(l, "voice.min")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:102
										__e = __c.Set("value", "normal"); if __e != nil { return }
//line settings.gox:102
										__e = __c.Set("selected", func() any { return vm == "normal" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:102
										__e = __c.Any(i18n.T(l, "voice.normal")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:103
										__e = __c.Set("value", "detailed"); if __e != nil { return }
//line settings.gox:103
										__e = __c.Set("selected", func() any { return vm == "detailed" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:103
										__e = __c.Any(i18n.T(l, "voice.detailed")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line settings.gox:106
								__e = __c.Set("class", "srow save-row"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:107
									__e = __c.Set("class", "btn primary"); if __e != nil { return }
//line settings.gox:107
									__e = __c.Set("type", "submit"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:107
									__e = __c.Any(i18n.T(l, "settings.save")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:108
								__e = __c.Any(saved.Bind(func(ok bool) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:109
								if ok {
									__e = __c.Init("span"); if __e != nil { return }
									{
//line settings.gox:110
										__e = __c.Set("class", "saved-msg"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:110
										__e = __c.Any(i18n.T(l, "settings.saved")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							return })
//line settings.gox:112
						})); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
					__e = __c.Init("details"); if __e != nil { return }
					{
//line settings.gox:117
						__e = __c.Set("class", "cycle-details"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("summary"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:118
							__e = __c.Any(i18n.T(l, "settings.cycle")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line settings.gox:121
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:122
								__e = __c.Any(i18n.T(l, "settings.program")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:123
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:123
								__e = __c.Modify(doors.AChange{On: s.selectProgram}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("option"); if __e != nil { return }
								{
//line settings.gox:124
									__e = __c.Set("value", "builtin"); if __e != nil { return }
//line settings.gox:124
									__e = __c.Set("selected", func() any { return active != "custom" }()); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Text("dani.cc"); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
//line settings.gox:125
								if hasCustom {
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:126
										__e = __c.Set("value", "custom"); if __e != nil { return }
//line settings.gox:126
										__e = __c.Set("selected", func() any { return active == "custom" }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:126
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
//line settings.gox:131
							__e = __c.Set("class", "srow"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:132
								__e = __c.Any(i18n.T(l, "home.day")); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("select"); if __e != nil { return }
							{
//line settings.gox:133
								__e = __c.Set("class", "lang-select compact"); if __e != nil { return }
//line settings.gox:133
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
							if n := r.Event().Number; n != nil {
								daySource.Update(ctx, int(*n))
							}
							return false
						}}); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:139
								for d := 1; d <= len(prog.Days); d++ {
									__e = __c.Init("option"); if __e != nil { return }
									{
//line settings.gox:140
										__e = __c.Set("value", d); if __e != nil { return }
//line settings.gox:140
										__e = __c.Set("selected", func() any { return d == 1 }()); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:140
										__e = __c.Any(d); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:145
						__e = (doors.ARawSubmit{On: s.saveProgram(daySource)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("form"); if __e != nil { return }
							{
//line settings.gox:145
								__e = __c.Set("class", "cycle-form"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:146
								__e = __c.Any(daySource.Bind(func(d int) gox.Elem {
							return gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
//line settings.gox:148
								idx := d - 1
								if idx < 0 {
									idx = 0
								}
								if idx >= len(prog.Days) {
									idx = len(prog.Days) - 1
								}
								rd := prog.Days[idx]

								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:157
									__e = __c.Set("class", "cyc-head"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:157
									__e = __c.Many(rd.Label, " · ", len(rd.Items), " ", i18n.T(l, "home.exercises")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("ol"); if __e != nil { return }
								{
//line settings.gox:158
									__e = __c.Set("class", "cyc-list"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:159
									for ii, it := range rd.Items {
										__e = __c.Init("li"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
											__e = __c.Init("span"); if __e != nil { return }
											{
//line settings.gox:161
												__e = __c.Set("class", "cyc-name"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line settings.gox:161
												__e = __c.Any(i18n.Exercise(l, it.ID)); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
											__e = __c.Init("span"); if __e != nil { return }
											{
//line settings.gox:162
												__e = __c.Set("class", "cyc-edit"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
												__e = __c.InitVoid("input"); if __e != nil { return }
												{
//line settings.gox:163
													__e = __c.Set("class", "cyc-val"); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("type", "number"); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("inputmode", "numeric"); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("min", "1"); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("max", "600"); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("name", "v" + strconv.Itoa(ii)); if __e != nil { return }
//line settings.gox:163
													__e = __c.Set("value", it.Value); if __e != nil { return }
												}
												__e = __c.Submit(); if __e != nil { return }
												__e = __c.Init("span"); if __e != nil { return }
												{
//line settings.gox:164
													__e = __c.Set("class", "cyc-unit"); if __e != nil { return }
													__e = __c.Submit(); if __e != nil { return }
//line settings.gox:164
													__e = __c.Any(unitWord(l, it.Unit, it.PerSide)); if __e != nil { return }
												}
												__e = __c.Close(); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
								}
								__e = __c.Close(); if __e != nil { return }
							return })
//line settings.gox:169
						})); if __e != nil { return }
								__e = __c.Init("div"); if __e != nil { return }
								{
//line settings.gox:170
									__e = __c.Set("class", "srow cyc-save"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.InitVoid("input"); if __e != nil { return }
									{
//line settings.gox:171
										__e = __c.Set("class", "text-input compact"); if __e != nil { return }
//line settings.gox:171
										__e = __c.Set("name", "pname"); if __e != nil { return }
//line settings.gox:171
										__e = __c.Set("type", "text"); if __e != nil { return }
//line settings.gox:171
										__e = __c.Set("maxlength", "40"); if __e != nil { return }
//line settings.gox:171
										__e = __c.Set("value", prog.Name); if __e != nil { return }
//line settings.gox:171
										__e = __c.Set("placeholder", i18n.T(l, "settings.version_name")); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("button"); if __e != nil { return }
									{
//line settings.gox:172
										__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line settings.gox:172
										__e = __c.Set("type", "submit"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line settings.gox:172
										__e = __c.Any(i18n.T(l, "settings.save_version")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
//line settings.gox:176
						if hasCustom {
//line settings.gox:177
							__e = (doors.AClick{On: s.deleteProgram}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line settings.gox:177
									__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:177
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:177
									__e = __c.Any(i18n.T(l, "settings.delete_version")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:181
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:182
							__e = __c.Any(i18n.T(l, "settings.program")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("a"); if __e != nil { return }
						{
//line settings.gox:183
							__e = __c.Set("class", "btn small"); if __e != nil { return }
//line settings.gox:183
							__e = __c.Modify(doors.ResourceBytes(program.ExportJSON())); if __e != nil { return }
//line settings.gox:183
							__e = __c.Set("name", "dani-program.json"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:183
							__e = __c.Any(i18n.T(l, "settings.download")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:186
						__e = __c.Set("class", "srow"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line settings.gox:187
							__e = __c.Any(i18n.T(l, "settings.account")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line settings.gox:188
						if s.sess.IsGuest {
//line settings.gox:189
							__e = (doors.ALink{Model: path.Path{Page: path.Register}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("a"); if __e != nil { return }
								{
//line settings.gox:189
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line settings.gox:189
									__e = __c.Any(i18n.T(l, "auth.register")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						} else  {
							__e = __c.Init("span"); if __e != nil { return }
							{
//line settings.gox:191
								__e = __c.Set("class", "email"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:191
								__e = __c.Any(s.sess.Email); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line settings.gox:195
						__e = __c.Set("class", "srow save-row"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line settings.gox:196
						__e = (doors.AClick{On: s.reset}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line settings.gox:196
								__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line settings.gox:196
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line settings.gox:196
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
//line settings.gox:201
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

// saveProgram stores the edited day into the user's custom program (built from
// the currently-active program), makes it active, and reloads.
func (s settingsPage) saveProgram(day doors.Source[int]) func(context.Context, doors.RequestRawForm) bool {
	return func(ctx context.Context, r doors.RequestRawForm) bool {
		pf, err := r.ParseForm(1 << 20)
		if err != nil {
			return false
		}
		prog := app.UserProgram(s.sess.UserID)
		if len(prog.Days) == 0 {
			return false
		}
		d := day.Get()
		if d < 1 {
			d = 1
		}
		if d > len(prog.Days) {
			d = len(prog.Days)
		}
		items := prog.Days[d-1].Items
		for i := range items {
			if v := pf.FormValue("v" + strconv.Itoa(i)); v != "" {
				if n, e := strconv.Atoi(v); e == nil && n >= 1 && n <= 600 {
					items[i].Value = n
				}
			}
		}
		name := strings.TrimSpace(pf.FormValue("pname"))
		if name == "" {
			name = "my dani.cc"
		}
		if len(name) > 40 {
			name = name[:40]
		}
		prog.Name = name
		_ = app.DB.SetProgramJSON(s.sess.UserID, prog.Marshal())
		_ = app.DB.SetActiveProgram(s.sess.UserID, "custom")
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
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
