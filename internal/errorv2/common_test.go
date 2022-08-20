package errorv2

import (
	stderrors "errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testError = stderrors.New("base error")

func TestPkgName(t *testing.T) {
	type tmp struct{}

	assert.True(t, strings.HasSuffix(reflect.TypeOf(tmp{}).PkgPath(), pkgName))
}

func TestUnwrap(t *testing.T) {
	assert.Nil(t, Unwrap(nil))
	assert.Nil(t, Unwrap(&messageError{cause: nil}))
	assert.Nil(t, Unwrap(testError))
	assert.Equal(t, testError, Unwrap(fmt.Errorf("%w%s", testError, "123")))
	assert.Equal(t, testError, Unwrap(WithStack(0, testError)))
	assert.Equal(t, testError, Unwrap(Mark(0, testError, "mark", []any{123})))
	assert.Equal(t, testError, Unwrap(WithMessage(0, testError, "mark", []any{123})))
}

func TestGetRuntimeStackPCList(t *testing.T) {
	depth := int(GetRuntimeStackPCListStartSize * 1.5)

	var fn func(int) []uintptr

	fn = func(i int) []uintptr {
		if i > 0 {
			return fn(i - 1)
		}
		return getRuntimeStackPCList(0)
	}

	pcList := fn(depth - 1)

	//for i, frame := range pcList {
	//	t.Logf("[%2d] %x -> %s", i, frame, funcName(runtime.FuncForPC(frame)))
	//}

	// depth + 3 (runtime.goexit, testing.tRunner, and this test function)
	assert.Equal(t, depth+3, len(pcList))
}
