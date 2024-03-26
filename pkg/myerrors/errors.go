package myerrors

import (
	"errors"
	"net/http"
)

var (
	ErrResNotFound   = errors.New("resource not found")
	ErrBadRequest    = errors.New("bad request")
	ErrInternalError = errors.New("something goes wrong")
)

type ErrorResponse struct {
	code   int
	Status string `json:"status"`
	Error  string `json:"error"`
}

var ErrorsMessage = map[error]*ErrorResponse{
	ErrResNotFound:   &ErrorResponse{http.StatusNotFound, "error", "Resource not found"},
	ErrBadRequest:    &ErrorResponse{http.StatusBadRequest, "error", "Bad request"},
	ErrInternalError: &ErrorResponse{http.StatusInternalServerError, "error", "Something goes wrong"},
}