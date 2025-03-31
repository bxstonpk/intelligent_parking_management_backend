package service

// User struct for user login and info
type UserLoginRequese struct {
	Password string `db:"password"`
	Username string `db:"username"`
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
	UserProfile  []byte `db:"user_profile"`
}

type UserUpdateInfoRequese struct {
	ID           int    `db:"id"`
	UserFullname string `db:"user_fullname"`
	UserBirthday string `db:"user_birthday"`
	UserGender   int    `db:"user_gender"`
}

type UserUpdatePasswordRequese struct {
	ID       int    `db:"id"`
	Password string `db:"password"`
}

type UserUpdateEmailRequese struct {
	ID    int    `db:"id"`
	Email string `db:"email"`
}

type UserUpdateUsernameRequese struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
}

type UserUpdateProfileRequese struct {
	ID          int    `db:"id"`
	UserProfile []byte `db:"user_profile"`
}

type UserResponse struct {
	ID           int    `db:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	UserFullname string `json:"user_fullname"`
	UserBirthday string `json:"user_birthday"`
	UserGender   int    `json:"user_gender"`
	UserProfile  []byte `json:"user_profile"`
}

type UserService interface {
	LoginUser(*UserLoginRequese) (*UserResponse, error)
	GetUser(int) (*UserResponse, error)
	RegisterUser(UserRegisterRequese) (int, error)
	UpdateUserInfo(UserUpdateInfoRequese) (*UserResponse, error)
	UpdateUserPassword(UserUpdatePasswordRequese) (*UserResponse, error)
	UpdateUserEmail(UserUpdateEmailRequese) (*UserResponse, error)
	UpdateUserUsername(UserUpdateUsernameRequese) (*UserResponse, error)
	UpdateUserProfile(UserUpdateProfileRequese) (*UserResponse, error)
	DeleteUser(int) (int, error)
}
