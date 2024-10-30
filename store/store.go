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

func Init(s *SqliteStore) error {
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
	return nil
}
