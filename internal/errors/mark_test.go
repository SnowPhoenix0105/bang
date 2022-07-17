package errors

import (
	"errors"
	"testing"
)

func TestErrorWithMark_Error(t *testing.T) {
	err := errors.New("test errors")
	err = Mark(1, err, "mark with %d", []any{123})
	t.Log(err.Error())
}
