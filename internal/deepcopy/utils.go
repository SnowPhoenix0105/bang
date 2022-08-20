package deepcopy

import (
	"reflect"
	"unsafe"
)

var uintptrSize uintptr = reflect.TypeOf(unsafe.Pointer(nil)).Size()

func isSimpleKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Bool:
		return true

	default:
		return false
	}
}
