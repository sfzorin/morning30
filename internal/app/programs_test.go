package app

import (
	"errors"
	"strings"
	"testing"
	"unicode/utf8"

	"morning30/internal/content"
	"morning30/internal/db"
)

func names(list []SelectedView) []string {
	out := make([]string, len(list))
	for i, s := range list {
		out[i] = s.Name
	}
	return out
}

func viewBySource(t *testing.T, list []SelectedView, source string) SelectedView {
	t.Helper()
	for _, s := range list {
		if s.Source == source {
			return s
		}
	}
	t.Fatalf("no selected program from %q in %v", source, names(list))
	return SelectedView{}
}

func advance(d *db.DB, userID int64, n int) {
	for i := 0; i < n; i++ {
		_ = d.AdvancePosition(userID)
	}
}

// A newcomer runs the morning workout without any selected program; picking
// a course before training adds only that course.
func TestNewcomerPicksCourse(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("new@example.com", "h", "ru", "N")
	if got := SelectedPrograms(u.ID); len(got) != 0 {
		t.Fatalf("newcomer has selected programs: %v", names(got))
	}
	if ProgramTitle(u.ID, "ru") != "Утренняя зарядка" || CurrentSource(u.ID) != content.DemoKey {
		t.Fatal("newcomer is not on the morning workout")
	}
	if err := SelectReady(u.ID, "plank30", "ru"); err != nil {
		t.Fatal(err)
	}
	list := SelectedPrograms(u.ID)
	if len(list) != 1 || !list[0].Current || list[0].Name != "Планка за 30 дней" || list[0].Day != 1 || list[0].Edited {
		t.Fatalf("after first pick: %+v", list)
	}
	if err := SelectReady(u.ID, "nope30", "ru"); !errors.Is(err, ErrUnknownCourse) {
		t.Fatalf("unknown key: %v", err)
	}
}

// Every selected program keeps its own day; the streak calendar is common.
func TestEachProgramKeepsItsOwnDay(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("days@example.com", "h", "ru", "D")
	advance(d, u.ID, 12) // twelve days on the morning workout
	_ = d.RecordActivity(u.ID, "2026-11-25")

	if err := SelectReady(u.ID, "plank30", "ru"); err != nil {
		t.Fatal(err)
	}
	if CurrentDay(u.ID) != 1 {
		t.Fatalf("plank starts on day %d", CurrentDay(u.ID))
	}
	advance(d, u.ID, 4)
	if err := SelectReady(u.ID, content.DemoKey, "ru"); err != nil {
		t.Fatal(err)
	}
	if CurrentDay(u.ID) != 13 {
		t.Fatalf("morning workout came back on day %d, want 13", CurrentDay(u.ID))
	}
	list := SelectedPrograms(u.ID)
	if len(list) != 2 || viewBySource(t, list, "plank30").Day != 5 || !viewBySource(t, list, content.DemoKey).Current {
		t.Fatalf("selected: %+v", list)
	}
	if ok, _ := d.HasActivity(u.ID, "2026-11-25"); !ok {
		t.Fatal("activity lost")
	}
}

// Edits belong to the program they were made in and survive switching.
func TestEditsSurviveSwitching(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("edit@example.com", "h", "en", "E")
	if err := SelectReady(u.ID, "abs30", "en"); err != nil {
		t.Fatal(err)
	}
	if err := AddToCurrent(u.ID, "C07", content.Main, content.AllDays(), 25); err != nil {
		t.Fatal(err)
	}
	if err := RenameProgram(u.ID, "  My   abs "); err != nil {
		t.Fatal(err)
	}
	edited := UserProgram(u.ID).Marshal()
	if err := SelectReady(u.ID, "plank30", "en"); err != nil {
		t.Fatal(err)
	}
	// Picking abs again from the catalog brings the user's version back.
	if err := SelectReady(u.ID, "abs30", "en"); err != nil {
		t.Fatal(err)
	}
	if UserProgram(u.ID).Marshal() != edited || ProgramTitle(u.ID, "en") != "My abs" {
		t.Fatal("edited abs program was not kept")
	}
	abs := viewBySource(t, SelectedPrograms(u.ID), "abs30")
	if !abs.Edited || abs.Name != "My abs" {
		t.Fatalf("abs view: %+v", abs)
	}
	if viewBySource(t, SelectedPrograms(u.ID), "plank30").Edited {
		t.Fatal("untouched plank marked edited")
	}
}

