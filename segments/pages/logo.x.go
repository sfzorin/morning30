// Managed by GoX v0.1.36

//line logo.gox:1
package pages

import (
	"fmt"

	"github.com/doors-dev/gox"

	"morning30/internal/brand"
)

// logoMark draws the morning30 sun with its 30 rays (internal/brand), the
// same mark as the app icon, for the login card.
//line logo.gox:13
func logoMark(size int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("svg"); if __e != nil { return }
		{
//line logo.gox:14
			__e = __c.Set("class", "logo-mark"); if __e != nil { return }
//line logo.gox:14
			__e = __c.Set("width", size); if __e != nil { return }
//line logo.gox:14
			__e = __c.Set("height", size); if __e != nil { return }
//line logo.gox:14
			__e = __c.Set("viewBox", "0 0 100 100"); if __e != nil { return }
//line logo.gox:14
			__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:14
			__e = __c.Set("focusable", "false"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line logo.gox:15
			for _, r := range brand.RayList() {
				__e = __c.Init("line"); if __e != nil { return }
				{
//line logo.gox:16
					__e = __c.Set("x1", f2(r.X1)); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("y1", f2(r.Y1)); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("x2", f2(r.X2)); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("y2", f2(r.Y2)); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("stroke", r.Color); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("stroke-width", f2(brand.RayWidth)); if __e != nil { return }
//line logo.gox:16
					__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Init("circle"); if __e != nil { return }
			{
//line logo.gox:18
				__e = __c.Set("cx", "50"); if __e != nil { return }
//line logo.gox:18
				__e = __c.Set("cy", "50"); if __e != nil { return }
//line logo.gox:18
				__e = __c.Set("r", f2(brand.SunR)); if __e != nil { return }
//line logo.gox:18
				__e = __c.Set("fill", brand.Sun); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line logo.gox:20
}

// flameMark is the streak flame: three layers of fire (red, orange, yellow).
//line logo.gox:23
func flameMark(size int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("svg"); if __e != nil { return }
		{
//line logo.gox:24
			__e = __c.Set("class", "flame"); if __e != nil { return }
//line logo.gox:24
			__e = __c.Set("width", size); if __e != nil { return }
//line logo.gox:24
			__e = __c.Set("height", size); if __e != nil { return }
//line logo.gox:24
			__e = __c.Set("viewBox", "0 0 24 24"); if __e != nil { return }
//line logo.gox:24
			__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:24
			__e = __c.Set("focusable", "false"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:25
				__e = __c.Set("d", "M12 1C14.5 5.5 20 8.5 20 15A8 8 0 0 1 4 15C4 11.5 5.8 9.4 7.4 7.8C7.6 10 8.6 11.4 9.8 12C9.4 8.4 10.2 4.2 12 1Z"); if __e != nil { return }
//line logo.gox:25
				__e = __c.Set("fill", "#ef4444"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:26
				__e = __c.Set("d", "M12.4 7.5C14.2 10.3 17 12.2 17 16A5 5 0 0 1 7 16C7 14 8 12.6 9.2 11.6C9.5 13 10.2 13.8 11 14.2C10.9 12 11.3 9.6 12.4 7.5Z"); if __e != nil { return }
//line logo.gox:26
				__e = __c.Set("fill", "#fb923c"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:27
				__e = __c.Set("d", "M12 13C13.4 14.4 14.4 15.6 14.4 17.2A2.4 2.4 0 0 1 9.6 17.2C9.6 15.8 10.6 14.5 12 13Z"); if __e != nil { return }
//line logo.gox:27
				__e = __c.Set("fill", "#fde047"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line logo.gox:29
}

func f2(v float64) string { return fmt.Sprintf("%.2f", v) }

var _ = gox.Elem(nil)
