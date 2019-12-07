package errors

import (
	"net/http"
)

var UnknownError = &Error{Code: http.StatusInternalServerError, Status: "error", Message: "an unknown error occurred"}

type Error struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func New(code int, cause error) *Error {
	return &Error{
		Code:    code,
		Status:  "error",
		Message: cause.Error(),
	}
}
