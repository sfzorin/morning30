package db

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

const mcpTokenPrefix = "m30_"

func hashMCPToken(tok string) string {
	sum := sha256.Sum256([]byte(tok))
	return hex.EncodeToString(sum[:])
}

// HasMCPToken reports whether the account has an active MCP bearer token.
func (d *DB) HasMCPToken(userID int64) bool {
	var h string
	_ = d.sql.QueryRow(`SELECT mcp_token_hash FROM users WHERE id = ?`, userID).Scan(&h)
	return h != ""
}

// IssueMCPToken creates a new bearer token for the user and stores only its hash.
// The plaintext is returned once; a later issue invalidates the previous token.
func (d *DB) IssueMCPToken(userID int64) (string, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	tok := mcpTokenPrefix + base64.RawURLEncoding.EncodeToString(raw)
	_, err := d.sql.Exec(`UPDATE users SET mcp_token_hash = ? WHERE id = ?`, hashMCPToken(tok), userID)
	if err != nil {
		return "", err
	}
	return tok, nil
}

// AccountByMCPToken resolves a per-user bearer token (plaintext) to that account.
func (d *DB) AccountByMCPToken(tok string) (Account, error) {
	tok = strings.TrimSpace(tok)
	if tok == "" || !strings.HasPrefix(tok, mcpTokenPrefix) {
		return Account{}, ErrAccountNotFound
	}
	var id int64
	err := d.sql.QueryRow(`SELECT id FROM users WHERE mcp_token_hash = ?`, hashMCPToken(tok)).Scan(&id)
	if err != nil {
		return Account{}, ErrAccountNotFound
	}
	return d.AccountByID(id)
}
