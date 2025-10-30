package errorz //nolint:revive // underscore in package name is okay

import (
	"fmt"
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
