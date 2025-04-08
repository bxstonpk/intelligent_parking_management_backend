package service

// User struct for user login and info
type UserLoginRequest struct {
	Password string `db:"password" validate:"required"`
	Username string `db:"username" validate:"required"`
}

type UserRegisterRequest struct {
	Email        string `db:"email" validate:"required,email"`
	Password     string `db:"password" validate:"required"`
	Username     string `db:"username" validate:"required"`
	UserFullname string `db:"user_fullname" validate:"required"`
	UserBirthday string `db:"user_birthday" validate:"required,datetime=2006-01-02"`
	UserGender   int    `db:"user_gender" validate:"required,numeric,min=1,max=2"`
	CreateAt     string `db:"create_at"`
	UserProfile  []byte `db:"user_profile"`
}

type UserUpdateInfoRequest struct {
	ID           int    `db:"id" validate:"required,numeric,min=1"`
	UserFullname string `db:"user_fullname"`
	UserBirthday string `db:"user_birthday"`
	UserGender   int    `db:"user_gender"`
}

type UserUpdatePasswordRequest struct {
	ID       int    `db:"id" validate:"required,numeric,min=1"`
	Password string `db:"password" validate:"required"`
}

type UserUpdateEmailRequest struct {
	ID    int    `db:"id" validate:"required,numeric,min=1"`
	Email string `db:"email" validate:"required,email"`
}

type UserUpdateUsernameRequest struct {
	ID       int    `db:"id" validate:"required,numeric,min=1"`
	Username string `db:"username" validate:"required"`
}

type UserUpdateProfileRequest struct {
	ID          int    `db:"id" validate:"required,numeric,min=1"`
	UserProfile []byte `db:"user_profile" validate:"required"`
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
	LoginUser(*UserLoginRequest) (*UserResponse, error)
	GetUser(int) (*UserResponse, error)
	RegisterUser(UserRegisterRequest) (int, error)
	UpdateUserInfo(UserUpdateInfoRequest) (*UserResponse, error)
	UpdateUserPassword(UserUpdatePasswordRequest) (*UserResponse, error)
	UpdateUserEmail(UserUpdateEmailRequest) (*UserResponse, error)
	UpdateUserUsername(UserUpdateUsernameRequest) (*UserResponse, error)
	UpdateUserProfile(UserUpdateProfileRequest) (*UserResponse, error)
	DeleteUser(int) (int, error)
}
