package berrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestWrapError struct {
	Inner error
}

func (e *TestWrapError) Error() string {
	return e.Inner.Error()
}

func (e *TestWrapError) Unwrap() error {
	return e.Inner
}

func TestCast(t *testing.T) {
	innerError := New("err")
	err := Mark(&TestWrapError{innerError}, "mark")
	target, found := Cast[*TestWrapError](err)
	require.True(t, found)
	assert.True(t, innerError == target.Inner)
}
