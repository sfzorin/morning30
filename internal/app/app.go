// Package app holds process-wide dependencies shared by the page/segment code.
package app

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"

	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
)

var (
	// ErrInvalidProgram is returned when uploaded program JSON is not a usable cycle.
	ErrInvalidProgram = errors.New("invalid program")
	// ErrUnknownCourse is returned when a "… in 30 days" course key matches nothing.
	ErrUnknownCourse = errors.New("unknown course")
	// ErrEmptyProgramName is returned when a program is renamed to nothing.
	ErrEmptyProgramName = errors.New("program name is empty")
	// ErrInvalidExercise is returned when a custom exercise document is unusable.
	ErrInvalidExercise = errors.New("invalid exercise")
)

// DB is the application database, set once at startup in main.
var DB *db.DB

// Today returns the server-local calendar date as YYYY-MM-DD.
func Today() string {
	return time.Now().Format("2006-01-02")
}

// ProgramByKey returns the program for "custom" (the stored one, if any) or a
// standard program by key, falling back to the default standard program.
func ProgramByKey(userID int64, key string) content.Resolved {
	if key == "custom" {
		if j, err := DB.GetProgramJSON(userID); err == nil && j != "" {
			if r, err := content.ParseResolved(j); err == nil && len(r.Days) > 0 {
				return r
			}
		}
	}
	if r, ok := content.StandardByKey(key); ok {
		return r
	}
	r, _ := content.StandardByKey(content.DefaultStandardKey())
	return r
}

// SeedBuiltins makes the global `exercises` table match the built-in library
// (authoritative upsert + prune), so it is the canonical base every user reads.
func SeedBuiltins() {
	ids := make([]string, 0, len(content.Library()))
	for _, ex := range content.Library() {
		d, ok := catalog.Builtin(ex.ID)
		if !ok {
			continue
		}
		if raw, err := DB.GetExercise(ex.ID); err == nil && raw != "" {
			var old catalog.Doc
			if json.Unmarshal([]byte(raw), &old) == nil {
				d.Poses = catalog.ClipPoses(old.Poses)
				if catalog.ValidAnimation(old.Animation) {
					d.Animation = old.Animation
				}
			}
		}
		_ = DB.SetExercise(ex.ID, d.Compact()) // upsert (code is canonical; poses kept)
		ids = append(ids, ex.ID)
	}
	_ = DB.PruneExercises(ids) // drop exercises removed from the library
}

// UserCatalog returns the user's exercise catalog: the global DB library overlaid
// with their custom exercises. Pages use it to resolve names/content/voice so
// custom exercises are selectable, shown and spoken.
func UserCatalog(userID int64) catalog.Catalog {
	base, _ := DB.AllExercises()
	j, _ := DB.GetCustomExercises(userID)
	return catalog.NewLayered(base, j)
}

// SaveCustomExercise adds or overrides a custom exercise in the user's library.
func SaveCustomExercise(userID int64, d catalog.Doc) error {
	d.ID = strings.TrimSpace(d.ID)
	if d.Slot == "" {
		d.Slot = string(content.Main)
	}
	if d.Unit == "" {
		d.Unit = string(content.Reps)
	}
	if !catalog.ValidDoc(d) {
		return ErrInvalidExercise
	}
	d.SVG = catalog.NormalizeFigure(d.SVG, d.ID)
	d.Poses = catalog.ClipPoses(d.Poses)
	c := UserCatalog(userID)
	c.Put(d)
	return DB.SetCustomExercises(userID, c.Marshal())
}

// AssignExercisePoses stores 0–2 pose catalog ids on an exercise. Custom
// overlay wins for that account; otherwise the global library row is updated
// (SeedBuiltins keeps those ids).
func AssignExercisePoses(userID int64, exerciseID string, poses []string) error {
	exerciseID = strings.TrimSpace(exerciseID)
	poses = catalog.ClipPoses(poses)
	cat := UserCatalog(userID)
	d, ok := cat.Doc(exerciseID)
	if !ok {
		return ErrInvalidExercise
	}
	d.Poses = poses
	if cat.IsCustom(exerciseID) {
		return SaveCustomExercise(userID, d)
	}
	return DB.SetExercise(exerciseID, d.Compact())
}

// SetLibraryPoses writes pose ids onto a global library exercise (MCP).
func SetLibraryPoses(exerciseID string, poses []string) error {
	exerciseID = strings.TrimSpace(exerciseID)
	raw, err := DB.GetExercise(exerciseID)
	if err != nil || raw == "" {
		return ErrInvalidExercise
	}
	var d catalog.Doc
	if json.Unmarshal([]byte(raw), &d) != nil {
		return ErrInvalidExercise
	}
	d.ID = exerciseID
	d.Poses = catalog.ClipPoses(poses)
	return DB.SetExercise(exerciseID, d.Compact())
}

