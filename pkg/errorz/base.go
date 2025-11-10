package errorz

import "net/http"

// Canonical base errors
var (
	ErrBadRequest         = New(CodeBadRequest, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrUnauthorized       = New(CodeUnauthorized, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	ErrForbidden          = New(CodeForbidden, http.StatusForbidden, http.StatusText(http.StatusForbidden))
	ErrNotFound           = New(CodeNotFound, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	ErrConflict           = New(CodeConflict, http.StatusConflict, http.StatusText(http.StatusConflict))
	ErrUnprocessable      = New(CodeUnprocessable, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
	ErrInternal           = New(CodeInternalError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	ErrServiceUnavailable = New(CodeServiceUnavailable, http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
)

// Derived semantic errors
var (
	ErrAlreadyExists      = New(CodeAlreadyExists, http.StatusConflict, http.StatusText(http.StatusConflict))
	ErrSomethingWentWrong = New(CodeSomethingWentWrong, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
)
