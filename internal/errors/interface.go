package errors

import "fmt"

type ErrorWrapper interface {
	Unwrap() error
}

type Stringer interface {
	String() string
}

type internalError interface {
	error
	Stringer
	fmt.Formatter
}

type stackTraceSpanNode interface {
	internalError
	ErrorWrapper
}
