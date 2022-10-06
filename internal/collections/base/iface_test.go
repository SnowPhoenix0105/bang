package base

import (
	"testing"

	"github.com/snowphoenix0105/bang/internal/types/mapping"

	"github.com/snowphoenix0105/bang/internal/types/slice"
)

func TestSlice(t *testing.T) {
	var backCollection BackCollection[int32]
	var readKV ReadableKVCollection[int, int32]

	backCollection = slice.New[int32]()
	readKV = slice.New[int32]()

	_ = backCollection
	_ = readKV
}

func TestMap(t *testing.T) {
	var kv ReadWriteKVCollection[int, int32]

	kv = mapping.Make[int, int32]()

	_ = kv
}
