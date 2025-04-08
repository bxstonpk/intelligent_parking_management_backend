package repository

type User struct {
	ID           int    `db:"id"`
	Email        string `db:"email" validate:"required,email"`
	Password     string `db:"password" validate:"required"`
	Username     string `db:"username" validate:"required"`
	UserFullname string `db:"user_fullname" validate:"required"`
	UserBirthday string `db:"user_birthday" validate:"required,datetime=2006-01-02"`
	UserGender   int    `db:"user_gender" validate:"required, numeric,min=1,max=2"`
	CreateAt     string `db:"create_at" validate:"required,datetime=2006-01-02 15:04:05"`
	UpdateAt     string `db:"update_at"`
	DeleteAt     string `db:"delete_at"`
	IsDelete     string `db:"is_delete"`
	UserProfile  []byte `db:"user_profile"`
}

type UserRepository interface {
	LoginUser(string, string) (*User, error)
	GetUser(int) (*User, error)
	RegisterUser(*User) (int, error)
	UpdateUserInfo(*User) (*User, error)
	UpdateUserPassword(*User) (*User, error)
	UpdateUserEmail(User) (*User, error)
	UpdateUserUsername(User) (*User, error)
	UpdateUserProfile(User) (*User, error)
	DeleteUser(int) (int, error)
}
