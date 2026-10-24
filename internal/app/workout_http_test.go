package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"morning30/internal/db"
)

func setupWorkoutHTTP(t *testing.T) db.WorkoutCompletion {
	t.Helper()
	d := openAppDB(t)
	u, err := d.CreateUser("http@example.com", "hash", "en", "Test")
	if err != nil {
		t.Fatal(err)
	}
	if err := d.CreateSession("valid", u.ID, time.Hour); err != nil {
		t.Fatal(err)
	}
	if err := d.CreateSession("expired", u.ID, -time.Hour); err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC()
	return db.WorkoutCompletion{
		EventID: "eea47236-50ec-4adb-9303-93fd090dc2d7", UserID: u.ID,
		Day: 1, Position: 0, CompletedAt: now.Format(time.RFC3339Nano), CompletedOn: now.Format("2006-01-02"),
	}
}

func workoutRequest(t *testing.T, event db.WorkoutCompletion, customize func(*http.Request)) *httptest.ResponseRecorder {
	t.Helper()
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodPost, "http://morning30.com/api/workouts/complete", strings.NewReader(string(body)))
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	// HTTPS Origin with a plain HTTP backend matches the production proxy setup.
	r.Header.Set("Origin", "https://morning30.com")
	r.AddCookie(&http.Cookie{Name: "session", Value: "valid"})
	if customize != nil {
		customize(r)
	}
	w := httptest.NewRecorder()
	CompleteWorkoutHTTP(w, r)
	return w
}

func TestWorkoutHTTPRetryPreservesOriginalDate(t *testing.T) {
	event := setupWorkoutHTTP(t)
	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	event.CompletedAt = yesterday.Format(time.RFC3339Nano)
	event.CompletedOn = yesterday.Format("2006-01-02")
	for range 2 {
		w := workoutRequest(t, event, nil)
		var result db.WorkoutResult
		if w.Code != http.StatusOK || json.Unmarshal(w.Body.Bytes(), &result) != nil || !result.Saved || result.EventID != event.EventID || result.Position != 1 || result.Streak != 1 {
			t.Fatalf("response = %d %s", w.Code, w.Body.String())
		}
	}
	if has, _ := DB.HasActivity(event.UserID, event.CompletedOn); !has {
		t.Fatal("original date was not credited")
	}
	if has, _ := DB.HasActivity(event.UserID, time.Now().UTC().Format("2006-01-02")); has {
		t.Fatal("retry incorrectly credited today's date")
	}
	event.Day = 2
	if w := workoutRequest(t, event, nil); w.Code != http.StatusConflict {
		t.Fatalf("conflicting ID response = %d %s", w.Code, w.Body.String())
	}
}

func TestWorkoutHTTPRejectsInvalidAuthenticationAndOrigin(t *testing.T) {
	for _, tc := range []struct {
		name   string
		change func(*http.Request)
		want   int
	}{
		{"missing session", func(r *http.Request) { r.Header.Del("Cookie") }, http.StatusUnauthorized},
		{"expired session", func(r *http.Request) { r.Header.Set("Cookie", "session=expired") }, http.StatusUnauthorized},
		{"unknown session", func(r *http.Request) { r.Header.Set("Cookie", "session=unknown") }, http.StatusUnauthorized},
		{"cross origin", func(r *http.Request) { r.Header.Set("Origin", "https://attacker.example") }, http.StatusForbidden},
		{"cross site", func(r *http.Request) { r.Header.Set("Sec-Fetch-Site", "cross-site") }, http.StatusForbidden},
		{"wrong method", func(r *http.Request) { r.Method = http.MethodGet }, http.StatusMethodNotAllowed},
		{"form submission", func(r *http.Request) { r.Header.Set("Content-Type", "application/x-www-form-urlencoded") }, http.StatusUnsupportedMediaType},
	} {
		t.Run(tc.name, func(t *testing.T) {
			event := setupWorkoutHTTP(t)
			w := workoutRequest(t, event, tc.change)
			if w.Code != tc.want || strings.Contains(w.Body.String(), `"saved":true`) {
				t.Fatalf("response = %d %s, want %d", w.Code, w.Body.String(), tc.want)
			}
			if has, _ := DB.HasActivity(event.UserID, event.CompletedOn); has {
				t.Fatal("rejected request credited activity")
			}
		})
	}
	t.Run("account mismatch", func(t *testing.T) {
		event := setupWorkoutHTTP(t)
		other, _ := DB.CreateUser("other@example.com", "hash", "en", "Other")
		event.UserID = other.ID
		if w := workoutRequest(t, event, nil); w.Code != http.StatusForbidden {
			t.Fatalf("account mismatch = %d %s", w.Code, w.Body.String())
		}
		if has, _ := DB.HasActivity(other.ID, event.CompletedOn); has {
			t.Fatal("wrong account received activity")
		}
	})
}

