package errors

/*
messageError implements WrapError
*/
type messageError struct {
	cause error
	msg   string
}

func (m *messageError) Error() string {
	return generateErrorMessage(m.cause, m.msg)
}

func (m *messageError) Unwrap() error {
	return m.cause
}

/*
stackError implements StackError
*/
type stackError struct {
	cause  error
	pcList []uintptr
}

func (s *stackError) Error() string {
	return s.cause.Error()
}

func (s *stackError) RawPCList() []uintptr {
	return s.pcList
}

func (s *stackError) StackTrace() StackTrace {
	return PCListToStackTrace(s.RawPCList())
}

func (s *stackError) Unwrap() error {
	return s.cause
}

/*
markError implements MarkError
*/
type markError struct {
	stackError
	mark string
}

func (m *markError) Error() string {
	return generateErrorMessage(m.cause, m.mark)
}

func (m *markError) Unwrap() error {
	return m.cause
}

func (m *markError) Mark() string {
	return m.mark
}
