package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shahinrahimi/telelancerbot/models"
)

type SqliteStore struct {
	l  *log.Logger
	db *sql.DB
}

func New(l *log.Logger) *SqliteStore {
	return &SqliteStore{l: l}
}

func Init(s *SqliteStore, rootID int64) error {
	if err := os.MkdirAll("db", 0755); err != nil {
		return fmt.Errorf("failed to create db directory: %v", err)
	}
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	if _, err := db.Exec(models.CREATE_TABLE_USERS); err != nil {
		return fmt.Errorf("failed to create table for users: %v", err)
	}
	s.db = db
	s.l.Println("DB Connected!")
	// create root user
	u, err := s.GetUser(rootID)
	if err != nil {
		if err != sql.ErrNoRows {
			s.l.Panicf("unexpected error to get root user: %v", err)
		}
		u = &models.User{
			ID:          rootID,
			IsAdmin:     true,
			IsConfirmed: true,
		}
		if err := s.InsertUser(u); err != nil {
			s.l.Panicf("unexpected error to insert root user: %v", err)
		}
		s.l.Printf("Root user created: %d", rootID)
	}
	if !u.IsAdmin {
		// update root user
		u.IsAdmin = true
		u.IsConfirmed = true
		if err := s.UpdateUser(u); err != nil {
			s.l.Panicf("unexpected error to update root user: %v", err)
		}
	}
	return nil
}
