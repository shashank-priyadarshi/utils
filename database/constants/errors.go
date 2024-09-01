package constants

import "errors"

type Error Type

func (t Error) Error() *error {
	err := errors.New(*t.String())
	return &err
}

func (t Error) String() *string {
	str := string(t)
	return &str
}

type ErrorFormat Error

var (
	CANNOTBEEMTPY ErrorFormat = "%v cannot be empty"
	UNSUPPORTED   ErrorFormat = "unsupported %v type"
)
