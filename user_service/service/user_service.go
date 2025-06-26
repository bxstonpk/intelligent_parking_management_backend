package service

import (
	"errors"
	"net/mail"
	"os"
	"strings"
	"time"
	"user_services/errs"
	"user_services/logs"
	"user_services/repository"
	"user_services/security"

	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) LoginUser(UserRegisterRequest *UserLoginRequest) (*UserResponse, error) {
	// Check user
	username := UserRegisterRequest.Username
	email := UserRegisterRequest.Email
	if username == "" && email == "" {
		logs.Error("username or email is empty")
		return nil, errs.NewBadRequestError("username or email is empty")
	}
	password := UserRegisterRequest.Password
	if password == "" {
		logs.Error("password is empty")
		return nil, errs.NewBadRequestError("password is empty")
	}

	// Check Empty email or usernaem
	userCheck := ""
	if UserRegisterRequest.Username != "" {
		userCheck = UserRegisterRequest.Username
	} else {
		userCheck = UserRegisterRequest.Email
	}

	user, err := s.userRepo.LoginUser(userCheck, UserRegisterRequest.Password)
	if err != nil {
		logs.Error(err)

		UserResponse := UserResponse{}

		if err == gorm.ErrRecordNotFound {
			return &UserResponse, errs.NewNotFoundError("user not found")
		}
		return &UserResponse, errs.NewUnexpectedError()
	}

	// Check password
	secretKey := os.Getenv("SECRET_KEY")
	if status := security.NewBcryptHasher(secretKey).CheckPasswordHash(UserRegisterRequest.Password, user.Password); !status {
		logs.Error(err)

		UserResponse := UserResponse{
			Message: "0",
		}

		return &UserResponse, errs.NewBadRequestError("invalid password")
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		UserFullname: user.UserFullname,
		UserBirthday: user.UserBirthday,
		UserGender:   user.UserGender,
		UserProfile:  string(user.UserProfile),
	}

	return &UserResponse, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		logs.Error(err)

		UserResponse := UserResponse{}

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return &UserResponse, errs.NewNotFoundError("user not found")
		}

		return &UserResponse, errs.NewUnexpectedError()
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		UserFullname: user.UserFullname,
		UserBirthday: user.UserBirthday,
		UserGender:   user.UserGender,
		UserProfile:  string(user.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) RegisterUser(UserRegisterRequest UserRegisterRequest) (int, error) {
	user := repository.Users{
		Email:        UserRegisterRequest.Email,
		Password:     UserRegisterRequest.Password,
		Username:     UserRegisterRequest.Username,
		UserFullname: UserRegisterRequest.UserFullname,
		UserBirthday: UserRegisterRequest.UserBirthday,
		UserGender:   UserRegisterRequest.UserGender,
		CreateAt:     UserRegisterRequest.CreateAt,
		UserProfile:  []byte(UserRegisterRequest.UserProfile),
	}

	// Validate email format
	if _, err := mail.ParseAddress(user.Email); err != nil {
		logs.Error(err)
		return 0, errs.NewBadRequestError("invalid email format")
	}
	// Validate username format
	if len(user.Username) < 3 || len(user.Username) > 20 {
		logs.Error("username length error")
		return 0, errs.NewBadRequestError("username must be between 5 and 20 characters")
	}
	// Validate password format
	if len(user.Password) < 6 || len(user.Password) > 20 {
		logs.Error("password length error")
		return 0, errs.NewBadRequestError("password must be between 8 and 20 characters")
	}
	// Validate userfullname userbirtday usergender is not empty
	if user.UserFullname == "" {
		logs.Error("userfullname is empty")
		return 0, errs.NewBadRequestError("userfullname is empty")
	}
	if user.UserBirthday == "" {
		logs.Error("user birthday is empty")
		return 0, errs.NewBadRequestError("user birthday is empty")
	}
	if user.UserProfile == nil {
		logs.Error("user profile is empty")
		return 0, errs.NewBadRequestError("user profile is empty")
	}
	if user.UserGender >= 0 || user.UserGender <= 2 {
		// Valid gender
		// 0 is female, 1 is male, 2 is other
	} else {
		logs.Error("user gender is not valid")
		return 0, errs.NewBadRequestError("user gender is not valid")
	}

	// hash password
	hashedPassword, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).HashPassword(user.Password)
	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}
	user.Password = hashedPassword

	status, err := s.userRepo.RegisterUser(&user)
	if err != nil {
		logs.Error(err)

		if strings.Contains(err.Error(), "uni_users_email") {
			logs.Error(err)
			return 0, errs.NewConflictError("email or username already exists")
		}

		return 0, errs.NewUnexpectedError()
	}
	return status, nil
}

