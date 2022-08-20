package bang

import (
	"testing"
	"unsafe"

	"github.com/snowphoenix0105/bang/pkg/bcaster"
	"github.com/stretchr/testify/assert"
)

func TestCast(t *testing.T) {
	type structA struct {
		A    string
		Ptr1 *int
		Ptr2 *int
	}

	type structB struct {
		B         string
		UintPtr   uintptr
		UnsafePtr unsafe.Pointer
	}

	integer := 1
	uptr := unsafe.Pointer(&integer)
	uiptr := uintptr(uptr)
	expect := structB{B: "1", UnsafePtr: uptr, UintPtr: uiptr}
	input := structA{A: "1", Ptr1: &integer, Ptr2: &integer}

	fn1, err := bcaster.New[structA, structB](bcaster.OptionIgnoreFieldName, bcaster.OptionEnablePtrDecay)
	assert.Nil(t, err)
	assert.Equal(t, expect, fn1(input))

	fn2, err := bcaster.New[*structA, *structB](bcaster.OptionIgnoreFieldName, bcaster.OptionEnablePtrDecay)
	assert.Nil(t, err)
	assert.Equal(t, &expect, fn2(&input))
}
