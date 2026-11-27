// Package brand holds the morning30 mark: a sun ringed by 30 rays, one per
// day of the cycle, most of them lit. It is described once here and drawn as
// inline SVG (site, login, course pages) and as PNG icons (cmd/genicons).
package brand

import (
	"fmt"
	"math"
	"strings"
)

// Geometry on a 100×100 canvas.
const (
	Background = "#0f172a" // app icon tile
	Sun        = "#fb923c"
	Unlit      = "#334155"

	SunR      = 17.0 // sun disc radius
	RayInner  = 28.0 // rays run from this radius…
	RayOuter  = 37.0 // …to this one
	RayWidth  = 3.6
	Rays      = 30 // one per day of the cycle
	LitRays   = 22 // the lit part, read clockwise from the top
	CenterXY  = 50.0
	TileRound = 22.0
)

// Ray is one of the 30 rays as a line segment with its colour.
type Ray struct {
	X1, Y1, X2, Y2 float64
	Color          string
}

// RayList returns the rays clockwise from twelve o'clock.
func RayList() []Ray {
	out := make([]Ray, Rays)
	for i := range out {
		a := (float64(i)*360/Rays - 90) * math.Pi / 180
		c := Sun
		if i >= LitRays {
			c = Unlit
		}
		out[i] = Ray{
			X1: CenterXY + RayInner*math.Cos(a), Y1: CenterXY + RayInner*math.Sin(a),
			X2: CenterXY + RayOuter*math.Cos(a), Y2: CenterXY + RayOuter*math.Sin(a),
			Color: c,
		}
	}
	return out
}

func shapes() string {
	var b strings.Builder
	for _, r := range RayList() {
		fmt.Fprintf(&b, `<line x1="%.2f" y1="%.2f" x2="%.2f" y2="%.2f" stroke="%s" stroke-width="%.1f" stroke-linecap="round"/>`, r.X1, r.Y1, r.X2, r.Y2, r.Color, RayWidth)
	}
	fmt.Fprintf(&b, `<circle cx="%.0f" cy="%.0f" r="%.0f" fill="%s"/>`, CenterXY, CenterXY, SunR, Sun)
	return b.String()
}

// Mark is the bare mark for use on the app's own dark background. class and
// size go on the <svg> element; it is decorative (aria-hidden).
func Mark(class string, size int) string {
	return fmt.Sprintf(`<svg class="%s" width="%d" height="%d" viewBox="0 0 100 100" aria-hidden="true" focusable="false">%s</svg>`, class, size, size, shapes())
}

// IconSVG is the standalone app icon: the mark on a rounded dark tile.
func IconSVG() string {
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" width="192" height="192"><rect width="100" height="100" rx="%.0f" fill="%s"/>%s</svg>`+"\n", TileRound, Background, shapes())
}
