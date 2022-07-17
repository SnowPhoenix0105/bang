package errors

import (
	"fmt"
)

func Markf(skip int, err error, format string, args []any) error {
	if err == nil {
		return nil
	}

	pcList := getRuntimeStackPCList(skip + 1)
	msg := formatMessage(format, args)

	return &errorWithMark{
		errorWithMarkData{
			cause:              err,
			runtimeStackPCList: pcList,
			msg:                msg,
		},
	}
}

type errorWithMark struct {
	errorWithMarkData
}

type errorWithMarkData struct {
	cause              error
	msg                string
	runtimeStackPCList []uintptr
}

func (e *errorWithMark) Unwrap() error {
	return e.cause
}

func (e *errorWithMark) Error() string {
	return generateErrorMessage(e.cause, e.msg)
}

func (e *errorWithMark) String() string {
	return e.Error()
}

func (e *errorWithMark) Format(state fmt.State, verb rune) {
	formatErrorWithFlag(e, state, verb)
}
