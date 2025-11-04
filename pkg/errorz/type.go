package errorz

import (
	"errors"
	"fmt"
)

// APIError is the canonical application error type.
// Code: machine-readable identifier (UPPER_SNAKE)
// Message: human readable safe description
// StatusCode: HTTP status for transport
// Details: optional non-sensitive context
// Cause: internal wrapped error (not serialized to clients)
type APIError struct {
	Code       string
	Message    string
	StatusCode int
	Details    map[string]any
	Cause      error
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *APIError) Unwrap() error { return e.Cause }

// New creates a base API error
func New(code string, status int, message string) *APIError {
	return &APIError{Code: code, StatusCode: status, Message: message, Details: map[string]any{}}
}

// Wrap clones the error and attaches a cause
func (e *APIError) Wrap(err error) *APIError {
	if err == nil || e == nil {
		return e
	}
	clone := &APIError{Code: e.Code, Message: e.Message, StatusCode: e.StatusCode, Details: map[string]any{}}
	for k, v := range e.Details {
		clone.Details[k] = v
	}
	clone.Cause = err
	return clone
}

// WithDetail adds a key/value detail
func (e *APIError) WithDetail(key string, val any) *APIError {
	if e == nil {
		return e
	}
	if e.Details == nil {
		e.Details = make(map[string]any)
	}
	e.Details[key] = val
	return e
}

// From converts arbitrary error to APIError (fallback INTERNAL_ERROR)
func From(err error) *APIError {
	if err == nil {
		return nil
	}
	var api *APIError
	if errors.As(err, &api) {
		return api
	}
	return ErrInternal.Wrap(err)
}
