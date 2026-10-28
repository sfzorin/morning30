package pages

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"morning30/internal/app"
	"morning30/internal/auth"
	"morning30/internal/content"
	"morning30/internal/db"
)

func TestStartResolvesCurrentCycleAfterScheduledRestart(t *testing.T) {
	database, err := db.Open(filepath.Join(t.TempDir(), "workout.db"))
	if err != nil {
		t.Fatal(err)
	}
	previous := app.DB
	app.DB = database
	t.Cleanup(func() { app.DB = previous; database.Close() })
	u, err := database.CreateUser("restart@example.com", "hash", "ru", "Test")
	if err != nil {
		t.Fatal(err)
	}
	for range 39 {
		if err := database.AdvancePosition(u.ID); err != nil {
			t.Fatal(err)
		}
	}
	page := workoutPage{sess: auth.Session{UserID: u.ID}, day: 0}
	if got := page.clampDay(); got != 10 {
		t.Fatalf("start day before restart = %d", got)
	}
	if err := database.ScheduleCycleRestart(context.Background(), u.ID, time.Now().Add(-time.Minute)); err != nil {
		t.Fatal(err)
	}
	if got := page.clampDay(); got != 1 {
		t.Fatalf("stale home must start new cycle day 1, got %d", got)
	}
	page.day = 7
	if got := page.clampDay(); got != 7 {
		t.Fatalf("manual calendar selection changed: %d", got)
	}
}

func TestRestSettingsDriveTheWholeWorkoutAndPlayerForEveryProgram(t *testing.T) {
	const legacyCustom = `{
		"name":"Custom", "warmupRounds":9,
		"warmup":[{"id":"W01","unit":"seconds","value":5,"slot":"main","rest":90}],
		"days":[{"day":1,"rounds":7,"warmup":[],"cooldown":[],"items":[
			{"id":"P01","unit":"reps","value":6,"slot":"warmup","rest":90},
			{"id":"C07","unit":"seconds","value":10,"slot":"main","rest":90,"sideRest":90}
		]}],
		"cooldown":[{"id":"CD01","unit":"seconds","value":5,"slot":"main","rest":90}]
	}`
	for _, program := range []string{"demo", "sergey", "custom"} {
		t.Run(program, func(t *testing.T) {
			database, err := db.Open(filepath.Join(t.TempDir(), "rest-settings.db"))
			if err != nil {
				t.Fatal(err)
			}
			previous := app.DB
			app.DB = database
			t.Cleanup(func() { app.DB = previous; database.Close() })
			u, err := database.CreateUser("rest@example.com", "hash", "ru", "Test")
			if err != nil {
				t.Fatal(err)
			}
			if err := database.SetActiveProgram(u.ID, program); err != nil {
				t.Fatal(err)
			}
			if program == "custom" {
				// Exercise the existing stored format, without pre-normalizing it.
				if err := database.SetProgramJSON(u.ID, legacyCustom); err != nil {
					t.Fatal(err)
				}
			}
			for _, day := range []string{"2026-10-27", "2026-10-28"} {
				if err := database.RecordActivity(u.ID, day); err != nil {
					t.Fatal(err)
				}
				if err := database.AdvancePosition(u.ID); err != nil {
					t.Fatal(err)
				}
			}
			page := workoutPage{sess: auth.Session{UserID: u.ID, Lang: "ru"}, day: 1}
			var first content.Workout
			var firstKey string
			for setting, rests := range []content.Rests{
				{Warmup: 5, Main: 10, Cooldown: 15},
				{Warmup: 20, Main: 35, Cooldown: 0},
			} {
				if err := database.UpdateSettings(u.ID, "Test", "ru", rests.Warmup, rests.Main, rests.Cooldown, "normal"); err != nil {
					t.Fatal(err)
				}
				wk := app.UserWorkout(u.ID, 1)
				payload := page.playerData()
				if payload.SideChangeSeconds != content.SideChangeSeconds {
					t.Fatal("player and estimate must use the same side-change preparation")
				}
				oldKey := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s\n%d,%d,%d", app.UserProgram(u.ID).Marshal(), rests.Warmup, rests.Main, rests.Cooldown))))
				if payload.ProgramKey == oldKey {
					t.Fatal("old combined timer checkpoints must not be interpreted as split work segments")
				}
				if program == "custom" && wk.EstSec != 92+4*rests.Main+rests.Cooldown {
					t.Fatalf("duration still contains an authored pause: %d", wk.EstSec)
				}
				if len(payload.Items) != len(wk.Items) {
					t.Fatal("player and duration estimate use different sequences")
				}
				var warmup, cooldown int
				rounds := map[int][]string{}
				for i, it := range wk.Items {
					got := payload.Items[i]
					if got.ID != it.ExerciseID || got.Rest != it.Rest || got.Round != it.Round || got.Value != it.Value || got.Slot != string(it.Slot) {
						t.Fatalf("player step %d differs from the workout: %#v / %#v", i, got, it)
					}
					wantRest := 0
					if i < len(wk.Items)-1 {
						switch wk.Items[i+1].Slot {
						case content.Warmup:
							wantRest = rests.Warmup
						case content.Main:
							wantRest = rests.Main
						case content.Cooldown:
							wantRest = rests.Cooldown
						}
					}
					if got.Rest != wantRest {
						t.Fatalf("setting %d, step %d: rest=%d want %d", setting, i, got.Rest, wantRest)
					}
					switch it.Slot {
					case content.Warmup:
						warmup++
						if it.Round != 1 {
							t.Fatal("warmup repeated")
						}
					case content.Cooldown:
						cooldown++
						if it.Round != 1 {
							t.Fatal("cooldown repeated")
						}
					case content.Main:
						rounds[it.Round] = append(rounds[it.Round], it.ExerciseID)
					}
				}
				if warmup == 0 || cooldown == 0 || len(rounds) != 2 || len(rounds[1]) == 0 || !reflect.DeepEqual(rounds[1], rounds[2]) {
					t.Fatalf("missing mandatory phase or full circuit: warmup=%d cooldown=%d rounds=%v", warmup, cooldown, rounds)
				}
				encoded, err := json.Marshal(payload)
				if err != nil || strings.Contains(string(encoded), `"sideRest"`) {
					t.Fatal("player still receives authored side pauses")
				}
				if setting == 0 {
					first, firstKey = wk, payload.ResumeKey
				} else {
					if payload.ResumeKey == firstKey {
						t.Fatal("rest change must invalidate the old workout checkpoint")
					}
					pauseDelta := 0
					for i, it := range wk.Items {
						pauseDelta += it.Rest - first.Items[i].Rest
					}
					if wk.EstSec-first.EstSec != pauseDelta {
						t.Fatalf("duration did not track actual pauses: got delta %d want %d", wk.EstSec-first.EstSec, pauseDelta)
					}
				}
			}
			if pos, err := database.Position(u.ID); err != nil || pos != 2 {
				t.Fatalf("settings changed cycle position: %d, %v", pos, err)
			}
			if streak, err := database.Streak(u.ID, "2026-10-28"); err != nil || streak != 2 {
				t.Fatalf("settings changed streak: %d, %v", streak, err)
			}
			if activity, err := database.ActivityInRange(u.ID, "2026-10-01", "2026-10-31"); err != nil || !reflect.DeepEqual(activity, map[string]bool{"2026-10-27": true, "2026-10-28": true}) {
				t.Fatalf("settings changed activity: %v, %v", activity, err)
			}
		})
	}
}
