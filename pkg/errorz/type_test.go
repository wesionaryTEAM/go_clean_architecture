package errorz_test

import (
	"errors"
	"net/http"
	"testing"

	"clean-architecture/pkg/errorz"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := errorz.New("CODE_X", http.StatusBadRequest, errorz.SeverityWarn)
	assert.Equal(t, "CODE_X", e.Code)
	assert.Equal(t, http.StatusBadRequest, e.StatusCode)
	assert.Equal(t, errorz.SeverityWarn, e.Severity)
	assert.Empty(t, e.Message)
	assert.Nil(t, e.Params)
}

func TestNewWithOptionalMessage(t *testing.T) {
	e := errorz.New("CODE_X", http.StatusBadRequest, errorz.SeverityWarn, "default message")
	assert.Equal(t, "default message", e.Message)
}

func TestWithMessageDoesNotMutateShared(t *testing.T) {
	orig := errorz.ErrBadRequest
	originalMessage := orig.Message
	withMsg := orig.WithMessage("custom message")

	assert.Equal(t, "custom message", withMsg.Message)
	assert.Equal(t, originalMessage, orig.Message, "shared package-level error must not be mutated")
	assert.Equal(t, orig.Code, withMsg.Code)
	assert.Equal(t, orig.StatusCode, withMsg.StatusCode)
	assert.Equal(t, orig.Severity, withMsg.Severity)
}

func TestWithParamAndWithParams(t *testing.T) {
	orig := errorz.ErrBadRequest

	one := orig.WithParam("field", "email")
	assert.Equal(t, "email", one.Params["field"])
	assert.Nil(t, orig.Params, "shared package-level error must not be mutated")

	many := orig.WithParams(map[string]any{"a": 1, "b": 2})
	assert.Equal(t, 1, many.Params["a"])
	assert.Equal(t, 2, many.Params["b"])
	assert.Nil(t, orig.Params, "shared package-level error must not be mutated")
}

func TestWrapSetsCauseAndUnwraps(t *testing.T) {
	cause := errors.New("boom")
	wrapped := errorz.ErrInternal.Wrap(cause)

	assert.Equal(t, cause, wrapped.Cause)
	assert.Equal(t, cause, errors.Unwrap(wrapped))
	assert.Nil(t, errorz.ErrInternal.Cause, "shared package-level error must not be mutated")
}

func TestFrom(t *testing.T) {
	api := errorz.ErrNotFound
	assert.Equal(t, api, errorz.From(api), "existing APIError should be returned as-is")

	plain := errors.New("random failure")
	converted := errorz.From(plain)
	assert.Equal(t, errorz.CodeInternalError, converted.Code)
	assert.Equal(t, plain, converted.Cause)

	assert.Nil(t, errorz.From(nil))
}
