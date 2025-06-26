package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"user_services/errs"
	"user_services/logs"
	"user_services/security"
	"user_services/service"

	"github.com/gorilla/mux"
)

func (h userHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var UserLoginRequese service.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&UserLoginRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(&UserLoginRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	// Generate JWT token here if needed and add it to the response
	token, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		logs.Error("Failed to generate JWT token: ")
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		logs.Error("Failed to encode user response: ")
		handlerError(w, err)
		return
	}
}

func (h userHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

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
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := h.userService.RegisterUser(UserRegisterRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) UpdateUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	var UserUpdateInfoRequese service.UserUpdateInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateInfoRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateInfoRequese.ID = id

	user, err := h.userService.UpdateUserInfo(UserUpdateInfoRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) UpdateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	var UserUpdatePasswordRequese service.UserUpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdatePasswordRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdatePasswordRequese.ID = id

	user, err := h.userService.UpdateUserPassword(UserUpdatePasswordRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) UpdateUserEmailHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	var UserUpdateEmailRequese service.UserUpdateEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateEmailRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateEmailRequese.ID = id

	user, err := h.userService.UpdateUserEmail(UserUpdateEmailRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) UpdateUserUsernameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	var UserUpdateUsernameRequese service.UserUpdateUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateUsernameRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateUsernameRequese.ID = id

	user, err := h.userService.UpdateUserUsername(UserUpdateUsernameRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Check authentication
	status := authenticationHandler(r)
	if !status {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	var UserUpdateProfileRequese service.UserUpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&UserUpdateProfileRequese); err != nil {
		logs.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserUpdateProfileRequese.ID = id

	user, err := h.userService.UpdateUserProfile(UserUpdateProfileRequese)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["userId"])

	// Authentication check
	s := authenticationHandler(r)
	if !s {
		handlerError(w, errs.NewForbiddenError("forbidden"))
		logs.Error("Authentication failed")
		return
	}

	status, err := h.userService.DeleteUser(id)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		handlerError(w, err)
		return
	}
}

func (h userHandler) CheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		handlerError(w, errs.NewBadRequestError("missing token"))
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	valid, err := service.NewSecurityService(os.Getenv("SECRET_KEY")).CheckToken(token)
	if err != nil {
		handlerError(w, err)
		return
	}

	if !valid {
		handlerError(w, errs.NewBadRequestError("invalid token"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token is valid"))
}
