package bang

import (
	"testing"
	"unsafe"

	"github.com/snowphoenix0105/bang/pkg/ptr"

	"github.com/snowphoenix0105/bang/pkg/bdeepcopy"

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

	fn1, err := bcaster.New[structA, structB](bcaster.Options.IgnoreFieldName(), bcaster.Options.EnablePtrDecay())
	assert.Nil(t, err)
	assert.Equal(t, expect, fn1(input))

	fn2, err := bcaster.New[*structA, *structB](bcaster.Options.IgnoreFieldName(), bcaster.Options.EnablePtrDecay())
	assert.Nil(t, err)
	assert.Equal(t, &expect, fn2(&input))
}

func TestDeepCopy(t *testing.T) {
	type class struct {
		A string
		B interface{}
	}

	copier := NewDeepCopier(bdeepcopy.Options.InterfaceSetNil(), bdeepcopy.Options.MapBitwiseCopyKey())
	cpy := DeepCopyOf(class{A: "1"}, bdeepcopy.Options.UseCopier(copier))
	assert.Equal(t, class{A: "1", B: nil}, cpy)

	intptr := ptr.To(4)
	input2 := &class{A: "3", B: intptr}
	cpy2 := DeepCopyOf(input2, bdeepcopy.Options.InterfaceBitwiseCopy())
	assert.Equal(t, input2, cpy2)
	assert.True(t, intptr == cpy2.B.(*int))

	input3 := class{A: "3", B: &class{A: "4", B: 5}}
	cpy3 := DeepCopyOf(&input3, bdeepcopy.Options.Unsafe.InterfaceDeepCopy())
	assert.Equal(t, input3, cpy3)
}
