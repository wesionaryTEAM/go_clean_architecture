package utils

import (
	"clean-architecture/lib"
	"errors"
	"fmt"
	"strings"

	smithy "github.com/aws/smithy-go"
)

type AWSError struct {
	OE               *smithy.OperationError
	StatusCode       string
	RequestID        string
	ExceptionType    string
	ExceptionMessage string
}

func MapAWSError(logger lib.Logger, err error) (awsError *AWSError) {

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
		}
	}()

	var oe *smithy.OperationError
	if errors.As(err, &oe) {
		errorStr := oe.Err.Error()
		errorData := strings.Split(errorStr, ",")
		if len(errorData) == 3 {
			awsError = &AWSError{
				OE:               oe,
				StatusCode:       strings.Split(strings.TrimSpace(errorData[0]), ": ")[1],
				RequestID:        strings.Split(strings.TrimSpace(errorData[1]), ": ")[1],
				ExceptionType:    strings.Split(strings.TrimSpace(errorData[2]), ": ")[0],
				ExceptionMessage: strings.Split(strings.TrimSpace(errorData[2]), "Exception: ")[1],
			}
			return
		}
	}
	return
}

func (e AWSError) String() string {
	return fmt.Sprintf(
		"StatusCode: %s, RequestID: %s, ExceptionType: %s, ExceptionMessage: %s",
		e.StatusCode,
		e.RequestID,
		e.ExceptionType,
		e.ExceptionMessage,
	)
}

func (e AWSError) Error() string {
	return strings.TrimSpace(e.ExceptionMessage)
}
