package errorz

import "net/http"

// Canonical base errors — message defaults to the standard HTTP status text.
var (
	ErrBadRequest         = New(CodeBadRequest, http.StatusBadRequest, SeverityWarn, http.StatusText(http.StatusBadRequest))
	ErrUnauthorized       = New(CodeUnauthorized, http.StatusUnauthorized, SeverityWarn, http.StatusText(http.StatusUnauthorized))
	ErrForbidden          = New(CodeForbidden, http.StatusForbidden, SeverityWarn, http.StatusText(http.StatusForbidden))
	ErrNotFound           = New(CodeNotFound, http.StatusNotFound, SeverityWarn, http.StatusText(http.StatusNotFound))
	ErrConflict           = New(CodeConflict, http.StatusConflict, SeverityWarn, http.StatusText(http.StatusConflict))
	ErrUnprocessable      = New(CodeUnprocessable, http.StatusUnprocessableEntity, SeverityWarn, http.StatusText(http.StatusUnprocessableEntity))
	ErrInternal           = New(CodeInternalError, http.StatusInternalServerError, SeverityError, http.StatusText(http.StatusInternalServerError))
	ErrServiceUnavailable = New(CodeServiceUnavailable, http.StatusServiceUnavailable, SeverityError, http.StatusText(http.StatusServiceUnavailable))
)

// Derived semantic errors
var (
	ErrAlreadyExists      = New(CodeAlreadyExists, http.StatusConflict, SeverityWarn, http.StatusText(http.StatusConflict))
	ErrSomethingWentWrong = New(CodeSomethingWentWrong, http.StatusInternalServerError, SeverityError, http.StatusText(http.StatusInternalServerError))
)
