package errs

import "net/http"

type AppError struct {
	Code int
	Msg  string
}

func (e AppError) Error() string {
	return e.Msg
}

func NewBadRequestError(msg string) error {
	return AppError{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}
}

func NewNotFoundError(msg string) error {
	return AppError{
		Code: http.StatusNotFound,
		Msg:  msg,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code: http.StatusInternalServerError,
		Msg:  "unexpected error",
	}
}

func NewForbiddenError(msg string) error {
	return AppError{
		Code: http.StatusForbidden,
		Msg:  msg,
	}
}

func NewConflictError(msg string) error {
	return AppError{
		Code: http.StatusConflict,
		Msg:  msg,
	}
}
