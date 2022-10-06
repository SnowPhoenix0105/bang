package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Depth())

	s.Push(114514)
	assert.Equal(t, 1, s.Depth())
	assert.Equal(t, 114514, s.Peak())

	s.Pop()
	assert.Equal(t, 0, s.Depth())
}
