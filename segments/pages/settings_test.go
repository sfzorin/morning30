package pages

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/doors-dev/doors"
	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/db"
)

func TestEditorPreservesExercisesWithoutPrivateRestsOrEmptyPhases(t *testing.T) {
	p, err := content.ParseResolved(`{"fixedDifficulty":true,"selfManagedSides":true,"warmup":[{"id":"w","value":45}],"cooldown":[{"id":"c","value":20}],"days":[{"rounds":1,"warmup":[],"cooldown":[],"items":[{"id":"a","value":6,"rest":0,"sideRest":10,"repSeconds":4}]}]}`)
	if err != nil {
		t.Fatal(err)
	}
	editProgramList(&p, 1, "day", func(items []content.RItem) []content.RItem { items[0].Value = 7; return items })
	editProgramList(&p, 1, "day-warmup", func(items []content.RItem) []content.RItem {
		return append(items, content.RItem{ID: "local", Value: 20})
	})
	editProgramList(&p, 1, "day-cooldown", func(items []content.RItem) []content.RItem { return nil })
	encoded := p.Marshal()
	if strings.Contains(encoded, `"rest"`) || strings.Contains(encoded, `"sideRest"`) {
		t.Fatal("editor exported private exercise pauses")
	}
	q, err := content.ParseResolved(encoded)
	if err != nil {
		t.Fatal(err)
	}
	it := q.Days[0].Items[0]
	if !q.FixedDifficulty || !q.SelfManagedSides || it.Value != 7 || it.RepSeconds != 4 {
		t.Fatalf("editor lost exercise data: %+v", it)
	}
	if len(q.Warmup) != 1 || q.Warmup[0].ID != "w" || len(q.WarmupFor(0)) != 2 || q.WarmupFor(0)[1].ID != "local" {
		t.Fatal("day phase edits changed shared phases")
	}
	if len(q.CooldownFor(0)) != 1 || q.CooldownFor(0)[0].ID != "c" {
		t.Fatal("editor removed the mandatory cooldown")
	}
	if q.WarmupRounds != 1 || q.Days[0].Rounds != 2 {
		t.Fatal("editor must retain one warmup and two main circuits")
	}
}

func TestEditorCannotDeleteTheLastExerciseOfAnyPhase(t *testing.T) {
	for _, target := range []string{"warmup", "day-warmup", "day", "cooldown", "day-cooldown"} {
		t.Run(target, func(t *testing.T) {
			p, err := content.ParseResolved(`{"warmup":[{"id":"w","value":45}],"cooldown":[{"id":"c","value":20}],"days":[{"items":[{"id":"a","value":6}]}]}`)
			if err != nil {
				t.Fatal(err)
			}
			editProgramList(&p, 1, target, func(items []content.RItem) []content.RItem {
				return append(append([]content.RItem{}, items...), items[0])
			})
			removeFirst := func(items []content.RItem) []content.RItem { return items[1:] }
			editProgramList(&p, 1, target, removeFirst)
			before := p.Marshal()
			editProgramList(&p, 1, target, removeFirst)
			if p.Marshal() != before {
				t.Fatal("deleting the last exercise must leave the phase intact")
			}
		})
	}
}

// The settings handler uses only Data; the embedded request keeps the fixture
// focused on the actual submitted form and its persisted/reactive result.
type settingsRequest struct {
	doors.RequestForm[settingsForm]
	form settingsForm
}

func (r settingsRequest) Data() settingsForm { return r.form }

func settingsTestPage(t *testing.T) (*db.DB, settingsPage, doors.Source[auth.Session]) {
	t.Helper()
	database, err := db.Open(filepath.Join(t.TempDir(), "settings.db"))
	if err != nil {
		t.Fatal(err)
	}
	previous := app.DB
	app.DB = database
	t.Cleanup(func() { app.DB = previous; database.Close() })
	u, err := database.CreateUser("settings@example.com", "hash", "ru", "Test")
	if err != nil {
		t.Fatal(err)
	}
	sess := auth.Session{UserID: u.ID, Name: u.Name, Lang: u.Lang, Rest: u.Rest, VoiceMode: u.VoiceMode, Voice: u.Voice}
	source := doors.NewSource(sess)
	return database, settingsPage{sess: sess, auth: source}, source
}

