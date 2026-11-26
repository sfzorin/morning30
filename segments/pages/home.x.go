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
		selected := app.SelectedPrograms(h.sess.UserID)
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
				__e = (doors.ALink{Model: path.Path{Page: path.Programs}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
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
					if !w.FixedDifficulty && h.sess.Level != 0 {
//line home.gox:79
						__e = __c.Many("  ·  ", i18n.T(l, "home.level"), " ", levelLabel(l, h.sess.Level)); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:80
				__e = (doors.ALink{Model: path.Path{Page: path.Today}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line home.gox:80
						__e = __c.Set("class", "day-num day-link"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:80
						__e = __c.Many(i18n.T(l, "home.day"), " ", curDay); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line home.gox:80
							__e = __c.Set("class", "chev"); if __e != nil { return }
//line home.gox:80
							__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("›"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
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
			__e = __c.Any(h.programs(selected)); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:93
				__e = __c.Set("class", "streak-row"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:94
					__e = __c.Set("class", "fire"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("🔥"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:95
					__e = __c.Set("class", "streak-n"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:95
					__e = __c.Any(streak); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:96
					__e = __c.Set("class", "streak-label"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:96
					__e = __c.Many(i18n.T(l, "home.streak"), " · ", streak, " ", i18n.T(l, "home.days")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:100
				__e = __c.Set("class", "calendar"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:101
				__e = __c.Any(h.month(monthOffset)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line home.gox:105
				__e = __c.Set("class", "cycle"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h3"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:106
					__e = __c.Many(i18n.T(l, "home.cycle"), " ", cycle); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:107
				if !w.FixedDifficulty {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line home.gox:108
						__e = __c.Set("class", "diff"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("div"); if __e != nil { return }
						{
//line home.gox:109
							__e = __c.Set("class", "diff-bar"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line home.gox:110
							__e = __c.Set("class", "diff-label"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line home.gox:110
							__e = __c.Any(i18n.T(l, "home.diff")); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Init("div"); if __e != nil { return }
				{
//line home.gox:113
					__e = __c.Set("class", "cyc-grid"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:114
					for d := 1; d <= content.TotalDays; d++ {
//line home.gox:115
						__e = __c.Any(h.cycleCell(d, curDay)); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("script"); if __e != nil { return }
			{
//line home.gox:119
				__e = __c.Set("src", assets.WorkoutHome()); if __e != nil { return }
				__e = __c.Set("inline", true); if __e != nil { return }
//line home.gox:119
				__e = __c.Set("data:resume", resumeData); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Raw(""); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:121
}

// programs offers quick switching: popular ready programs until the user has
// another selected program, then their other selected programs. One tap switches.
//line home.gox:125
func (h homePage) programs(selected []app.SelectedView) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:127
		l := i18n.Lang(h.sess.Lang)
		var others []app.SelectedView
		for _, p := range selected {
			if !p.Current {
				others = append(others, p)
			}
		}
		// Until there is another selected program to switch to, offer the
		// popular ready programs (the current one highlighted).
		newcomer := len(others) == 0
		current := app.CurrentSource(h.sess.UserID)

		__e = __c.Init("section"); if __e != nil { return }
		{
//line home.gox:139
			__e = __c.Set("class", "courses"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:140
				__e = __c.Set("class", "courses-head"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h3"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:141
					if newcomer {
//line home.gox:141
						__e = __c.Any(i18n.T(l, "home.programs_ready")); if __e != nil { return }
					} else  {
//line home.gox:141
						__e = __c.Any(i18n.T(l, "home.switch")); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:142
				__e = (doors.ALink{Model: path.Path{Page: path.Programs}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line home.gox:142
						__e = __c.Set("class", "all-link"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:142
						__e = __c.Any(i18n.T(l, "home.all_programs")); if __e != nil { return }
						__e = __c.Text(" ›"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:144
				__e = __c.Set("class", "course-list"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:145
				if newcomer {
//line home.gox:146
					for _, key := range content.PopularKeys() {
//line home.gox:148
						p, _ := content.ReadyByKey(key)
						first, peak := content.MinutesSpan(p.Build())

//line home.gox:151
						__e = (doors.AClick{On: h.pickReady(key)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line home.gox:151
								__e = __c.Set("class", courseClass(key == current)); if __e != nil { return }
//line home.gox:151
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:152
									__e = __c.Set("class", "course-title"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:152
									__e = __c.Any(p.TitleIn(h.sess.Lang)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:153
									__e = __c.Set("class", "course-meta"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:153
									__e = __c.Many(first, "–", peak, " ", i18n.T(l, "home.min")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
					}
				} else  {
//line home.gox:157
					for _, p := range others {
//line home.gox:158
						__e = (doors.AClick{On: h.pickSelected(p.ID)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line home.gox:158
								__e = __c.Set("class", "course-card"); if __e != nil { return }
//line home.gox:158
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:159
									__e = __c.Set("class", "course-title"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:159
									__e = __c.Any(p.Name); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line home.gox:160
									__e = __c.Set("class", "course-meta"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line home.gox:160
									__e = __c.Many(i18n.T(l, "programs.day"), " ", p.Day); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
					}
				}
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:166
}

func courseClass(on bool) string {
	if on {
		return "course-card on"
	}
	return "course-card"
}

// pickReady switches to a ready program and reloads the home screen.
func (h homePage) pickReady(key string) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		if app.SelectReady(h.sess.UserID, key, h.sess.Lang) == nil {
			_ = r.After(doors.ActionLocationReload{})
		}
		return false
	}
}

// pickSelected switches to one of the user's selected programs.
func (h homePage) pickSelected(id int64) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		if app.SelectByID(h.sess.UserID, id) == nil {
			_ = r.After(doors.ActionLocationReload{})
		}
		return false
	}
}

// month renders the date calendar for the current month ± offset, reactively.
//line home.gox:196
func (h homePage) month(monthOffset doors.Source[int]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:197
		__e = __c.Any(monthOffset.Bind(func(off int) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line home.gox:199
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
//line home.gox:212
				__e = __c.Set("class", "cal-head"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:213
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				monthOffset.Mutate(ctx, func(o int) int { return o - 1 })
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line home.gox:216
						__e = __c.Set("class", "cal-nav"); if __e != nil { return }
//line home.gox:216
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("‹"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("span"); if __e != nil { return }
				{
//line home.gox:217
					__e = __c.Set("class", "cal-title"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line home.gox:217
					__e = __c.Many(i18n.Month(l, month), " ", year); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line home.gox:218
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				monthOffset.Mutate(ctx, func(o int) int { return o + 1 })
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line home.gox:221
						__e = __c.Set("class", "cal-nav"); if __e != nil { return }
//line home.gox:221
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
//line home.gox:223
				__e = __c.Set("class", "cal-week"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:224
				for wd := 0; wd < 7; wd++ {
					__e = __c.Init("span"); if __e != nil { return }
					{
						__e = __c.Submit(); if __e != nil { return }
//line home.gox:225
						__e = __c.Any(i18n.Weekday(l, wd)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line home.gox:228
				__e = __c.Set("class", "cal-grid month"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:229
				for b := 0; b < lead; b++ {
					__e = __c.Init("div"); if __e != nil { return }
					{
//line home.gox:230
						__e = __c.Set("class", "cal-blank"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
//line home.gox:232
				for dd := 1; dd <= daysIn; dd++ {
//line home.gox:233
					__e = __c.Any(h.dayCell(year, month, dd, acts, todayStr)); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })
//line home.gox:236
	})); if __e != nil { return }
	return })
//line home.gox:237
}

// dayCell renders one date cell, marked when the user worked out that day.
//line home.gox:240
func (h homePage) dayCell(year, month, dd int, acts map[string]bool, todayStr string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:242
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
//line home.gox:251
			__e = __c.Set("class", cls); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line home.gox:252
			if acts[ds] {
				__e = __c.Text("🔥"); if __e != nil { return }
			} else  {
//line home.gox:255
				__e = __c.Any(dd); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line home.gox:258
}

// cycleCell renders one of the 30 program days; sequential, looping each cycle.
//line home.gox:261
func (h homePage) cycleCell(d, current int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line home.gox:263
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

//line home.gox:278
		__e = (doors.ALink{Model: path.Path{Page: path.Day, Day: d}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("a"); if __e != nil { return }
			{
//line home.gox:278
				__e = __c.Set("class", cls); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line home.gox:279
				if done {
					__e = __c.Text("✓"); if __e != nil { return }
				} else  {
//line home.gox:282
					__e = __c.Any(d); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line home.gox:285
}
