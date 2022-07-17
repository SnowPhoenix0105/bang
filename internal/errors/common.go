package errors

import (
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
)

const (
	GetRuntimeStackPCListStartSize = 32

	ErrorMessageSplitter = ": "

	pkgName = "errors"
)

func formatMessage(skip int, format string, args []any) string {
	// TODO
	return fmt.Sprintf(format, args...)
}

func generateErrorMessage(cause error, msg string) string {
	return msg + ErrorMessageSplitter + cause.Error()
}

func getRuntimeStackPCList(skip int) []uintptr {
	bufferSize := GetRuntimeStackPCListStartSize
	for {
		pcBuffer := make([]uintptr, bufferSize)
		cnt := runtime.Callers(skip+2, pcBuffer)
		if cnt < bufferSize {
			return pcBuffer[:cnt]
		}
		bufferSize *= 2
	}
}

func Unwrap(err error) error {
	if wrapper, ok := err.(ErrorWrapper); ok {
		if subError := wrapper.Unwrap(); subError != nil {
			return subError
		}
	}
	return nil
}

func funcName(fn *runtime.Func) string {
	name := fn.Name()
	i := strings.LastIndexByte(name, '/')
	name = name[i+1:]
	return name
}

func formatErrorWithFlag(err internalError, state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			formatStackTrace(err, state)
			return
		}

		if state.Flag('#') {
			// *errorWithXXX is a fmt.Formatter, but errorWithXXX (not ptr) isn't.
			fmt.Fprintf(state, "&%#v", reflect.ValueOf(err).Elem().Interface())

			return
		}

		io.WriteString(state, err.Error())

	case 's':
		io.WriteString(state, err.String())
	}
}
