package user

import (
	_ "errors"
	_ "fmt"

	sqlite "profileyou/api/config/database"

	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

func (u *User) Create() {
	db := sqlite.New()

	connect, err := db.DB()
	if err != nil {
		panic(err)
	}

	db.Create(u)

	connect.Close()
}

// func (u *User) PasswordMatches(plainText string) (bool, error) {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(plainText))
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
// 			// invalid password
// 			return false, nil
// 		default:
// 			return false, err
// 		}
// 	}

// 	return true, nil
// }

// func New(userId string, email string, password string) (*User, error) {
// 	createdUserId, err := NewUserId(userId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	createdEmail, err := newEmail(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	createdPassword, err := newPassword(password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := User{
// 		userId:   *createdUserId,
// 		email:    *createdEmail,
// 		password: *createdPassword,
// 	}

// 	return &user, nil
// }

// // Create Constructor
// func Create(email string, password string) (*User, error) {
// 	userId := uuid.New().String()
// 	user, err := New(userId, email, password)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, err
// }

// // Getter
// func (u User) GetUserId() string {
// 	return string(u.userId)
// }

// func (u User) GetEmail() string {
// 	return string(u.email)
// }

// func (u User) GetPassword() string {
// 	return string(u.password)
// }

// // value constructors
// func NewUserId(value string) (*userId, error) {
// 	if value == "" {
// 		err := fmt.Errorf("%s", "empty arg:userId NewUserId()")
// 		return nil, err
// 	}

// 	userId := userId(value)

// 	return &userId, nil
// }

// func newEmail(value string) (*email, error) {
// 	if value == "" {
// 		err := fmt.Errorf("%s", "empty arg:email newEmail()")
// 		return nil, err
// 	}

// 	email := email(value)

// 	return &email, nil
// }

// func newPassword(value string) (*password, error) {
// 	if value == "" {
// 		err := fmt.Errorf("%s", "empty arg:password newPassword()")
// 		return nil, err
// 	}

// 	password := password(value)

// 	return &password, nil
// }
