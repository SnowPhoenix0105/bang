package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	s1 := arr[0:3]
	s2 := arr[0:3:4]

	assert.Equal(t, len(arr), cap(s1))
	assert.Equal(t, 4, cap(s2))
	// assert.True(t, s1 == s2)
}