func TestWorkoutHTTPRejectsMalformedOrOversizedBody(t *testing.T) {
	event := setupWorkoutHTTP(t)
	valid, _ := json.Marshal(event)
	for _, body := range []string{
		`{`, string(valid) + `{}`, strings.TrimSuffix(string(valid), "}") + `,"unknown":true}`,
		`{"eventId":"` + strings.Repeat("a", 5000) + `"}`, `null`,
	} {
		r := httptest.NewRequest(http.MethodPost, "http://morning30.com/api/workouts/complete", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		r.AddCookie(&http.Cookie{Name: "session", Value: "valid"})
		w := httptest.NewRecorder()
		CompleteWorkoutHTTP(w, r)
		if w.Code < 400 || strings.Contains(w.Body.String(), `"saved":true`) {
			t.Fatalf("malformed body response = %d %s", w.Code, w.Body.String())
		}
	}
}

func TestValidateCompletionDatesAndBounds(t *testing.T) {
	now := time.Date(2026, 10, 24, 12, 0, 0, 0, time.UTC)
	base := db.WorkoutCompletion{EventID: strings.Repeat("a", 32), UserID: 1, Day: 1, CompletedAt: now.Format(time.RFC3339Nano), CompletedOn: "2026-10-24"}
	for _, tc := range []struct {
		name   string
		change func(*db.WorkoutCompletion)
		valid  bool
	}{
		{"yesterday locally", func(e *db.WorkoutCompletion) { e.CompletedOn = "2026-10-23" }, true},
		{"tomorrow locally", func(e *db.WorkoutCompletion) { e.CompletedOn = "2026-10-25" }, true},
		{"two days away", func(e *db.WorkoutCompletion) { e.CompletedOn = "2026-10-26" }, false},
		{"invalid calendar date", func(e *db.WorkoutCompletion) { e.CompletedOn = "2026-02-30" }, false},
		{"future boundary", func(e *db.WorkoutCompletion) { e.CompletedAt = now.Add(5 * time.Minute).Format(time.RFC3339Nano) }, true},
		{"future", func(e *db.WorkoutCompletion) {
			e.CompletedAt = now.Add(5*time.Minute + time.Second).Format(time.RFC3339Nano)
		}, false},
		{"old boundary", func(e *db.WorkoutCompletion) {
			e.CompletedAt = now.Add(-31 * 24 * time.Hour).Format(time.RFC3339Nano)
			e.CompletedOn = "2026-09-23"
		}, true},
		{"too old", func(e *db.WorkoutCompletion) {
			e.CompletedAt = now.Add(-31*24*time.Hour - time.Second).Format(time.RFC3339Nano)
			e.CompletedOn = "2026-09-23"
		}, false},
		{"not UTC", func(e *db.WorkoutCompletion) { e.CompletedAt = "2026-10-24T13:00:00+01:00" }, false},
		{"short ID", func(e *db.WorkoutCompletion) { e.EventID = "short" }, false},
		{"invalid ID characters", func(e *db.WorkoutCompletion) { e.EventID = strings.Repeat("a", 31) + "/" }, false},
		{"negative position", func(e *db.WorkoutCompletion) { e.Position = -1 }, false},
		{"day zero", func(e *db.WorkoutCompletion) { e.Day = 0 }, false},
		{"day 31", func(e *db.WorkoutCompletion) { e.Day = 31 }, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			event := base
			tc.change(&event)
			if err := validateCompletion(&event, now); (err == nil) != tc.valid {
				t.Fatalf("validation error = %v, want valid %v", err, tc.valid)
			}
		})
	}
}

func TestWorkoutHTTPDatabaseFailureIsNotAcknowledged(t *testing.T) {
	event := setupWorkoutHTTP(t)
	_ = DB.Close()
	w := workoutRequest(t, event, nil)
	if w.Code != http.StatusInternalServerError || strings.Contains(w.Body.String(), `"saved":true`) {
		t.Fatalf("database failure = %d %s", w.Code, w.Body.String())
	}
}
