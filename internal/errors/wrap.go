package errors

import (
	"fmt"
	"io"
)

func Wrap(err error, format string, args []any) error {
	if err == nil {
		return nil
	}

	msg := formatMessage(format, args)

	return &errorWithMessage{
		errorWithMessageData{
			cause: err,
			msg:   msg,
		},
	}
}

type errorWithMessage struct {
	errorWithMessageData
}

type errorWithMessageData struct {
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
	switch verb {
	case 'v':
		if state.Flag('+') {
			formatStackTrace(e, state)
			return
		}

		if state.Flag('#') {
			fmt.Fprintf(state, "&errors.errorWithMessage{%#v}", e.errorWithMessageData)
			return
		}

		io.WriteString(state, e.Error())

	case 's':
		io.WriteString(state, e.String())
	}
}
