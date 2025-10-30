package errorz

var (
	ErrUnauthorizedAccess        = ErrUnauthorized.JoinError("Unauthorized access")
	ErrForbiddenAccess           = ErrForbidden.JoinError("Forbidden access")
	ErrInvalidToken              = ErrBadRequest.JoinError("Invalid token")
	ErrInvalidUUID               = ErrBadRequest.JoinError("Invalid UUID")
	ErrRecordNotFound            = ErrNotFound.JoinError("Record not found")
	ErrInvalidUserNameOrPassword = ErrBadRequest.JoinError("Invalid username and password")
	ErrExtensionMismatch         = ErrBadRequest.JoinError("file extension not supported")
	ErrThumbExtensionMismatch    = ErrBadRequest.JoinError("file extension not supported for thumbnail")
	ErrFileRead                  = ErrBadRequest.JoinError("file read error")
)