// Removing an edited program needs consent; removing the current one moves
// to another selected program, or to the morning workout when none is left.
func TestUnselect(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("rm@example.com", "h", "en", "R")
	_ = SelectReady(u.ID, "plank30", "en")
	_ = SelectReady(u.ID, "abs30", "en")
	_ = RenameProgram(u.ID, "Mine")
	abs := viewBySource(t, SelectedPrograms(u.ID), "abs30")
	if err := Unselect(u.ID, abs.ID, false); !errors.Is(err, ErrHasEdits) {
		t.Fatalf("edited program removed without consent: %v", err)
	}
	if err := Unselect(u.ID, abs.ID, true); err != nil {
		t.Fatal(err)
	}
	list := SelectedPrograms(u.ID)
	if len(list) != 1 || list[0].Source != "plank30" || !list[0].Current {
		t.Fatalf("after removing current: %+v", list)
	}
	if err := UnstarReady(u.ID, "plank30", false); err != nil {
		t.Fatal(err)
	}
	if len(SelectedPrograms(u.ID)) != 0 || HasCustomProgram(u.ID) || CurrentSource(u.ID) != content.DemoKey {
		t.Fatal("last program removed but not back on the morning workout")
	}
	other, _ := d.CreateUser("other@example.com", "h", "en", "O")
	_ = SelectReady(other.ID, "slim30", "en")
	theirs := SelectedPrograms(other.ID)[0]
	if err := Unselect(u.ID, theirs.ID, true); !errors.Is(err, db.ErrSelectedNotFound) {
		t.Fatalf("removed another account's program: %v", err)
	}
}

// Starring adds without switching; the star is idempotent.
func TestStarReady(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("star@example.com", "h", "de", "S")
	id1, err := StarReady(u.ID, "posture30", "de")
	if err != nil {
		t.Fatal(err)
	}
	id2, _ := StarReady(u.ID, "posture30", "de")
	if id1 != id2 || len(SelectedPrograms(u.ID)) != 1 {
		t.Fatal("star added twice")
	}
	if CurrentSource(u.ID) != content.DemoKey || SelectedPrograms(u.ID)[0].Current {
		t.Fatal("starring switched the program")
	}
	if err := SelectByID(u.ID, id1); err != nil || CurrentSource(u.ID) != "posture30" {
		t.Fatalf("select by id: %v", err)
	}
}

// Existing accounts move into the new model without losing anything: the
// personal program keeps its day, earlier "saved programs" become selected.
func TestMigrationFromPersonalProgramAndSavedList(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("old@example.com", "h", "ru", "O")
	mine := content.ResolveDemo()
	mine.Name = "Моя зарядка"
	mine.Days[0].Items[0].Value = 17
	_ = d.SetProgramJSON(u.ID, mine.Marshal())
	_ = d.SetActiveProgram(u.ID, "custom")
	advance(d, u.ID, 9)
	plank, _ := content.CourseByKey("plank30")
	if err := d.PutSavedProgramForTest(u.ID, "Планка за 30 дней", plank.Resolve().Marshal()); err != nil {
		t.Fatal(err)
	}

	list := SelectedPrograms(u.ID)
	if len(list) != 2 {
		t.Fatalf("migrated: %+v", list)
	}
	own := list[0]
	if !own.Current || own.Name != "Моя зарядка" || own.Day != 10 || own.Source != "" || !own.Edited {
		t.Fatalf("own program: %+v", own)
	}
	if kept := list[1]; kept.Source != "plank30" || kept.Edited || kept.Day != 1 {
		t.Fatalf("kept program: %+v", kept)
	}
	if UserProgram(u.ID).Marshal() != mine.Marshal() || CurrentDay(u.ID) != 10 {
		t.Fatal("migration changed the running program")
	}
	again := SelectedPrograms(u.ID) // idempotent
	if len(again) != 2 {
		t.Fatalf("migration repeated: %+v", again)
	}
}

