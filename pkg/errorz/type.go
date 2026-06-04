package errorz

import (
	"errors"
	"fmt"
)

// Severity classifies how serious an error is, for client handling and telemetry.
type Severity string

const (
	SeverityInfo  Severity = "info"
	SeverityWarn  Severity = "warn"
	SeverityError Severity = "error"
)

// APIError is the canonical application error type.
//
// Code:       machine-readable identifier (UPPER_SNAKE); the client localizes from this.
// StatusCode: HTTP status for transport; never serialized into the response body.
// Severity:   how serious the error is (info|warn|error).
// Message:    optional human-readable message set by the backend; omitted when empty.
// Params:     optional context / i18n interpolation values; omitted when empty.
// Cause:      internal wrapped error for logging/Sentry/Unwrap; never serialized.
type APIError struct {
	Code       string
	StatusCode int
	Severity   Severity
	Message    string
	Params     map[string]any
	Cause      error
}

// Error renders the error for logs: "CODE", optionally followed by the custom
// message and/or the internal cause. The custom Message is included so it is
// not lost from log lines and Sentry events once a backend sets it.
func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	s := e.Code
	if e.Message != "" {
		s += ": " + e.Message
	}
	if e.Cause != nil {
		s += fmt.Sprintf(" (cause: %v)", e.Cause)
	}
	return s
}

// Unwrap exposes the internal Cause so errors.Is/As can walk the chain.
func (e *APIError) Unwrap() error { return e.Cause }

// New creates a base API error. An optional custom message may be supplied as
// the final argument; existing three-argument calls remain valid.
func New(code string, status int, severity Severity, message ...string) *APIError {
	e := &APIError{Code: code, StatusCode: status, Severity: severity}
	if len(message) > 0 {
		e.Message = message[0]
	}
	return e
}

// clone returns a copy so package-level error vars are never mutated by builders.
func (e *APIError) clone() *APIError {
	if e == nil {
		return nil
	}
	c := *e
	if e.Params != nil {
		c.Params = make(map[string]any, len(e.Params))
		for k, v := range e.Params {
			c.Params[k] = v
		}
	}
	return &c
}

// WithMessage returns a copy with an optional custom message set.
func (e *APIError) WithMessage(msg string) *APIError {
	if e == nil {
		return e
	}
	c := e.clone()
	c.Message = msg
	return c
}

// WithParam returns a copy with a single param set.
func (e *APIError) WithParam(key string, val any) *APIError {
	if e == nil {
		return e
	}
	c := e.clone()
	if c.Params == nil {
		c.Params = make(map[string]any)
	}
	c.Params[key] = val
	return c
}

// WithParams returns a copy with the given params merged in.
func (e *APIError) WithParams(params map[string]any) *APIError {
	if e == nil || params == nil {
		return e
	}
	c := e.clone()
	if c.Params == nil {
		c.Params = make(map[string]any, len(params))
	}
	for k, v := range params {
		c.Params[k] = v
	}
	return c
}

// Wrap returns a copy with an internal cause attached.
func (e *APIError) Wrap(err error) *APIError {
	if e == nil || err == nil {
		return e
	}
	c := e.clone()
	c.Cause = err
	return c
}

// From converts an arbitrary error into an *APIError (fallback ErrInternal).
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
