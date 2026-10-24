package app

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"regexp"
	"time"

	"morning30/internal/db"
)

var completionEventID = regexp.MustCompile(`^[A-Za-z0-9_-]{32,128}$`)
var workoutOriginProtection = http.NewCrossOriginProtection()

// CompleteWorkoutHTTP stays available after a Doors page instance expires.
// Authentication is always checked against the persistent session cookie.
func CompleteWorkoutHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	fail := func(status int, message string) {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		fail(http.StatusMethodNotAllowed, "method_not_allowed")
		return
	}
	if err := workoutOriginProtection.Check(r); err != nil {
		fail(http.StatusForbidden, "cross_origin")
		return
	}
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		fail(http.StatusUnsupportedMediaType, "json_required")
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		fail(http.StatusUnauthorized, "sign_in_required")
		return
	}
	u, ok, err := DB.SessionUser(cookie.Value)
	if err != nil {
		log.Printf("workout completion session: %v", err)
		fail(http.StatusInternalServerError, "save_failed")
		return
	}
	if !ok {
		fail(http.StatusUnauthorized, "sign_in_required")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var event db.WorkoutCompletion
	if err := decoder.Decode(&event); err != nil {
		fail(http.StatusBadRequest, "invalid_completion")
		return
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		fail(http.StatusBadRequest, "invalid_completion")
		return
	}
	if event.UserID != u.ID {
		fail(http.StatusForbidden, "account_mismatch")
		return
	}
	if err := validateCompletion(&event, time.Now()); err != nil {
		fail(http.StatusBadRequest, "invalid_completion")
		return
	}
	result, err := DB.CompleteWorkout(r.Context(), event)
	if errors.Is(err, db.ErrCompletionConflict) {
		fail(http.StatusConflict, "event_conflict")
		return
	}
	if err != nil {
		log.Printf("workout completion save: %v", err)
		fail(http.StatusInternalServerError, "save_failed")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func validateCompletion(event *db.WorkoutCompletion, now time.Time) error {
	invalid := errors.New("invalid completion")
	if !completionEventID.MatchString(event.EventID) || event.Day < 1 || event.Day > 30 || event.Position < 0 || event.Position > math.MaxInt-30 {
		return invalid
	}
	completed, err := time.Parse(time.RFC3339Nano, event.CompletedAt)
	if err != nil {
		return invalid
	}
	_, offset := completed.Zone()
	if offset != 0 || completed.After(now.Add(5*time.Minute)) || completed.Before(now.Add(-31*24*time.Hour)) {
		return invalid
	}
	on, err := time.Parse("2006-01-02", event.CompletedOn)
	if err != nil {
		return invalid
	}
	utcDay := time.Date(completed.Year(), completed.Month(), completed.Day(), 0, 0, 0, 0, time.UTC)
	if on.Before(utcDay.AddDate(0, 0, -1)) || on.After(utcDay.AddDate(0, 0, 1)) {
		return invalid
	}
	event.CompletedAt = completed.UTC().Format(time.RFC3339Nano)
	return nil
}
