package errors

import (
	std "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testError = std.New("base error")

func TestWithStack(t *testing.T) {
	assert.Nil(t, WithStack(0, nil))

	err := WithStack(0, testError)
	assert.NotNil(t, err)

	_, ok := err.(stackTraceSpanNode)
	assert.True(t, ok)

	t.Log(fmt.Sprintf("%#v", err))
	t.Log(fmt.Sprintf("%+v", err))
	t.Log(fmt.Sprintf("%v", err))
	t.Log(fmt.Sprintf("%s", err))
}
