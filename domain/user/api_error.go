package user

import (
	"clean-architecture/pkg/errorz"
	"net/http"
)

var (
	ErrInvalidUserID = errorz.NewAPIError(http.StatusBadRequest, "Invalid user ID")
)
