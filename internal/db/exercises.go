package db

import "strings"

// SeedExercise inserts a built-in exercise Doc if it's not already present
// (idempotent — existing rows are preserved).
func (d *DB) SeedExercise(id, doc string) error {
	_, err := d.sql.Exec(`INSERT OR IGNORE INTO exercises(id, doc) VALUES(?, ?)`, id, doc)
	return err
}

// PruneExercises deletes global exercises whose ID is not in keep (removes
// exercises dropped from the built-in library).
func (d *DB) PruneExercises(keep []string) error {
	if len(keep) == 0 {
		return nil
	}
	ph := strings.TrimSuffix(strings.Repeat("?,", len(keep)), ",")
	args := make([]any, len(keep))
	for i, k := range keep {
		args[i] = k
	}
	_, err := d.sql.Exec(`DELETE FROM exercises WHERE id NOT IN (`+ph+`)`, args...)
	return err
}

// AllExercises returns the global exercise library as id -> Doc JSON.
func (d *DB) AllExercises() (map[string]string, error) {
	rows, err := d.sql.Query(`SELECT id, doc FROM exercises`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	m := make(map[string]string)
	for rows.Next() {
		var id, doc string
		if err := rows.Scan(&id, &doc); err != nil {
			return nil, err
		}
		m[id] = doc
	}
	return m, rows.Err()
}

// GetExercise returns one global exercise Doc JSON ("" if absent).
func (d *DB) GetExercise(id string) (string, error) {
	var doc string
	err := d.sql.QueryRow(`SELECT doc FROM exercises WHERE id = ?`, id).Scan(&doc)
	return doc, err
}

// SetExercise inserts or replaces a global exercise Doc (for future global edits).
func (d *DB) SetExercise(id, doc string) error {
	_, err := d.sql.Exec(
		`INSERT INTO exercises(id, doc, updated_at) VALUES(?, ?, datetime('now'))
		 ON CONFLICT(id) DO UPDATE SET doc = excluded.doc, updated_at = datetime('now')`,
		id, doc)
	return err
}
