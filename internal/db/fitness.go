package db

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

// FitnessProfile is the stored "About you" answers.
type FitnessProfile struct {
	Child    bool
	Age      int
	Height   int // cm
	Weight   int // kg
	Concerns string
}

// FitnessTest is one stored self-test.
type FitnessTest struct {
	ID                                 int64
	PushUps, Squats, PlankSec, PullUps int
	Score                              int
	TakenAt                            time.Time
}

// FitnessProfile returns the user's answers; ok is false when there are none.
func (d *DB) FitnessProfile(userID int64) (p FitnessProfile, ok bool, err error) {
	var child int
	err = d.sql.QueryRow(`SELECT child, age, height_cm, weight_kg, concerns FROM fitness_profiles WHERE user_id = ?`, userID).
		Scan(&child, &p.Age, &p.Height, &p.Weight, &p.Concerns)
	if errors.Is(err, sql.ErrNoRows) {
		return FitnessProfile{}, false, nil
	}
	if err != nil {
		return FitnessProfile{}, false, err
	}
	p.Child = child != 0
	return p, true, nil
}

// SaveFitnessProfile stores the user's answers, replacing earlier ones.
func (d *DB) SaveFitnessProfile(userID int64, p FitnessProfile) error {
	child := 0
	if p.Child {
		child = 1
	}
	_, err := d.sql.Exec(`INSERT INTO fitness_profiles (user_id, child, age, height_cm, weight_kg, concerns, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(user_id) DO UPDATE SET child = excluded.child, age = excluded.age, height_cm = excluded.height_cm,
			weight_kg = excluded.weight_kg, concerns = excluded.concerns, updated_at = excluded.updated_at`,
		userID, child, p.Age, p.Height, p.Weight, strings.TrimSpace(p.Concerns), time.Now().UTC().Format(time.RFC3339))
	return err
}

// AddFitnessTest stores a self-test and returns its id.
func (d *DB) AddFitnessTest(userID int64, t FitnessTest) (int64, error) {
	res, err := d.sql.Exec(`INSERT INTO fitness_tests (user_id, push_ups, squats, plank_sec, pull_ups, score, taken_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		userID, t.PushUps, t.Squats, t.PlankSec, t.PullUps, t.Score, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// FitnessTests returns the user's latest self-tests, newest first.
func (d *DB) FitnessTests(userID int64, limit int) ([]FitnessTest, error) {
	rows, err := d.sql.Query(`SELECT id, push_ups, squats, plank_sec, pull_ups, score, taken_at FROM fitness_tests
		WHERE user_id = ? ORDER BY id DESC LIMIT ?`, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []FitnessTest
	for rows.Next() {
		var t FitnessTest
		var at string
		if err := rows.Scan(&t.ID, &t.PushUps, &t.Squats, &t.PlankSec, &t.PullUps, &t.Score, &at); err != nil {
			return nil, err
		}
		t.TakenAt, _ = time.Parse(time.RFC3339, at)
		out = append(out, t)
	}
	return out, rows.Err()
}
