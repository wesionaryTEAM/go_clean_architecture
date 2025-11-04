package errorz

import "net/http"

var (
	ErrUnauthorizedAccess        = New("UNAUTHORIZED_ACCESS", ErrUnauthorized.StatusCode, http.StatusText(ErrUnauthorized.StatusCode))
	ErrForbiddenAccess           = New("FORBIDDEN_ACCESS", ErrForbidden.StatusCode, http.StatusText(ErrForbidden.StatusCode))
	ErrInvalidToken              = New("INVALID_TOKEN", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrInvalidUUID               = New("INVALID_UUID", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrRecordNotFound            = New("RECORD_NOT_FOUND", ErrNotFound.StatusCode, http.StatusText(ErrNotFound.StatusCode))
	ErrInvalidUserNameOrPassword = New("INVALID_USERNAME_PASSWORD", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrExtensionMismatch         = New("EXTENSION_MISMATCH", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrThumbExtensionMismatch    = New("THUMB_EXTENSION_MISMATCH", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
	ErrFileRead                  = New("FILE_READ_ERROR", ErrBadRequest.StatusCode, http.StatusText(ErrBadRequest.StatusCode))
)
