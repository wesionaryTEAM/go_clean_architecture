package user

import (
	"clean-architecture/pkg/errorz"
	"net/http"
)

var (
	ErrInvalidUserID     = errorz.New(errorz.CodeInvalidUserID, http.StatusBadRequest, "Invalid user ID")
	ErrUserAlreadyExists = errorz.New(errorz.CodeUserAlreadyExists, http.StatusConflict, "User already exists")
)
