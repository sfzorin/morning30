package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"
)

// ScheduleCycleRestart replaces the pending restart instant without changing
// progress or activity. The instant is stored in UTC, independently of server timezone.
func (d *DB) ScheduleCycleRestart(ctx context.Context, userID int64, at time.Time) error {
	if userID <= 0 || at.IsZero() {
		return errors.New("invalid cycle restart schedule")
	}
	stamp, err := at.UTC().MarshalText()
	if err != nil {
		return fmt.Errorf("cycle restart instant: %w", err)
	}
	result, err := d.sql.ExecContext(ctx, `UPDATE users SET cycle_restart_on = ? WHERE id = ?`, string(stamp), userID)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// ScheduledCycleRestart returns the pending instant, or zero when none is set.
func (d *DB) ScheduledCycleRestart(ctx context.Context, userID int64) (time.Time, error) {
	var scheduled string
	if err := d.sql.QueryRowContext(ctx, `SELECT cycle_restart_on FROM users WHERE id = ?`, userID).Scan(&scheduled); err != nil {
		return time.Time{}, err
	}
	if scheduled == "" {
		return time.Time{}, nil
	}
	at, err := time.Parse(time.RFC3339Nano, scheduled)
	if err != nil {
		return time.Time{}, fmt.Errorf("stored cycle restart instant: %w", err)
	}
	return at, nil
}

// ApplyDueCycleRestart advances to day one of the next cycle when the scheduled
// instant has arrived. It never removes activity or completion receipts. Callers
// must supply server time, not a completion timestamp supplied by the client.
func (d *DB) ApplyDueCycleRestart(ctx context.Context, userID int64, now time.Time, totalDays int) (position int, changed bool, err error) {
	tx, err := d.sql.BeginTx(ctx, nil)
	if err != nil {
		return 0, false, err
	}
	defer tx.Rollback()
	position, changed, err = applyDueCycleRestart(ctx, tx, userID, now, totalDays)
	if err != nil {
		return 0, false, err
	}
	if err := tx.Commit(); err != nil {
		return 0, false, err
	}
	return position, changed, nil
}

// Completion handling uses the same transaction so a receipt captured before a
// due restart cannot advance the newly started cycle, even without a page visit.
func applyDueCycleRestart(ctx context.Context, tx *sql.Tx, userID int64, now time.Time, totalDays int) (int, bool, error) {
	if userID <= 0 || now.IsZero() || totalDays <= 0 {
		return 0, false, errors.New("invalid cycle restart request")
	}
	var position int
	var scheduled string
	if err := tx.QueryRowContext(ctx, `SELECT position, cycle_restart_on FROM users WHERE id = ?`, userID).Scan(&position, &scheduled); err != nil {
		return 0, false, err
	}
	if position < 0 {
		return 0, false, errors.New("invalid cycle position")
	}
	if scheduled == "" {
		return position, false, nil
	}
	at, err := time.Parse(time.RFC3339Nano, scheduled)
	if err != nil {
		return 0, false, fmt.Errorf("stored cycle restart instant: %w", err)
	}
	if now.Before(at) {
		return position, false, nil
	}
	// Move forward even at an exact boundary: position also distinguishes stale
	// completions from new workouts, so reusing the same value would be unsafe.
	advance := totalDays - position%totalDays
	if position > math.MaxInt-advance {
		return 0, false, errors.New("cycle restart position overflow")
	}
	next := position + advance
	if _, err := tx.ExecContext(ctx, `UPDATE users SET position = ?, cycle_restart_on = '' WHERE id = ?`, next, userID); err != nil {
		return 0, false, err
	}
	return next, true, nil
}
