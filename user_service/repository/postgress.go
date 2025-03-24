package repository

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) postgresUserRepository {
	return postgresUserRepository{db: db}
}

func (r postgresUserRepository) LoginUser(email_username string, pasword string) (*User, error) {
	// Check user login without password hashing
	user := User{}
	query := "SELECT * FROM users WHERE (email = $1 OR username = $1) AND password = $2"
	err := r.db.Get(&user, query, email_username, pasword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r postgresUserRepository) GetUser(id int) (*User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE id = $1"
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r postgresUserRepository) RegisterUser(user User) (int, error) {
	query := "INSERT INTO users (email, password, username, user_fullname, user_birthday, user_gender, create_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	result, err := r.db.Exec(query, user.Email, user.Password, user.Username, user.UserFullname, user.UserBirthday, user.UserGender, user.CreateAt)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (r postgresUserRepository) UpdateUserInfo(user User) (*User, error) {
	UpdateUser := User{}
	query := "UPDATE users SET user_fullname = $1, user_birthday = $2, user_gender = $3, update_at = $4 WHERE id = $5"
	err := r.db.Get(&UpdateUser, query, user.UserFullname, user.UserBirthday, user.UserGender, user.UpdateAt, user.ID)
	if err != nil {
		return nil, err
	}
	return &UpdateUser, nil
}

func (r postgresUserRepository) UpdateUserPassword(user User) (*User, error) {
	UpdateUser := User{}
	query := "UPDATE users SET password = $1, update_at = $2 WHERE id = $3"
	err := r.db.Get(&UpdateUser, query, user.Password, user.UpdateAt, user.ID)
	if err != nil {
		return nil, err
	}
	return &UpdateUser, nil
}

func (r postgresUserRepository) UpdateUserEmail(user User) (*User, error) {
	Updateuser := User{}
	query := "UPDATE users SET email = $1, update_at = $2 WHERE id = $3"
	err := r.db.Get(&Updateuser, query, user.Email, user.UpdateAt, user.ID)
	if err != nil {
		return nil, err
	}
	return &Updateuser, nil
}

func (r postgresUserRepository) UpdateUserUsername(user User) (*User, error) {
	UpdateUser := User{}
	query := "UPDATE users SET username = $1, update_at = $2 WHERE id = $3"
	err := r.db.Get(&UpdateUser, query, user.Username, user.UpdateAt, user.ID)
	if err != nil {
		return nil, err
	}
	return &UpdateUser, nil
}

func (r postgresUserRepository) DeleteUser(id int) (int, error) {
	query := "UPDATE users SET is_delete = 1, delete_at = NOW() WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