// ExerciseFrames returns 1 or 2 image URLs for the player. Pose catalog stills
// win when present; otherwise the legacy built-in PNG.
func ExerciseFrames(cat catalog.Catalog, id, sex string) []string {
	if sex != "f" {
		sex = "m"
	}
	var out []string
	for _, pid := range cat.PoseIDs(id) {
		if DB != nil && DB.HasPose(pid) {
			out = append(out, "/poses/"+url.PathEscape(pid)+"?sex="+url.QueryEscape(sex))
		}
	}
	if len(out) == 0 {
		if figure := cat.Figure(id, sex); figure != "" {
			return []string{figure}
		}
	}
	return out
}

// ExerciseDemo resolves optional animation metadata, retaining stills as fallback.
type ExerciseDemo struct {
	Src      string `json:"src"`
	OtherSrc string `json:"otherSrc,omitempty"`
	Frames   int    `json:"frames"`
	Kind     string `json:"kind"`
}

func ExerciseAnimation(cat catalog.Catalog, id, sex string) *ExerciseDemo {
	d, ok := cat.Doc(id)
	if !ok || d.Animation == nil || !catalog.ValidAnimation(d.Animation) || DB == nil || !DB.HasPose(d.Animation.Pose) {
		return nil
	}
	if sex != "f" {
		sex = "m"
	}
	otherSrc := ""
	if d.Animation.OtherPose != "" {
		if !DB.HasPose(d.Animation.OtherPose) {
			return nil
		}
		otherSrc = "/poses/" + url.PathEscape(d.Animation.OtherPose) + "?sex=" + url.QueryEscape(sex)
	}
	return &ExerciseDemo{
		Src:      "/poses/" + url.PathEscape(d.Animation.Pose) + "?sex=" + url.QueryEscape(sex),
		OtherSrc: otherSrc,
		Frames:   d.Animation.Frames, Kind: d.Animation.Kind,
	}
}

// SnapshotAll copies leftover named built-in plans into personal files for every
// existing account. Idempotent; used at process start and by MCP.
func SnapshotAll() {
	if DB == nil {
		return
	}
	accs, err := DB.ListAccounts()
	if err != nil {
		return
	}
	for _, a := range accs {
		SnapshotLegacy(a.ID)
	}
}

// HasCustomProgram reports whether the user saved a custom program.
func HasCustomProgram(userID int64) bool {
	j, err := DB.GetProgramJSON(userID)
	return err == nil && j != ""
}

// SnapshotLegacy copies a previously named built-in plan (Sergey / Vlad / Ilya)
// into the user's personal file. Registered accounts still on the old default
// (`builtin`) get the Sergey cycle. Guests and new demo accounts are left on
// the live demo. Cycle position, streaks and activity are never touched.
func SnapshotLegacy(userID int64) {
	if DB == nil || HasCustomProgram(userID) {
		return
	}
	k, _ := DB.ActiveProgram(userID)
	u, err := DB.UserByID(userID)
	if err != nil {
		return
	}
	legacyKey := k
	switch k {
	case "ilya", "vlad", "sergey":
		// keep
	case "builtin", "":
		if u.IsGuest {
			return
		}
		legacyKey = "sergey"
	default:
		return
	}
	r, ok := content.LegacyNamed(legacyKey)
	if !ok {
		return
	}
	_ = DB.SetProgramJSON(userID, r.Marshal())
	_ = DB.SetActiveProgram(userID, "custom")
}

// SetUserProgram stores a personal cycle from JSON and makes it active.
// It does not reset cycle position, streaks, or activity.
func SetUserProgram(userID int64, raw string) error {
	r, err := content.ParseResolved(strings.TrimSpace(raw))
	if err != nil || len(r.Days) == 0 {
		return ErrInvalidProgram
	}
	if strings.TrimSpace(r.Name) == "" {
		r.Name = "My program"
	}
	if err := DB.SetProgramJSON(userID, r.Marshal()); err != nil {
		return err
	}
	return DB.SetActiveProgram(userID, "custom")
}

// ReplaceUserProgram swaps in a different personal cycle from JSON (a course,
// an imported file, an agent upload). Unlike SetUserProgram, which saves edits
// to the same program, it first keeps the program being replaced whenever that
// program is the user's own work, so it can be brought back.
// It does not reset cycle position, streaks, or activity.
func ReplaceUserProgram(userID int64, raw string) error {
	r, err := content.ParseResolved(strings.TrimSpace(raw))
	if err != nil || len(r.Days) == 0 {
		return ErrInvalidProgram
	}
	if strings.TrimSpace(r.Name) == "" {
		r.Name = "My program"
	}
	return DB.ReplaceProgram(userID, r.Marshal(), "custom", ownProgramName(userID))
}

