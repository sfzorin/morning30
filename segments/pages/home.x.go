// Managed by GoX v0.1.36

//line home.gox:1
package pages

import (
	"context"
	"fmt"
	"time"

	"morning30/assets"
	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// homePage is the dashboard: today's session, the date-based streak calendar,
// and the separate looping 30-day cycle track.
type homePage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

//line home.gox:27
func (h homePage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:29
		l := i18n.Lang(h.sess.Lang)
		today := app.Today()
		pos := app.CurrentPosition(h.sess.UserID)
		streak, _ := app.DB.Streak(h.sess.UserID, today)
		doneToday, _ := app.DB.HasActivity(h.sess.UserID, today)

		cycle := pos/content.TotalDays + 1
		curDay := pos%content.TotalDays + 1
		w := app.UserWorkout(h.sess.UserID, curDay)
		mainRounds := 1
		for _, item := range w.Items {
			if item.Slot == content.Main && item.Round > mainRounds {
				mainRounds = item.Round
			}
		}

		startLabel := i18n.T(l, "home.start")
		if pos > 0 {
			startLabel = i18n.T(l, "home.continue")
		}
		if doneToday {
			startLabel = i18n.T(l, "home.repeat")
		}

		greeting := i18n.T(l, "home.morning")
		if h.sess.Name != "" {
			greeting += ", " + h.sess.Name
		}
		greeting += "!"

		monthOffset := doors.NewSource(0)
		hasCustom := app.HasCustomProgram(h.sess.UserID)
		resumeData := map[string]any{"userId": h.sess.UserID, "position": pos, "programKey": app.WorkoutProgramKey(h.sess.UserID), "streak": streak, "restartAt": app.CycleRestartAt(h.sess.UserID)}

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line home.gox:63
			__e = __c.Any(i18n.T(l, "app.name")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line home.gox:64
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line home.gox:65
			__e = __c.Any(header{sess: h.sess, auth: h.auth, path: h.path, streak: streak, title: greeting}); if __e != nil { return }
//line home.gox:67
			if h.sess.IsGuest {
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:68
					__e = __c.Set("class", "guest-note"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:69
						__e = __c.Any(i18n.T(l, "guest.note")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line home.gox:70
					__e = (doors.ALink{Model: path.Path{Page: path.Register}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("a"); if __e != nil { return }
						{
//line home.gox:70
							__e = __c.Set("class", "btn primary small"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line home.gox:70
							__e = __c.Any(i18n.T(l, "auth.register")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:74
				__e = __c.Set("class", "hero"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:75
				if doneToday {
					__e = __c.Init("p"); if __e != nil { return }
					{
//line home.gox:76
						__e = __c.Set("class", "done-banner"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:76
						__e = __c.Any(i18n.T(l, "home.done_today")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
//line home.gox:78
				__e = (doors.ALink{Model: path.Path{Page: path.Settings}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line home.gox:78
						__e = __c.Set("class", "hero-program"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:78
						__e = __c.Any(app.ProgramTitle(h.sess.UserID, h.sess.Lang)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:79
					__e = __c.Set("class", "hero-cycle"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:79
					__e = __c.Many(i18n.T(l, "home.cycle"), " ", cycle); if __e != nil { return }
//line home.gox:79
					if !w.FixedDifficulty {
//line home.gox:79
						__e = __c.Any("  ·  " + levelLabel(l, h.sess.Level)); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:80
					__e = __c.Set("class", "day-num"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:80
					__e = __c.Many(i18n.T(l, "home.day"), " ", curDay); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:81
					__e = __c.Set("class", "day-meta"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:82
						__e = __c.Many(i18n.T(l, "home.rounds"), ": ", mainRounds); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:83
						__e = __c.Many(len(w.Items), " ", i18n.T(l, "home.exercises")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:84
						__e = __c.Many(i18n.T(l, "home.est"), " ", minutes(w.EstSec), " ", i18n.T(l, "home.min")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:86
				__e = (doors.ALink{Model: path.Path{Page: path.Start}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line home.gox:86
						__e = __c.Set("class", "btn primary big"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:86
						__e = __c.Any(startLabel); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("a"); if __e != nil { return }
				{
//line home.gox:87
					__e = __c.Set("class", "btn resume-workout hidden"); if __e != nil { return }
//line home.gox:87
					__e = __c.Set("href", "/"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:87
					__e = __c.Any(i18n.T(l, "home.resume_workout")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("p"); if __e != nil { return }
				{
//line home.gox:88
					__e = __c.Set("class", "home-save-status hidden"); if __e != nil { return }
//line home.gox:88
					__e = __c.Set("role", "status"); if __e != nil { return }
//line home.gox:88
					__e = __c.Set("data-pending", i18n.T(l, "done.save_pending")); if __e != nil { return }
//line home.gox:88
					__e = __c.Set("data-saved", i18n.T(l, "done.saved")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line home.gox:91
			if !hasCustom {
//line home.gox:92
				__e = __c.Any(h.courses(hasCustom)); if __e != nil { return }
			}
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:95
				__e = __c.Set("class", "streak-row"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:96
					__e = __c.Set("class", "fire"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("🔥"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:97
					__e = __c.Set("class", "streak-n"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:97
					__e = __c.Any(streak); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:98
					__e = __c.Set("class", "streak-label"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:98
					__e = __c.Many(i18n.T(l, "home.streak"), " · ", streak, " ", i18n.T(l, "home.days")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:102
				__e = __c.Set("class", "calendar"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:103
				__e = __c.Any(h.month(monthOffset)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:107
				__e = __c.Set("class", "cycle"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h3"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:108
					__e = __c.Many(i18n.T(l, "home.cycle"), " ", cycle); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:109
				if !w.FixedDifficulty {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line home.gox:110
						__e = __c.Set("class", "diff"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line home.gox:111
							__e = __c.Set("class", "diff-bar"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line home.gox:112
							__e = __c.Set("class", "diff-label"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line home.gox:112
							__e = __c.Any(i18n.T(l, "home.diff")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:115
					__e = __c.Set("class", "cyc-grid"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:116
					for d := 1; d <= content.TotalDays; d++ {
//line home.gox:117
						__e = __c.Any(h.cycleCell(d, curDay)); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line home.gox:121
			if hasCustom {
//line home.gox:122
				__e = __c.Any(h.courses(hasCustom)); if __e != nil { return }
			}
			__e = __c.Init("script"); if __e != nil { return }
			{
//line home.gox:124
				__e = __c.Set("src", assets.WorkoutHome()); if __e != nil { return }
				__e = __c.Set("inline", true); if __e != nil { return }
//line home.gox:124
				__e = __c.Set("data:resume", resumeData); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Raw(""); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:126
}

// courses lists the "… in 30 days" series. Tapping a course asks first; only
// "Start course" copies it into this account's program (position and streak
// stay). Accounts still on the demo see the list right under today's session,
// accounts with their own program at the bottom of the page.
//line home.gox:132
func (h homePage) courses(hasCustom bool) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:134
		l := i18n.Lang(h.sess.Lang)
		restW, restM, restC := app.DB.Rests(h.sess.UserID)
		rests := content.Rests{Warmup: restW, Main: restM, Cooldown: restC}
		picked := doors.NewSource("")

		__e = __c.Init("section"); if __e != nil { return }
		{
//line home.gox:139
			__e = __c.Set("class", "courses"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("h3"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:140
				__e = __c.Any(i18n.T(l, "home.courses")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line home.gox:141
			__e = __c.Any(picked.Bind(func(key string) gox.Elem {
			return gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:142
					__e = __c.Set("class", "course-list"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:143
					for _, c := range content.Series30() {
//line home.gox:145
						r := c.Resolve()
						first := minutes(r.Workout(1, h.sess.Level, rests).EstSec)
						peak := minutes(r.Workout(content.TotalDays-1, h.sess.Level, rests).EstSec)
						cls := "course-card"
						if c.Key == key {
							cls += " on"
						}

//line home.gox:153
						__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
						picked.Update(ctx, c.Key)
						return false
					}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line home.gox:156
								__e = __c.Set("class", cls); if __e != nil { return }
//line home.gox:156
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:157
									__e = __c.Set("class", "course-title"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:157
									__e = __c.Any(c.TitleIn(h.sess.Lang)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:158
									__e = __c.Set("class", "course-meta"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:158
									__e = __c.Many(first, "–", peak, " ", i18n.T(l, "home.min")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:162
				if key != "" {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line home.gox:163
						__e = __c.Set("class", "confirm-reset course-confirm"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line home.gox:164
							__e = __c.Set("class", "confirm-q calm"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line home.gox:164
							__e = __c.Any(i18n.T(l, "settings.course_note")); if __e != nil { return }
//line home.gox:164
							if hasCustom {
//line home.gox:164
								__e = __c.Any(" "); if __e != nil { return }
								__e = __c.Init("b"); if __e != nil { return }
								{
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:164
									__e = __c.Any(i18n.T(l, "settings.course_replace")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line home.gox:165
							__e = __c.Set("class", "confirm-actions"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line home.gox:166
							__e = (doors.AClick{On: h.useCourse(key)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line home.gox:166
									__e = __c.Set("class", "btn primary small"); if __e != nil { return }
//line home.gox:166
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:166
									__e = __c.Any(i18n.T(l, "settings.course_use")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
//line home.gox:167
							__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
							picked.Update(ctx, "")
							return false
						}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line home.gox:170
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line home.gox:170
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:170
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
//line home.gox:174
		})); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:176
}

// useCourse installs a 30-day course as this account's program and reloads.
func (h homePage) useCourse(key string) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		if app.UseCourse(h.sess.UserID, key, h.sess.Lang) != nil {
			return false
		}
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
}

// month renders the date calendar for the current month ± offset, reactively.
//line home.gox:190
func (h homePage) month(monthOffset doors.Source[int]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:191
		__e = __c.Any(monthOffset.Bind(func(off int) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line home.gox:193
			now := time.Now()
			base := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			disp := base.AddDate(0, off, 0)
			year := disp.Year()
			month := int(disp.Month())
			daysIn := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local).Day()
			lead := (int(disp.Weekday()) + 6) % 7 // Monday-first leading blanks
			first := disp.Format("2006-01-02")
			last := time.Date(year, time.Month(month), daysIn, 0, 0, 0, 0, time.Local).Format("2006-01-02")
			acts, _ := app.DB.ActivityInRange(h.sess.UserID, first, last)
			todayStr := app.Today()
			l := i18n.Lang(h.sess.Lang)

			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:206
				__e = __c.Set("class", "cal-head"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:207
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				monthOffset.Mutate(ctx, func(o int) int { return o - 1 })
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line home.gox:210
						__e = __c.Set("class", "cal-nav"); if __e != nil { return }
//line home.gox:210
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("‹"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:211
					__e = __c.Set("class", "cal-title"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:211
					__e = __c.Many(i18n.Month(l, month), " ", year); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:212
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				monthOffset.Mutate(ctx, func(o int) int { return o + 1 })
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line home.gox:215
						__e = __c.Set("class", "cal-nav"); if __e != nil { return }
//line home.gox:215
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("›"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:217
				__e = __c.Set("class", "cal-week"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:218
				for wd := 0; wd < 7; wd++ {
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:219
						__e = __c.Any(i18n.Weekday(l, wd)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:222
				__e = __c.Set("class", "cal-grid month"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:223
				for b := 0; b < lead; b++ {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line home.gox:224
						__e = __c.Set("class", "cal-blank"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
//line home.gox:226
				for dd := 1; dd <= daysIn; dd++ {
//line home.gox:227
					__e = __c.Any(h.dayCell(year, month, dd, acts, todayStr)); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })
//line home.gox:230
	})); if __e != nil { return }
	return })
//line home.gox:231
}

// dayCell renders one date cell, marked when the user worked out that day.
//line home.gox:234
func (h homePage) dayCell(year, month, dd int, acts map[string]bool, todayStr string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:236
		ds := fmt.Sprintf("%04d-%02d-%02d", year, month, dd)
		cls := "cal-day"
		if acts[ds] {
			cls += " act"
		}
		if ds == todayStr {
			cls += " today"
		}

		__e = __c.Init("div"); if __e != nil { return }
		{
//line home.gox:245
			__e = __c.Set("class", cls); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line home.gox:246
			if acts[ds] {
				__e = __c.Text("🔥"); if __e != nil { return }
			} else  {
//line home.gox:249
				__e = __c.Any(dd); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:252
}

// cycleCell renders one of the 30 program days; sequential, looping each cycle.
//line home.gox:255
func (h homePage) cycleCell(d, current int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:257
		done := d < current
		isCurrent := d == current
		locked := d > current
		cls := "cell"
		if done {
			cls += " done"
		}
		if isCurrent {
			cls += " current"
		}
		if locked {
			cls += " locked"
		}

//line home.gox:272
		__e = (doors.ALink{Model: path.Path{Page: path.Day, Day: d}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("a"); if __e != nil { return }
			{
//line home.gox:272
				__e = __c.Set("class", cls); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:273
				if done {
					__e = __c.Text("✓"); if __e != nil { return }
				} else  {
//line home.gox:276
					__e = __c.Any(d); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line home.gox:279
}
