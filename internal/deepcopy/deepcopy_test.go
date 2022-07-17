package deepcopy

import (
	"testing"

	"github.com/snowphoenix0105/bang/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestDeepCopySimple(t *testing.T) {
	type InnerClass struct {
		Str          string
		Integer      int
		StrPtr       *string
		IntPtr       *int
		IntArray     [4]int
		IntArrayPtr  *[4]int
		IntPtrArray  [4]*int
		IntSlice     []int
		IntSlicePtr  *[]int
		IntPtrSlice  []*int
		IntMapInt    map[int]int
		IntMapIntPtr map[int]*int
		IntPtrMapInt map[*int]int
	}

	type Class struct {
		InnerClass
		InnerClassObj InnerClass
		InnerClassPtr *InnerClass
	}

	generator := 0
	genInt := func() int {
		generator++
		return generator
	}

	keys := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	obj := Class{
		InnerClass: InnerClass{
			Str:          ".Str",
			Integer:      genInt(),
			StrPtr:       ptr.To(".StrPtr"),
			IntPtr:       ptr.To(genInt()),
			IntArray:     [4]int{genInt(), genInt(), genInt(), genInt()},
			IntArrayPtr:  &[4]int{genInt(), genInt(), genInt(), genInt()},
			IntPtrArray:  [4]*int{ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt())},
			IntSlice:     []int{genInt(), genInt()},
			IntSlicePtr:  &[]int{genInt(), genInt(), genInt()},
			IntPtrSlice:  []*int{ptr.To(genInt())},
			IntMapInt:    map[int]int{1: genInt(), 2: genInt()},
			IntMapIntPtr: map[int]*int{1: ptr.To(genInt()), 2: ptr.To(genInt())},
			IntPtrMapInt: map[*int]int{&keys[0]: genInt(), &keys[1]: genInt()},
		},
		InnerClassObj: InnerClass{
			Str:          "InnerClassObj.Str",
			Integer:      genInt(),
			StrPtr:       ptr.To("InnerClassObj.StrPtr"),
			IntPtr:       ptr.To(genInt()),
			IntArray:     [4]int{genInt(), genInt(), genInt(), genInt()},
			IntArrayPtr:  &[4]int{genInt(), genInt(), genInt(), genInt()},
			IntPtrArray:  [4]*int{ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt())},
			IntSlice:     []int{genInt(), genInt()},
			IntSlicePtr:  &[]int{genInt(), genInt(), genInt()},
			IntPtrSlice:  []*int{ptr.To(genInt())},
			IntMapInt:    map[int]int{1: genInt(), 2: genInt()},
			IntMapIntPtr: map[int]*int{1: ptr.To(genInt()), 2: ptr.To(genInt())},
			IntPtrMapInt: map[*int]int{&keys[0]: genInt(), &keys[1]: genInt()},
		},
		InnerClassPtr: &InnerClass{
			Str:          "InnerClassPtr.Str",
			Integer:      genInt(),
			StrPtr:       ptr.To("InnerClassPtr.StrPtr"),
			IntPtr:       ptr.To(genInt()),
			IntArray:     [4]int{genInt(), genInt(), genInt(), genInt()},
			IntArrayPtr:  &[4]int{genInt(), genInt(), genInt(), genInt()},
			IntPtrArray:  [4]*int{ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt()), ptr.To(genInt())},
			IntSlice:     []int{genInt(), genInt()},
			IntSlicePtr:  &[]int{genInt(), genInt(), genInt()},
			IntPtrSlice:  []*int{ptr.To(genInt())},
			IntMapInt:    map[int]int{1: genInt(), 2: genInt()},
			IntMapIntPtr: map[int]*int{1: ptr.To(genInt()), 2: ptr.To(genInt())},
			IntPtrMapInt: map[*int]int{&keys[0]: genInt(), &keys[1]: genInt()},
		},
	}

	cpy := Of(obj)

	// ensure they are equal
	assert.Equal(t, obj, cpy)

	// then, ensure the two object are independent
}
