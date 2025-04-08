package handler

import (
	"user_services/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}
