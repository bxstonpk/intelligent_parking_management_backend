package service

// User struct for user login and info
type UserLoginResponse struct {
	ID           int    `db:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	UserFullname string `json:"user_fullname"`
	UserBirthday string `json:"user_birthday"`
	UserGender   int    `json:"user_gender"`
	UserProfile  []byte `json:"user_profile"`
}

type UserRegisterRequese struct {
	ID           int    `db:"id"`
	Email        string `db:"email"`
	Password     string `db:"password"`
	Username     string `db:"username"`
	UserFullname string `db:"user_fullname"`
	UserBirthday string `db:"user_birthday"`
	UserGender   int    `db:"user_gender"`
	CreateAt     string `db:"create_at"`
	UpdateAt     string `db:"update_at"`
	DeleteAt     string `db:"delete_at"`
	IsDelete     int    `db:"is_delete"`
	UserProfile  []byte `db:"user_profile"`
}

type UserService interface {
	LoginUser(string, string) (*UserLoginResponse, error)
}
