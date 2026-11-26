package db

// SavedProgram is a program kept by the earlier "saved programs" list. Those
// rows are moved into selected programs on first use (TakeSavedPrograms).
type SavedProgram struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	SavedAt string `json:"savedAt"` // UTC, RFC 3339
}

// PutSavedProgramForTest stores a row in the earlier "saved programs" table,
// so migration tests can start from data written by the previous release.
func (d *DB) PutSavedProgramForTest(userID int64, name, program string) error {
	_, err := d.sql.Exec(`INSERT INTO saved_programs (user_id, name, program_json, saved_at) VALUES (?, ?, ?, '2026-11-24T00:00:00Z')`,
		userID, name, program)
	return err
}
