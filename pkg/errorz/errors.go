package errorz //nolint:revive // underscore in package name is okay

// Error code constants
const (
	// Base HTTP errors
	CodeBadRequest         = "BAD_REQUEST"
	CodeUnauthorized       = "UNAUTHORIZED"
	CodeForbidden          = "FORBIDDEN"
	CodeNotFound           = "NOT_FOUND"
	CodeConflict           = "CONFLICT"
	CodeUnprocessable      = "UNPROCESSABLE"
	CodeInternalError      = "INTERNAL_ERROR"
	CodeServiceUnavailable = "SERVICE_UNAVAILABLE"

	// Derived semantic errors
	CodeAlreadyExists      = "ALREADY_EXISTS"
	CodeSomethingWentWrong = "SOMETHING_WENT_WRONG"
	CodeCustomError        = "CUSTOM_ERROR"

	// Authentication and Authorization
	CodeUnauthorizedAccess = "UNAUTHORIZED_ACCESS"
	CodeForbiddenAccess    = "FORBIDDEN_ACCESS"
	CodeInvalidToken       = "INVALID_TOKEN"

	// Validation
	CodeInvalidUUID               = "INVALID_UUID"
	CodeInvalidUserNameOrPassword = "INVALID_USERNAME_PASSWORD"
	CodeInvalidUserID             = "INVALID_USER_ID"

	// Data
	CodeRecordNotFound    = "RECORD_NOT_FOUND"
	CodeUserAlreadyExists = "USER_ALREADY_EXISTS"

	// File operations
	CodeExtensionMismatch      = "EXTENSION_MISMATCH"
	CodeThumbExtensionMismatch = "THUMB_EXTENSION_MISMATCH"
	CodeFileRead               = "FILE_READ_ERROR"
)
