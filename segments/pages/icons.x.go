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
		class := "ic"
		if filled {
			class = "ic fill"
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
			} else if name == "next" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:21
					__e = __c.Set("d", "M9 6l6 6-6 6"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:22
			} else if name == "close" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:23
					__e = __c.Set("d", "M18 6L6 18"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:23
					__e = __c.Set("d", "M6 6l12 12"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:24
			} else if name == "plus" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:25
					__e = __c.Set("d", "M12 5v14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:25
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:26
			} else if name == "minus" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:27
					__e = __c.Set("d", "M5 12h14"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:28
			} else if name == "edit" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:29
					__e = __c.Set("d", "M4 20h4L18.5 9.5a2.83 2.83 0 0 0-4-4L4 16v4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:29
					__e = __c.Set("d", "M13.5 6.5l4 4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:30
			} else if name == "star" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:31
					__e = __c.Set("d", "M12 17.75l-6.17 3.25 1.18-6.88-5-4.87 6.9-1L12 2l3.09 6.25 6.9 1-5 4.87 1.18 6.88z"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:32
			} else if name == "info" {
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:33
					__e = __c.Set("cx", "12"); if __e != nil { return }
//line icons.gox:33
					__e = __c.Set("cy", "12"); if __e != nil { return }
//line icons.gox:33
					__e = __c.Set("r", "9"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:33
					__e = __c.Set("d", "M12 8h.01"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:33
					__e = __c.Set("d", "M11 12h1v4h1"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:34
			} else if name == "replace" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:35
					__e = __c.Set("d", "M7 10h14l-4-4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:35
					__e = __c.Set("d", "M17 14H3l4 4"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
//line icons.gox:36
			} else if name == "settings" {
				__e = __c.Init("path"); if __e != nil { return }
				{
//line icons.gox:37
					__e = __c.Set("d", "M10.33 4.32c.43-1.76 2.92-1.76 3.35 0a1.72 1.72 0 0 0 2.57 1.07c1.54-.94 3.31.83 2.37 2.37a1.72 1.72 0 0 0 1.07 2.57c1.76.43 1.76 2.92 0 3.35a1.72 1.72 0 0 0-1.07 2.57c.94 1.54-.83 3.31-2.37 2.37a1.72 1.72 0 0 0-2.57 1.07c-.43 1.76-2.92 1.76-3.35 0a1.72 1.72 0 0 0-2.57-1.07c-1.54.94-3.31-.83-2.37-2.37a1.72 1.72 0 0 0-1.07-2.57c-1.76-.43-1.76-2.92 0-3.35a1.72 1.72 0 0 0 1.07-2.57c-.94-1.54.83-3.31 2.37-2.37 1 .61 2.3.07 2.57-1.07z"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("circle"); if __e != nil { return }
				{
//line icons.gox:37
					__e = __c.Set("cx", "12"); if __e != nil { return }
//line icons.gox:37
					__e = __c.Set("cy", "12"); if __e != nil { return }
//line icons.gox:37
					__e = __c.Set("r", "3"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line icons.gox:40
}

var _ = gox.Elem(nil)
