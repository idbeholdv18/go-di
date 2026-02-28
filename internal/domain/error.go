package domain

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

var (
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)

func NewUnauthorizedError() *Error {
	return &Error{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
		Code:    ErrUnauthorized.Error(),
	}
}

func SendError(w http.ResponseWriter, r *http.Request, errorFabric func() *Error) {
	json.NewEncoder(w).Encode(errorFabric())
}
