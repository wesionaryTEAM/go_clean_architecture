package errorz

import "net/http"

// Canonical base errors
var (
	ErrBadRequest         = New("BAD_REQUEST", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrUnauthorized       = New("UNAUTHORIZED", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	ErrForbidden          = New("FORBIDDEN", http.StatusForbidden, http.StatusText(http.StatusForbidden))
	ErrNotFound           = New("NOT_FOUND", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	ErrConflict           = New("CONFLICT", http.StatusConflict, http.StatusText(http.StatusConflict))
	ErrUnprocessable      = New("UNPROCESSABLE", http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
	ErrInternal           = New("INTERNAL_ERROR", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	ErrServiceUnavailable = New("SERVICE_UNAVAILABLE", http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
)

// Derived semantic errors
var (
	ErrAlreadyExists      = New("ALREADY_EXISTS", http.StatusConflict, http.StatusText(http.StatusConflict))
	ErrSomethingWentWrong = New("SOMETHING_WENT_WRONG", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
)
