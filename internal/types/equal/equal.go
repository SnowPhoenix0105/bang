package equal

import (
	"reflect"
	"unsafe"
)

func Of[T any](left, right T) bool {
	return Interface(left, right)
}

func OfDeeply[T any](left, right T) bool {
	return InterfaceDeeply(left, right)
}

func Interface(left, right interface{}) bool {
	if left == nil || right == nil {
		return left == right
	}

	return Reflect(reflect.ValueOf(left), reflect.ValueOf(right))
}

func Reflect(left, right reflect.Value) bool {
	if !left.IsValid() || !right.IsValid() {
		return false
	}

	typLeft := left.Type()
	typRight := right.Type()
	if typLeft != typRight {
		return false
	}

	switch typLeft.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return left.Int() == right.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return left.Uint() == right.Uint()

	case reflect.String:
		return left.String() == right.String()

	case reflect.Bool:
		return left.Bool() == right.Bool()

	case reflect.Float32, reflect.Float64:
		return left.Float() == right.Float()

	case reflect.Complex64, reflect.Complex128:
		return left.Complex() == right.Complex()

	case reflect.Ptr, reflect.UnsafePointer:
		return left.UnsafePointer() == right.UnsafePointer()

	case reflect.Struct:
		numField := typLeft.NumField()
		for i := 0; i < numField; i++ {
			if !Reflect(left.Field(i), right.Field(i)) {
				return false
			}
		}

		return true

	case reflect.Slice:
		length := left.Len()
		if length != right.Len() {
			return false
		}

		if length == 0 {
			return true
		}

		return indexEqual(length, left, right)

	case reflect.Array:
		length := typLeft.Len()
		if length == 0 {
			return true
		}

		return indexEqual(length, left, right)

	case reflect.Map:
		length := left.Len()
		if length != right.Len() {
			return false
		}

		if length == 0 {
			return true
		}

		iterLeft := left.MapRange()
		for iterLeft.Next() {
			valueRight := right.MapIndex(iterLeft.Key())
			if !valueRight.IsValid() {
				return false
			}

			if !Reflect(iterLeft.Value(), valueRight) {
				return false
			}
		}

		return true

	default:
		return reflect.DeepEqual(left.Interface(), right.Interface())
	}
}

func indexEqual(length int, left, right reflect.Value) bool {
	for i := 0; i < length; i++ {
		if !Reflect(left.Index(i), right.Index(i)) {
			return false
		}
	}

	return true
}

func InterfaceDeeply(left, right interface{}) bool {
	return reflect.DeepEqual(left, right)
}

func Bool(left, right bool) bool {
	return left == right
}

func Int(left, right int) bool {
	return left == right
}

func Int8(left, right int8) bool {
	return left == right
}

func Int16(left, right int16) bool {
	return left == right
}

func Int32(left, right int32) bool {
	return left == right
}

func Int64(left, right int64) bool {
	return left == right
}

func Uint(left, right uint) bool {
	return left == right
}

func Uint8(left, right uint8) bool {
	return left == right
}

func Uint16(left, right uint16) bool {
	return left == right
}

func Uint32(left, right uint32) bool {
	return left == right
}

func Uint64(left, right uint64) bool {
	return left == right
}

func Uintptr(left, right uintptr) bool {
	return left == right
}

func String(left, right string) bool {
	return left == right
}

func Float32(left, right float32) bool {
	return left == right
}

func Float64(left, right float64) bool {
	return left == right
}

func Complex64(left, right complex64) bool {
	return left == right
}

func Complex128(left, right complex128) bool {
	return left == right
}

func UnsafePointer(left, right unsafe.Pointer) bool {
	return left == right
}
