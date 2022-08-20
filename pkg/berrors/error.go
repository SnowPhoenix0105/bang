package berrors

import (
	inner "github.com/snowphoenix0105/bang/internal/errors"
)

type (
	Frame            = inner.Frame
	StackTrace       = inner.StackTrace
	MarkedFrame      = inner.MarkedFrame
	MarkedStackTrace = inner.MarkedStackTrace

	WrapError  = inner.WrapError
	StackError = inner.StackError
	MarkError  = inner.MarkError
)

func WithMessage(err error, format string, args ...any) error {
	return inner.WithMessage(1, err, format, args)
}

func WithStack(err error) error {
	return inner.WithStack(1, err)
}

// Wrap = WithMessage + WithStack, not recommend, use Mark instead
func Wrap(err error, format string, args ...any) error {
	return inner.Wrap(1, err, format, args)
}

func Mark(err error, format string, args ...any) error {
	return inner.Mark(1, err, format, args)
}

func CollectMarkedStackTraceFromError(err error) (marked MarkedStackTrace, baseError error) {
	return inner.CollectMarkedStackTraceFromError(err)
}

// CollectStackTrace returns the deepest StackTrace and the baseError of it
func CollectStackTrace(err error) (stack StackTrace, baseError error) {
	return inner.CollectStackTrace(err)
}

func FormatErrorWithFullStackTrace(err error) string {
	return inner.FormatErrorWithFullStackTrace(err)
}

// PCListToStackTrace cast the []uintptr to StackTrace (they share the memory)
func PCListToStackTrace(pcList []uintptr) StackTrace {
	return inner.PCListToStackTrace(pcList)
}

// StackTraceToPCList cast the StackTrace to []uintptr (they share the memory)
func StackTraceToPCList(trace StackTrace) []uintptr {
	return inner.StackTraceToPCList(trace)
}
