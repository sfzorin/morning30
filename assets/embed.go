// Package assets manages embedded static files for the application.
// Files are embedded at compile time using Go's embed directive,
// then served via Doors middleware or used as managed resources.
package assets

import (
	"embed"
	"io/fs"
)

// Style is the embedded CSS content served as a Doors managed resource.
//
//go:embed style.css
var Style []byte

// Player is the client-side workout engine, served as a managed inline script.
//
//go:embed player.js
var Player []byte

// ExUpload wires the exercise-JSON file picker to its textarea (read via
// FileReader so the server only ever receives text). Served as an inline script.
//
//go:embed exupload.js
var ExUpload []byte

// LeapLibrary is the localized exercise library imported from the Leap Fitness
// "How to Do" videos (names + cards in 7 languages + source video). Parsed by
// the internal/library package.
//
//go:embed leap_library.json
var LeapLibrary []byte

// LeapVideoMap maps existing app exercise IDs to a matching demo video.
//
//go:embed leap_video_map.json
var LeapVideoMap []byte

// static holds the static/ directory tree (icons, manifest, sw, exercise SVGs),
// embedded recursively and exposed via Static() for the UseFS middleware.
//
//go:embed static
var static embed.FS

// Static returns an fs.FS rooted at the static/ subdirectory.
func Static() fs.FS {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
