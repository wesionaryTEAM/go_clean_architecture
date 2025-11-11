package errorz

import "net/http"

var (
	ErrUnauthorizedAccess        = New(CodeUnauthorizedAccess, ErrUnauthorized.StatusCode, http.StatusText(ErrUnauthorized.StatusCode))
	ErrForbiddenAccess           = New(CodeForbiddenAccess, ErrForbidden.StatusCode, http.StatusText(ErrForbidden.StatusCode))
	ErrInvalidToken              = New(CodeInvalidToken, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrInvalidUUID               = New(CodeInvalidUUID, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrRecordNotFound            = New(CodeRecordNotFound, ErrNotFound.StatusCode, http.StatusText(ErrNotFound.StatusCode))
	ErrInvalidUserNameOrPassword = New(CodeInvalidUserNameOrPassword, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrExtensionMismatch         = New(CodeExtensionMismatch, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrThumbExtensionMismatch    = New(CodeThumbExtensionMismatch, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrFileRead                  = New(CodeFileRead, ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
)
