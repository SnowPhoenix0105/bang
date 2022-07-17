package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMark(t *testing.T) {
	assert.Nil(t, Markf(0, nil, "mark with %d", []any{123}))

	err := Markf(0, testError, "mark with %d%s%f", []any{123, "123", 1.23})
	assert.NotNil(t, err)

	_, ok := err.(stackTraceSpanNode)
	assert.True(t, ok)

	t.Log(fmt.Sprintf("%#v", err))
	t.Log(fmt.Sprintf("%+v", err))
	t.Log(fmt.Sprintf("%v", err))
	t.Log(fmt.Sprintf("%s", err))
}
