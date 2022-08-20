package caster

import (
	"reflect"
	"testing"

	"github.com/snowphoenix0105/bang/pkg/ptr"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type Int int

	fnInt, err := New[int, Int]()
	assert.Nil(t, err)
	assert.True(t, reflect.TypeOf(Int(0)) == reflect.TypeOf(fnInt(0)))
	assert.Equal(t, Int(2), fnInt(2))

	fnPtr, err := New[*int, *Int]()
	assert.Nil(t, err)
	assert.True(t, reflect.TypeOf((*Int)(nil)) == reflect.TypeOf(fnPtr(ptr.To(0))))
	assert.Equal(t, ptr.To(Int(2)), fnPtr(ptr.To(2)))
	assert.False(t, ptr.To(Int(2)) == fnPtr(ptr.To(2)))
	assert.Equal(t, Int(3), *fnPtr(ptr.To(3)))
}
