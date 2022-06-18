package deepcopy

import (
	"fmt"
	"reflect"
	"unsafe"
)

func produceInterface(config *Config, origin interface{}) interface{} {
	src := reflect.ValueOf(origin)
	if isSimpleKind(src.Kind()) {
		return origin
	}

	dst := reflect.New(src.Type()).Elem()
	produce(config, dst, src)

	return dst.Interface()
}

/*
produce copy the src object to dst object deeply with their reflect-objects. Requires
`dst.CanSet() && dst.Type() == src.Type() && dst.IsZero()`
*/
func produce(option *Config, dst, src reflect.Value) {
	if DEBUG {
		if !dst.CanSet() {
			panic("dst is not settable")
		}
		if dst.Type() != src.Type() {
			panic(fmt.Sprintf("dst.Type[%s] != src.Type[%s]", dst.Type().Name(), src.Type().Name()))
		}
		if !dst.IsZero() {
			panic(fmt.Sprintf("dst=reflect.ValueOf(%#v) is not zero", dst.Interface()))
		}
	}

	// deal with zero-value in advance.
	if src.IsZero() {
		return
	}

	switch src.Kind() {
	default:
		// it's simple type, just copy it.
		dst.Set(src)

	case reflect.Pointer:
		ptr := reflect.New(src.Type().Elem())
		produce(option, ptr.Elem(), src.Elem())

		dst.Set(ptr)
		return

	case reflect.Array:
		if isSimpleKind(src.Type().Elem().Kind()) {
			dst.Set(src)
			return
		}

		length := src.Len()

		for i := 0; i < length; i++ {
			produce(option, dst.Index(i), src.Index(i))
		}

		return

	case reflect.Slice:
		length := src.Len()
		newSlice := reflect.MakeSlice(src.Type(), length, src.Cap())

		for i := 0; i < length; i++ {
			produce(option, newSlice.Index(i), src.Index(i))
		}

		dst.Set(newSlice)
		return

	case reflect.Map:
		typ := src.Type()
		newMap := reflect.MakeMapWithSize(typ, src.Len())

		// two settable reflect-object to store the copied key and value of each iteration.
		kTmp := reflect.New(typ.Key()).Elem()
		vTmp := reflect.New(typ.Elem()).Elem()

		iter := src.MapRange()
		for iter.Next() {
			produce(option, kTmp, iter.Key())
			produce(option, vTmp, iter.Value())
			newMap.SetMapIndex(kTmp, vTmp)
		}

		dst.Set(newMap)
		return

	case reflect.Struct:
		// copy simple-kind fields and unexported fields
		dst.Set(src)

		numField := src.NumField()
		for i := 0; i < numField; i++ {
			srcField := src.Field(i)
			if isSimpleKind(srcField.Kind()) {
				// simple-kind fields has already been copied
				continue
			}

			produce(option, dst.Field(i), srcField)
		}

		return

	case reflect.Interface:
		if option.InterfaceStrategy == InterfaceStrategySetNil {
			// leave dst zero
			return
		}

		dst.Set(src)

		if option.InterfaceStrategy == InterfaceStrategyBitwiseCopy {
			return
		}

		srcElem := src.Elem()
		if srcElem.IsZero() {
			return
		}

		ptr := reflect.New(srcElem.Type())
		ptrElem := ptr.Elem()
		produce(option, ptrElem, srcElem)

		// when an interface{} binds to a pointer-type object, it stores the value of the object directly.
		// when it binds to a non-pointer-type object, it stores the pointer to the object.
		var srcPtr unsafe.Pointer
		if srcElem.Kind() != reflect.Pointer {
			srcPtr = ptr.UnsafePointer()
		} else {
			srcPtr = ptrElem.UnsafePointer()
		}

		dstPtr := unsafe.Pointer(uintptr(dst.Addr().UnsafePointer()) + uintptrSize)
		*(*uintptr)(dstPtr) = *(*uintptr)(srcPtr)
	}
}
