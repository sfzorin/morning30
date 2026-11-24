package db

import (
	"database/sql"
	"errors"
	"time"
)

// SavedProgramsMax is how many replaced programs an account keeps.
const SavedProgramsMax = 10

// ErrSavedProgramNotFound is returned when a saved program id matches nothing
// for that account.
var ErrSavedProgramNotFound = errors.New("saved program not found")

// SavedProgram is a personal program that was replaced and kept so the user
// can bring it back.
type SavedProgram struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	SavedAt string `json:"savedAt"` // UTC, RFC 3339
}

// ReplaceProgram installs a program ("" clears the personal file) together
// with its active key, in one transaction. When keepAs is not empty, the
// personal program being replaced is first kept under that name.
// Cycle position, streak and activity are not touched.
func (d *DB) ReplaceProgram(userID int64, program, active, keepAs string) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if keepAs != "" {
		if err := keepCurrentProgram(tx, userID, keepAs); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(`UPDATE users SET program_json = ?, active_program = ? WHERE id = ?`, program, active, userID); err != nil {
		return err
	}
	return tx.Commit()
}

// RestoreSavedProgram makes a saved program the account's personal program
// again and removes it from the saved list. When keepAs is not empty, the
// program it replaces is kept under that name, so the two can be swapped back.
func (d *DB) RestoreSavedProgram(userID, savedID int64, keepAs string) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var program string
	err = tx.QueryRow(`SELECT program_json FROM saved_programs WHERE id = ? AND user_id = ?`, savedID, userID).Scan(&program)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrSavedProgramNotFound
	}
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM saved_programs WHERE id = ? AND user_id = ?`, savedID, userID); err != nil {
		return err
	}
	if keepAs != "" {
		if err := keepCurrentProgram(tx, userID, keepAs); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(`UPDATE users SET program_json = ?, active_program = 'custom' WHERE id = ?`, program, userID); err != nil {
		return err
	}
	return tx.Commit()
}

// SavedPrograms lists the account's kept programs, newest first.
func (d *DB) SavedPrograms(userID int64) ([]SavedProgram, error) {
	rows, err := d.sql.Query(`SELECT id, name, saved_at FROM saved_programs WHERE user_id = ? ORDER BY id DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []SavedProgram{}
	for rows.Next() {
		var p SavedProgram
		if err := rows.Scan(&p.ID, &p.Name, &p.SavedAt); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// keepCurrentProgram copies the account's personal program into the saved
// list. An identical copy already there is replaced rather than duplicated,
// and only the newest SavedProgramsMax are kept.
func keepCurrentProgram(tx *sql.Tx, userID int64, name string) error {
	var current string
	if err := tx.QueryRow(`SELECT program_json FROM users WHERE id = ?`, userID).Scan(&current); err != nil {
		return err
	}
	if current == "" {
		return nil
	}
	if _, err := tx.Exec(`DELETE FROM saved_programs WHERE user_id = ? AND program_json = ?`, userID, current); err != nil {
		return err
	}
	if _, err := tx.Exec(
		`INSERT INTO saved_programs (user_id, name, program_json, saved_at) VALUES (?, ?, ?, ?)`,
		userID, name, current, time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}
	_, err := tx.Exec(
		`DELETE FROM saved_programs WHERE user_id = ? AND id NOT IN
		   (SELECT id FROM saved_programs WHERE user_id = ? ORDER BY id DESC LIMIT ?)`,
		userID, userID, SavedProgramsMax,
	)
	return err
}
