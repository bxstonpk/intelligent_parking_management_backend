package repository

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) postgresUserRepository {
	return postgresUserRepository{db: db}
}

func (r postgresUserRepository) LoginUser(email_username string, pasword string) (*User, error) {
	// Check user login without password hashing
	user := User{}
	result := r.db.First(&user, "(username = ? OR email = ?) AND password = ?", email_username, email_username, pasword)
	if result.Error != nil {
		log.Printf("Error: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (r postgresUserRepository) GetUser(id int) (*User, error) {
	user := User{}
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Fatalf("Error: %v", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (r postgresUserRepository) RegisterUser(user *User) (int, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		log.Fatalf("Error: %v", result.Error)
		return 0, result.Error
	}
	return 1, nil
}

func (r postgresUserRepository) UpdateUserInfo(user *User) (*User, error) {
	UserResponse := User{}
	result := r.db.Model(User{}).Where("id = ?", user.ID).Updates(user).Scan(&UserResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse, nil
}

func (r postgresUserRepository) UpdateUserPassword(user *User) (*User, error) {
	UserResponse := User{}
	result := r.db.Model(User{}).Where("id = ?", user.ID).Update("password", user.Password).Scan(&UserResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse, nil
}

func (r postgresUserRepository) UpdateUserEmail(user User) (*User, error) {
	UserResponse := User{}
	result := r.db.Model(User{}).Where("id = ?", user.ID).Update("email", user.Email).Scan(&UserResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse, nil
}

func (r postgresUserRepository) UpdateUserUsername(user User) (*User, error) {
	UserResponse := User{}
	/* query := "UPDATE users SET username = $1, update_at = $2 WHERE id = $3 RETURNING *" */
	result := r.db.Model(User{}).Where("id = ?", user.ID).Update("username", user.Username).Scan(&UserResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse, nil
}

func (r postgresUserRepository) UpdateUserProfile(user User) (*User, error) {
	UserResponse := User{}
	/* query := "UPDATE users SET user_profile = $1, update_at = $2 WHERE id = $3 RETURNING *" */
	result := r.db.Model(User{}).Where("id = ?", user.ID).Update("user_profile", user.UserProfile).Scan(&UserResponse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse, nil
}

func (r postgresUserRepository) DeleteUser(id int) (int, error) {
	/* query := "UPDATE users SET is_delete = 1, delete_at = NOW() WHERE id = $1" */
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
