package repository

type Users struct {
	ID           int    `db:"id" gorm:"primaryKey" validate:"required"`
	Email        string `db:"email" gorm:"unique" validate:"required,email"`
	Password     string `db:"password" validate:"required"`
	Username     string `db:"username" gorm:"unique" validate:"required"`
	UserFullname string `db:"user_fullname" validate:"required"`
	UserBirthday string `db:"user_birthday" validate:"required,datetime=2006-01-02"`
	UserGender   int    `db:"user_gender" validate:"required, numeric,min=1,max=2"`
	CreateAt     string `db:"create_at" validate:"required,datetime=2006-01-02 15:04:05"`
	UpdateAt     string `db:"update_at"`
	DeleteAt     string `db:"delete_at"`
	UserProfile  []byte `db:"user_profile"`
}

type UserRepository interface {
	LoginUser(string, string) (*Users, error)
	GetUser(int) (*Users, error)
	RegisterUser(*Users) (int, error)
	UpdateUserInfo(*Users) (*Users, error)
	UpdateUserPassword(*Users) (*Users, error)
	UpdateUserEmail(Users) (*Users, error)
	UpdateUserUsername(Users) (*Users, error)
	UpdateUserProfile(Users) (*Users, error)
	DeleteUser(int) (int, error)
}
