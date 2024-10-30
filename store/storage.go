package store

import "github.com/shahinrahimi/telelancerbot/models"

type Storage interface {
	GetUsers() ([]*models.User, error)
	GetUser(id int64) (*models.User, error)
	InsertUser(u *models.User) error
	UpdateUser(u *models.User) error
	DeleteUser(id int64) error
}
