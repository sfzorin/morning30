package app

import (
	"encoding/json"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"

	"morning30/internal/db"
	"morning30/internal/fitness"
)

// The fitness profile: "About you" answers, self-tests and the program they
// recommend. See internal/fitness for the scoring.

// FitnessProfile returns what the user told about themselves, with the figure
// setting as the sex; ok is false when they never filled it in.
func FitnessProfile(userID int64) (fitness.Profile, bool) {
	stored, ok, err := DB.FitnessProfile(userID)
	if err != nil {
		log.Printf("fitness profile %d: %v", userID, err)
	}
	sex := DB.GetSex(userID)
	return fitness.Profile{
		Child: stored.Child, Sex: sex, Age: stored.Age, Height: stored.Height, Weight: stored.Weight,
		Concerns: fitness.ParseConcerns(stored.Concerns),
	}, ok
}

// SaveFitnessProfile stores the answers and sets the figure to match them.
func SaveFitnessProfile(userID int64, p fitness.Profile) error {
	clamp := func(v, lo, hi int) int {
		if v < lo || v > hi {
			return 0
		}
		return v
	}
	if err := DB.SaveFitnessProfile(userID, db.FitnessProfile{
		Child: p.Child, Age: clamp(p.Age, 3, 110), Height: clamp(p.Height, 80, 250), Weight: clamp(p.Weight, 15, 300),
		Concerns: strings.Join(fitness.ParseConcerns(strings.Join(p.Concerns, ",")), ","),
	}); err != nil {
		return err
	}
	return DB.SetSex(userID, FigureFor(p.Sex, p.Child))
}

// FitnessScore is a stored test as the screens show it.
type FitnessScore struct {
	Index   int
	Level   string // low | fair | good | great
	Results fitness.Results
}

// LatestFitness returns the newest test and the one before it, if any.
func LatestFitness(userID int64) (latest, previous FitnessScore, hasLatest, hasPrevious bool) {
	tests, err := DB.FitnessTests(userID, 2)
	if err != nil || len(tests) == 0 {
		return
	}
	score := func(t db.FitnessTest) FitnessScore {
		return FitnessScore{Index: t.Score, Level: fitness.Level(t.Score),
			Results: fitness.Results{PushUps: t.PushUps, Squats: t.Squats, PlankSec: t.PlankSec, PullUps: t.PullUps}}
	}
	latest, hasLatest = score(tests[0]), true
	if len(tests) > 1 {
		previous, hasPrevious = score(tests[1]), true
	}
	return
}

// RecordFitnessTest scores and stores a self-test.
func RecordFitnessTest(userID int64, r fitness.Results) (int, bool, error) {
	r = r.Clamp()
	p, _ := FitnessProfile(userID)
	index, ok := fitness.Index(p, r)
	if !ok {
		return 0, false, nil
	}
	_, err := DB.AddFitnessTest(userID, db.FitnessTest{PushUps: r.PushUps, Squats: r.Squats, PlankSec: r.PlankSec, PullUps: r.PullUps, Score: index})
	return index, true, err
}

// FitnessPlan is the program the profile and the latest test recommend.
func FitnessPlan(userID int64) fitness.Plan {
	p, _ := FitnessProfile(userID)
	latest, _, has, _ := LatestFitness(userID)
	return fitness.Recommend(p, latest.Index, has)
}

// ApplyFitnessPlan switches to the recommended program and sets its starting
// difficulty and rounds. The figure keeps what the person chose.
func ApplyFitnessPlan(userID int64, lang string) error {
	plan := FitnessPlan(userID)
	sex := DB.GetSex(userID)
	if err := SelectReady(userID, plan.Program, lang); err != nil {
		return err
	}
	if err := DB.SetLevel(userID, plan.Level); err != nil {
		return err
	}
	if err := DB.SetRounds(userID, plan.Rounds); err != nil {
		return err
	}
	return DB.SetSex(userID, sex)
}

// FitnessTestHTTP stores the results the self-test page sends:
// POST /api/fitness/test {"pushUps":12,"squats":34,"plankSec":72,"pullUps":-1}.
func FitnessTestHTTP(w http.ResponseWriter, r *http.Request) {
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
	if mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type")); err != nil || mediaType != "application/json" {
		fail(http.StatusUnsupportedMediaType, "json_required")
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		fail(http.StatusUnauthorized, "sign_in_required")
		return
	}
	u, ok, err := DB.SessionUser(cookie.Value)
	if err != nil || !ok {
		fail(http.StatusUnauthorized, "sign_in_required")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1024)
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var res fitness.Results
	if err := dec.Decode(&res); err != nil || dec.Decode(new(any)) != io.EOF {
		fail(http.StatusBadRequest, "invalid_results")
		return
	}
	index, done, err := RecordFitnessTest(u.ID, res)
	if err != nil {
		log.Printf("fitness test %d: %v", u.ID, err)
		fail(http.StatusInternalServerError, "save_failed")
		return
	}
	if !done {
		fail(http.StatusBadRequest, "no_results")
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]int{"index": index})
}

// NeedsFitnessProfile reports whether a newcomer should see "About you" first:
// no answers yet and no workout done.
func NeedsFitnessProfile(userID int64) bool {
	if _, ok, err := DB.FitnessProfile(userID); err != nil || ok {
		return false
	}
	pos, _ := DB.Position(userID)
	return pos == 0
}
