package bcaster

import (
	"reflect"
	"unsafe"

	"github.com/snowphoenix0105/bang/internal/caster"
)

type Option interface {
	apply(conf *caster.Config)
}

var (
	OptionUseDefault Option = wrap(func(conf *caster.Config) {
		*conf = *caster.NewDefaultConfig()
	})
	OptionDeepCopy Option = wrap(func(conf *caster.Config) {
		conf.CastFunc = deepCopyCastFunc
	})
	OptionIgnoreFieldName Option = wrap(func(conf *caster.Config) {
		conf.FieldNameStrategy = caster.FieldNameStrategyIgnore
	})
	OptionIgnoreFieldNameCase Option = wrap(func(conf *caster.Config) {
		conf.FieldNameStrategy = caster.FieldNameStrategyIgnoreCase
	})
	OptionEnablePtrToUintptr Option = wrap(func(conf *caster.Config) {
		conf.EnablePtrToUintptr = true
	})
	OptionEnablePtrToUnsafePtr Option = wrap(func(conf *caster.Config) {
		conf.EnablePtrToUnsafePtr = true
	})
	OptionEnablePtrDecay Option = wrap(func(conf *caster.Config) {
		conf.EnablePtrToUintptr = true
		conf.EnablePtrToUnsafePtr = true
	})
)

type optionWrap struct {
	fn func(conf *caster.Config)
}

func wrap(fn func(conf *caster.Config)) optionWrap {
	return optionWrap{fn: fn}
}

func (w optionWrap) apply(conf *caster.Config) {
	w.fn(conf)
}

func deepCopyCastFunc(outType reflect.Type, addr unsafe.Pointer) interface{} {
	// TODO

	return nil
}
