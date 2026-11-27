// Managed by GoX v0.1.36

//line exercise_ui.gox:1
package pages

import (
	"context"
	"strconv"

	"morning30/internal/app"
	"morning30/internal/content"
	"morning30/internal/i18n"

	"github.com/doors-dev/doors"
	"github.com/doors-dev/gox"
)

// Shared pieces of the day screen and the library: an exercise card with the
// four figures, a day selection and dose labels.

var figureLabels = map[string]string{
	"m": "sex.male", "f": "sex.female", "b": "sex.boy", "g": "sex.girl",
}

//line exercise_ui.gox:22
func figureTile(l i18n.Lang, figure string) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("span"); if __e != nil { return }
		{
//line exercise_ui.gox:23
			__e = __c.Set("class", "figure-art"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.InitVoid("img"); if __e != nil { return }
			{
//line exercise_ui.gox:24
				__e = __c.Set("src", "/static/characters/" + figure + ".webp"); if __e != nil { return }
//line exercise_ui.gox:24
				__e = __c.Set("alt", ""); if __e != nil { return }
//line exercise_ui.gox:24
				__e = __c.Set("width", "288"); if __e != nil { return }
//line exercise_ui.gox:24
				__e = __c.Set("height", "288"); if __e != nil { return }
			}
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("span"); if __e != nil { return }
			{
//line exercise_ui.gox:25
				__e = __c.Set("class", "figure-check"); if __e != nil { return }
//line exercise_ui.gox:25
				__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Text("✓"); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
		__e = __c.Init("span"); if __e != nil { return }
		{
//line exercise_ui.gox:27
			__e = __c.Set("class", "figure-name"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:27
			__e = __c.Any(i18n.T(l, figureLabels[figure])); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line exercise_ui.gox:28
}

//line exercise_ui.gox:30
func figureCards(l i18n.Lang, selected string, choose func(context.Context, string)) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line exercise_ui.gox:31
			__e = __c.Set("class", "figure-grid"); if __e != nil { return }
//line exercise_ui.gox:31
			__e = __c.Set("role", "group"); if __e != nil { return }
//line exercise_ui.gox:31
			__e = __c.Set("aria-label", i18n.T(l, "settings.sex")); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:32
			for _, figure := range app.Figures {
//line exercise_ui.gox:33
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				choose(ctx, figure)
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line exercise_ui.gox:36
						__e = __c.Set("class", "figure-choice"); if __e != nil { return }
//line exercise_ui.gox:36
						__e = __c.Set("type", "button"); if __e != nil { return }
//line exercise_ui.gox:36
						__e = __c.Set("aria-pressed", strconv.FormatBool(selected == figure)); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:37
						__e = __c.Any(figureTile(l, figure)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line exercise_ui.gox:41
}

// doseLabel formats a dose: "× 12", "45 сек", "5 дыханий", plus per side.
func doseLabel(l i18n.Lang, unit string, value int, perSide bool) string {
	var s string
	switch content.Unit(unit) {
	case content.Seconds:
		s = i18n.Tf(l, "fmt.sec", value)
	case content.Breaths:
		s = i18n.Tn(l, "fmt.breaths", value)
	default:
		s = "× " + strconv.Itoa(value)
	}
	if perSide {
		s += " " + i18n.T(l, "today.per_side")
	}
	return s
}

// whereState is a day selection: "today", "all", or "pick" with a day mask.
type whereState struct {
	Mode string // today | all | pick
	Mask uint32 // bit d-1 set = day d, for "pick"
}

// days turns a selection into the days an edit touches.
func (w whereState) days(today int) content.Days {
	switch w.Mode {
	case "today":
		return content.OneDay(today)
	case "pick":
		var list []int
		for d := 1; d <= content.TotalDays; d++ {
			if w.Mask&(1<<uint(d-1)) != 0 {
				list = append(list, d)
			}
		}
		return content.Days{List: list}
	}
	return content.AllDays()
}

// wherePicker renders "Today / All days / Choose" and, when choosing, a grid
// of the 30 days. update stores a new selection.
//line exercise_ui.gox:85
func wherePicker(l i18n.Lang, w whereState, update func(context.Context, whereState)) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line exercise_ui.gox:86
			__e = __c.Set("class", "seg3"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:87
			for _, opt := range []struct{ Mode, Key string }{{"today", "days.today"}, {"all", "days.all"}, {"pick", "days.pick"}} {
//line exercise_ui.gox:88
				__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
				update(ctx, whereState{Mode: opt.Mode, Mask: w.Mask})
				return false
			}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
					ctx := __c.Context(); _ = ctx
					__e = __c.Init("button"); if __e != nil { return }
					{
//line exercise_ui.gox:91
						__e = __c.Set("type", "button"); if __e != nil { return }
//line exercise_ui.gox:91
						__e = __c.Set("class", segClass(w.Mode == opt.Mode)); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:91
						__e = __c.Any(i18n.T(l, opt.Key)); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				return })); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
//line exercise_ui.gox:94
		if w.Mode == "pick" {
			__e = __c.Init("div"); if __e != nil { return }
			{
//line exercise_ui.gox:95
				__e = __c.Set("class", "day-pick"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:96
				for d := 1; d <= content.TotalDays; d++ {
//line exercise_ui.gox:98
					bit := uint32(1) << uint(d-1)

//line exercise_ui.gox:100
					__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
					update(ctx, whereState{Mode: "pick", Mask: w.Mask ^ bit})
					return false
				}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
						ctx := __c.Context(); _ = ctx
						__e = __c.Init("button"); if __e != nil { return }
						{
//line exercise_ui.gox:103
							__e = __c.Set("type", "button"); if __e != nil { return }
//line exercise_ui.gox:103
							__e = __c.Set("class", segClass(w.Mask&bit != 0)); if __e != nil { return }
							__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:103
							__e = __c.Any(d); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					return })); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
		}
	return })
//line exercise_ui.gox:107
}

// toTop shows a newly opened view from its start.
func toTop(r doors.RequestPointer) {
	_ = r.After(doors.ActionScroll{Selector: "main", Options: map[string]any{"block": "start"}})
}

func segClass(on bool) string {
	if on {
		return "seg on"
	}
	return "seg"
}

// exerciseCard shows an exercise: picture with a figure switcher, technique,
// what is correct and the common mistakes. onAdd (optional) adds it to the program.
type exerciseCard struct {
	userID int64
	lang   string
	id     string
	figure doors.Source[string]
	onBack func(context.Context)
	onAdd  func(context.Context)
}

//line exercise_ui.gox:132
func (c exerciseCard) Main() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line exercise_ui.gox:134
		l := i18n.Lang(c.lang)
		cat := app.UserCatalog(c.userID)
		d, _ := cat.Detail(l, c.id)

		__e = __c.Init("div"); if __e != nil { return }
		{
//line exercise_ui.gox:138
			__e = __c.Set("class", "lib-head"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:139
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			c.onBack(ctx)
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line exercise_ui.gox:142
					__e = __c.Set("class", "lib-back"); if __e != nil { return }
//line exercise_ui.gox:142
					__e = __c.Set("type", "button"); if __e != nil { return }
//line exercise_ui.gox:142
					__e = __c.Set("aria-label", i18n.T(l, "nav.close")); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:142
					__e = __c.Any(icon("back", false)); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
			__e = __c.Init("h1"); if __e != nil { return }
			{
				__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:143
				__e = __c.Any(cat.Name(l, c.id)); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
//line exercise_ui.gox:145
		__e = __c.Any(c.figure.Bind(func(fig string) gox.Elem {
		return gox.Elem(func(__c gox.Cursor) (__e error) {
			ctx := __c.Context(); _ = ctx
//line exercise_ui.gox:146
			if src := firstFrame(app.ExerciseFrames(cat, c.id, fig)); src != "" {
				__e = __c.InitVoid("img"); if __e != nil { return }
				{
//line exercise_ui.gox:147
					__e = __c.Set("class", "card-img"); if __e != nil { return }
//line exercise_ui.gox:147
					__e = __c.Set("src", src); if __e != nil { return }
//line exercise_ui.gox:147
					__e = __c.Set("alt", cat.Name(l, c.id)); if __e != nil { return }
				}
				__e = __c.Submit(); if __e != nil { return }
			}
//line exercise_ui.gox:149
			__e = __c.Any(figureCards(l, fig, func(ctx context.Context, figure string) { c.figure.Update(ctx, figure) })); if __e != nil { return }
		return })
//line exercise_ui.gox:150
	})); if __e != nil { return }
//line exercise_ui.gox:151
		if d.Desc != "" {
			__e = __c.Init("p"); if __e != nil { return }
			{
//line exercise_ui.gox:152
				__e = __c.Set("class", "lib-desc"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:152
				__e = __c.Any(d.Desc); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line exercise_ui.gox:154
		if len(d.HowTo) > 0 {
			__e = __c.Init("div"); if __e != nil { return }
			{
//line exercise_ui.gox:155
				__e = __c.Set("class", "lib-sec"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:156
					__e = __c.Any(i18n.T(l, "library.howto")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ol"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:157
					for _, line := range d.HowTo {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:157
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line exercise_ui.gox:160
		if len(d.Correct) > 0 {
			__e = __c.Init("div"); if __e != nil { return }
			{
//line exercise_ui.gox:161
				__e = __c.Set("class", "lib-sec lib-correct"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:162
					__e = __c.Any(i18n.T(l, "library.correct")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ul"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:163
					for _, line := range d.Correct {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:163
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line exercise_ui.gox:166
		if len(d.Wrong) > 0 {
			__e = __c.Init("div"); if __e != nil { return }
			{
//line exercise_ui.gox:167
				__e = __c.Set("class", "lib-sec lib-mistakes"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("h2"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:168
					__e = __c.Any(i18n.T(l, "library.mistakes")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("ul"); if __e != nil { return }
				{
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:169
					for _, line := range d.Wrong {
						__e = __c.Init("li"); if __e != nil { return }
						{
							__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:169
							__e = __c.Any(line); if __e != nil { return }
						}
						__e = __c.Close(); if __e != nil { return }
					}
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
//line exercise_ui.gox:172
		if c.onAdd != nil {
//line exercise_ui.gox:173
			__e = (doors.AClick{On: func(ctx context.Context, r doors.RequestPointer) bool {
			c.onAdd(ctx)
			return false
		}}).Proxy(__c, gox.Elem(func(__c gox.Cursor) (__e error) {
				ctx := __c.Context(); _ = ctx
				__e = __c.Init("button"); if __e != nil { return }
				{
//line exercise_ui.gox:176
					__e = __c.Set("class", "btn primary wide"); if __e != nil { return }
//line exercise_ui.gox:176
					__e = __c.Set("type", "button"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
//line exercise_ui.gox:176
					__e = __c.Any(icon("plus", false)); if __e != nil { return }
//line exercise_ui.gox:176
					__e = __c.Any(i18n.T(l, "library.add_to_program")); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			return })); if __e != nil { return }
		}
	return })
//line exercise_ui.gox:178
}

var _ = gox.Elem(nil)
