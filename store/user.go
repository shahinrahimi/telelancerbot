package store

import (
	"database/sql"
	"fmt"

	"github.com/shahinrahimi/telelancerbot/models"
)

func (s *SqliteStore) InsertUser(u *models.User) error {
	if _, err := s.db.Exec(models.INSERT_USER, u.ToArgs()...); err != nil {
		return fmt.Errorf("failed to insert user to DB: %v", err)
	}
	return nil
}

func (s *SqliteStore) GetUser(id int64) (*models.User, error) {
	u := &models.User{ID: id}
	if err := s.db.QueryRow(models.SELECT_USER, id).Scan(u.ToFelids()...); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get user %d from DB: %v", id, err)
	}
	return u, nil
}

func (s *SqliteStore) GetUsers() ([]*models.User, error) {
	rows, err := s.db.Query(models.SELECT_USERS)
	if err != nil {
		return nil, fmt.Errorf("failed to get users from DB: %v", err)
	}
	defer rows.Close()
	var us []*models.User
	for rows.Next() {
		u := &models.User{}
		if err := rows.Scan(u.ToFelids()...); err != nil {
			s.l.Printf("failed to scan user from DB: %v", err)
			continue
		}

		us = append(us, u)

	}
	return us, nil
}
func (s *SqliteStore) UpdateUser(u *models.User) error {
	if _, err := s.db.Exec(models.UPDATE_USER, u.ToUpdatedArgs()...); err != nil {
		return fmt.Errorf("failed to update user %d in DB: %v", u.ID, err)
	}
	return nil
}

func (s *SqliteStore) DeleteUser(id int64) error {
	if _, err := s.db.Exec(models.DELETE_USER, id); err != nil {
		return fmt.Errorf("failed to delete user %d from DB: %v", id, err)
	}
	return nil
}
