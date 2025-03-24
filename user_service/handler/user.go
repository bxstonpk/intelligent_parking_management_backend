package handler

import (
	"net/http"
	"user_services/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	println("LoginUser")
}
