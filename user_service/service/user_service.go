package service

import (
	"database/sql"
	"errors"
	"log"
	"time"
	"user_services/repository"

	"github.com/lib/pq"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) LoginUser(UserRegisterRequese *UserLoginRequese) (*UserResponse, error) {
	user, err := s.userRepo.LoginUser(UserRegisterRequese.Username, UserRegisterRequese.Password)
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

func (s userService) RegisterUser(UserRegisterRequese UserRegisterRequese) (int, error) {
	user := repository.User{
		Email:        UserRegisterRequese.Email,
		Password:     UserRegisterRequese.Password,
		Username:     UserRegisterRequese.Username,
		UserFullname: UserRegisterRequese.UserFullname,
		UserBirthday: UserRegisterRequese.UserBirthday,
		UserGender:   UserRegisterRequese.UserGender,
		CreateAt:     UserRegisterRequese.CreateAt,
		UserProfile:  UserRegisterRequese.UserProfile,
	}

	status, err := s.userRepo.RegisterUser(user)
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

func (s userService) UpdateUserInfo(UserUpdateInfoRequese UserUpdateInfoRequese) (*UserResponse, error) {
	user := repository.User{
		ID:           UserUpdateInfoRequese.ID,
		UserFullname: UserUpdateInfoRequese.UserFullname,
		UserBirthday: UserUpdateInfoRequese.UserBirthday,
		UserGender:   UserUpdateInfoRequese.UserGender,
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserInfo(user)
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

func (s userService) UpdateUserPassword(UserUpdatePasswordRequese UserUpdatePasswordRequese) (*UserResponse, error) {
	user := repository.User{
		ID:       UserUpdatePasswordRequese.ID,
		Password: UserUpdatePasswordRequese.Password,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	userResponse, err := s.userRepo.UpdateUserPassword(user)
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

func (s userService) UpdateUserEmail(UserUpdateEmailRequese UserUpdateEmailRequese) (*UserResponse, error) {
	user := repository.User{
		ID:    UserUpdateEmailRequese.ID,
		Email: UserUpdateEmailRequese.Email,
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

func (s userService) UpdateUserUsername(UserUpdateUsernameRequese UserUpdateUsernameRequese) (*UserResponse, error) {
	user := repository.User{
		ID:       UserUpdateUsernameRequese.ID,
		Username: UserUpdateUsernameRequese.Username,
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

func (s userService) UpdateUserProfile(UserUpdateProfileRequese UserUpdateProfileRequese) (*UserResponse, error) {
	user := repository.User{
		ID:          UserUpdateProfileRequese.ID,
		UserProfile: UserUpdateProfileRequese.UserProfile,
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
