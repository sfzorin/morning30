// Managed by GoX v0.1.36

//line library.gox:1
package pages

import (
	"context"
	"errors"

	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// libraryPage is the exercise library: find an exercise, look at its card,
// and add it to the current program ("+": part, days and amount).
type libraryPage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

// libState is the whole library screen in one comparable value.
type libState struct {
	Query   string
	Group   string
	NoJumps bool
	Card    string // exercise whose card is open
	Add     string // exercise being added
	Part    string
	Where   whereState
	Value   int
	Err     string // i18n key of the last error
	Done    string // name of the exercise just added
}

// list returns to the list, keeping the search and filters.
func (s libState) list() libState {
	return libState{Query: s.Query, Group: s.Group, NoJumps: s.NoJumps}
}

//line library.gox:44
func (p libraryPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:46
		l := i18n.Lang(p.sess.Lang)
		st := doors.NewSource(libState{})
		figure := doors.NewSource(app.DB.GetSex(p.sess.UserID))

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:50
			__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line library.gox:51
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line library.gox:52
				__e = __c.Set("class", "library"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:53
				__e = __c.Any(st.Bind(func(s libState) gox.Elem {
				return gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line library.gox:54
					if s.Add != "" {
//line library.gox:55
						__e = __c.Any(p.addSheet(l, st, s)); if __e != nil { return }
//line library.gox:56
					} else if s.Card != "" {
//line library.gox:57
						__e = __c.Any(exerciseCard{userID: p.sess.UserID, lang: p.sess.Lang, id: s.Card, figure: figure,
						onBack: func(ctx context.Context) { st.Update(ctx, s.list()) },
						onAdd:  func(ctx context.Context) { st.Update(ctx, p.startAdd(s, s.Card)) }}); if __e != nil { return }
					} else  {
//line library.gox:61
						__e = __c.Any(p.list(l, st, s)); if __e != nil { return }
					}
				return })
//line library.gox:63
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:66
}

// startAdd opens the add sheet with the exercise's own part and a default amount.
func (p libraryPage) startAdd(s libState, id string) libState {
	next := s.list()
	next.Add = id
	next.Part = string(content.Main)
	next.Where = whereState{Mode: "all"}
	next.Value = 10
	if def, ok := app.UserCatalog(p.sess.UserID).Def(id); ok {
		if def.Slot != "" {
			next.Part = string(def.Slot)
		}
		next.Value = content.DefaultValue(def.Unit)
	}
	return next
}

//line library.gox:84
func (p libraryPage) list(l i18n.Lang, st doors.Source[libState], s libState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:86
		items := app.Library(p.sess.UserID, p.sess.Lang, s.Group, s.Query, s.NoJumps)
		all := app.Library(p.sess.UserID, p.sess.Lang, "", "", false)
		used := map[string]bool{}
		for _, e := range all {
			used[e.Group] = true
		}
		fig := app.DB.GetSex(p.sess.UserID)
		group := ""

		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:95
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:96
			__e = (doors.ALink{Model: path.Path{Page: path.Editor}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("a"); if __e != nil { return }
				{
//line library.gox:96
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line library.gox:96
					__e = __c.Set("aria-label", i18n.T(l, "settings.editor")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("←"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:97
				__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line library.gox:99
		if s.Done != "" {
			__e = __c.Init("p"); if __e != nil { return }
			{
//line library.gox:100
				__e = __c.Set("class", "ok-msg"); if __e != nil { return }
//line library.gox:100
				__e = __c.Set("role", "status"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:100
				__e = __c.Any(i18n.Tf(l, "library.added", s.Done, app.ProgramTitle(p.sess.UserID, p.sess.Lang))); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.InitVoid("input"); if __e != nil { return }
		{
//line library.gox:102
			__e = __c.Set("class", "text-input search"); if __e != nil { return }
//line library.gox:102
			__e = __c.Set("type", "search"); if __e != nil { return }
//line library.gox:102
			__e = __c.Set("value", s.Query); if __e != nil { return }
//line library.gox:102
			__e = __c.Set("placeholder", i18n.T(l, "library.search")); if __e != nil { return }
//line library.gox:102
			__e = __c.Set("aria-label", i18n.T(l, "library.search")); if __e != nil { return }
//line library.gox:103
			__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
			st.Update(ctx, libState{Query: r.Event().Value, Group: s.Group, NoJumps: s.NoJumps})
			return false
		}}); if __e != nil { return }
		}
		__e = __c.Submit(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:107
			__e = __c.Set("class", "chips"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:108
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, libState{Query: s.Query, NoJumps: s.NoJumps})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line library.gox:111
					__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:111
					__e = __c.Set("class", chipClass(s.Group == "")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:111
					__e = __c.Any(i18n.T(l, "programs.all")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
//line library.gox:112
			for _, g := range app.LibraryGroups() {
//line library.gox:113
				if used[g] {
//line library.gox:114
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					st.Update(ctx, libState{Query: s.Query, Group: g, NoJumps: s.NoJumps})
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line library.gox:117
							__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:117
							__e = __c.Set("class", chipClass(s.Group == g)); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:117
							__e = __c.Any(i18n.T(l, "grp."+g)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:121
			__e = __c.Set("class", "toggles"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:122
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, libState{Query: s.Query, Group: s.Group, NoJumps: !s.NoJumps})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line library.gox:125
					__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:125
					__e = __c.Set("class", toggleClass(s.NoJumps)); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line library.gox:125
						__e = __c.Set("class", "sw"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line library.gox:125
					__e = __c.Any(i18n.T(l, "filter.nojumps")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line library.gox:127
		if len(items) == 0 {
			__e = __c.Init("p"); if __e != nil { return }
			{
//line library.gox:128
				__e = __c.Set("class", "lib-intro"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:128
				__e = __c.Any(i18n.T(l, "library.nothing")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:130
			__e = __c.Set("class", "lib-list"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:131
			for _, e := range items {
//line library.gox:132
				if s.Group == "" && e.Group != group {
//line library.gox:133
					group = e.Group
					__e = __c.Init("h2"); if __e != nil { return }
					{
//line library.gox:134
						__e = __c.Set("class", "lib-grp"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line library.gox:134
						__e = __c.Any(i18n.T(l, "grp."+e.Group)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Init("div"); if __e != nil { return }
				{
//line library.gox:136
					__e = __c.Set("class", "lib-item"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:137
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					next := s.list()
					next.Card = e.ID
					st.Update(ctx, next)
					toTop(r)
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line library.gox:143
							__e = __c.Set("class", "lib-open"); if __e != nil { return }
//line library.gox:143
							__e = __c.Set("type", "button"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:144
							if src := firstFrameOf(p.sess.UserID, e.ID, fig); src != "" {
								__e = __c.InitVoid("img"); if __e != nil { return }
								{
//line library.gox:145
									__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
//line library.gox:145
									__e = __c.Set("loading", "lazy"); if __e != nil { return }
//line library.gox:145
									__e = __c.Set("src", src); if __e != nil { return }
//line library.gox:145
									__e = __c.Set("alt", ""); if __e != nil { return }
								}
								__e = __c.Submit(); if __e != nil { return }
							} else  {
								__e = __c.Init("span"); if __e != nil { return }
								{
//line library.gox:147
									__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
							__e = __c.Init("span"); if __e != nil { return }
							{
//line library.gox:149
								__e = __c.Set("class", "lib-item-name"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line library.gox:149
								__e = __c.Any(e.Name); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
//line library.gox:151
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					st.Update(ctx, p.startAdd(s, e.ID))
					toTop(r)
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line library.gox:155
							__e = __c.Set("class", "add-btn"); if __e != nil { return }
//line library.gox:155
							__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:155
							__e = __c.Set("aria-label", i18n.T(l, "library.add_to_program")); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Text("+"); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:159
}

// addSheet asks where, on which days and how much, then adds the exercise.
//line library.gox:162
func (p libraryPage) addSheet(l i18n.Lang, st doors.Source[libState], s libState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:164
		cat := app.UserCatalog(p.sess.UserID)
		def, _ := cat.Def(s.Add)
		step := 1
		if def.Unit == content.Seconds {
			step = 5
		}
		with := func(change func(*libState)) libState {
			next := s
			next.Err = ""
			change(&next)
			return next
		}
		today := app.CurrentDay(p.sess.UserID)

		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:178
			__e = __c.Set("class", "sheet"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:179
				__e = __c.Any(cat.Name(l, s.Add)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("p"); if __e != nil { return }
			{
//line library.gox:180
				__e = __c.Set("class", "lib-intro"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:180
				__e = __c.Any(i18n.Tf(l, "library.into", app.ProgramTitle(p.sess.UserID, p.sess.Lang))); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("h2"); if __e != nil { return }
			{
//line library.gox:181
				__e = __c.Set("class", "lib-grp"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:181
				__e = __c.Any(i18n.T(l, "library.part")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:182
				__e = __c.Set("class", "seg3"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:183
				for _, part := range []struct{ Slot, Key string }{{"warmup", "workout.warmup"}, {"main", "workout.main"}, {"cooldown", "workout.cooldown"}} {
//line library.gox:184
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					st.Update(ctx, with(func(n *libState) { n.Part = part.Slot }))
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line library.gox:187
							__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:187
							__e = __c.Set("class", segClass(s.Part == part.Slot)); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:187
							__e = __c.Any(i18n.T(l, part.Key)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("h2"); if __e != nil { return }
			{
//line library.gox:190
				__e = __c.Set("class", "lib-grp"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:190
				__e = __c.Any(i18n.T(l, "library.days")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line library.gox:191
			__e = __c.Any(wherePicker(l, s.Where, func(ctx context.Context, w whereState) {
			st.Update(ctx, with(func(n *libState) { n.Where = w }))
		})); if __e != nil { return }
			__e = __c.Init("h2"); if __e != nil { return }
			{
//line library.gox:194
				__e = __c.Set("class", "lib-grp"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:194
				__e = __c.Any(i18n.T(l, "library.amount")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:195
				__e = __c.Set("class", "stepper"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:196
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, with(func(n *libState) {
					if n.Value-step >= 1 {
						n.Value -= step
					}
				}))
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:203
						__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:203
						__e = __c.Set("aria-label", "−"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("−"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("b"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:204
					__e = __c.Any(doseLabel(l, string(def.Unit), s.Value, def.PerSide)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line library.gox:205
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, with(func(n *libState) { n.Value += step }))
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:208
						__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:208
						__e = __c.Set("aria-label", "+"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("+"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line library.gox:210
			if s.Err != "" {
				__e = __c.Init("p"); if __e != nil { return }
				{
//line library.gox:211
					__e = __c.Set("class", "err"); if __e != nil { return }
//line library.gox:211
					__e = __c.Set("role", "alert"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:211
					__e = __c.Any(i18n.T(l, s.Err)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:213
				__e = __c.Set("class", "preview-actions"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:214
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				err := app.AddToCurrent(p.sess.UserID, s.Add, content.Slot(s.Part), s.Where.days(today), s.Value)
				if err != nil {
					key := "today.edit_failed"
					if errors.Is(err, content.ErrNoDays) {
						key = "library.no_days"
					}
					st.Update(ctx, with(func(n *libState) { n.Err = key }))
					return false
				}
				done := s.list()
				done.Done = cat.Name(l, s.Add)
				st.Update(ctx, done)
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:228
						__e = __c.Set("class", "btn primary"); if __e != nil { return }
//line library.gox:228
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line library.gox:228
						__e = __c.Any(i18n.T(l, "library.add")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
//line library.gox:229
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, s.list())
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:232
						__e = __c.Set("class", "btn"); if __e != nil { return }
//line library.gox:232
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line library.gox:232
						__e = __c.Any(i18n.T(l, "settings.cancel")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:235
}

var _ = gox.Elem(nil)
