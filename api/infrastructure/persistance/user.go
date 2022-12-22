package persistance

import (
	"fmt"
	"profileyou/api/domain/model/user"
	"profileyou/api/domain/repository"

	"gorm.io/gorm"
)

type userPersistance struct {
	Conn *gorm.DB
}

func NewUserPersistance(conn *gorm.DB) repository.UserRepository {
	return &userPersistance{Conn: conn}
}

// Login Authenticate

// GetUserByEmail returns one use, by email.
func (up *userPersistance) GetUserByEmail(email string) (result *user.User, err error) {
	fmt.Printf("GET USER BY EMAIL %v -----SEARCHING....\n", email)
	var user user.User
	if result := up.Conn.Where("email = ?", email).First(&user); result.Error != nil {
		fmt.Println("NOT FOUND THE USER!")
		err := result.Error
		return nil, err
	}

	fmt.Printf("....FOUND THE USER! %v\n", user)
	return &user, nil
}

// GetUserByID returns one use, by ID.
func (up *userPersistance) GetUserByID(id int) (result *user.User, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	// defer cancel()

	// query := `select id, email, first_name, last_name, password,
	// 		created_at, updated_at from users where id = $1`

	var user user.User
	if result := up.Conn.Where("id = ?", id).First(&user); result.Error != nil {
		err := result.Error
		return nil, err
	}
	// row := m.DB.QueryRowContext(ctx, query, id)

	// err := row.Scan(
	// 	&user.ID,
	// 	&user.Email,
	// 	&user.FirstName,
	// 	&user.LastName,
	// 	&user.Password,
	// 	&user.CreatedAt,
	// 	&user.UpdatedAt,
	// )
	// result_user, err := dto.AdaptUser(&user)
	// if err != nil {
	// 	return nil, err
	// }

	// if err != nil {
	// 	return nil, err
	// }

	return result, nil
}

func (up *userPersistance) Create(u user.User) error {

	if result := up.Conn.Create(&u); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}
