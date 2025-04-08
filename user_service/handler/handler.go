package handler

import (
	"fmt"
	"net/http"
	"user_services/errs"
	"user_services/logs"
	"user_services/middleware"

	"github.com/gorilla/mux"
)

func handlerError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errs.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}

func authenticationHandler(r *http.Request) bool {
	// user_id from JWT context
	userIDFromJWT, ok := r.Context().Value(middleware.UserIDKey).(string)
	logs.Info("userIDFromJWT: " + userIDFromJWT)
	if !ok {
		logs.Error("Failed to get userID from JWT context")
		return false
	}

	// user_id from URL
	vars := mux.Vars(r)
	userIDFromRequest := vars["userId"]

	// Authentication: user from JWT equal URL Request
	if userIDFromJWT != userIDFromRequest {
		logs.Error("User ID from JWT does not match user ID from request")
		return false
	}
	return true
}
