package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"user_services/service"

	"github.com/gorilla/mux"
)

func (h userHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var UserLoginRequese service.UserLoginRequese
	if err := json.NewDecoder(r.Body).Decode(&UserLoginRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(&UserLoginRequese)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	user, err := h.userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h userHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var UserRegisterRequese service.UserRegisterRequese
	if err := json.NewDecoder(r.Body).Decode(&UserRegisterRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.userService.RegisterUser(UserRegisterRequese)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
