// Managed by GoX v0.1.36

//line library.gox:1
package pages

import (
	"context"
	"strings"

	"morning30/internal/auth"
	"morning30/internal/i18n"
	"morning30/internal/library"
	"morning30/path"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// libState is the library browser's position: a group key, then an exercise id.
// "" group = the group list; group set, id "" = that group's exercises; id set =
// the exercise detail (card). One Source drives all three levels.
type libState struct {
	Group string
	ID    string
}

// libraryPage is the browse-only exercise library imported from Leap Fitness:
// grouped lists down to a localized card.
type libraryPage struct {
	sess auth.Session
	auth doors.Source[auth.Session]
	path doors.Source[path.Path]
}

//line library.gox:32
func (p libraryPage) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:34
		l := i18n.Lang(p.sess.Lang)
		lang := string(l)
		st := doors.NewSource(libState{})

		__e = __c.Init("title"); if __e != nil { return }
		{
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:38
			__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("main"); if __e != nil { return }
		{
//line library.gox:39
			__e = __c.Set("class", "screen"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("section"); if __e != nil { return }
			{
//line library.gox:40
				__e = __c.Set("class", "library"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:41
				__e = __c.Any(st.Bind(func(s libState) gox.Elem {
				return gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
//line library.gox:42
					if s.ID != "" {
//line library.gox:43
						__e = __c.Any(p.detail(l, lang, st, s)); if __e != nil { return }
					} else  {
//line library.gox:45
						if s.Group != "" {
//line library.gox:46
							__e = __c.Any(p.groupView(l, lang, st, s.Group)); if __e != nil { return }
						} else  {
//line library.gox:48
							__e = __c.Any(p.groupList(l, st)); if __e != nil { return }
						}
					}
				return })
//line library.gox:51
			})); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:54
}

// groupList is level 1: every browse category with its exercise count.
//line library.gox:57
func (p libraryPage) groupList(l i18n.Lang, st doors.Source[libState]) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:58
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:59
				__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
//line library.gox:60
			__e = (doors.ALink{Model: path.Path{Page: path.Home}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("a"); if __e != nil { return }
				{
//line library.gox:60
					__e = __c.Set("class", "lib-x"); if __e != nil { return }
//line library.gox:60
					__e = __c.Set("aria-label", i18n.T(l, "workout.quit")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("✕"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("p"); if __e != nil { return }
		{
//line library.gox:62
			__e = __c.Set("class", "lib-intro"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:62
			__e = __c.Any(i18n.T(l, "library.intro")); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:63
			__e = __c.Set("class", "lib-groups"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:64
			for _, g := range library.Groups() {
//line library.gox:65
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, libState{Group: g.Key})
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:69
						__e = __c.Set("class", "lib-grow"); if __e != nil { return }
//line library.gox:69
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line library.gox:70
							__e = __c.Set("class", "lib-grow-name"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:70
							__e = __c.Any(g.Label); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
						__e = __c.Init("span"); if __e != nil { return }
						{
//line library.gox:71
							__e = __c.Set("class", "lib-grow-count"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:71
							__e = __c.Any(len(g.Cards)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:75
}

// groupView is level 2: the exercises in one group.
//line library.gox:78
func (p libraryPage) groupView(l i18n.Lang, lang string, st doors.Source[libState], group string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:80
		var cards []*library.Card
		label := group
		for _, g := range library.Groups() {
			if g.Key == group {
				cards = g.Cards
				label = g.Label
				break
			}
		}

		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:90
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:91
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			st.Update(ctx, libState{})
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line library.gox:94
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line library.gox:94
					__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:94
					__e = __c.Set("aria-label", "←"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Text("←"); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:95
				__e = __c.Any(label); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("div"); if __e != nil { return }
		{
//line library.gox:97
			__e = __c.Set("class", "lib-list"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line library.gox:98
			for _, c := range cards {
//line library.gox:99
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, libState{Group: group, ID: c.ID})
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:103
						__e = __c.Set("class", "lib-item"); if __e != nil { return }
//line library.gox:103
						__e = __c.Set("type", "button"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line library.gox:104
						if thumb := localThumb(c.Source.Thumbnail); thumb != "" {
							__e = __c.InitVoid("img"); if __e != nil { return }
							{
//line library.gox:105
								__e = __c.Set("class", "lib-thumb"); if __e != nil { return }
//line library.gox:105
								__e = __c.Set("loading", "lazy"); if __e != nil { return }
//line library.gox:105
								__e = __c.Set("src", thumb); if __e != nil { return }
//line library.gox:105
								__e = __c.Set("alt", ""); if __e != nil { return }
							}
							__e = __c.Submit(); if __e != nil { return }
						}
						__e = __c.Init("span"); if __e != nil { return }
						{
//line library.gox:107
							__e = __c.Set("class", "lib-item-name"); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:107
							__e = __c.Any(c.Name(lang)); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line library.gox:111
}

// detail is level 3: the localized card.
//line library.gox:114
func (p libraryPage) detail(l i18n.Lang, lang string, st doors.Source[libState], s libState) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line library.gox:116
		c, ok := library.ByID(s.ID)

//line library.gox:118
		if !ok {
			__e = __c.Init("p"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:119
				__e = __c.Any(i18n.T(l, "library.title")); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		} else  {
//line library.gox:122
			t := c.Text(lang)

			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:124
				__e = __c.Set("class", "lib-head"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:125
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				st.Update(ctx, libState{Group: s.Group})
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line library.gox:128
						__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line library.gox:128
						__e = __c.Set("type", "button"); if __e != nil { return }
//line library.gox:128
						__e = __c.Set("aria-label", "←"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
						__e = __c.Text("←"); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
				__e = __c.Init("h1"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:129
					__e = __c.Any(c.Name(lang)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("p"); if __e != nil { return }
			{
//line library.gox:131
				__e = __c.Set("class", "lib-desc"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line library.gox:131
				__e = __c.Any(t.ShortDescription); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:133
				__e = __c.Set("class", "lib-sec"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:134
					__e = __c.Any(i18n.T(l, "library.howto")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ul"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:135
					for _, line := range t.HowTo {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:135
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:137
				__e = __c.Set("class", "lib-sec lib-correct"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:138
					__e = __c.Any(i18n.T(l, "library.correct")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ul"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:139
					for _, line := range t.Correct {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:139
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("div"); if __e != nil { return }
			{
//line library.gox:141
				__e = __c.Set("class", "lib-sec lib-mistakes"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:142
					__e = __c.Any(i18n.T(l, "library.mistakes")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ul"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line library.gox:143
					for _, line := range t.Mistakes {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line library.gox:143
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line library.gox:146
}

func localThumb(url string) string {
	u := strings.ToLower(url)
	if u == "" || strings.Contains(u, "youtu") || strings.Contains(u, "ytimg") {
		return ""
	}
	return url
}
