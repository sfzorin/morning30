// Command genicons draws the morning30 icon (internal/brand) for the site,
// the iPhone app and the Android app:
//
//	assets/static/icon.svg, icon-180.png, icon-192.png, icon-512.png
//	apps/ios/…/AppIcon.appiconset/icon-1024.png
//	apps/android/…/mipmap-*/ic_launcher.png
//
// PNGs are full-bleed squares; the phones round the corners themselves.
// Run from the repository root: go run ./cmd/genicons
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"morning30/internal/brand"
)

func hex(s string) color.RGBA {
	v, _ := strconv.ParseUint(s[1:], 16, 32)
	return color.RGBA{uint8(v >> 16), uint8(v >> 8), uint8(v), 0xff}
}

// segDist is the distance from p to the segment a–b.
func segDist(px, py, ax, ay, bx, by float64) float64 {
	dx, dy := bx-ax, by-ay
	t := ((px-ax)*dx + (py-ay)*dy) / (dx*dx + dy*dy)
	t = math.Max(0, math.Min(1, t))
	return math.Hypot(px-(ax+t*dx), py-(ay+t*dy))
}

// colorAt samples the mark at a point of the 100×100 canvas.
func colorAt(x, y float64, rays []brand.Ray, bg, sun color.RGBA) color.RGBA {
	if math.Hypot(x-brand.CenterXY, y-brand.CenterXY) <= brand.SunR {
		return sun
	}
	for _, r := range rays {
		if segDist(x, y, r.X1, r.Y1, r.X2, r.Y2) <= brand.RayWidth/2 {
			return hex(r.Color)
		}
	}
	return bg
}

// draw renders the icon with 4×4 supersampling for smooth edges.
func draw(size int) *image.RGBA {
	const ss = 4
	bg, sun, rays := hex(brand.Background), hex(brand.Sun), brand.RayList()
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	scale := 100 / float64(size)
	for py := 0; py < size; py++ {
		for px := 0; px < size; px++ {
			var r, g, b int
			for sy := 0; sy < ss; sy++ {
				for sx := 0; sx < ss; sx++ {
					c := colorAt((float64(px)+(float64(sx)+.5)/ss)*scale, (float64(py)+(float64(sy)+.5)/ss)*scale, rays, bg, sun)
					r, g, b = r+int(c.R), g+int(c.G), b+int(c.B)
				}
			}
			n := ss * ss
			img.Set(px, py, color.RGBA{uint8(r / n), uint8(g / n), uint8(b / n), 0xff})
		}
	}
	return img
}

func write(size int, path string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, draw(size)); err != nil {
		panic(err)
	}
}

func main() {
	if err := os.WriteFile(filepath.Join("assets", "static", "icon.svg"), []byte(brand.IconSVG()), 0o644); err != nil {
		panic(err)
	}
	for _, s := range []int{180, 192, 512} {
		write(s, filepath.Join("assets", "static", "icon-"+strconv.Itoa(s)+".png"))
	}
	write(1024, filepath.Join("apps", "ios", "Morning30", "Assets.xcassets", "AppIcon.appiconset", "icon-1024.png"))
	for dir, s := range map[string]int{"mdpi": 48, "hdpi": 72, "xhdpi": 96, "xxhdpi": 144, "xxxhdpi": 192} {
		write(s, filepath.Join("apps", "android", "app", "src", "main", "res", "mipmap-"+dir, "ic_launcher.png"))
	}
	println("wrote icons")
}
