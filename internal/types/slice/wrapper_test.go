package slice

import (
	"testing"
	"unsafe"

	"github.com/snowphoenix0105/bang/internal/math"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	s := Raw(1)

	ws := Wrap(s)
	ws.Append(2, 3)
	assert.Equal(t, Raw(1, 2, 3), ws.Raw())
	assert.Equal(t, Raw(1), s)

	wsp := WrapPtr(&s)
	wsp.Append(2, 3)
	assert.Equal(t, Raw(1, 2, 3), wsp.Raw())
	assert.Equal(t, Raw(1, 2, 3), s)
	assert.Equal(t, unsafe.Pointer((*Slice[int])(wsp)), unsafe.Pointer((*[]int)(wsp.RawPtr())))
}

func TestSlice_Append(t *testing.T) {
	s := Make[int]()
	assert.Equal(t, []int{}, s.Raw())

	s.Append(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, s.Raw())

	idx, ok := s.Find(2)
	assert.True(t, ok)
	assert.Equal(t, 1, idx)

	_, ok = s.Find(4)
	assert.False(t, ok)

	val, exists := s.TryGet(2)
	assert.True(t, exists)
	assert.Equal(t, 3, val)

	_, exists = s.TryGet(3)
	assert.False(t, exists)
}

func TestSlice_Set(t *testing.T) {
	s := Make[int]()

	assert.False(t, s.SetLen(2))
	assert.False(t, s.TrySet(1, 2))

	s.Append(1, 2, 3)
	ok := s.TrySet(1, 3)
	assert.True(t, ok)
	assert.Equal(t, []int{1, 3, 3}, s.Raw())
	assert.False(t, s.Contains(2))
	assert.True(t, s.Contains(3))
	assert.False(t, s.ContainsKey(3))
	assert.True(t, s.ContainsKey(2))
}

func TestSlice_Back(t *testing.T) {
	s := NewWithCap[int](2)
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, 2, s.Cap())

	s.PushBack(-1)
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, -1, s.GetBack())

	s.PushBack(-2)
	assert.Equal(t, 2, s.Len())
	assert.Equal(t, -2, s.GetBack())

	s.PopBack()
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, -1, s.GetBack())

	s.PopBack()
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, 0, s.GetBack())
}

func TestSlice_All(t *testing.T) {
	s := NewWithVal(1, 2, 3, 4, 5)
	assert.True(t, s.All(math.GreaterThanZero[int]))
	assert.True(t, s.Any(math.GreaterThanZero[int]))

	s2 := NewWithVal(-1, 0, 1, 2)
	assert.False(t, s2.All(math.GreaterThanZero[int]))
	assert.True(t, s2.Any(math.GreaterThanZero[int]))

	s3 := MakeWithVal(1, 2, 114514)
	s4 := NewWithCap[int](4)
	s3.ForAll(func(e int) {
		s4.Append(e)
	})
	assert.Equal(t, s3.Raw(), s4.Raw())

	s5 := New[int]()
	s3.ForEach(func(e int) bool {
		if e > 10 {
			return false
		}
		s5.Append(e)
		return true
	})

	assert.Equal(t, []int{1, 2}, s5.Raw())
}
