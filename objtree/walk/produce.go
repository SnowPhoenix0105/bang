package walk

import (
	"reflect"
	"unsafe"
)

func getPtrFromReflect(ref reflect.Value) uintptr {
	type hackValue struct {
		typ  uintptr
		ptr  uintptr
		flag uintptr
	}
	return (*hackValue)(unsafe.Pointer(&ref)).ptr
}

func invokeWalkCallback(path *pathRecorder, ref reflect.Value, callback NodeCallback) error {
	ctx := NodeContext{
		Path:    path.String(),
		Address: getPtrFromReflect(ref),
		Type:    ref.Type(),
		Reflect: ref,
	}
	return callback(&ctx)
}

func produce(config *Config, node reflect.Value) {

}
