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

	"github.com/golang-jwt/jwt"
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
	token, refreshToken, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).GenerateJWT(strconv.Itoa(user.ID))
	if err != nil {
		logs.Error("Failed to generate JWT token: ")
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Refresh-Token", refreshToken)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		logs.Error("Failed to encode user response: ")
		handlerError(w, err)
		return
	}
}

func (h userHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Extract the refresh token from the request
	refreshToken := r.Header.Get("Refresh-Token")

	// Validate the refresh token
	if refreshToken == "" {
		http.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}

	// Verify the refresh token
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		// Use the secret key to validate the token
		return refreshToken, nil
	})
	if err != nil || !token.Valid {
		logs.Error("Invalid refresh token")
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	clams := token.Claims.(jwt.MapClaims)
	userId, ok := clams["user_id"].(string)
	if !ok {
		logs.Error("Invalid token claims")
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}
	// Generate a new JWT token
	newToken, _, err := security.NewBcryptHasher(os.Getenv("SECRET_KEY")).GenerateJWT(userId)
	if err != nil {
		logs.Error("Failed to generate new token: ")
		http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+newToken)
	if err := json.NewEncoder(w).Encode(newToken); err != nil {
		logs.Error("Failed to encode new token response: ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
