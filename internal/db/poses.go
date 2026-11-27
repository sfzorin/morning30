package db

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const maxPoseBytes = 4 << 20 // 4 MiB per still

var poseIDRe = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{0,63}$`)

// ErrPose is returned when a pose id, sex, or image cannot be stored.
var ErrPose = errors.New("invalid pose")

// Pose is one still in the shared illustration catalog (no image bytes).
type Pose struct {
	ID    string `json:"id"`
	Kind  string `json:"kind,omitempty"`
	Label string `json:"label,omitempty"`
	HasM  bool   `json:"hasM"`
	HasF  bool   `json:"hasF"`
	HasB  bool   `json:"hasB"` // boy
	HasG  bool   `json:"hasG"` // girl
}

// NormalizePoseID lowercases and trims a catalog slug. Empty if illegal.
func NormalizePoseID(id string) string {
	s := strings.ToLower(strings.TrimSpace(id))
	if !poseIDRe.MatchString(s) {
		return ""
	}
	return s
}

// Figures: m (man), f (woman), b (boy), g (girl).

// NormalizeSex maps a figure code to m, f, b or g; anything else is "m".
func NormalizeSex(sex string) string {
	switch s := strings.ToLower(strings.TrimSpace(sex)); s {
	case "f", "b", "g":
		return s
	}
	return "m"
}

// figureFallback is the order in which stills are tried for a figure: a child
// without its own still is shown as the adult of the same sex.
var figureFallback = map[string][]string{
	"m": {"m", "f"},
	"f": {"f", "m"},
	"b": {"b", "m", "g", "f"},
	"g": {"g", "f", "b", "m"},
}

// SniffImageMIME returns image/png, image/jpeg, image/webp, or "".
func SniffImageMIME(b []byte) string {
	switch {
	case bytes.HasPrefix(b, []byte("\x89PNG\r\n\x1a\n")):
		return "image/png"
	case bytes.HasPrefix(b, []byte{0xff, 0xd8, 0xff}):
		return "image/jpeg"
	case len(b) >= 12 && bytes.Equal(b[:4], []byte("RIFF")) && bytes.Equal(b[8:12], []byte("WEBP")):
		return "image/webp"
	default:
		return ""
	}
}

func clipLabel(s string) string {
	s = strings.TrimSpace(s)
	if utf8.RuneCountInString(s) > 80 {
		r := []rune(s)
		s = string(r[:80])
	}
	return s
}

func clipKind(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s != "" && !poseIDRe.MatchString(s) {
		return ""
	}
	return s
}

// PutPose upserts catalog metadata and stores one still for figure m|f|b|g.
func (d *DB) PutPose(id, kind, label, sex string, data []byte) error {
	id = NormalizePoseID(id)
	if id == "" {
		return fmt.Errorf("%w: id", ErrPose)
	}
	if len(data) == 0 || len(data) > maxPoseBytes {
		return fmt.Errorf("%w: image size", ErrPose)
	}
	mime := SniffImageMIME(data)
	if mime == "" {
		return fmt.Errorf("%w: image type", ErrPose)
	}
	sex = NormalizeSex(sex)
	kind = clipKind(kind)
	label = clipLabel(label)

	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(
		`INSERT INTO poses(id, kind, label, updated_at) VALUES(?, ?, ?, datetime('now'))
		 ON CONFLICT(id) DO UPDATE SET
		   kind = CASE WHEN excluded.kind != '' THEN excluded.kind ELSE poses.kind END,
		   label = CASE WHEN excluded.label != '' THEN excluded.label ELSE poses.label END,
		   updated_at = datetime('now')`,
		id, kind, label)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`INSERT INTO pose_images(pose_id, sex, mime, data) VALUES(?, ?, ?, ?)
		 ON CONFLICT(pose_id, sex) DO UPDATE SET mime = excluded.mime, data = excluded.data`,
		id, sex, mime, data)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// ListPoses returns the catalog, grouped-ready (kind then id).
func (d *DB) ListPoses() ([]Pose, error) {
	rows, err := d.sql.Query(`
		SELECT p.id, p.kind, p.label,
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'm'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'f'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'b'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'g')
		FROM poses p
		ORDER BY p.kind, p.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]Pose, 0)
	for rows.Next() {
		var p Pose
		var hasM, hasF, hasB, hasG int
		if err := rows.Scan(&p.ID, &p.Kind, &p.Label, &hasM, &hasF, &hasB, &hasG); err != nil {
			return nil, err
		}
		p.HasM, p.HasF, p.HasB, p.HasG = hasM != 0, hasF != 0, hasB != 0, hasG != 0
		out = append(out, p)
	}
	return out, rows.Err()
}

// HasPose reports whether any still exists for id.
func (d *DB) HasPose(id string) bool {
	id = NormalizePoseID(id)
	if id == "" {
		return false
	}
	var n int
	_ = d.sql.QueryRow(`SELECT COUNT(*) FROM pose_images WHERE pose_id = ?`, id).Scan(&n)
	return n > 0
}

// HasPoseFigure checks only the requested variant, without reading the image blob.
func (d *DB) HasPoseFigure(id, sex string) bool {
	var exists bool
	err := d.sql.QueryRow(`SELECT EXISTS(SELECT 1 FROM pose_images WHERE pose_id = ? AND sex = ?)`,
		NormalizePoseID(id), NormalizeSex(sex)).Scan(&exists)
	return err == nil && exists
}

// PoseImage returns bytes for sex, falling back to the other sex.
func (d *DB) PoseImage(id, sex string) (mime string, data []byte, err error) {
	id = NormalizePoseID(id)
	if id == "" {
		return "", nil, sql.ErrNoRows
	}
	for _, s := range figureFallback[NormalizeSex(sex)] {
		err = d.sql.QueryRow(
			`SELECT mime, data FROM pose_images WHERE pose_id = ? AND sex = ?`,
			id, s).Scan(&mime, &data)
		if err == nil {
			return mime, data, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return "", nil, err
		}
	}
	return "", nil, sql.ErrNoRows
}

// GetPose returns catalog metadata for one id.
func (d *DB) GetPose(id string) (Pose, error) {
	id = NormalizePoseID(id)
	if id == "" {
		return Pose{}, sql.ErrNoRows
	}
	var p Pose
	var hasM, hasF, hasB, hasG int
	err := d.sql.QueryRow(`
		SELECT p.id, p.kind, p.label,
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'm'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'f'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'b'),
		       EXISTS(SELECT 1 FROM pose_images i WHERE i.pose_id = p.id AND i.sex = 'g')
		FROM poses p WHERE p.id = ?`, id).Scan(&p.ID, &p.Kind, &p.Label, &hasM, &hasF, &hasB, &hasG)
	if err != nil {
		return Pose{}, err
	}
	p.HasM, p.HasF, p.HasB, p.HasG = hasM != 0, hasF != 0, hasB != 0, hasG != 0
	return p, nil
}
