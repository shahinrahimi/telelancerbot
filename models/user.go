package models

import "time"

type KeyUser struct{}

type User struct {
	ID          int64     `json:"id"`       // is telegram user id or chat id
	Username    string    `json:"username"` // is telegram username
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	IsConfirmed bool      `json:"is_confirmed"`
	IsAdmin     bool      `json:"is_admin"`
}

const (
	CREATE_TABLE_USERS = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT,
		is_confirmed BOOLEAN DEFAULT FALSE,
		is_admin BOOLEAN DEFAULT FALSE,
		updated_at DATETIME DEFAULT CURRENT_TIME,
		created_at DATETIME DEFAULT CURRENT_TIME
	)`
	SELECT_COUNT_USERS string = "SELECT COUNT(*) FROM users"
	SELECT_USERS       string = "SELECT * FROM users"
	SELECT_USER        string = "SELECT * FROM users WHERE id = ?"
	INSERT_USER        string = "INSERT INTO users (id, username, is_admin) VALUES (?, ?, ?)"
	UPDATE_USER        string = "UPDATE users SET is_confirmed = ?, is_admin = ?, updated_at = CURRENT_TIME WHERE id = ?"
	DELETE_USER        string = "DELETE FROM users WHERE id = ?"
)

// use for inserting to DB
func (u *User) ToArgs() []interface{} {
	return []interface{}{u.ID, u.Username, u.IsAdmin}
}

// use for scanning from DB
func (u *User) ToFelids() []interface{} {
	return []interface{}{&u.ID, &u.Username, &u.IsConfirmed, &u.IsAdmin, &u.UpdatedAt, &u.CreatedAt}
}

// use for updating record to DB
func (u *User) ToUpdatedArgs() []interface{} {
	return []interface{}{u.IsConfirmed, u.IsAdmin, u.ID}
}
