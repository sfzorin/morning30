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

func f2(v float64) string { return fmt.Sprintf("%.2f", v) }

var _ = gox.Elem(nil)
