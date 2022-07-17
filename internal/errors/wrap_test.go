package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	assert.Nil(t, Wrap(0, nil, "wrap with %d", []any{123}))

	err := Wrap(0, testError, "wrap with %d%f%s", []any{123, 2.34, "tmp"})
	assert.NotNil(t, err)

	_, ok := err.(stackTraceSpanNode)
	assert.True(t, ok)

	t.Log(fmt.Sprintf("%#v", err))
	t.Log(fmt.Sprintf("%+v", err))
	t.Log(fmt.Sprintf("%v", err))
	t.Log(fmt.Sprintf("%s", err))
}
