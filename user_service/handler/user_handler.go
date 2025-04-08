package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"user_services/security"
	"user_services/service"

	"github.com/gorilla/mux"
)

func (h userHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var UserLoginRequese service.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&UserLoginRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(&UserLoginRequese)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate JWT token here if needed and add it to the response
	token, err := security.NewBcryptHasher(h.secretKey).GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
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
	var UserRegisterRequese service.UserRegisterRequest
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

func (h userHandler) UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])
	var UserUpdateInfoRequese service.UserUpdateInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateInfoRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateInfoRequese.ID = id

	user, err := h.userService.UpdateUserInfo(UserUpdateInfoRequese)
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

func (h userHandler) UpdateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])
	var UserUpdatePasswordRequese service.UserUpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdatePasswordRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdatePasswordRequese.ID = id

	user, err := h.userService.UpdateUserPassword(UserUpdatePasswordRequese)
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

func (h userHandler) UpdateUserEmailHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])
	var UserUpdateEmailRequese service.UserUpdateEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateEmailRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateEmailRequese.ID = id

	user, err := h.userService.UpdateUserEmail(UserUpdateEmailRequese)
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

func (h userHandler) UpdateUserUsernameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])
	var UserUpdateUsernameRequese service.UserUpdateUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateUsernameRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateUsernameRequese.ID = id

	user, err := h.userService.UpdateUserUsername(UserUpdateUsernameRequese)
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

func (h userHandler) UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])
	var UserUpdateProfileRequese service.UserUpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateProfileRequese); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateProfileRequese.ID = id

	user, err := h.userService.UpdateUserProfile(UserUpdateProfileRequese)
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

func (h userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	status, err := h.userService.DeleteUser(id)
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
