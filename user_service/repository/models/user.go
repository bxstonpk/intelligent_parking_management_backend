package repository

type User struct {
	ID           int    `db:"id"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	UserFullname string `db:"user_fullname"`
	UserBirthday string `db:"user_birthday"`
	UserGender   int    `db:"user_gender"`
	CreateAt     string `db:"create_at"`
	UpdateAt     string `db:"update_at"`
	DeleteAt     string `db:"delete_at"`
	IsDelete     int    `db:"is_delete"`
}

type UserRepository interface {
	LoginUser(string, string) (*User, error)
	GetUser(int) (*User, error)
	RegisterUser(User) error
	UpdateUser(User) error
	DeleteUser(int) error
}
