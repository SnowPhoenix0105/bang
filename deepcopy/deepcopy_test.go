package deepcopy

import (
	"testing"
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
		IntMapInt    map[int]string
		IntMapIntPtr map[int]*int
		IntPtrMapInt map[*int]int
	}

	type Class struct {
		InnerClass
		InnerClassObj InnerClass
		InnerClassPtr *InnerClass
	}
}
