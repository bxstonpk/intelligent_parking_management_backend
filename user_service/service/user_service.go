package service

import (
	"log"
	"user_services/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) LoginUser(username string, password string) (*UserLoginResponse, error) {
	user, err := s.userRepo.LoginUser(username, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	UserResponse := UserLoginResponse{
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
