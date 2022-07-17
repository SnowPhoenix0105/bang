package berrors

import (
	inner "github.com/snowphoenix0105/bang/internal/errors"
)

func FormatStackTrace(err error) string {
	return inner.FormatStackTrace(err)
}

func WithStack(err error) error {
	return inner.WithStack(1, err)
}

func Mark(err error, format string, args ...any) error {
	return inner.Mark(1, err, format, args)
}

func Wrap(err error, format string, args ...any) error {
	return inner.Wrap(1, err, format, args)
}
