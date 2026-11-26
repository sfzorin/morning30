// Managed by GoX v0.1.36

//line programs.gox:1
package pages

import (
	"context"
	"strconv"
	"strings"

	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/i18n"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// programsPage is the program catalog with the user's selected programs on
// top. One tap on a program switches to it; the star adds or removes it from
// the selected ones; the arrow opens a preview with the 30-day plan.
type programsPage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

type progState struct {
	Cat        string // category filter ("" = all)
	NoJumps    bool
	NoEquip    bool
	Preview    string // catalog key shown in preview
	Rename     bool   // renaming the current program
	ConfirmID  int64  // selected program waiting for "discard edits"
	ConfirmKey string // catalog program waiting for "discard edits"
	Tick       int    // bumped to redraw after a change made in place
}

//line programs.gox:38
func (p programsPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line programs.gox:40
		l := i18n.Lang(p.sess.Lang)
		cat := content.CatStart
		if src, ok := content.ReadyByKey(app.CurrentSource(p.sess.UserID)); ok {
			cat = src.Category
		}
		st := doors.NewSource(progState{Cat: cat})

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:47
			__e = __c.Any(i18n.T(l, "programs.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line programs.gox:48
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line programs.gox:49
				__e = __c.Set("class", "library programs"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:50
				__e = __c.Any(st.Bind(func(s progState) gox.Elem {
				return gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line programs.gox:51
					if s.Preview != "" {
//line programs.gox:52
						__e = __c.Any(p.preview(l, st, s)); if __e != nil { return }
					} else  {
//line programs.gox:54
						__e = __c.Any(p.catalog(l, st, s)); if __e != nil { return }
					}
				return })
//line programs.gox:56
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line programs.gox:59
}

//line programs.gox:61
func (p programsPage) catalog(l i18n.Lang, st doors.Source[progState], s progState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line programs.gox:63
		selected := app.SelectedPrograms(p.sess.UserID)
		starred := map[string]app.SelectedView{}
		for _, sp := range selected {
			if sp.Source != "" {
				starred[sp.Source] = sp
			}
		}
		used := map[string]bool{}
		for _, rp := range content.Catalog() {
			used[rp.Category] = true
		}

		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:75
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:76
			__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("a"); if __e != nil { return }
				{
//line programs.gox:76
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line programs.gox:76
					__e = __c.Set("aria-label", i18n.T(l, "nav.home")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line programs.gox:76
					__e = __c.Any(icon("back", false)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:77
				__e = __c.Any(i18n.T(l, "programs.title")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line programs.gox:80
		if len(selected) > 0 {
			__e = __c.Init("h2"); if __e != nil { return }
			{
//line programs.gox:81
				__e = __c.Set("class", "lib-grp"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:81
				__e = __c.Any(i18n.T(l, "programs.selected")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line programs.gox:82
				__e = __c.Set("class", "sel-list"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:83
				for _, sp := range selected {
//line programs.gox:84
					if s.ConfirmID == sp.ID {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line programs.gox:85
							__e = __c.Set("class", "sel-row confirm"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line programs.gox:86
								__e = __c.Any(i18n.Tf(l, "programs.remove_confirm", sp.Name)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line programs.gox:87
								__e = __c.Set("class", "confirm-actions"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line programs.gox:88
								__e = (doors.AClick{On: p.unselect(sp.ID, true)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line programs.gox:88
										__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line programs.gox:88
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:88
										__e = __c.Any(i18n.T(l, "programs.remove")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
//line programs.gox:89
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip})
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line programs.gox:92
										__e = __c.Set("class", "btn small"); if __e != nil { return }
//line programs.gox:92
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:92
										__e = __c.Any(i18n.T(l, "settings.cancel")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
//line programs.gox:95
					} else if sp.Current && s.Rename {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line programs.gox:96
							__e = __c.Set("class", "sel-row on"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.InitVoid("input"); if __e != nil { return }
							{
//line programs.gox:97
								__e = __c.Set("class", "text-input"); if __e != nil { return }
//line programs.gox:97
								__e = __c.Set("type", "text"); if __e != nil { return }
//line programs.gox:97
								__e = __c.Set("maxlength", app.ProgramNameMax); if __e != nil { return }
//line programs.gox:97
								__e = __c.Set("value", sp.Name); if __e != nil { return }
//line programs.gox:97
								__e = __c.Set("aria-label", i18n.T(l, "programs.rename")); if __e != nil { return }
								__e = __c.Set("autofocus", true); if __e != nil { return }
//line programs.gox:98
								__e = __c.Modify(doors.AChange{On: func(ctx context.Context, r doors.RequestEvent[doors.ChangeEvent]) bool {
								if app.RenameProgram(p.sess.UserID, r.Event().Value) == nil {
									_ = r.After(doors.ActionLocationReload{})
								}
								return false
							}}); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					} else  {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line programs.gox:106
							__e = __c.Set("class", selRowClass(sp.Current)); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line programs.gox:107
							__e = (doors.AClick{On: p.pickSelected(sp.ID)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line programs.gox:107
									__e = __c.Set("class", "sel-main"); if __e != nil { return }
//line programs.gox:107
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line programs.gox:108
										__e = __c.Set("class", "sel-name"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:108
										__e = __c.Any(sp.Name); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
//line programs.gox:110
							if sp.Current {
//line programs.gox:111
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip, Rename: true})
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line programs.gox:114
										__e = __c.Set("class", "icon-mini"); if __e != nil { return }
//line programs.gox:114
										__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:114
										__e = __c.Set("aria-label", i18n.T(l, "programs.rename")); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:114
										__e = __c.Any(icon("edit", false)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
							__e = __c.Init("span"); if __e != nil { return }
							{
//line programs.gox:116
								__e = __c.Set("class", "sel-day"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line programs.gox:116
								__e = __c.Any(i18n.Tf(l, "fmt.on_day", sp.Day)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
//line programs.gox:117
							__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
							if sp.Edited {
								st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip, ConfirmID: sp.ID})
								return false
							}
							if app.Unselect(p.sess.UserID, sp.ID, false) == nil {
								_ = r.After(doors.ActionLocationReload{})
							}
							return false
						}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line programs.gox:126
									__e = __c.Set("class", "star on"); if __e != nil { return }
//line programs.gox:126
									__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:126
									__e = __c.Set("aria-label", i18n.T(l, "programs.unstar")); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line programs.gox:126
									__e = __c.Any(icon("star", true)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:133
			__e = __c.Set("class", "chips"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:134
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, progState{NoJumps: s.NoJumps, NoEquip: s.NoEquip})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line programs.gox:137
					__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:137
					__e = __c.Set("class", chipClass(s.Cat == "")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line programs.gox:137
					__e = __c.Any(i18n.T(l, "programs.all")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
//line programs.gox:138
			for _, c := range content.Categories() {
//line programs.gox:139
				if used[c] {
//line programs.gox:140
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					st.Update(ctx, progState{Cat: c, NoJumps: s.NoJumps, NoEquip: s.NoEquip})
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line programs.gox:143
							__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:143
							__e = __c.Set("class", chipClass(s.Cat == c)); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line programs.gox:143
							__e = __c.Any(i18n.T(l, "cat."+c)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:147
			__e = __c.Set("class", "toggles"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:148
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, progState{Cat: s.Cat, NoJumps: !s.NoJumps, NoEquip: s.NoEquip})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line programs.gox:151
					__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:151
					__e = __c.Set("class", toggleClass(s.NoJumps)); if __e != nil { return }
//line programs.gox:151
					__e = __c.Set("aria-pressed", strconv.FormatBool(s.NoJumps)); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line programs.gox:151
						__e = __c.Set("class", "sw"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line programs.gox:151
					__e = __c.Any(i18n.T(l, "filter.nojumps")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
//line programs.gox:152
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: !s.NoEquip})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line programs.gox:155
					__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:155
					__e = __c.Set("class", toggleClass(s.NoEquip)); if __e != nil { return }
//line programs.gox:155
					__e = __c.Set("aria-pressed", strconv.FormatBool(s.NoEquip)); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("span"); if __e != nil { return }
					{
//line programs.gox:155
						__e = __c.Set("class", "sw"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
//line programs.gox:155
					__e = __c.Any(i18n.T(l, "filter.noequipment")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:158
			__e = __c.Set("class", "prog-list"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:159
			for _, rp := range content.Catalog() {
//line programs.gox:161
				r := rp.Build()
				tags := content.ProgramTags(r)
				first, peak := content.MinutesSpan(r)
				show := (s.Cat == "" || rp.Category == s.Cat) && !(s.NoJumps && tags.Jumps) && !(s.NoEquip && tags.Equipment)
				mine, isStarred := starred[rp.Key]

//line programs.gox:167
				if show {
//line programs.gox:168
					if s.ConfirmKey == rp.Key {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line programs.gox:169
							__e = __c.Set("class", "sel-row confirm"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
							__e = __c.Init("span"); if __e != nil { return }
							{
								__e = __c.Submit(); if __e != nil { return }
//line programs.gox:170
								__e = __c.Any(i18n.Tf(l, "programs.remove_confirm", mine.Name)); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
							__e = __c.Init("div"); if __e != nil { return }
							{
//line programs.gox:171
								__e = __c.Set("class", "confirm-actions"); if __e != nil { return }
								__e = __c.Submit(); if __e != nil { return }
//line programs.gox:172
								__e = (doors.AClick{On: p.unselect(mine.ID, true)}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line programs.gox:172
										__e = __c.Set("class", "btn danger small"); if __e != nil { return }
//line programs.gox:172
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:172
										__e = __c.Any(i18n.T(l, "programs.remove")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
//line programs.gox:173
								__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
								st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip})
								return false
							}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
									ctx := __c.Context(); _ = ctx
									__e = __c.Init("button"); if __e != nil { return }
									{
//line programs.gox:176
										__e = __c.Set("class", "btn small"); if __e != nil { return }
//line programs.gox:176
										__e = __c.Set("type", "button"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:176
										__e = __c.Any(i18n.T(l, "settings.cancel")); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
								return })); if __e != nil { return }
							}
							__e = __c.Close(); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					} else  {
						__e = __c.Init("div"); if __e != nil { return }
						{
//line programs.gox:180
							__e = __c.Set("class", "prog-row"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line programs.gox:181
							__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
							st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip, Preview: rp.Key})
							toTop(r)
							return false
						}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line programs.gox:185
									__e = __c.Set("class", "prog-main"); if __e != nil { return }
//line programs.gox:185
									__e = __c.Set("type", "button"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line programs.gox:186
										__e = __c.Set("class", "prog-top"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
										__e = __c.Init("b"); if __e != nil { return }
										{
											__e = __c.Submit(); if __e != nil { return }
//line programs.gox:186
											__e = __c.Any(rp.TitleIn(p.sess.Lang)); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
										__e = __c.Init("span"); if __e != nil { return }
										{
//line programs.gox:186
											__e = __c.Set("class", "m"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line programs.gox:186
											__e = __c.Any(i18n.Tf(l, "fmt.min_range", first, peak)); if __e != nil { return }
										}
										__e = __c.Close(); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
									__e = __c.Init("span"); if __e != nil { return }
									{
//line programs.gox:187
										__e = __c.Set("class", "prog-about"); if __e != nil { return }
										__e = __c.Submit(); if __e != nil { return }
//line programs.gox:187
										__e = __c.Any(rp.AboutIn(p.sess.Lang)); if __e != nil { return }
									}
									__e = __c.Close(); if __e != nil { return }
//line programs.gox:188
									if tags.Jumps || tags.Equipment {
										__e = __c.Init("span"); if __e != nil { return }
										{
//line programs.gox:189
											__e = __c.Set("class", "prog-tags"); if __e != nil { return }
											__e = __c.Submit(); if __e != nil { return }
//line programs.gox:190
											if tags.Jumps {
												__e = __c.Init("span"); if __e != nil { return }
												{
//line programs.gox:190
													__e = __c.Set("class", "tag"); if __e != nil { return }
													__e = __c.Submit(); if __e != nil { return }
//line programs.gox:190
													__e = __c.Any(i18n.T(l, "tag.jumps")); if __e != nil { return }
												}
												__e = __c.Close(); if __e != nil { return }
											}
//line programs.gox:191
											if tags.Equipment {
												__e = __c.Init("span"); if __e != nil { return }
												{
//line programs.gox:191
													__e = __c.Set("class", "tag"); if __e != nil { return }
													__e = __c.Submit(); if __e != nil { return }
//line programs.gox:191
													__e = __c.Any(i18n.T(l, "tag.equipment")); if __e != nil { return }
												}
												__e = __c.Close(); if __e != nil { return }
											}
										}
										__e = __c.Close(); if __e != nil { return }
									}
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
//line programs.gox:195
							__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
							if isStarred {
								if mine.Edited {
									st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip, ConfirmKey: rp.Key})
									return false
								}
								if app.Unselect(p.sess.UserID, mine.ID, false) == nil {
									_ = r.After(doors.ActionLocationReload{})
								}
								return false
							}
							if _, err := app.StarReady(p.sess.UserID, rp.Key, p.sess.Lang); err == nil {
								_ = r.After(doors.ActionLocationReload{})
							}
							return false
						}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
								ctx := __c.Context(); _ = ctx
								__e = __c.Init("button"); if __e != nil { return }
								{
//line programs.gox:210
									__e = __c.Set("class", starClass(isStarred)); if __e != nil { return }
//line programs.gox:210
									__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:210
									__e = __c.Set("aria-label", starLabel(l, isStarred)); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
//line programs.gox:210
									__e = __c.Any(icon("star", isStarred)); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							return })); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line programs.gox:216
}

// preview shows one ready program before choosing it.
//line programs.gox:219
func (p programsPage) preview(l i18n.Lang, st doors.Source[progState], s progState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line programs.gox:221
		rp, _ := content.ReadyByKey(s.Preview)
		r := rp.Build()
		tags := content.ProgramTags(r)
		first, peak := content.MinutesSpan(r)
		cat := app.UserCatalog(p.sess.UserID)
		w, m, c := app.DB.Rests(p.sess.UserID)
		rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
		isStarred := false
		for _, sp := range app.SelectedPrograms(p.sess.UserID) {
			isStarred = isStarred || sp.Source == rp.Key
		}

		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:233
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:234
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line programs.gox:237
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line programs.gox:237
					__e = __c.Set("type", "button"); if __e != nil { return }
//line programs.gox:237
					__e = __c.Set("aria-label", i18n.T(l, "programs.title")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line programs.gox:237
					__e = __c.Any(icon("back", false)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:238
				__e = __c.Any(rp.TitleIn(p.sess.Lang)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("p"); if __e != nil { return }
		{
//line programs.gox:240
			__e = __c.Set("class", "lib-desc"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:240
			__e = __c.Any(rp.AboutIn(p.sess.Lang)); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:241
			__e = __c.Set("class", "pills"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line programs.gox:242
				__e = __c.Set("class", "pill"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:242
				__e = __c.Many(i18n.Tf(l, "fmt.min_range", first, peak), " ", i18n.T(l, "landing.a_day")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line programs.gox:243
				__e = __c.Set("class", "pill"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:243
				if tags.Jumps {
//line programs.gox:243
					__e = __c.Any(i18n.T(l, "tag.jumps")); if __e != nil { return }
				} else  {
//line programs.gox:243
					__e = __c.Any(i18n.T(l, "tag.nojumps")); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line programs.gox:244
				__e = __c.Set("class", "pill"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line programs.gox:244
				if tags.Equipment {
//line programs.gox:244
					__e = __c.Any(i18n.T(l, "tag.equipment")); if __e != nil { return }
				} else  {
//line programs.gox:244
					__e = __c.Any(i18n.T(l, "tag.noequipment")); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line programs.gox:246
			__e = __c.Set("class", "preview-actions"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:247
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			if app.SelectReady(p.sess.UserID, rp.Key, p.sess.Lang) == nil {
				p.path.Update(ctx, path.Path{Page: path.Home})
			}
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line programs.gox:252
					__e = __c.Set("class", "btn primary"); if __e != nil { return }
//line programs.gox:252
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line programs.gox:252
					__e = __c.Any(i18n.T(l, "programs.start")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
//line programs.gox:253
			if !isStarred {
//line programs.gox:254
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				if _, err := app.StarReady(p.sess.UserID, rp.Key, p.sess.Lang); err == nil {
					st.Update(ctx, progState{Cat: s.Cat, NoJumps: s.NoJumps, NoEquip: s.NoEquip, Preview: rp.Key, Tick: s.Tick + 1})
				}
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line programs.gox:259
						__e = __c.Set("class", "btn"); if __e != nil { return }
//line programs.gox:259
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line programs.gox:259
						__e = __c.Any(icon("star", false)); if __e != nil { return }
//line programs.gox:259
						__e = __c.Any(i18n.T(l, "programs.star")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			} else  {
				__e = __c.Init("span"); if __e != nil { return }
				{
//line programs.gox:261
					__e = __c.Set("class", "pill on"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line programs.gox:261
					__e = __c.Any(icon("star", true)); if __e != nil { return }
//line programs.gox:261
					__e = __c.Any(i18n.T(l, "programs.in_selected")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("h2"); if __e != nil { return }
		{
//line programs.gox:264
			__e = __c.Set("class", "lib-grp"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:264
			__e = __c.Any(i18n.T(l, "landing.plan")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("ol"); if __e != nil { return }
		{
//line programs.gox:265
			__e = __c.Set("class", "plan-days"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line programs.gox:266
			for i, d := range r.Days {
//line programs.gox:268
				var items []string
				for _, it := range d.Items {
					items = append(items, cat.Name(l, it.ID)+" "+doseLabel(l, string(it.Unit), it.Value, false))
				}
				est := r.Workout(i+1, 0, rests).EstSec

				__e = __c.Init("li"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("div"); if __e != nil { return }
					{
//line programs.gox:275
						__e = __c.Set("class", "plan-head"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("b"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line programs.gox:276
							__e = __c.Any(i18n.Tf(l, "fmt.day", d.Day)); if __e != nil { return }
//line programs.gox:276
							if content.IsRecoveryDay(d.Day) {
								__e = __c.Init("span"); if __e != nil { return }
								{
//line programs.gox:276
									__e = __c.Set("class", "light"); if __e != nil { return }
									__e = __c.Submit(); if __e != nil { return }
									__e = __c.Text(" · "); if __e != nil { return }
//line programs.gox:276
									__e = __c.Any(i18n.T(l, "landing.light")); if __e != nil { return }
								}
								__e = __c.Close(); if __e != nil { return }
							}
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line programs.gox:277
							__e = __c.Set("class", "m"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line programs.gox:277
							__e = __c.Many("≈ ", i18n.Tf(l, "fmt.min", (est+59)/60)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("p"); if __e != nil { return }
					{
//line programs.gox:279
						__e = __c.Set("class", "m"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line programs.gox:279
						__e = __c.Any(strings.Join(items, ", ")); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line programs.gox:283
}

func selRowClass(current bool) string {
	if current {
		return "sel-row on"
	}
	return "sel-row"
}

func chipClass(on bool) string {
	if on {
		return "chip on"
	}
	return "chip"
}

func toggleClass(on bool) string {
	if on {
		return "toggle on"
	}
	return "toggle"
}

func starClass(on bool) string {
	if on {
		return "star on"
	}
	return "star"
}

func starLabel(l i18n.Lang, on bool) string {
	if on {
		return i18n.T(l, "programs.unstar")
	}
	return i18n.T(l, "programs.star")
}

func (p programsPage) pickSelected(id int64) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		if app.SelectByID(p.sess.UserID, id) == nil {
			_ = r.After(doors.ActionLocationReload{})
		}
		return false
	}
}

func (p programsPage) unselect(id int64, discard bool) func(context.Context, doors.RequestPointer) bool {
	return func(ctx context.Context, r doors.RequestPointer) bool {
		if app.Unselect(p.sess.UserID, id, discard) == nil {
			_ = r.After(doors.ActionLocationReload{})
		}
		return false
	}
}

var _ = gox.Elem(nil)
