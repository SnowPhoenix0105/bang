package bcaster

import (
	"reflect"
	"unsafe"

	inner "github.com/snowphoenix0105/bang/internal/caster"
)

type Option interface {
	Apply(conf *inner.Config)
}

type OptionInstance struct{}

var Options OptionInstance

func (o OptionInstance) UseDefault() Option {
	return OptionFunc(func(conf *inner.Config) {
		*conf = *inner.NewDefaultConfig()
	})
}
func (o OptionInstance) DeepCopy() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.CastFunc = deepCopyCastFunc
	})
}
func (o OptionInstance) IgnoreFieldName() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.FieldNameStrategy = inner.FieldNameStrategyIgnore
	})
}
func (o OptionInstance) IgnoreFieldNameCase() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.FieldNameStrategy = inner.FieldNameStrategyIgnoreCase
	})
}
func (o OptionInstance) EnablePtrToUintptr() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.EnablePtrToUintptr = true
	})
}
func (o OptionInstance) EnablePtrToUnsafePtr() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.EnablePtrToUnsafePtr = true
	})
}
func (o OptionInstance) EnablePtrDecay() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.EnablePtrToUintptr = true
		conf.EnablePtrToUnsafePtr = true
	})
}

type OptionFuncWrapper struct {
	fn func(conf *inner.Config)
}

func OptionFunc(fn func(conf *inner.Config)) OptionFuncWrapper {
	return OptionFuncWrapper{fn: fn}
}

func (w OptionFuncWrapper) Apply(conf *inner.Config) {
	w.fn(conf)
}

func deepCopyCastFunc(outType reflect.Type, addr unsafe.Pointer) interface{} {
	// TODO

	return nil
}
