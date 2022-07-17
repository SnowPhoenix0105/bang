package errors

import (
	std "errors"
	"testing"
)

func testFuncNew() error {
	return std.New("new error")
}

func testFuncStack() error {
	err := testFuncNew()
	return WithStack(err)
}

func testFuncMark() error {
	err := testFuncStack()
	err = Mark(err, "this is a marked multi-line message {\n\t\"test\": true\n}")
	err = Mark(err, "this is a marked single-line message")
	return err
}

func testFuncWrap() error {
	err := testFuncMark()
	err = Wrap(err, "this is a wrapped message")
	return err
}

func TestFormatStackTrace(t *testing.T) {
	err := testFuncWrap()
	t.Log("\n" + FormatStackTrace(err))
}
