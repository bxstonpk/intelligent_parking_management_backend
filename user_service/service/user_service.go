package service

import (
	"database/sql"
	"errors"
	"log"
	"net/mail"
	"time"
	"user_services/repository"

	"github.com/lib/pq"
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
	log.Println(username)
	if username == "" {
		return nil, errors.New("username is empty")
	}
	password := UserRegisterRequest.Password
	log.Println(password)
	if password == "" {
		return nil, errors.New("password is empty")
	}

	user, err := s.userRepo.LoginUser(UserRegisterRequest.Username, UserRegisterRequest.Password)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}

		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		UserFullname: user.UserFullname,
		UserBirthday: user.UserBirthday,
		UserGender:   user.UserGender,
		UserProfile:  user.UserProfile,
	}

	return &UserResponse, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		UserFullname: user.UserFullname,
		UserBirthday: user.UserBirthday,
		UserGender:   user.UserGender,
		UserProfile:  user.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) RegisterUser(UserRegisterRequest UserRegisterRequest) (int, error) {
	user := repository.User{
		Email:        UserRegisterRequest.Email,
		Password:     UserRegisterRequest.Password,
		Username:     UserRegisterRequest.Username,
		UserFullname: UserRegisterRequest.UserFullname,
		UserBirthday: UserRegisterRequest.UserBirthday,
		UserGender:   UserRegisterRequest.UserGender,
		CreateAt:     UserRegisterRequest.CreateAt,
		UserProfile:  UserRegisterRequest.UserProfile,
	}

	status, err := s.userRepo.RegisterUser(&user)
	if err != nil {
		log.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return 0, errors.New("username or email already exists")
			}
		}
		return 0, err
	}
	return status, nil
}

func (s userService) UpdateUserInfo(UserUpdateInfoRequest UserUpdateInfoRequest) (*UserResponse, error) {
	user := repository.User{
		ID:           UserUpdateInfoRequest.ID,
		UserFullname: UserUpdateInfoRequest.UserFullname,
		UserBirthday: UserUpdateInfoRequest.UserBirthday,
		UserGender:   UserUpdateInfoRequest.UserGender,
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserInfo(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  userResponse.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserPassword(UserUpdatePasswordRequest UserUpdatePasswordRequest) (*UserResponse, error) {
	user := repository.User{
		ID:       UserUpdatePasswordRequest.ID,
		Password: UserUpdatePasswordRequest.Password,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserPassword(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  userResponse.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserEmail(UserUpdateEmailRequest UserUpdateEmailRequest) (*UserResponse, error) {
	user := repository.User{
		ID:    UserUpdateEmailRequest.ID,
		Email: UserUpdateEmailRequest.Email,
	}

	// Validate email format
	if _, err := mail.ParseAddress(user.Email); err != nil {
		log.Println(err)
		return nil, errors.New("invalid email format")
	}

	userResponse, err := s.userRepo.UpdateUserEmail(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  userResponse.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserUsername(UserUpdateUsernameRequest UserUpdateUsernameRequest) (*UserResponse, error) {
	user := repository.User{
		ID:       UserUpdateUsernameRequest.ID,
		Username: UserUpdateUsernameRequest.Username,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserUsername(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  userResponse.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) UpdateUserProfile(UserUpdateProfileRequest UserUpdateProfileRequest) (*UserResponse, error) {
	user := repository.User{
		ID:          UserUpdateProfileRequest.ID,
		UserProfile: UserUpdateProfileRequest.UserProfile,
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserProfile(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserResponse{
		ID:           userResponse.ID,
		Email:        userResponse.Email,
		Username:     userResponse.Username,
		UserFullname: userResponse.UserFullname,
		UserBirthday: userResponse.UserBirthday,
		UserGender:   userResponse.UserGender,
		UserProfile:  userResponse.UserProfile,
	}
	return &UserResponse, nil
}

func (s userService) DeleteUser(id int) (int, error) {
	_, err := s.userRepo.DeleteUser(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return 1, nil
}
