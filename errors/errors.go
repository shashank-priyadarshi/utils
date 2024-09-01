package errors

import (
	"errors"
	"fmt"
)

type Error string

var (
	InsufficientParameters Error = "insufficient parameters: expected: %d, received: %d"
	InvalidParameterType   Error = "invalid parameter type for %s: expected: %T, received: %T"
	InvalidParameterValue  Error = "invalid value for parameter %s: expected: %v, received: %v"
	OperationFailed        Error = "operation failed: err: %v"
)

func (e Error) Error(args ...interface{}) error {
	return errors.New(e.String(args...))
}

func (e Error) String(args ...interface{}) string {
	return fmt.Sprintf(string(e), args...)
}
