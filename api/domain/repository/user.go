package repository

import (
	"profileyou/api/domain/model/user"
)

type UserRepository interface {
	GetUserByEmail(email string) (result *user.User, err error)
	GetUserByID(id int) (result *user.User, err error)
	Create(u user.User) error
}
