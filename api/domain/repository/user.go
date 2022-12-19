package repository

import (
	"profileyou/api/domain/model/user"
)

type UserRepository interface {
	// GetKeyword(id string) (result *user.User, err error)
	// GetUsers() (result []*user.User, err error)
	// Create(k *user.User) error
	// Update(k *user.User) error
	// Delete(k *user.User) error
	GetUserByEmail(email string) (result *user.User, err error)
	GetUserByID(id int) (*user.User, error)
}
