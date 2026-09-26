package db

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

// ErrEmailTaken is returned when registering an email that already exists.
var ErrEmailTaken = errors.New("email already registered")

func scanUser(row interface {
	Scan(dest ...any) error
}) (User, error) {
	var u User
	var voice, guest, jokes int
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Lang, &u.Rest, &voice, &guest, &u.VoiceMode, &u.Level, &jokes, &u.Created)
	u.Voice = voice != 0
	u.IsGuest = guest != 0
	u.Jokes = jokes != 0
	return u, err
}

const userCols = `id, email, name, lang, rest_seconds, voice, is_guest, voice_mode, level, jokes, created_at`

// CreateUser inserts a new account. Returns ErrEmailTaken on duplicate email.
func (d *DB) CreateUser(email, passHash, lang, name string) (User, error) {
	res, err := d.sql.Exec(
		// rest_warmup/rest_cooldown set explicitly so new users get the current
		// defaults even on a DB whose column default predates them.
		`INSERT INTO users (email, pass_hash, lang, name, rest_warmup, rest_cooldown, active_program) VALUES (?, ?, ?, ?, 10, 10, 'demo')`,
		email, passHash, lang, name,
	)
	if err != nil {
		// modernc sqlite reports UNIQUE violations in the error text.
		return User{}, ErrEmailTaken
	}
	id, _ := res.LastInsertId()
	return d.UserByID(id)
}

// CreateGuest inserts an anonymous account identified only by a cookie session.
// email must be a unique placeholder (e.g. "guest:<token>").
func (d *DB) CreateGuest(email, lang string) (User, error) {
	res, err := d.sql.Exec(
		`INSERT INTO users (email, pass_hash, lang, is_guest, rest_warmup, rest_cooldown, active_program) VALUES (?, '', ?, 1, 10, 10, 'demo')`,
		email, lang,
	)
	if err != nil {
		return User{}, err
	}
	id, _ := res.LastInsertId()
	return d.UserByID(id)
}

// ConvertGuest upgrades a guest account into a real one in place, preserving all
// progress. Returns ErrEmailTaken if the email already belongs to someone else.
func (d *DB) ConvertGuest(userID int64, email, passHash, name string) error {
	if _, err := d.sql.Exec(
		`UPDATE users SET email = ?, pass_hash = ?, name = ?, is_guest = 0 WHERE id = ?`,
		email, passHash, name, userID,
	); err != nil {
		return ErrEmailTaken
	}
	return nil
}

// UserByID loads a user by primary key.
func (d *DB) UserByID(id int64) (User, error) {
	row := d.sql.QueryRow(`SELECT `+userCols+` FROM users WHERE id = ?`, id)
	return scanUser(row)
}

// UserByEmail loads a user by email (ok=false if not found).
func (d *DB) UserByEmail(email string) (User, bool, error) {
	row := d.sql.QueryRow(`SELECT `+userCols+` FROM users WHERE email = ?`, email)
	u, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, false, nil
	}
	if err != nil {
		return User{}, false, err
	}
	return u, true, nil
}

