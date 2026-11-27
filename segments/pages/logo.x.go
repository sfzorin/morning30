// Managed by GoX v0.1.36

//line logo.gox:1
package pages

import (
	"fmt"
	"math"

	"github.com/doors-dev/gox"

	"morning30/internal/brand"
)

// logoMark draws the morning30 sun with its 30 rays (internal/brand), the
// same mark as the app icon, for the login card.
//line logo.gox:14
func logoMark(size int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("svg"); if __e != nil { return }
		{
//line logo.gox:15
			__e = __c.Set("class", "logo-mark"); if __e != nil { return }
//line logo.gox:15
			__e = __c.Set("width", size); if __e != nil { return }
//line logo.gox:15
			__e = __c.Set("height", size); if __e != nil { return }
//line logo.gox:15
			__e = __c.Set("viewBox", "0 0 100 100"); if __e != nil { return }
//line logo.gox:15
			__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:15
			__e = __c.Set("focusable", "false"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
//line logo.gox:16
			for _, r := range brand.RayList() {
				__e = __c.Init("line"); if __e != nil { return }
				{
//line logo.gox:17
					__e = __c.Set("x1", f2(r.X1)); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("y1", f2(r.Y1)); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("x2", f2(r.X2)); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("y2", f2(r.Y2)); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("stroke", r.Color); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("stroke-width", f2(brand.RayWidth)); if __e != nil { return }
//line logo.gox:17
					__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Init("circle"); if __e != nil { return }
			{
//line logo.gox:19
				__e = __c.Set("cx", "50"); if __e != nil { return }
//line logo.gox:19
				__e = __c.Set("cy", "50"); if __e != nil { return }
//line logo.gox:19
				__e = __c.Set("r", f2(brand.SunR)); if __e != nil { return }
//line logo.gox:19
				__e = __c.Set("fill", brand.Sun); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line logo.gox:21
}

// flameMark is the streak flame: three layers of fire (red, orange, yellow).
//line logo.gox:24
func flameMark(size int) gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("svg"); if __e != nil { return }
		{
//line logo.gox:25
			__e = __c.Set("class", "flame"); if __e != nil { return }
//line logo.gox:25
			__e = __c.Set("width", size); if __e != nil { return }
//line logo.gox:25
			__e = __c.Set("height", size); if __e != nil { return }
//line logo.gox:25
			__e = __c.Set("viewBox", "0 0 24 24"); if __e != nil { return }
//line logo.gox:25
			__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:25
			__e = __c.Set("focusable", "false"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:26
				__e = __c.Set("d", "M12 1C14.5 5.5 20 8.5 20 15A8 8 0 0 1 4 15C4 11.5 5.8 9.4 7.4 7.8C7.6 10 8.6 11.4 9.8 12C9.4 8.4 10.2 4.2 12 1Z"); if __e != nil { return }
//line logo.gox:26
				__e = __c.Set("fill", "#ef4444"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:27
				__e = __c.Set("d", "M12.4 7.5C14.2 10.3 17 12.2 17 16A5 5 0 0 1 7 16C7 14 8 12.6 9.2 11.6C9.5 13 10.2 13.8 11 14.2C10.9 12 11.3 9.6 12.4 7.5Z"); if __e != nil { return }
//line logo.gox:27
				__e = __c.Set("fill", "#fb923c"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("path"); if __e != nil { return }
			{
//line logo.gox:28
				__e = __c.Set("d", "M12 13C13.4 14.4 14.4 15.6 14.4 17.2A2.4 2.4 0 0 1 9.6 17.2C9.6 15.8 10.6 14.5 12 13Z"); if __e != nil { return }
//line logo.gox:28
				__e = __c.Set("fill", "#fde047"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line logo.gox:30
}

// trophyMark is the finished-workout trophy: a gold cup with a star, 30
// slowly turning rays behind it (as in the logo) and a few twinkles.
//line logo.gox:34
func trophyMark() gox.Elem {
	return gox.Elem(func(__c gox.Cursor) (__e error) {
		ctx := __c.Context(); _ = ctx
		__e = __c.Init("div"); if __e != nil { return }
		{
//line logo.gox:35
			__e = __c.Set("class", "trophy"); if __e != nil { return }
			__e = __c.Submit(); if __e != nil { return }
			__e = __c.Init("svg"); if __e != nil { return }
			{
//line logo.gox:36
				__e = __c.Set("class", "trophy-rays"); if __e != nil { return }
//line logo.gox:36
				__e = __c.Set("viewBox", "0 0 170 170"); if __e != nil { return }
//line logo.gox:36
				__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:36
				__e = __c.Set("focusable", "false"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
//line logo.gox:37
				for i := 0; i < 30; i++ {
//line logo.gox:38
					x1, y1, x2, y2, op, w := trophyRay(i)
					__e = __c.Init("line"); if __e != nil { return }
					{
//line logo.gox:39
						__e = __c.Set("x1", x1); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("y1", y1); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("x2", x2); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("y2", y2); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("stroke", "#fbbf24"); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("stroke-opacity", op); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("stroke-width", w); if __e != nil { return }
//line logo.gox:39
						__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
			}
			__e = __c.Close(); if __e != nil { return }
			__e = __c.Init("svg"); if __e != nil { return }
			{
//line logo.gox:42
				__e = __c.Set("class", "trophy-cup"); if __e != nil { return }
//line logo.gox:42
				__e = __c.Set("viewBox", "0 0 170 170"); if __e != nil { return }
//line logo.gox:42
				__e = __c.Set("aria-hidden", "true"); if __e != nil { return }
//line logo.gox:42
				__e = __c.Set("focusable", "false"); if __e != nil { return }
				__e = __c.Submit(); if __e != nil { return }
				__e = __c.Init("g"); if __e != nil { return }
				{
//line logo.gox:43
					__e = __c.Set("transform", "translate(85 88)"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:44
						__e = __c.Set("d", "M-44 -46 C-66 -46 -70 -16 -42 -4"); if __e != nil { return }
//line logo.gox:44
						__e = __c.Set("fill", "none"); if __e != nil { return }
//line logo.gox:44
						__e = __c.Set("stroke", "#d97706"); if __e != nil { return }
//line logo.gox:44
						__e = __c.Set("stroke-width", "9"); if __e != nil { return }
//line logo.gox:44
						__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:45
						__e = __c.Set("d", "M44 -46 C66 -46 70 -16 42 -4"); if __e != nil { return }
//line logo.gox:45
						__e = __c.Set("fill", "none"); if __e != nil { return }
//line logo.gox:45
						__e = __c.Set("stroke", "#b45309"); if __e != nil { return }
//line logo.gox:45
						__e = __c.Set("stroke-width", "9"); if __e != nil { return }
//line logo.gox:45
						__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:46
						__e = __c.Set("d", "M-44 -50 H44 C44 -12 30 10 0 14 C-30 10 -44 -12 -44 -50Z"); if __e != nil { return }
//line logo.gox:46
						__e = __c.Set("fill", "#fbbf24"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:47
						__e = __c.Set("d", "M0 -50 H44 C44 -12 30 10 0 14Z"); if __e != nil { return }
//line logo.gox:47
						__e = __c.Set("fill", "#f59e0b"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:48
						__e = __c.Set("d", "M-34 -44 C-34 -24 -28 -10 -16 0"); if __e != nil { return }
//line logo.gox:48
						__e = __c.Set("fill", "none"); if __e != nil { return }
//line logo.gox:48
						__e = __c.Set("stroke", "#fef3c7"); if __e != nil { return }
//line logo.gox:48
						__e = __c.Set("stroke-width", "5"); if __e != nil { return }
//line logo.gox:48
						__e = __c.Set("stroke-linecap", "round"); if __e != nil { return }
//line logo.gox:48
						__e = __c.Set("opacity", ".9"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:49
						__e = __c.Set("x", "-48"); if __e != nil { return }
//line logo.gox:49
						__e = __c.Set("y", "-56"); if __e != nil { return }
//line logo.gox:49
						__e = __c.Set("width", "96"); if __e != nil { return }
//line logo.gox:49
						__e = __c.Set("height", "10"); if __e != nil { return }
//line logo.gox:49
						__e = __c.Set("rx", "5"); if __e != nil { return }
//line logo.gox:49
						__e = __c.Set("fill", "#fcd34d"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:50
						__e = __c.Set("x", "0"); if __e != nil { return }
//line logo.gox:50
						__e = __c.Set("y", "-56"); if __e != nil { return }
//line logo.gox:50
						__e = __c.Set("width", "48"); if __e != nil { return }
//line logo.gox:50
						__e = __c.Set("height", "10"); if __e != nil { return }
//line logo.gox:50
						__e = __c.Set("rx", "5"); if __e != nil { return }
//line logo.gox:50
						__e = __c.Set("fill", "#fbbf24"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:51
						__e = __c.Set("d", "M0 -38 L5 -27 L17 -26 L8 -18 L11 -6 L0 -12 L-11 -6 L-8 -18 L-17 -26 L-5 -27Z"); if __e != nil { return }
//line logo.gox:51
						__e = __c.Set("fill", "#fff7d6"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:52
						__e = __c.Set("d", "M-7 14 H7 L10 32 H-10Z"); if __e != nil { return }
//line logo.gox:52
						__e = __c.Set("fill", "#d97706"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("path"); if __e != nil { return }
					{
//line logo.gox:53
						__e = __c.Set("d", "M0 14 H7 L10 32 H0Z"); if __e != nil { return }
//line logo.gox:53
						__e = __c.Set("fill", "#b45309"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:54
						__e = __c.Set("x", "-26"); if __e != nil { return }
//line logo.gox:54
						__e = __c.Set("y", "32"); if __e != nil { return }
//line logo.gox:54
						__e = __c.Set("width", "52"); if __e != nil { return }
//line logo.gox:54
						__e = __c.Set("height", "11"); if __e != nil { return }
//line logo.gox:54
						__e = __c.Set("rx", "3"); if __e != nil { return }
//line logo.gox:54
						__e = __c.Set("fill", "#f59e0b"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:55
						__e = __c.Set("x", "0"); if __e != nil { return }
//line logo.gox:55
						__e = __c.Set("y", "32"); if __e != nil { return }
//line logo.gox:55
						__e = __c.Set("width", "26"); if __e != nil { return }
//line logo.gox:55
						__e = __c.Set("height", "11"); if __e != nil { return }
//line logo.gox:55
						__e = __c.Set("rx", "3"); if __e != nil { return }
//line logo.gox:55
						__e = __c.Set("fill", "#d97706"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:56
						__e = __c.Set("x", "-34"); if __e != nil { return }
//line logo.gox:56
						__e = __c.Set("y", "43"); if __e != nil { return }
//line logo.gox:56
						__e = __c.Set("width", "68"); if __e != nil { return }
//line logo.gox:56
						__e = __c.Set("height", "14"); if __e != nil { return }
//line logo.gox:56
						__e = __c.Set("rx", "4"); if __e != nil { return }
//line logo.gox:56
						__e = __c.Set("fill", "#92400e"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
					__e = __c.Init("rect"); if __e != nil { return }
					{
//line logo.gox:57
						__e = __c.Set("x", "-20"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("y", "47"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("width", "40"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("height", "5"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("rx", "2"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("fill", "#fbbf24"); if __e != nil { return }
//line logo.gox:57
						__e = __c.Set("opacity", ".85"); if __e != nil { return }
						__e = __c.Submit(); if __e != nil { return }
					}
					__e = __c.Close(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line logo.gox:59
					__e = __c.Set("class", "twinkle t1"); if __e != nil { return }
//line logo.gox:59
					__e = __c.Set("d", "M34 36 L36.5 43.5 L44 46 L36.5 48.5 L34 56 L31.5 48.5 L24 46 L31.5 43.5Z"); if __e != nil { return }
//line logo.gox:59
					__e = __c.Set("fill", "#ffffff"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line logo.gox:60
					__e = __c.Set("class", "twinkle t2"); if __e != nil { return }
//line logo.gox:60
					__e = __c.Set("d", "M138 56 L140 62 L146 64 L140 66 L138 72 L136 66 L130 64 L136 62Z"); if __e != nil { return }
//line logo.gox:60
					__e = __c.Set("fill", "#fde68a"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
				__e = __c.Init("path"); if __e != nil { return }
				{
//line logo.gox:61
					__e = __c.Set("class", "twinkle t3"); if __e != nil { return }
//line logo.gox:61
					__e = __c.Set("d", "M128 20 L129.5 24.5 L134 26 L129.5 27.5 L128 32 L126.5 27.5 L122 26 L126.5 24.5Z"); if __e != nil { return }
//line logo.gox:61
					__e = __c.Set("fill", "#ffffff"); if __e != nil { return }
					__e = __c.Submit(); if __e != nil { return }
				}
				__e = __c.Close(); if __e != nil { return }
			}
			__e = __c.Close(); if __e != nil { return }
		}
		__e = __c.Close(); if __e != nil { return }
	return })
//line logo.gox:64
}

// trophyRay places ray i of 30 around the cup; every other ray is fainter.
func trophyRay(i int) (x1, y1, x2, y2, opacity, width string) {
	a := float64(i) * 12 * math.Pi / 180
	op, w := "0.22", "7"
	if i%2 == 1 {
		op, w = "0.10", "4"
	}
	return f2(85 + 50*math.Cos(a)), f2(88 + 50*math.Sin(a)), f2(85 + 82*math.Cos(a)), f2(88 + 82*math.Sin(a)), op, w
}

func f2(v float64) string { return fmt.Sprintf("%.2f", v) }

var _ = gox.Elem(nil)
