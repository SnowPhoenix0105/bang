package errors

import (
	"fmt"
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
	return strings.ReplaceAll(msg, "\n", " ") + ErrorMessageSplitter + cause.Error()
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

func funcName(fn *runtime.Func) string {
	name := fn.Name()
	i := strings.LastIndexByte(name, '/')
	name = name[i+1:]
	return name
}

func Unwrap(err error) error {
	if wrapper, ok := err.(WrapError); ok {
		return wrapper.Unwrap()
	}
	return nil
}