func TestImportAddsNewProgramAndSetUpdatesCurrent(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("imp@example.com", "h", "en", "I")
	advance(d, u.ID, 3)
	plan := content.ResolveDemo()
	plan.Name = "Imported"
	plan.Days[0].Items[0].Value = 23
	if err := ImportProgram(u.ID, plan.Marshal()); err != nil {
		t.Fatal(err)
	}
	if CurrentDay(u.ID) != 1 || ProgramTitle(u.ID, "en") != "Imported" {
		t.Fatal("import did not start the new program at day 1")
	}
	if viewBySource(t, SelectedPrograms(u.ID), content.DemoKey).Day != 4 {
		t.Fatal("morning workout day lost on import")
	}
	if err := ImportProgram(u.ID, "{broken"); !errors.Is(err, ErrInvalidProgram) {
		t.Fatalf("broken import: %v", err)
	}
	// An agent's update edits the current program in place and keeps its day.
	advance(d, u.ID, 2)
	plan.Days[1].Items[0].Value = 31
	if err := SetUserProgram(u.ID, plan.Marshal()); err != nil {
		t.Fatal(err)
	}
	if CurrentDay(u.ID) != 3 || UserProgram(u.ID).Days[1].Items[0].Value != 31 || len(SelectedPrograms(u.ID)) != 2 {
		t.Fatal("update did not edit in place")
	}
}

func TestRenameClipsAndRejectsEmpty(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("name@example.com", "h", "ru", "N")
	if err := RenameProgram(u.ID, "   "); !errors.Is(err, ErrEmptyProgramName) || HasCustomProgram(u.ID) {
		t.Fatalf("empty rename: %v", err)
	}
	long := strings.Repeat("Планка ", 10)
	if err := RenameProgram(u.ID, long); err != nil {
		t.Fatal(err)
	}
	got := ProgramTitle(u.ID, "ru")
	if n := []rune(got); len(n) > ProgramNameMax || !utf8.ValidString(got) || !strings.HasPrefix(long, got) {
		t.Fatalf("clipped name = %q", got)
	}
	// Renaming the morning workout made it a selected program of its own.
	if list := SelectedPrograms(u.ID); len(list) != 1 || list[0].Source != content.DemoKey || !list[0].Edited {
		t.Fatalf("renamed demo: %+v", list)
	}
}

func TestEditCurrentReportsProblems(t *testing.T) {
	d := openAppDB(t)
	u, _ := d.CreateUser("ops@example.com", "h", "en", "O")
	if err := AddToCurrent(u.ID, "NOPE", content.Main, content.AllDays(), 0); !errors.Is(err, ErrUnknownExercise) {
		t.Fatalf("unknown exercise: %v", err)
	}
	if n, err := ReplaceInCurrent(u.ID, "C01", "P08", content.OneDay(1)); err != nil || n != 1 {
		t.Fatalf("replace: %d %v", n, err)
	}
	if n, err := RemoveFromCurrent(u.ID, "P08", content.Main, content.OneDay(1)); err != nil || n == 0 {
		t.Fatalf("remove: %d %v", n, err)
	}
	if _, err := RemoveFromCurrent(u.ID, "ZZ1", "", content.AllDays()); !errors.Is(err, content.ErrNotInProgram) {
		t.Fatalf("remove missing: %v", err)
	}
}

func TestFigures(t *testing.T) {
	for _, c := range []struct {
		cur  string
		kids bool
		want string
	}{{"m", true, "b"}, {"f", true, "g"}, {"b", true, "b"}, {"g", false, "f"}, {"b", false, "m"}, {"f", false, "f"}, {"x", false, "m"}} {
		if got := FigureFor(c.cur, c.kids); got != c.want {
			t.Errorf("FigureFor(%q, %v) = %q, want %q", c.cur, c.kids, got, c.want)
		}
	}
	d := openAppDB(t)
	u, _ := d.CreateUser("fig@example.com", "h", "en", "F")
	for _, f := range Figures {
		_ = d.SetSex(u.ID, f)
		if d.GetSex(u.ID) != f {
			t.Fatalf("figure %s not stored", f)
		}
	}
}
