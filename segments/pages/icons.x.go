// Managed by GoX v0.1.36

//line icons.gox:1
package pages

import (
	"github.com/doors-dev/gox"
)

// icon draws a small outline icon (24×24 grid, current text colour) so
// buttons never rely on text symbols or emoji that render differently on
// every phone. filled paints the shape (used for a selected star).
//line icons.gox:10
func icon(name string, filled bool) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
//line icons.gox:12
		class := "ic ic-" + name // the name lets right-to-left layouts mirror arrows
		if filled {
			class += " fill"
		}

		__e = __c.Init("svg"); if __e != nil { return }
		{
//line icons.gox:17
			__e = __c.Set("class", class); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("viewBox", "0 0 24 24"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("fill", "none"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("stroke", "currentColor"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("stroke-width", "2"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("stroke-linejoin", "round"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line icons.gox:17
			__e = __c.Set("focusable", "false"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line icons.gox:18
			if name == "back" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:19
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:19
					__e = __c.Set("d", "M5 12l6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:19
					__e = __c.Set("d", "M5 12l6-6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:20
			} else if name == "prev" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:21
					__e = __c.Set("d", "M15 6l-6 6 6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:22
			} else if name == "next" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:23
					__e = __c.Set("d", "M9 6l6 6-6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:24
			} else if name == "to" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:25
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:25
					__e = __c.Set("d", "M13 6l6 6-6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:26
			} else if name == "trophy" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("d", "M8 21h8"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("d", "M12 17v4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("d", "M7 4h10"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("d", "M17 4v8a5 5 0 0 1-10 0V4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("cx", "5"); if __e != nil { return }
//line icons.gox:27
					__e = __c.Set("cy", "9"); if __e != nil { return }
//line icons.gox:27
					__e = __c.Set("r", "2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("cx", "19"); if __e != nil { return }
//line icons.gox:27
					__e = __c.Set("cy", "9"); if __e != nil { return }
//line icons.gox:27
					__e = __c.Set("r", "2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:28
			} else if name == "check" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:29
					__e = __c.Set("d", "M5 12l5 5L20 7"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:30
			} else if name == "up" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:31
					__e = __c.Set("d", "M12 19V5"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:31
					__e = __c.Set("d", "M6 11l6-6 6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:32
			} else if name == "down" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:33
					__e = __c.Set("d", "M12 5v14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:33
					__e = __c.Set("d", "M6 13l6 6 6-6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:34
			} else if name == "close" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:35
					__e = __c.Set("d", "M18 6L6 18"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:35
					__e = __c.Set("d", "M6 6l12 12"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:36
			} else if name == "plus" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:37
					__e = __c.Set("d", "M12 5v14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:37
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:38
			} else if name == "minus" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:39
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:40
			} else if name == "edit" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:41
					__e = __c.Set("d", "M4 20h4L18.5 9.5a2.83 2.83 0 0 0-4-4L4 16v4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:41
					__e = __c.Set("d", "M13.5 6.5l4 4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:42
			} else if name == "star" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:43
					__e = __c.Set("d", "M12 17.75l-6.17 3.25 1.18-6.88-5-4.87 6.9-1L12 2l3.09 6.25 6.9 1-5 4.87 1.18 6.88z"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:44
			} else if name == "info" {
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:45
					__e = __c.Set("cx", "12"); if __e != nil { return }
//line icons.gox:45
					__e = __c.Set("cy", "12"); if __e != nil { return }
//line icons.gox:45
					__e = __c.Set("r", "9"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:45
					__e = __c.Set("d", "M12 8h.01"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:45
					__e = __c.Set("d", "M11 12h1v4h1"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:46
			} else if name == "replace" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:47
					__e = __c.Set("d", "M7 10h14l-4-4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:47
					__e = __c.Set("d", "M17 14H3l4 4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:48
			} else if name == "sparkles" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:49
					__e = __c.Set("d", "M16 18a2 2 0 0 1 2 2a2 2 0 0 1 2-2a2 2 0 0 1-2-2a2 2 0 0 1-2 2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:49
					__e = __c.Set("d", "M16 6a2 2 0 0 1 2 2a2 2 0 0 1 2-2a2 2 0 0 1-2-2a2 2 0 0 1-2 2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:49
					__e = __c.Set("d", "M9 18a6 6 0 0 1 6-6a6 6 0 0 1-6-6a6 6 0 0 1-6 6a6 6 0 0 1 6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:50
			} else if name == "settings" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M4 6h8"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M16 6h4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("cx", "14"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("cy", "6"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("r", "2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M4 12h2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M10 12h10"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("cx", "8"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("cy", "12"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("r", "2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M4 18h11"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("d", "M19 18h1"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:51
					__e = __c.Set("cx", "17"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("cy", "18"); if __e != nil { return }
//line icons.gox:51
					__e = __c.Set("r", "2"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line icons.gox:54
}

var _ = gox.Elem(nil)
