package errors

import (
	"fmt"
)

func WithStack(skip int, err error) error {
	if err == nil {
		return nil
	}

	pcList := getRuntimeStackPCList(skip + 1)

	return &errorWithStack{
		runtimeStackPCList: pcList,
		cause:              err,
	}
}

type errorWithStack struct {
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
	// TODO
}
