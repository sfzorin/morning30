// Managed by GoX v0.1.36

//line today.gox:1
package pages

import (
	"context"
	"errors"
	"strconv"

	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// todayPage lists the exercises of the user's next day. Each exercise opens a
// menu: its card, replace it with a similar one, or remove it.
type todayPage struct {
	sess auth.Session
	path doors.Source[path.Path]
}

// todayState is the whole screen state in one comparable value.
type todayState struct {
	Open    string     // "part:index" of the item whose menu is open
	Card    string     // exercise id whose card is shown
	Replace string     // exercise id being replaced
	Remove  string     // "part:id" being removed (asks where)
	Where   whereState // days the replace/remove applies to
	Err     string     // i18n key of the last error
}

//line today.gox:35
func (t todayPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line today.gox:37
		l := i18n.Lang(t.sess.Lang)
		st := doors.NewSource(todayState{Where: whereState{Mode: "all"}})
		figure := doors.NewSource(app.DB.GetSex(t.sess.UserID))

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line today.gox:41
			__e = __c.Any(i18n.Tf(l, "fmt.day", app.CurrentDay(t.sess.UserID))); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line today.gox:42
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line today.gox:43
				__e = __c.Set("class", "library today"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:44
				__e = __c.Any(st.Bind(func(s todayState) gox.Elem {
				return gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line today.gox:45
					if s.Card != "" {
//line today.gox:46
						__e = __c.Any(exerciseCard{userID: t.sess.UserID, lang: t.sess.Lang, id: s.Card, figure: figure,
						onBack: func(ctx context.Context) { st.Update(ctx, todayState{Where: s.Where}) }}); if __e != nil { return }
//line today.gox:48
					} else if s.Replace != "" {
//line today.gox:49
						__e = __c.Any(t.replace(l, st, s)); if __e != nil { return }
					} else  {
//line today.gox:51
						__e = __c.Any(t.list(l, st, s)); if __e != nil { return }
					}
				return })
//line today.gox:53
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line today.gox:56
}

// list is the day's plan by part.
//line today.gox:59
func (t todayPage) list(l i18n.Lang, st doors.Source[todayState], s todayState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line today.gox:61
		plan := app.Day(t.sess.UserID, t.sess.Lang, 0)
		cat := app.UserCatalog(t.sess.UserID)
		parts := []struct{ Slot, Key string }{
			{"warmup", "workout.warmup"}, {"main", "workout.main"}, {"cooldown", "workout.cooldown"},
		}
		fig := app.DB.GetSex(t.sess.UserID)

		__e = __c.Init("div"); if __e != nil { return }
		{
//line today.gox:68
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line today.gox:69
			__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("a"); if __e != nil { return }
				{
//line today.gox:69
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line today.gox:69
					__e = __c.Set("aria-label", i18n.T(l, "nav.home")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line today.gox:69
					__e = __c.Any(icon("back", false)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:70
				__e = __c.Many(i18n.Tf(l, "fmt.day", plan.Day), " · ", plan.Program); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("p"); if __e != nil { return }
		{
//line today.gox:72
			__e = __c.Set("class", "lib-intro"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line today.gox:72
			__e = __c.Many(i18n.Tn(l, "fmt.exercises", plan.Exercises), " · ≈ ", i18n.Tf(l, "fmt.min", plan.EstMinutes)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line today.gox:73
		if s.Err != "" {
			__e = __c.Init("p"); if __e != nil { return }
			{
//line today.gox:74
				__e = __c.Set("class", "err"); if __e != nil { return }
//line today.gox:74
				__e = __c.Set("role", "alert"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:74
				__e = __c.Any(i18n.T(l, s.Err)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line today.gox:76
		for _, part := range parts {
			__e = __c.Init("h2"); if __e != nil { return }
			{
//line today.gox:77
				__e = __c.Set("class", "lib-grp"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:78
				__e = __c.Any(i18n.T(l, part.Key)); if __e != nil { return }
//line today.gox:79
				if part.Slot == "main" {
//line today.gox:79
					__e = __c.Many(" · ", i18n.Tn(l, "fmt.rounds", plan.MainRounds)); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line today.gox:81
				__e = __c.Set("class", "lib-list"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:82
				for i, it := range plan.Items {
//line today.gox:83
					if it.Part == part.Slot {
//line today.gox:85
						key := part.Slot + ":" + strconv.Itoa(i)
						open := s.Open == key

//line today.gox:88
						__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
						next := todayState{Where: s.Where}
						if !open {
							next.Open = key
						}
						st.Update(ctx, next)
						return false
					}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
							ctx := __c.Context(); _ = ctx
							__e = __c.Init("button"); if __e != nil { return }
							{
//line today.gox:95
								__e = __c.Set("class", itemClass(open)); if __e != nil { return }
//line today.gox:95
								__e = __c.Set("type", "button"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line today.gox:96
								if src := firstFrame(app.ExerciseFrames(cat, it.ID, fig)); src != "" {
									__e = __c.InitVoid("img"); if __e != nil { return }
									{
//line today.gox:97
										__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
//line today.gox:97
										__e = __c.Set("loading", "lazy"); if __e != nil { return }
//line today.gox:97
										__e = __c.Set("src", src); if __e != nil { return }
//line today.gox:97
										__e = __c.Set("alt", ""); if __e != nil { return }
									}
									__e = __c.Submit(); if __e != nil { return }
								} else  {
									__e = __c.Init("span"); if __e != nil { return }
									{
//line today.gox:99
										__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Init("span"); if __e != nil { return }
								{
//line today.gox:101
									__e = __c.Set("class", "lib-item-name"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line today.gox:101
									__e = __c.Any(it.Name); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
								__e = __c.Init("span"); if __e != nil { return }
								{
//line today.gox:102
									__e = __c.Set("class", "lib-item-dose"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line today.gox:102
									__e = __c.Any(doseLabel(l, it.Unit, it.Value, it.PerSide)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						return })); if __e != nil { return }
//line today.gox:104
						if open {
							__e = __c.Init("div"); if __e != nil { return }
							{
//line today.gox:105
								__e = __c.Set("class", "item-menu"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line today.gox:106
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								st.Update(ctx, todayState{Card: it.ID, Where: s.Where})
								toTop(r)
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line today.gox:110
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line today.gox:110
										__e = __c.Any(icon("info", false)); if __e != nil { return }
//line today.gox:110
										__e = __c.Any(i18n.T(l, "today.card")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
//line today.gox:111
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								st.Update(ctx, todayState{Replace: it.ID, Where: whereState{Mode: "all"}})
								toTop(r)
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line today.gox:115
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line today.gox:115
										__e = __c.Any(icon("replace", false)); if __e != nil { return }
//line today.gox:115
										__e = __c.Any(i18n.T(l, "today.replace")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
//line today.gox:116
								if s.Remove == part.Slot+":"+it.ID {
									__e = __c.Init("div"); if __e != nil { return }
									{
//line today.gox:117
										__e = __c.Set("class", "remove-ask"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line today.gox:118
											__e = __c.Any(i18n.T(l, "today.remove_where")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
//line today.gox:119
										__e = (doors.AClick{On: t.remove(st, s, it.ID, part.Slot, content.OneDay(plan.Day))}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line today.gox:119
												__e = __c.Set("class", "danger"); if __e != nil { return }
//line today.gox:119
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line today.gox:119
												__e = __c.Any(i18n.T(l, "days.today")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										return })); if __e != nil { return }
//line today.gox:120
										__e = (doors.AClick{On: t.remove(st, s, it.ID, part.Slot, content.AllDays())}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
											ctx := __c.Context(); _ = ctx
											__e = __c.Init("button"); if __e != nil { return }
											{
//line today.gox:120
												__e = __c.Set("class", "danger"); if __e != nil { return }
//line today.gox:120
												__e = __c.Set("type", "button"); if __e != nil { return }
												__e = __c.Submit(); if __e != nil { return }
//line today.gox:120
												__e = __c.Any(i18n.T(l, "days.all")); if __e != nil { return }
											}
											__e = __c.Close(); if __e != nil { return }
										return })); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								} else  {
//line today.gox:123
									__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
									st.Update(ctx, todayState{Open: key, Remove: part.Slot + ":" + it.ID, Where: s.Where})
									return false
								}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
										ctx := __c.Context(); _ = ctx
										__e = __c.Init("button"); if __e != nil { return }
										{
//line today.gox:126
											__e = __c.Set("class", "danger"); if __e != nil { return }
//line today.gox:126
											__e = __c.Set("type", "button"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line today.gox:126
											__e = __c.Any(icon("close", false)); if __e != nil { return }
//line today.gox:126
											__e = __c.Any(i18n.T(l, "today.remove")); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									return })); if __e != nil { return }
								}
							}
							__e = __c.Close(); if __e != nil { return }
						}
					}
				}
			}
			__e = __c.Close(); if __e != nil { return }
//line today.gox:133
			if part.Slot == "main" {
//line today.gox:134
				__e = (doors.ALink{Model: path.Path{Page: path.Library}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("a"); if __e != nil { return }
					{
//line today.gox:134
						__e = __c.Set("class", "add-link"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line today.gox:134
						__e = __c.Any(icon("plus", false)); if __e != nil { return }
//line today.gox:134
						__e = __c.Any(i18n.T(l, "today.add")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
		}
//line today.gox:137
		__e = (doors.ALink{Model: path.Path{Page: path.Start}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
			__e = __c.Init("a"); if __e != nil { return }
			{
//line today.gox:137
				__e = __c.Set("class", "btn primary big"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:137
				__e = __c.Any(i18n.T(l, "home.start")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		return })); if __e != nil { return }
	return })
//line today.gox:138
}

func itemClass(open bool) string {
	if open {
		return "lib-item open"
	}
	return "lib-item"
}

// remove takes an exercise out of today or of every day.
func (t todayPage) remove(st doors.Source[todayState], s todayState, id, part string, days content.Days) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		_, err := app.RemoveFromCurrent(t.sess.UserID, id, content.Slot(part), days)
		if err != nil {
			key := "today.edit_failed"
			if errors.Is(err, content.ErrEmptyPhase) {
				key = "today.part_empty"
			}
			st.Update(ctx, todayState{Err: key, Where: s.Where})
			return false
		}
		_ = r.After(doors.ActionLocationReload{})
		return false
	}
}

// replace offers easier, similar and harder exercises for one exercise.
//line today.gox:165
func (t todayPage) replace(l i18n.Lang, st doors.Source[todayState], s todayState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line today.gox:167
		cat := app.UserCatalog(t.sess.UserID)
		sim, _ := app.SimilarExercises(t.sess.UserID, t.sess.Lang, s.Replace)
		fig := app.DB.GetSex(t.sess.UserID)
		groups := []struct {
			Key  string
			List []app.ExerciseInfo
		}{{"replace.easier", sim.Easier}, {"replace.same", sim.Same}, {"replace.harder", sim.Harder}}
		today := app.CurrentDay(t.sess.UserID)

		__e = __c.Init("div"); if __e != nil { return }
		{
//line today.gox:176
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line today.gox:177
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, todayState{Where: s.Where})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line today.gox:180
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line today.gox:180
					__e = __c.Set("type", "button"); if __e != nil { return }
//line today.gox:180
					__e = __c.Set("aria-label", i18n.T(l, "nav.close")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line today.gox:180
					__e = __c.Any(icon("back", false)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:181
				__e = __c.Many(i18n.T(l, "today.replace"), ": ", cat.Name(l, s.Replace)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("h2"); if __e != nil { return }
		{
//line today.gox:183
			__e = __c.Set("class", "lib-grp"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line today.gox:183
			__e = __c.Any(i18n.T(l, "replace.where")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line today.gox:184
		__e = __c.Any(wherePicker(l, s.Where, func(ctx context.Context, w whereState) {
		st.Update(ctx, todayState{Replace: s.Replace, Where: w})
	})); if __e != nil { return }
//line today.gox:187
		if s.Err != "" {
			__e = __c.Init("p"); if __e != nil { return }
			{
//line today.gox:188
				__e = __c.Set("class", "err"); if __e != nil { return }
//line today.gox:188
				__e = __c.Set("role", "alert"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line today.gox:188
				__e = __c.Any(i18n.T(l, s.Err)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line today.gox:190
		for _, g := range groups {
//line today.gox:191
			if len(g.List) > 0 {
				__e = __c.Init("h2"); if __e != nil { return }
				{
//line today.gox:192
					__e = __c.Set("class", "lib-grp"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line today.gox:192
					__e = __c.Any(i18n.T(l, g.Key)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("div"); if __e != nil { return }
				{
//line today.gox:193
					__e = __c.Set("class", "lib-list"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line today.gox:194
					for _, e := range g.List {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line today.gox:195
							__e = __c.Set("class", "lib-item"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line today.gox:196
							if src := firstFrame(app.ExerciseFrames(cat, e.ID, fig)); src != "" {
								__e = __c.InitVoid("img"); if __e != nil { return }
								{
//line today.gox:197
									__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
//line today.gox:197
									__e = __c.Set("loading", "lazy"); if __e != nil { return }
//line today.gox:197
									__e = __c.Set("src", src); if __e != nil { return }
//line today.gox:197
									__e = __c.Set("alt", ""); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							} else  {
								__e = __c.Init("span"); if __e != nil { return }
								{
//line today.gox:199
									__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Init("span"); if __e != nil { return }
							{
//line today.gox:201
								__e = __c.Set("class", "lib-item-name"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line today.gox:201
								__e = __c.Any(e.Name); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line today.gox:202
							__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
							days := s.Where.days(today)
							if _, err := app.ReplaceInCurrent(t.sess.UserID, s.Replace, e.ID, days); err != nil {
								st.Update(ctx, todayState{Replace: s.Replace, Where: s.Where, Err: "today.edit_failed"})
								return false
							}
							_ = r.After(doors.ActionLocationReload{})
							return false
						}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line today.gox:210
									__e = __c.Set("class", "btn small"); if __e != nil { return }
//line today.gox:210
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line today.gox:210
									__e = __c.Any(i18n.T(l, "replace.pick")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
	return })
//line today.gox:216
}

var _ = gox.Elem(nil)
