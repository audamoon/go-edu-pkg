package reqerror

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
	Code   int
	Status string `json:"status"`
	Error  string `json:"error"`
}

var ErrorsMessage = map[error]*ErrorResponse{
	ErrResNotFound:   {http.StatusNotFound, "error", "Resource not found"},
	ErrBadRequest:    {http.StatusBadRequest, "error", "Bad request"},
	ErrInternalError: {http.StatusInternalServerError, "error", "Something goes wrong"},
}