func (s userService) UpdateUserInfo(UserUpdateInfoRequest UserUpdateInfoRequest) (*UserResponse, error) {
	user := repository.Users{
		ID:           UserUpdateInfoRequest.ID,
		UserFullname: UserUpdateInfoRequest.UserFullname,
		UserBirthday: UserUpdateInfoRequest.UserBirthday,
		UserGender:   UserUpdateInfoRequest.UserGender,
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	// Validate userfullname userbirtday usergender is not empty
	if user.UserFullname == "" {
		logs.Error("userfullname is empty")
		return nil, errs.NewBadRequestError("userfullname is empty")
	}
	if user.UserBirthday == "" {
		logs.Error("user birthday is empty")
		return nil, errs.NewBadRequestError("user birthday is empty")
	}
	if user.UserGender >= 0 || user.UserGender <= 2 {
		// Valid gender
		// 0 is female, 1 is male, 2 is other
	} else {
		logs.Error("user gender is not valid")
		return nil, errs.NewBadRequestError("user gender is not valid")
	}

	userResponse, err := s.userRepo.UpdateUserInfo(&user)
	if err != nil {
		logs.Error(err)
		userResponse := UserResponse{
			Message: "0",
		}

		return &userResponse, errs.NewBadRequestError("invalid user id")
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  string(userResponse.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserPassword(UserUpdatePasswordRequest UserUpdatePasswordRequest) (*UserResponse, error) {
	user := repository.Users{
		ID:       UserUpdatePasswordRequest.ID,
		Password: UserUpdatePasswordRequest.Password,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Validate password format
	if len(user.Password) < 6 || len(user.Password) > 20 {
		logs.Error("password length error")
		return nil, errs.NewBadRequestError("password must be between 8 and 20 characters")
	}

	// hash password
	hashPassword, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).HashPassword(user.Password)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	user.Password = hashPassword

	userResponse, err := s.userRepo.UpdateUserPassword(&user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  string(userResponse.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserEmail(UserUpdateEmailRequest UserUpdateEmailRequest) (*UserResponse, error) {
	user := repository.Users{
		ID:    UserUpdateEmailRequest.ID,
		Email: UserUpdateEmailRequest.Email,
	}

	// Validate email format
	if _, err := mail.ParseAddress(user.Email); err != nil {
		logs.Error(err)
		return nil, errors.New("invalid email format")
	}

	userResponse, err := s.userRepo.UpdateUserEmail(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  string(userResponse.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserUsername(UserUpdateUsernameRequest UserUpdateUsernameRequest) (*UserResponse, error) {
	user := repository.Users{
		ID:       UserUpdateUsernameRequest.ID,
		Username: UserUpdateUsernameRequest.Username,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Validate username format
	if len(user.Username) < 3 || len(user.Username) > 20 {
		logs.Error("username length error")
		return nil, errs.NewBadRequestError("username must be between 5 and 20 characters")
	}

	userResponse, err := s.userRepo.UpdateUserUsername(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  string(userResponse.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserProfile(UserUpdateProfileRequest UserUpdateProfileRequest) (*UserResponse, error) {
	user := repository.Users{
		ID:          UserUpdateProfileRequest.ID,
		UserProfile: []byte(UserUpdateProfileRequest.UserProfile),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserProfile(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	UserResponse := UserResponse{
		Message:      "1",
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  string(userResponse.UserProfile),
	}
	return &UserResponse, nil
}

func (s userService) DeleteUser(id int) (int, error) {
	_, err := s.userRepo.DeleteUser(id)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	return 1, nil
}

func (s userService) CheckToken(token string) (bool, error) {
	if token == "" {
		logs.Error("token is empty")
		return false, errs.NewBadRequestError("token is empty")
	}

	// Check if the token is valid
	_, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).ValidateToken(token)
	if err != nil {
		logs.Error("invalid token")
		return false, errs.NewBadRequestError("invalid token")
	}
	return true, nil
}
