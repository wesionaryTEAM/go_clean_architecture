package errorz

import (
	"fmt"
	"net/http"
)

var (
	ErrBadRequest         = NewAPIError(http.StatusBadRequest, "Bad Request")
	ErrUnauthorized       = NewAPIError(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden          = NewAPIError(http.StatusForbidden, "Forbidden")
	ErrNotFound           = NewAPIError(http.StatusNotFound, "Not Found")
	ErrConflict           = NewAPIError(http.StatusConflict, "Conflict")
	ErrUnprocessable      = NewAPIError(http.StatusUnprocessableEntity, "Unable to process the contained instructions")
	ErrInternal           = NewAPIError(http.StatusInternalServerError, "Internal Server Error")
	ErrServiceUnavailable = NewAPIError(http.StatusServiceUnavailable, "Service Unavailable")
	ErrAlreadyExists      = JoinError("Already Exists", ErrConflict)
	ErrSomethingWentWrong = JoinError("something went wrong", ErrInternal)
)

func JoinError(message string, base error) error {
	if base.Error() == "" {
		return fmt.Errorf("%v%w", message, base)
	}
	return fmt.Errorf("%v %w", message, base)
}
