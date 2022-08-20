package errors

import "testing"

func TestInterfaceImplementation(t *testing.T) {
	var (
		stackInterface StackError
		wrapInterface  WrapError
		markInterface  MarkError
	)

	wrapInterface = &messageError{}
	stackInterface = &stackError{}
	markInterface = &markError{}

	_ = stackInterface
	_ = wrapInterface
	_ = markInterface
}
