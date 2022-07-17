package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	GetRuntimeStackPCListStartSize = 32

	ErrorMessageSplitter = ": "
)

func formatMessage(format string, args []any) string {
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
		cnt := runtime.Callers(skip+1, pcBuffer)
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

func commonStackFrameCount(left, right []uintptr) int {
	leftIndex := len(left) - 1
	rightIndex := len(right) - 1

	cnt := 0
	for leftIndex >= 0 && rightIndex >= 0 && left[leftIndex] == right[rightIndex] {
		cnt++
		leftIndex--
		rightIndex--
	}

	return cnt
}

func funcName(fn *runtime.Func) string {
	name := fn.Name()
	i := strings.LastIndexByte(name, '/')
	name = name[i+1:]
	return name
}
