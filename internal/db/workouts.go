package db

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"time"
)

var ErrCompletionConflict = errors.New("completion event already has different data")

// WorkoutCompletion is an immutable receipt. Position is captured when the
// workout starts, so two sessions started at the same position cannot advance twice.
type WorkoutCompletion struct {
	EventID     string `json:"eventId"`
	UserID      int64  `json:"userId"`
	Day         int    `json:"day"`
	Position    int    `json:"position"`
	CompletedAt string `json:"completedAt"`
	CompletedOn string `json:"completedOn"`
}

type WorkoutResult struct {
	Saved    bool   `json:"saved"`
	EventID  string `json:"eventId"`
	Streak   int    `json:"streak"`
	Position int    `json:"position"`
}

// CompleteWorkout records a completion, its activity date, and any cycle advance
// in one transaction. An identical retry returns success without changing progress.
func (d *DB) CompleteWorkout(ctx context.Context, event WorkoutCompletion) (WorkoutResult, error) {
	var result WorkoutResult
	if event.UserID <= 0 || event.EventID == "" || event.Day < 1 || event.Day > 30 || event.Position < 0 || event.Position > math.MaxInt-30 {
		return result, errors.New("invalid workout completion")
	}
	if _, err := time.Parse("2006-01-02", event.CompletedOn); err != nil {
		return result, err
	}
	tx, err := d.sql.BeginTx(ctx, nil)
	if err != nil {
		return result, err
	}
	defer tx.Rollback()
	// A scheduled restart moves to a fresh cycle before considering an old
	// completion's captured position. Its activity date is still credited below.
	if _, _, err := applyDueCycleRestart(ctx, tx, event.UserID, time.Now(), 30); err != nil {
		return result, err
	}
	inserted, err := tx.ExecContext(ctx, `INSERT INTO workout_completions
		(user_id, event_id, day, position, completed_at, completed_on) VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(user_id, event_id) DO NOTHING`,
		event.UserID, event.EventID, event.Day, event.Position, event.CompletedAt, event.CompletedOn)
	if err != nil {
		return result, err
	}
	n, err := inserted.RowsAffected()
	if err != nil {
		return result, err
	}
	if n == 0 {
		previous := WorkoutCompletion{UserID: event.UserID, EventID: event.EventID}
		if err := tx.QueryRowContext(ctx, `SELECT day, position, completed_at, completed_on
			FROM workout_completions WHERE user_id = ? AND event_id = ?`, event.UserID, event.EventID).
			Scan(&previous.Day, &previous.Position, &previous.CompletedAt, &previous.CompletedOn); err != nil {
			return result, err
		}
		if previous != event {
			return result, ErrCompletionConflict
		}
	} else {
		if _, err := tx.ExecContext(ctx, `INSERT INTO activity (user_id, done_on) VALUES (?, ?)
			ON CONFLICT(user_id, done_on) DO NOTHING`, event.UserID, event.CompletedOn); err != nil {
			return result, err
		}
		currentDay := event.Position%30 + 1
		if event.Day >= currentDay {
			nextPosition := event.Position - event.Position%30 + event.Day
			if _, err := tx.ExecContext(ctx, `UPDATE users SET position = ? WHERE id = ? AND position = ?`,
				nextPosition, event.UserID, event.Position); err != nil {
				return result, err
			}
		}
	}
	if err := tx.QueryRowContext(ctx, `SELECT position FROM users WHERE id = ?`, event.UserID).Scan(&result.Position); err != nil {
		return result, err
	}
	result.Streak, err = completionStreak(ctx, tx, event.UserID, event.CompletedOn)
	if err != nil {
		return result, err
	}
	if err := tx.Commit(); err != nil {
		return WorkoutResult{}, err
	}
	result.Saved, result.EventID = true, event.EventID
	return result, nil
}

// Read inside the transaction so a failed streak query cannot acknowledge an
// otherwise partially completed operation. The receipt's local date is retained.
func completionStreak(ctx context.Context, tx *sql.Tx, userID int64, on string) (int, error) {
	cursor, err := time.Parse("2006-01-02", on)
	if err != nil {
		return 0, err
	}
	rows, err := tx.QueryContext(ctx, `SELECT done_on FROM activity WHERE user_id = ? AND done_on <= ? ORDER BY done_on DESC`, userID, on)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	streak := 0
	for rows.Next() {
		var date string
		if err := rows.Scan(&date); err != nil {
			return 0, err
		}
		if streak == 0 && date != cursor.Format("2006-01-02") {
			cursor = cursor.AddDate(0, 0, -1)
		}
		if date != cursor.Format("2006-01-02") {
			break
		}
		streak++
		cursor = cursor.AddDate(0, 0, -1)
	}
	return streak, rows.Err()
}
