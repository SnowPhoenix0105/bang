package errors

import (
	inner "github.com/snowphoenix0105/bang/internal/errors"
)

func FormatStackTrace(err error) string {
	return inner.FormatStackTrace(err)
}

func WithStack(err error) error {
	return inner.WithStack(1, err)
}

func Markf(err error, format string, args ...any) error {
	return inner.Markf(1, err, format, args)
}

func Wrapf(err error, format string, args ...any) error {
	return inner.Wrapf(1, err, format, args)
}
