package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestMapWithPtrKey shows that when use a pointer as the key of map, it compares only the address stored in the ptr,
instead of the object that the ptr referencing to.
*/
func TestMapWithPtrKey(t *testing.T) {
	key1 := 1
	key2 := 1
	m := map[*int]int{
		&key1: 1,
	}

	_, exist := m[&key2]
	assert.False(t, exist)

	m[&key2] = 2
	assert.Equal(t, 1, m[&key1])
	assert.Equal(t, 2, m[&key2])
}
