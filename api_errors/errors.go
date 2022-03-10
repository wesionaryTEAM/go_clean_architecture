package api_errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnauthorizedAccess     = errors.New("unauthorized access")
	ErrInvalidUUID            = errors.New("invalid uuid")
	ErrExtensionMismatch      = errors.New("file extension not supported")
	ErrThumbExtensionMismatch = errors.New("file extension not supported for thumbnail")
	ErrFileRead               = errors.New("file read error")
)

// for dynamic error
type ErrTokenVerification struct {
	id string
}

func NewErrTokenVerification(id string) error {
	return ErrTokenVerification{
		id: id,
	}
}

func (e ErrTokenVerification) Error() string {
	return fmt.Sprintf("error verifying id token %v\n", e.id)
}