// ownProgramName returns the name of the account's personal program when it is
// worth keeping: the user (or an agent) made or edited it. It is "" when there
// is no personal program or it is an untouched copy of the demo or of a course,
// which can always be started again.
func ownProgramName(userID int64) string {
	raw, err := DB.GetProgramJSON(userID)
	if err != nil || raw == "" {
		return ""
	}
	r, err := content.ParseResolved(raw)
	if err != nil {
		return "My program" // unreadable here, but still the user's data
	}
	name := CleanProgramName(r.Name)
	if name == "" {
		name = "My program"
	}
	r.Name = ""
	body := r.Marshal()
	readyMade := []content.Resolved{content.ResolveDemo()}
	for _, c := range content.Series30() {
		readyMade = append(readyMade, c.Resolve())
	}
	for _, ready := range readyMade {
		ready.Name = ""
		if ready.Marshal() == body {
			return ""
		}
	}
	return name
}

// KeepsCurrentProgram reports whether replacing the account's program now
// would put the current one into the saved list.
func KeepsCurrentProgram(userID int64) bool { return ownProgramName(userID) != "" }

// SavedPrograms lists programs this account replaced and can bring back.
func SavedPrograms(userID int64) []db.SavedProgram {
	list, err := DB.SavedPrograms(userID)
	if err != nil {
		return nil
	}
	return list
}

// RestoreSavedProgram brings a saved program back as the account's program.
// The program it replaces is kept in turn when it is the user's own work.
func RestoreSavedProgram(userID, savedID int64) error {
	return DB.RestoreSavedProgram(userID, savedID, ownProgramName(userID))
}

// ProgramNameMax is the longest program name, in characters.
const ProgramNameMax = 40

// CleanProgramName trims a user-typed program name and clips it on whole
// characters. An empty result means "no name given".
func CleanProgramName(name string) string {
	name = strings.Join(strings.Fields(name), " ")
	if runes := []rune(name); len(runes) > ProgramNameMax {
		name = strings.TrimSpace(string(runes[:ProgramNameMax]))
	}
	return name
}

// ProgramTitle is the name shown for the account's active program: the
// personal file's own name, or the localized demo label.
func ProgramTitle(userID int64, lang string) string {
	if HasCustomProgram(userID) {
		if name := strings.TrimSpace(UserProgram(userID).Name); name != "" {
			return name
		}
	}
	return i18n.T(i18n.Lang(lang), "program.demo")
}

// RenameProgram gives the account's program a name of the user's choosing.
// Exercises, doses, cycle position and streak are untouched. An account still
// on the demo gets its own copy of the demo under that name.
func RenameProgram(userID int64, name string) error {
	name = CleanProgramName(name)
	if name == "" {
		return ErrEmptyProgramName
	}
	r := UserProgram(userID)
	r.Name = name
	return SetUserProgram(userID, r.Marshal())
}

// UseCourse makes a "… in 30 days" course this account's personal program,
// titled in the user's language. From then on it is an ordinary personal file
// the user may edit. Like any program change, it keeps the cycle position,
// streak and activity calendar. The program it replaces is kept in the saved
// list when it is the user's own work.
func UseCourse(userID int64, key, lang string) error {
	c, ok := content.CourseByKey(key)
	if !ok {
		return ErrUnknownCourse
	}
	r := c.Resolve()
	r.Name = c.TitleIn(lang)
	return ReplaceUserProgram(userID, r.Marshal())
}

// RestoreDemo clears the personal file so the user runs the public demo.
// Progress is left as-is, and the user's own program is kept in the saved list.
func RestoreDemo(userID int64) error {
	return DB.ReplaceProgram(userID, "", content.DefaultStandardKey(), ownProgramName(userID))
}

// ActiveProgramKey returns the user's selected program key: "custom" when a saved
// custom program is active, otherwise a valid standard program key (the default
// when the stored value is empty or a legacy/unknown key).
func ActiveProgramKey(userID int64) string {
	SnapshotLegacy(userID)
	k, _ := DB.ActiveProgram(userID)
	if k == "custom" && HasCustomProgram(userID) {
		return "custom"
	}
	if _, ok := content.StandardByKey(k); ok {
		return k
	}
	return content.DefaultStandardKey()
}

// UserProgram returns the user's active program (a standard set or custom).
func UserProgram(userID int64) content.Resolved {
	return ProgramByKey(userID, ActiveProgramKey(userID))
}

// UserWorkout builds a day's workout from the active program (built-in or custom),
// scaled by the user's universal difficulty level, with the user's per-phase rests.
func UserWorkout(userID int64, day int) content.Workout {
	w, m, c := DB.Rests(userID)
	rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
	return UserProgram(userID).Workout(day, DB.GetLevel(userID), rests)
}

// UserWorkoutBase builds the day's workout at level 0 (base, unscaled values) and
// returns the user's difficulty level separately, so the client can scale main
// exercises on the fly (the in-workout difficulty nudge).
func UserWorkoutBase(userID int64, day int) (content.Workout, int) {
	w, m, c := DB.Rests(userID)
	rests := content.Rests{Warmup: w, Main: m, Cooldown: c}
	return UserProgram(userID).Workout(day, 0, rests), DB.GetLevel(userID)
}
