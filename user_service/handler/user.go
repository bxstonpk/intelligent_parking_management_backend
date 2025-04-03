package handler

import (
	"user_services/service"
)

type userHandler struct {
	userService service.UserService
	secretKey   string
}

func NewUserHandler(userService service.UserService, secretKey string) userHandler {
	return userHandler{userService: userService, secretKey: secretKey}
}
