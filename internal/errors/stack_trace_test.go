package errors

import (
	std "errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFuncNewErrorf() error {
	return fmt.Errorf("inner fmt.Errorf: %w", std.New("new error"))
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

func TestStackTraceMode(t *testing.T) {
	err := testFuncStdErrorf()
	trace := stackTrace{}
	trace.applyError(err)

	assert.Equal(t, stackTraceErrorMessageModeCompatible, trace.errorMessageMode)

	err = testFuncWrap()
	trace.applyError(err)

	assert.Equal(t, stackTraceErrorMessageModeNormal, trace.errorMessageMode)

	err = testFuncNewErrorf()
	trace.applyError(err)

	assert.Equal(t, stackTraceErrorMessageModeCompatible, trace.errorMessageMode)
}

func TestFormatStackTrace(t *testing.T) {
	err := testFuncStdErrorf()
	msg := FormatStackTrace(err)
	t.Log("testFuncStdErrorf->\n" + msg)
	assert.True(t, strings.HasPrefix(msg, err.Error()))

	err = testFuncWrap()
	msg = FormatStackTrace(err)
	t.Log("testFuncWrap->\n" + msg)
	assert.False(t, strings.HasPrefix(msg, err.Error()))

	err = testFuncNewErrorf()
	msg = FormatStackTrace(err)
	t.Log("testFuncNewErrorf->\n" + msg)
	assert.Equal(t, err.Error(), msg)
}
