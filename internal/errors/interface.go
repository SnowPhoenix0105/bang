package errors

type WrapError interface {
	error
	Unwrap() error
}
type StackError interface {
	WrapError
	StackTrace() StackTrace
}

type MarkError interface {
	StackError
	Mark() string
}
