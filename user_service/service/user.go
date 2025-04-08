package service

// User struct for user login and info
type UserLoginRequest struct {
	Password string `validate:"required" json:"password"`
	Username string `json:"username"`
	Email    string `db:"email" json:"email"`
}

type UserRegisterRequest struct {
	Email        string `validate:"required,email" json:"email"`
	Password     string `validate:"required" json:"password"`
	Username     string `validate:"required" json:"username"`
	UserFullname string `validate:"required" json:"user_fullname"`
	UserBirthday string `validate:"required,datetime=2006-01-02" json:"user_birthday"`
	UserGender   int    `validate:"required,numeric,min=1,max=2" json:"user_gender"`
	CreateAt     string
	UserProfile  string `json:"user_profile"`
}

type UserUpdateInfoRequest struct {
	ID           int    `validate:"required,numeric,min=1" json:"id"`
	UserFullname string `json:"user_fullname"`
	UserBirthday string `json:"user_birthday"`
	UserGender   int    `json:"user_gender"`
}

type UserUpdatePasswordRequest struct {
	ID       int    `validate:"required,numeric,min=1" json:"id"`
	Password string `validate:"required" json:"password"`
}

type UserUpdateEmailRequest struct {
	ID    int    `validate:"required,numeric,min=1" json:"id"`
	Email string `validate:"required,email" json:"email"`
}

type UserUpdateUsernameRequest struct {
	ID       int    `validate:"required,numeric,min=1" json:"id"`
	Username string `validate:"required" json:"username"`
}

type UserUpdateProfileRequest struct {
	ID          int    `validate:"required,numeric,min=1" json:"id"`
	UserProfile string `validate:"required" json:"user_profile"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	UserFullname string `json:"user_fullname"`
	UserBirthday string `json:"user_birthday"`
	UserGender   int    `json:"user_gender"`
	UserProfile  string `json:"user_profile"`
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
