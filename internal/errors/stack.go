package errors

import (
	"fmt"
	"io"
)

func WithStack(skip int, err error) error {
	if err == nil {
		return nil
	}

	pcList := getRuntimeStackPCList(skip + 1)

	return &errorWithStack{
		errorWithStackData{
			runtimeStackPCList: pcList,
			cause:              err,
		},
	}
}

type errorWithStack struct {
	errorWithStackData
}

type errorWithStackData struct {
	cause              error
	runtimeStackPCList []uintptr
}

func (e *errorWithStack) Unwrap() error {
	return e.cause
}

func (e *errorWithStack) Error() string {
	return e.cause.Error()
}

func (e *errorWithStack) String() string {
	return e.Error()
}

func (e *errorWithStack) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			formatStackTrace(e, state)
			return
		}

		if state.Flag('#') {
			fmt.Fprintf(state, "&errors.errorWithStack{%#v}", e.errorWithStackData)
			return
		}

		io.WriteString(state, e.Error())

	case 's':
		io.WriteString(state, e.String())
	}
}
