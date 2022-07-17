package errors

import (
	"fmt"
)

func Wrapf(skip int, err error, format string, args []any) error {
	if err == nil {
		return nil
	}

	msg := formatMessage(format, args)

	return &errorWithMessage{
		cause: err,
		msg:   msg,
	}
}

type errorWithMessage struct {
	cause error
	msg   string
}

func (e *errorWithMessage) Unwrap() error {
	return e.cause
}

func (e *errorWithMessage) Error() string {
	return generateErrorMessage(e.cause, e.msg)
}

func (e *errorWithMessage) String() string {
	return e.Error()
}

func (e *errorWithMessage) Format(state fmt.State, verb rune) {
	formatErrorWithFlag(e, state, verb)
}
