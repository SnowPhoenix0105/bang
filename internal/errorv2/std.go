package errorv2

import (
	stderrors "errors"
)

func New(msg string) error {
	return stderrors.New(msg)
}

func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

func As(err error, target any) bool {
	return stderrors.As(err, target)
}

func Cast[T any](err error) (T, bool) {
	var target T
	ok := stderrors.As(err, &target)
	return target, ok
}

func Unwrap(err error) error {
	if wrapper, ok := err.(WrapError); ok {
		return wrapper.Unwrap()
	}
	return nil
}
