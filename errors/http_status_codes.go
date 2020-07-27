package errors

import (
	"net/http"
)

const (
	BadRequest = HTTPErrorType(iota)
	Unauthorized
	Forbidden
	NotFound
	Conflict
	InternalError
	Unavailable
)

type HTTPErrorType uint

// GetType returns the status code for the error type
func GetStatusCode(httpErrorType HTTPErrorType) int {
	switch httpErrorType {
	case BadRequest:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case Conflict:
		return http.StatusConflict
	case InternalError:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
