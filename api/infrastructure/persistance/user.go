package persistance

import (
	"profileyou/api/domain/model/user"
	"profileyou/api/domain/repository"
	"profileyou/api/infrastructure/dto"

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
	// ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	// defer cancel()

	// query := `select id, email, first_name, last_name, password,
	// 		created_at, updated_at from users where email = $1`

	var user dto.User
	if result := up.Conn.Where("email = ?", email).First(&user); result.Error != nil {
		err := result.Error
		return nil, err
	}
	result_user, err := dto.AdaptUser(&user)
	if err != nil {
		return nil, err
	}
	// row := kp.Conn.QueryRowContext(ctx, query, email)

	// err := row.Scan(
	// 	&user.ID,
	// 	&user.Email,
	// 	&user.FirstName,
	// 	&user.LastName,
	// 	&user.Password,
	// 	&user.CreatedAt,
	// 	&user.UpdatedAt,
	// )

	if err != nil {
		return nil, err
	}

	return result_user, nil
}

// GetUserByID returns one use, by ID.
func (up *userPersistance) GetUserByID(userId int) (*user.User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	// defer cancel()

	// query := `select id, email, first_name, last_name, password,
	// 		created_at, updated_at from users where id = $1`

	var user dto.User
	if result := up.Conn.Where("user_id = ?", userId).First(&user); result.Error != nil {
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
	result_user, err := dto.AdaptUser(&user)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return result_user, nil
}
