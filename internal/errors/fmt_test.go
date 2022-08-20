package errors

import (
	stderrors "errors"
	"fmt"
	"testing"
)

func testFuncNewErrorf() error {
	return fmt.Errorf("inner fmt.Errorf: %w", stderrors.New("new error"))
}

func testFuncStack() error {
	err := testFuncNewErrorf()
	return WithStack(0, err)
}

func testFuncMark() error {
	err := testFuncStack()
	err = Mark(0, err, "this is a marked multi-line message {\n\t\"test\": true\n}", []any{})
	err = Mark(0, err, "this is a marked single-line message", []any{})
	return err
}

func testFuncWrap() error {
	err := testFuncMark()
	err = Wrap(0, err, "this is a wrapped message", []any{})
	return err
}

func testFuncStdErrorf() error {
	err := testFuncWrap()
	err = fmt.Errorf("fmt.Errorf error: %w", err)
	return err
}

func TestFormatErrorWithFullStackTrace(t *testing.T) {
	err := testFuncStdErrorf()

	t.Log(FormatErrorWithFullStackTrace(err))
}
