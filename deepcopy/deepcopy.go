package deepcopy

import "reflect"

var defaultOption Option = *NewDefaultOption()

func SetDefaultOption(option *Option) {
	defaultOption = *option
}

type DeepCopier struct {
	option Option
}

func Of[T any](origin T) T {
	return InterfaceOf(origin).(T)
}

func InterfaceOf(origin interface{}) interface{} {
	src := reflect.ValueOf(origin)
	if isSimpleKind(src.Kind()) {
		return origin
	}

	dst := reflect.New(src.Type()).Elem()
	produce(&defaultOption, dst, src)

	return dst.Interface()
}
