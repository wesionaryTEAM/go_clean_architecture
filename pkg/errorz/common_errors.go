package errorz

import "net/http"

var (
	ErrUnauthorizedAccess        = New(CodeUnauthorizedAccess, ErrUnauthorized.StatusCode, SeverityWarn, http.StatusText(ErrUnauthorized.StatusCode))
	ErrForbiddenAccess           = New(CodeForbiddenAccess, ErrForbidden.StatusCode, SeverityWarn, http.StatusText(ErrForbidden.StatusCode))
	ErrInvalidToken              = New(CodeInvalidToken, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
	ErrInvalidUUID               = New(CodeInvalidUUID, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
	ErrRecordNotFound            = New(CodeRecordNotFound, ErrNotFound.StatusCode, SeverityWarn, http.StatusText(ErrNotFound.StatusCode))
	ErrInvalidUserNameOrPassword = New(CodeInvalidUserNameOrPassword, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
	ErrExtensionMismatch         = New(CodeExtensionMismatch, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
	ErrThumbExtensionMismatch    = New(CodeThumbExtensionMismatch, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
	ErrFileRead                  = New(CodeFileRead, ErrBadRequest.StatusCode, SeverityWarn, http.StatusText(ErrBadRequest.StatusCode))
)
