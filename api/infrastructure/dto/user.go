package dto

import (
	"profileyou/api/domain/model/user"
	"time"
)

type User struct {
	ID        int
	UserId    string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func ConvertUSer(k *user.User) *User {
	return &User{
		UserId:   k.GetUserId(),
		Email:    k.GetEmail(),
		Password: k.GetPassword(),
	}
}

func AdaptUser(converted_user *User) (*user.User, error) {
	user, err := user.New(
		converted_user.UserId,
		converted_user.Email,
		converted_user.Password,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func AdaptUsers(converted_users []*User) ([]*user.User, error) {
	var users []*user.User

	for _, converted_user := range converted_users {
		user, err := user.New(
			converted_user.UserId,
			converted_user.Email,
			converted_user.Password,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
