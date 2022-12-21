package usecase

import (
	"fmt"
	"profileyou/api/domain/model/user"
	"profileyou/api/domain/repository"
)

type UserUseCase interface {
	CreateUser(email string, password string) error
	GetUserForAuth(email string) (result *user.User, err error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu *userUseCase) CreateUser(email string, password string) error {
	user := user.User{Email: email, Password: password}
	fmt.Printf("useCase: %v\n", user)
	err := uu.userRepository.Create(user)
	// err := uu.userRepository.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUseCase) GetUserForAuth(email string) (result *user.User, err error) {
	fmt.Printf("User email %v has requested to log in!\n", email)
	current_user, err := uu.userRepository.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Email is null")
		return nil, err
	}

	return current_user, nil
}
