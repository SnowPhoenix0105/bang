package errors

func WithStack(skip int, err error) error {
	return &stackError{
		cause:  err,
		pcList: getRuntimeStackPCList(skip + 1),
	}
}

func WithMessage(skip int, err error, format string, args []any) error {
	return &messageError{
		cause: err,
		msg:   formatMessage(skip+1, format, args),
	}
}

// Wrap = WithMessage + WithStack, not recommend, use Mark instead
func Wrap(skip int, err error, format string, args []any) error {
	return &stackError{
		cause: &messageError{
			cause: err,
			msg:   formatMessage(skip+1, format, args),
		},
		pcList: getRuntimeStackPCList(skip + 1),
	}
}

func Mark(skip int, err error, format string, args []any) error {
	return &markError{
		stackError: stackError{
			cause:  err,
			pcList: getRuntimeStackPCList(skip + 1),
		},
		mark: formatMessage(skip+1, format, args),
	}
}