func TestSettingsSavePersistsAllThreeRestsAndWorkoutUsesThem(t *testing.T) {
	database, page, session := settingsTestPage(t)
	status := doors.NewSource("")
	for _, rests := range []content.Rests{{Warmup: 0, Main: 5, Cooldown: 10}, {Warmup: 30, Main: 35, Cooldown: 0}, {Warmup: -5, Main: 95, Cooldown: 80}} {
		page.save(status)(context.Background(), settingsRequest{form: settingsForm{
			Name: "Updated", RestWarmup: rests.Warmup, Rest: rests.Main, RestCooldown: rests.Cooldown, VoiceMode: "min", Sex: "f",
		}})
		want := rests.Normalize()
		warmup, main, cooldown := database.Rests(page.sess.UserID)
		if (content.Rests{Warmup: warmup, Main: main, Cooldown: cooldown}) != want {
			t.Fatalf("saved rests = %d/%d/%d, want %+v", warmup, main, cooldown, want)
		}
		if status.Get() != "settings.saved" || session.Get().Rest != want.Main || session.Get().Name != "Updated" || session.Get().VoiceMode != "min" {
			t.Fatalf("successful save did not update the visible session: status %q session %+v", status.Get(), session.Get())
		}
		workout := app.UserWorkout(page.sess.UserID, 1)
		for i, item := range workout.Items {
			pause := 0
			if i+1 < len(workout.Items) {
				switch workout.Items[i+1].Slot {
				case content.Warmup:
					pause = want.Warmup
				case content.Main:
					pause = want.Main
				case content.Cooldown:
					pause = want.Cooldown
				}
			}
			if item.Rest != pause {
				t.Fatalf("exercise %d rest=%d want %d", i, item.Rest, pause)
			}
		}
	}
}

func TestSettingsSaveFailureDoesNotClaimSuccessOrChangeSession(t *testing.T) {
	database, page, session := settingsTestPage(t)
	before := session.Get()
	if err := database.Close(); err != nil {
		t.Fatal(err)
	}
	status := doors.NewSource("settings.saved")
	page.save(status)(context.Background(), settingsRequest{form: settingsForm{Name: "Changed", Rest: 35, VoiceMode: "off"}})
	if status.Get() != "settings.save_failed" {
		t.Fatalf("failed save showed %q", status.Get())
	}
	if session.Get() != before {
		t.Fatal("failed save changed the reactive session")
	}
}

type importedProgramForm struct {
	doors.ParsedForm
	raw string
}

func (f importedProgramForm) FormValue(key string) string {
	if key == "progjson" {
		return f.raw
	}
	return ""
}

type importProgramRequest struct {
	doors.RequestRawForm
	raw      string
	parseErr error
	reloaded bool
}

func (r *importProgramRequest) ParseForm(int) (doors.ParsedForm, error) {
	return importedProgramForm{raw: r.raw}, r.parseErr
}

func (r *importProgramRequest) After(doors.Actions) error {
	r.reloaded = true
	return nil
}

func TestInvalidProgramImportShowsErrorAndPreservesSavedProgramAndProgress(t *testing.T) {
	database, page, _ := settingsTestPage(t)
	if err := app.SetUserProgram(page.sess.UserID, content.ResolveDemo().Marshal()); err != nil {
		t.Fatal(err)
	}
	if err := database.AdvancePosition(page.sess.UserID); err != nil {
		t.Fatal(err)
	}
	if err := database.RecordActivity(page.sess.UserID, "2026-10-28"); err != nil {
		t.Fatal(err)
	}
	before, err := database.GetProgramJSON(page.sess.UserID)
	if err != nil {
		t.Fatal(err)
	}
	for _, invalid := range []importProgramRequest{
		{raw: "{"},
		{raw: "   "},
		{parseErr: errors.New("form too large")},
		{raw: `{"days":[{"items":[{"id":"P01","unit":"reps","value":6,"round":1},{"id":"P01","unit":"reps","value":8,"round":2}]}]}`},
	} {
		status := doors.NewSource("")
		page.importProgram(status)(context.Background(), &invalid)
		if status.Get() != "settings.import_failed" || invalid.reloaded {
			t.Fatalf("invalid import did not retain the form with an error: %q, reload=%v", status.Get(), invalid.reloaded)
		}
		after, err := database.GetProgramJSON(page.sess.UserID)
		if err != nil || after != before {
			t.Fatalf("invalid import overwrote the saved program: %v", err)
		}
		if active, err := database.ActiveProgram(page.sess.UserID); err != nil || active != "custom" {
			t.Fatalf("invalid import changed active program: %q %v", active, err)
		}
		if position, err := database.Position(page.sess.UserID); err != nil || position != 1 {
			t.Fatalf("invalid import changed position: %d %v", position, err)
		}
		if streak, err := database.Streak(page.sess.UserID, "2026-10-28"); err != nil || streak != 1 {
			t.Fatalf("invalid import changed streak: %d %v", streak, err)
		}
	}
}