// PassHash returns the stored bcrypt hash for an email (ok=false if no user).
func (d *DB) PassHash(email string) (hash string, ok bool, err error) {
	row := d.sql.QueryRow(`SELECT pass_hash FROM users WHERE email = ?`, email)
	err = row.Scan(&hash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return hash, true, nil
}

// UpdateSettings updates per-user preferences, including the three per-phase
// rest durations. voiceMode is off|min|normal|detailed.
func (d *DB) UpdateSettings(userID int64, name, lang string, restWarmup, restMain, restCooldown int, voiceMode string) error {
	v := 1
	if voiceMode == "off" {
		v = 0
	}
	_, err := d.sql.Exec(
		`UPDATE users SET name = ?, lang = ?, rest_seconds = ?, rest_warmup = ?, rest_cooldown = ?, voice = ?, voice_mode = ? WHERE id = ?`,
		name, lang, restMain, restWarmup, restCooldown, v, voiceMode, userID,
	)
	return err
}

// Rests returns the user's per-phase rest seconds (warm-up, main, cool-down),
// falling back to the defaults if the row is missing.
func (d *DB) Rests(userID int64) (warmup, main, cooldown int) {
	warmup, main, cooldown = 10, 20, 10
	_ = d.sql.QueryRow(
		`SELECT rest_warmup, rest_seconds, rest_cooldown FROM users WHERE id = ?`, userID,
	).Scan(&warmup, &main, &cooldown)
	return
}

// SafetyAck reports whether the user acknowledged the safety disclaimer.
func (d *DB) SafetyAck(userID int64) (bool, error) {
	var n int
	err := d.sql.QueryRow(`SELECT safety_ack FROM users WHERE id = ?`, userID).Scan(&n)
	return n != 0, err
}

// AckSafety records that the user acknowledged the safety disclaimer.
func (d *DB) AckSafety(userID int64) error {
	_, err := d.sql.Exec(`UPDATE users SET safety_ack = 1 WHERE id = ?`, userID)
	return err
}

// GetProgramJSON returns the user's custom program JSON ("" = built-in).
func (d *DB) GetProgramJSON(userID int64) (string, error) {
	var s string
	err := d.sql.QueryRow(`SELECT program_json FROM users WHERE id = ?`, userID).Scan(&s)
	return s, err
}

// SetProgramJSON stores (or clears, with "") the user's custom program JSON.
func (d *DB) SetProgramJSON(userID int64, j string) error {
	_, err := d.sql.Exec(`UPDATE users SET program_json = ? WHERE id = ?`, j, userID)
	return err
}

// ActiveProgram returns the stored active-program key ("custom" or a standard
// program key); the app layer normalizes legacy/unknown values.
func (d *DB) ActiveProgram(userID int64) (string, error) {
	var s string
	err := d.sql.QueryRow(`SELECT active_program FROM users WHERE id = ?`, userID).Scan(&s)
	return s, err
}

// SetActiveProgram stores which program is active ("custom" or a standard key).
func (d *DB) SetActiveProgram(userID int64, which string) error {
	_, err := d.sql.Exec(`UPDATE users SET active_program = ? WHERE id = ?`, which, userID)
	return err
}

// GetLevel returns the user's universal difficulty level (clamped −3..+3).
func (d *DB) GetLevel(userID int64) int {
	var lvl int
	_ = d.sql.QueryRow(`SELECT level FROM users WHERE id = ?`, userID).Scan(&lvl)
	return clampLevel(lvl)
}

// SetLevel stores the user's universal difficulty level (clamped −3..+3).
func (d *DB) SetLevel(userID int64, level int) error {
	_, err := d.sql.Exec(`UPDATE users SET level = ? WHERE id = ?`, clampLevel(level), userID)
	return err
}

func clampLevel(l int) int {
	if l < -3 {
		return -3
	}
	if l > 3 {
		return 3
	}
	return l
}

// GetSex returns the user's figure preference ("m" or "f"; defaults to "m").
func (d *DB) GetSex(userID int64) string {
	var s string
	_ = d.sql.QueryRow(`SELECT sex FROM users WHERE id = ?`, userID).Scan(&s)
	if s != "f" {
		return "m"
	}
	return "f"
}

// SetSex stores the user's figure preference (anything but "f" is stored as "m").
func (d *DB) SetSex(userID int64, sex string) error {
	if sex != "f" {
		sex = "m"
	}
	_, err := d.sql.Exec(`UPDATE users SET sex = ? WHERE id = ?`, sex, userID)
	return err
}

// GetCustomExercises returns the user's custom exercise library JSON ("" = none).
func (d *DB) GetCustomExercises(userID int64) (string, error) {
	var s string
	err := d.sql.QueryRow(`SELECT custom_exercises FROM users WHERE id = ?`, userID).Scan(&s)
	return s, err
}

// SetCustomExercises stores (or clears, with "") the user's custom exercise library.
func (d *DB) SetCustomExercises(userID int64, j string) error {
	_, err := d.sql.Exec(`UPDATE users SET custom_exercises = ? WHERE id = ?`, j, userID)
	return err
}

// SetLang updates only the language (used by the quick switcher).
func (d *DB) SetLang(userID int64, lang string) error {
	_, err := d.sql.Exec(`UPDATE users SET lang = ? WHERE id = ?`, lang, userID)
	return err
}

// Account is the MCP/settings identity of a user (no secrets).
type Account struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Guest         bool   `json:"guest"`
	ActiveProgram string `json:"activeProgram"`
	HasProgram    bool   `json:"hasProgram"`
}

var (
	// ErrAccountNotFound is returned when FindAccount matches nobody.
	ErrAccountNotFound = errors.New("account not found")
	// ErrAccountAmbiguous is returned when a name matches more than one account.
	ErrAccountAmbiguous = errors.New("account name is ambiguous")
)

// ListAccounts returns every user row (guests included) for MCP listing.
func (d *DB) ListAccounts() ([]Account, error) {
	rows, err := d.sql.Query(
		`SELECT id, email, name, is_guest, COALESCE(active_program, ''),
		        CASE WHEN program_json IS NULL OR program_json = '' THEN 0 ELSE 1 END
		 FROM users ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Account
	for rows.Next() {
		var a Account
		var guest, has int
		if err := rows.Scan(&a.ID, &a.Email, &a.Name, &guest, &a.ActiveProgram, &has); err != nil {
			return nil, err
		}
		a.Guest = guest != 0
		a.HasProgram = has != 0
		out = append(out, a)
	}
	return out, rows.Err()
}

// AccountByID loads the MCP identity for a user id.
func (d *DB) AccountByID(id int64) (Account, error) {
	var a Account
	var guest, has int
	err := d.sql.QueryRow(
		`SELECT id, email, name, is_guest, COALESCE(active_program, ''),
		        CASE WHEN program_json IS NULL OR program_json = '' THEN 0 ELSE 1 END
		 FROM users WHERE id = ?`, id,
	).Scan(&a.ID, &a.Email, &a.Name, &guest, &a.ActiveProgram, &has)
	if errors.Is(err, sql.ErrNoRows) {
		return Account{}, ErrAccountNotFound
	}
	if err != nil {
		return Account{}, err
	}
	a.Guest = guest != 0
	a.HasProgram = has != 0
	return a, nil
}

// FindAccount resolves an MCP "account" argument: numeric id, email, or unique name.
func (d *DB) FindAccount(who string) (Account, error) {
	who = strings.TrimSpace(who)
	if who == "" {
		return Account{}, ErrAccountNotFound
	}
	if id, err := strconv.ParseInt(who, 10, 64); err == nil {
		return d.AccountByID(id)
	}
	if u, ok, err := d.UserByEmail(who); err != nil {
		return Account{}, err
	} else if ok {
		return d.AccountByID(u.ID)
	}
	rows, err := d.sql.Query(`SELECT id FROM users WHERE lower(name) = lower(?)`, who)
	if err != nil {
		return Account{}, err
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return Account{}, err
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return Account{}, err
	}
	switch len(ids) {
	case 1:
		return d.AccountByID(ids[0])
	case 0:
		return Account{}, ErrAccountNotFound
	default:
		return Account{}, ErrAccountAmbiguous
	}
}
