package user

import (
	"clean-architecture/pkg/errorz"
	"net/http"
)

var (
	ErrInvalidUserID     = errorz.New("INVALID_USER_ID", http.StatusBadRequest, "Invalid user ID")
	ErrUserAlreadyExists = errorz.New("USER_ALREADY_EXISTS", http.StatusConflict, "User already exists")
)
