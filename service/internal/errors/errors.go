package api

import (
	"net/http"
)

var unknownError = &Error{Code: http.StatusInternalServerError, Status: "error", Message: "an unknown error occurred"}

type Error struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Status:  "error",
		Message: message,
	}
}
