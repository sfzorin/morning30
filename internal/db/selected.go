package db

import (
	"database/sql"
	"errors"
	"time"
)

// Selected programs: a user's own shortlist. Each row keeps its own copy of a
// program (with the user's edits), its own cycle day and its own workout
// settings: difficulty, number of main rounds and the three rests. The program the user
// is running right now also lives in users.program_json / users.position, the
// "active slot" that the player, completions and cycle restarts already use —
// as do users.level, users.main_rounds and the users.rest_* columns. Switching
// programs saves the slot into its row and loads the next row into it.

// ErrSelectedNotFound is returned when a selected program id is not the user's.
var ErrSelectedNotFound = errors.New("selected program not found")

// SelectedProgram is one entry of a user's selected programs.
type SelectedProgram struct {
	ID       int64  `json:"id"`
	Source   string `json:"source"` // catalog key it came from, "" for the user's own
	Name     string `json:"name"`
	Program  string `json:"-"`
	Position int    `json:"position"` // completed sessions; the day is position%30+1
	AddedAt  string `json:"addedAt"`
}

// ProgramEpoch counts program switches. It is part of the key that resumes an
// interrupted workout, so switching away makes that workout stale for good,
// even after switching back.
func (d *DB) ProgramEpoch(userID int64) int {
	var n int
	_ = d.sql.QueryRow(`SELECT program_epoch FROM users WHERE id = ?`, userID).Scan(&n)
	return n
}

// CurrentSelected returns the id of the selected program in the active slot
// (0 when the user runs the plain morning workout without a row).
func (d *DB) CurrentSelected(userID int64) int64 {
	var id int64
	_ = d.sql.QueryRow(`SELECT current_selected FROM users WHERE id = ?`, userID).Scan(&id)
	return id
}

