package caster

import (
	"reflect"
	"unsafe"

	"github.com/snowphoenix0105/bang/internal/collections/stack"
)

var defaultConfig Config

func init() {
	defaultConfig = *NewDefaultConfig()
}

type Func[InType any, OutType any] func(InType) OutType

func MustNew[InType any, OutType any](config *Config) Func[InType, OutType] {
	fn, err := New[InType, OutType](config)
	if err != nil {
		panic(err)
	}
	return fn
}

func New[InType any, OutType any](config *Config) (Func[InType, OutType], error) {
	if config == nil {
		config = &defaultConfig
	}

	var inInstance InType
	var outInstance OutType
	inType := reflect.TypeOf(inInstance)
	outType := reflect.TypeOf(outInstance)

	inPath := stack.NewString()
	outPath := stack.NewString()
	if err := check(config, inPath, outPath, inType, outType); err != nil {
		return nil, err
	}

	fn := func(in InType) OutType {
		return reflect.NewAt(outType, unsafe.Pointer(&in)).Elem().Interface().(OutType)
	}

	return fn, nil
}
