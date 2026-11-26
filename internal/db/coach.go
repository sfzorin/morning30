package db

import (
	"time"
)

// CoachMessage is one line of the AI coach conversation.
type CoachMessage struct {
	ID        int64
	ProgramID int64
	Role      string // user | coach
	Text      string
	Changes   string // JSON
	Undo      string // JSON snapshot before the edits ("" = nothing to undo)
	Undone    bool
	Created   time.Time
}

// AddCoachMessage stores one message and returns its id.
func (d *DB) AddCoachMessage(userID int64, m CoachMessage) (int64, error) {
	res, err := d.sql.Exec(`INSERT INTO coach_messages (user_id, program_id, role, text, changes, undo, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, userID, m.ProgramID, m.Role, m.Text, m.Changes, m.Undo, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// CoachMessages returns the latest messages of one program's thread, oldest first.
func (d *DB) CoachMessages(userID, programID int64, limit int) ([]CoachMessage, error) {
	rows, err := d.sql.Query(`SELECT id, program_id, role, text, changes, undo, undone, created_at FROM (
		SELECT * FROM coach_messages WHERE user_id = ? AND program_id = ? ORDER BY id DESC LIMIT ?) ORDER BY id`,
		userID, programID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []CoachMessage
	for rows.Next() {
		var m CoachMessage
		var undone int
		var at string
		if err := rows.Scan(&m.ID, &m.ProgramID, &m.Role, &m.Text, &m.Changes, &m.Undo, &undone, &at); err != nil {
			return nil, err
		}
		m.Undone = undone != 0
		m.Created, _ = time.Parse(time.RFC3339, at)
		out = append(out, m)
	}
	return out, rows.Err()
}

// CoachMessage returns one message of the user.
func (d *DB) CoachMessage(userID, id int64) (CoachMessage, error) {
	var m CoachMessage
	var undone int
	var at string
	err := d.sql.QueryRow(`SELECT id, program_id, role, text, changes, undo, undone, created_at FROM coach_messages
		WHERE user_id = ? AND id = ?`, userID, id).
		Scan(&m.ID, &m.ProgramID, &m.Role, &m.Text, &m.Changes, &m.Undo, &undone, &at)
	m.Undone = undone != 0
	m.Created, _ = time.Parse(time.RFC3339, at)
	return m, err
}

// MarkCoachUndone records that a reply's edits were undone.
func (d *DB) MarkCoachUndone(userID, id int64) error {
	_, err := d.sql.Exec(`UPDATE coach_messages SET undone = 1 WHERE user_id = ? AND id = ?`, userID, id)
	return err
}

// CoachRequestsSince counts the user's own messages since t (the daily limit).
func (d *DB) CoachRequestsSince(userID int64, t time.Time) int {
	var n int
	_ = d.sql.QueryRow(`SELECT COUNT(*) FROM coach_messages WHERE user_id = ? AND role = 'user' AND created_at >= ?`,
		userID, t.UTC().Format(time.RFC3339)).Scan(&n)
	return n
}