// SelectedPrograms lists a user's selected programs in the order they were added.
// The active row's Program and Position are read from the active slot.
func (d *DB) SelectedPrograms(userID int64) ([]SelectedProgram, error) {
	var current int64
	var live string
	var livePos int
	if err := d.sql.QueryRow(`SELECT current_selected, program_json, position FROM users WHERE id = ?`, userID).
		Scan(&current, &live, &livePos); err != nil {
		return nil, err
	}
	rows, err := d.sql.Query(`SELECT id, source, name, program_json, position, added_at
		FROM selected_programs WHERE user_id = ? ORDER BY id`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []SelectedProgram{}
	for rows.Next() {
		var p SelectedProgram
		if err := rows.Scan(&p.ID, &p.Source, &p.Name, &p.Program, &p.Position, &p.AddedAt); err != nil {
			return nil, err
		}
		if p.ID == current {
			if live != "" {
				p.Program = live
			}
			p.Position = livePos
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// AddSelected stores a new selected program and returns its id.
func (d *DB) AddSelected(userID int64, source, name, program string, position int) (int64, error) {
	res, err := d.sql.Exec(`INSERT INTO selected_programs (user_id, source, name, program_json, position, added_at)
		VALUES (?, ?, ?, ?, ?, ?)`, userID, source, name, program, position, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// AdoptActive records the program in the active slot as a selected program and
// makes it current, without touching its cycle day. It is a no-op when the
// slot already has a row.
func (d *DB) AdoptActive(userID int64, source, name, program string) (int64, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	var current int64
	var position int
	if err := tx.QueryRow(`SELECT current_selected, position FROM users WHERE id = ?`, userID).Scan(&current, &position); err != nil {
		return 0, err
	}
	if current != 0 {
		return current, tx.Commit()
	}
	st, err := readSlot(tx, userID)
	if err != nil {
		return 0, err
	}
	res, err := tx.Exec(`INSERT INTO selected_programs (user_id, source, name, program_json, position, added_at,
		level, rounds, rest_warmup, rest_main, rest_cooldown) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID, source, name, program, position, time.Now().UTC().Format(time.RFC3339), st.level, st.rounds, st.warm, st.main, st.cool)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if _, err := tx.Exec(`UPDATE users SET program_json = ?, active_program = 'custom', current_selected = ? WHERE id = ?`,
		program, id, userID); err != nil {
		return 0, err
	}
	return id, tx.Commit()
}

// slotSettings are the per-program workout settings held in the active slot.
type slotSettings struct{ level, rounds, warm, main, cool int }

func readSlot(tx *sql.Tx, userID int64) (slotSettings, error) {
	var st slotSettings
	err := tx.QueryRow(`SELECT level, main_rounds, rest_warmup, rest_seconds, rest_cooldown FROM users WHERE id = ?`, userID).
		Scan(&st.level, &st.rounds, &st.warm, &st.main, &st.cool)
	return st, err
}

// saveSlot stores the active slot's program, day and settings into its row.
func saveSlot(tx *sql.Tx, userID, rowID int64, program string, position int) error {
	st, err := readSlot(tx, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE selected_programs SET program_json = ?, position = ?, level = ?, rounds = ?,
		rest_warmup = ?, rest_main = ?, rest_cooldown = ? WHERE id = ? AND user_id = ?`,
		program, position, st.level, st.rounds, st.warm, st.main, st.cool, rowID, userID)
	return err
}

// SwitchSelected saves the active slot into its row and loads another selected
// program into the slot, with that program's own cycle day.
func (d *DB) SwitchSelected(userID, targetID int64) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var current int64
	var live string
	var position int
	if err := tx.QueryRow(`SELECT current_selected, program_json, position FROM users WHERE id = ?`, userID).
		Scan(&current, &live, &position); err != nil {
		return err
	}
	var program string
	var targetPos int
	err = tx.QueryRow(`SELECT program_json, position FROM selected_programs WHERE id = ? AND user_id = ?`, targetID, userID).
		Scan(&program, &targetPos)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrSelectedNotFound
	}
	if err != nil {
		return err
	}
	if current == targetID {
		return tx.Commit()
	}
	if current != 0 {
		if err := saveSlot(tx, userID, current, live, position); err != nil {
			return err
		}
	}
	// The target's own settings; rests it never stored keep the current ones.
	var level, rounds int
	var warm, main, cool sql.NullInt64
	if err := tx.QueryRow(`SELECT level, rounds, rest_warmup, rest_main, rest_cooldown FROM selected_programs WHERE id = ? AND user_id = ?`,
		targetID, userID).Scan(&level, &rounds, &warm, &main, &cool); err != nil {
		return err
	}
	// A new epoch makes any interrupted workout of the previous program stale.
	if _, err := tx.Exec(`UPDATE users SET program_json = ?, active_program = 'custom', position = ?, current_selected = ?,
		program_epoch = program_epoch + 1, level = ?, main_rounds = ?,
		rest_warmup = COALESCE(?, rest_warmup), rest_seconds = COALESCE(?, rest_seconds), rest_cooldown = COALESCE(?, rest_cooldown)
		WHERE id = ?`,
		program, targetPos, targetID, level, rounds, warm, main, cool, userID); err != nil {
		return err
	}
	return tx.Commit()
}

// ClearActive returns the active slot to the plain morning workout at day 1,
// saving the slot's program into its row first.
func (d *DB) ClearActive(userID int64, demoKey string) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var current int64
	var live string
	var position int
	if err := tx.QueryRow(`SELECT current_selected, program_json, position FROM users WHERE id = ?`, userID).
		Scan(&current, &live, &position); err != nil {
		return err
	}
	if current != 0 {
		if err := saveSlot(tx, userID, current, live, position); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(`UPDATE users SET program_json = '', active_program = ?, position = 0, current_selected = 0,
		program_epoch = program_epoch + 1, level = 0, main_rounds = 2 WHERE id = ?`,
		demoKey, userID); err != nil {
		return err
	}
	return tx.Commit()
}

// SaveActive stores an edited program in the active slot and in its row.
func (d *DB) SaveActive(userID int64, name, program string) error {
	tx, err := d.sql.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var current int64
	if err := tx.QueryRow(`SELECT current_selected FROM users WHERE id = ?`, userID).Scan(&current); err != nil {
		return err
	}
	if _, err := tx.Exec(`UPDATE users SET program_json = ?, active_program = 'custom' WHERE id = ?`, program, userID); err != nil {
		return err
	}
	if current != 0 {
		if _, err := tx.Exec(`UPDATE selected_programs SET name = ?, program_json = ? WHERE id = ? AND user_id = ?`,
			name, program, current, userID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

// DeleteSelected removes a selected program that is not in the active slot.
func (d *DB) DeleteSelected(userID, id int64) error {
	if d.CurrentSelected(userID) == id {
		return errors.New("cannot delete the active program")
	}
	res, err := d.sql.Exec(`DELETE FROM selected_programs WHERE id = ? AND user_id = ?`, id, userID)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrSelectedNotFound
	}
	return nil
}

// TakeSavedPrograms moves the programs kept by the earlier "saved programs"
// list into selected programs, in one transaction. source names the catalog
// program a kept program is an untouched copy of ("" for the user's own).
func (d *DB) TakeSavedPrograms(userID int64, source func(program string) string) (int, error) {
	tx, err := d.sql.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	rows, err := tx.Query(`SELECT name, program_json, saved_at FROM saved_programs WHERE user_id = ? ORDER BY id`, userID)
	if err != nil {
		return 0, err
	}
	type kept struct{ name, program, at string }
	var list []kept
	for rows.Next() {
		var k kept
		if err := rows.Scan(&k.name, &k.program, &k.at); err != nil {
			rows.Close()
			return 0, err
		}
		list = append(list, k)
	}
	rows.Close()
	if len(list) == 0 {
		return 0, tx.Commit()
	}
	for _, k := range list {
		if _, err := tx.Exec(`INSERT INTO selected_programs (user_id, source, name, program_json, position, added_at)
			VALUES (?, ?, ?, ?, 0, ?)`, userID, source(k.program), k.name, k.program, k.at); err != nil {
			return 0, err
		}
	}
	if _, err := tx.Exec(`DELETE FROM saved_programs WHERE user_id = ?`, userID); err != nil {
		return 0, err
	}
	return len(list), tx.Commit()
}

// SetSelectedSource records which catalog program a row is a copy of.
func (d *DB) SetSelectedSource(userID, id int64, source string) error {
	_, err := d.sql.Exec(`UPDATE selected_programs SET source = ? WHERE id = ? AND user_id = ?`, source, id, userID)
	return err
}
