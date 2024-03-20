package utilities

import "fmt"

var (
	InsufficientParameters error = NewError("insufficient parameters")
	InvalidParameter       error = NewError("invalid parameter")

	OperationFailed error = NewError("operation failed")

	UnsupportedType error = NewError("unsupported type")
)

type UtilityError struct {
	errorMessage string
}

func NewError(args ...string) *UtilityError {
	var errorMessage string

	for _, arg := range args {
		errorMessage = fmt.Sprintf("%s: %s", errorMessage, arg)
	}

	return &UtilityError{errorMessage: errorMessage}
}

func (u *UtilityError) Error() string {
	return u.errorMessage
}
