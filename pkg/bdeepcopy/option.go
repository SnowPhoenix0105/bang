package bdeepcopy

import (
	inner "github.com/snowphoenix0105/bang/internal/deepcopy"
)

type Option interface {
	Apply(config *inner.Config)
}

type OptionUnsafeInstance struct{}

type OptionInstance struct {
	Unsafe OptionUnsafeInstance
}

var Options OptionInstance

func buildConfig(origin inner.Config, options []Option) *inner.Config {
	cpy := origin

	for _, opt := range options {
		opt.Apply(&cpy)
	}

	return &cpy
}

type useCopier struct {
	copier *DeepCopier
}

func (uc *useCopier) Apply(conf *inner.Config) {
	*conf = uc.copier.config
}

func (o OptionInstance) UseCopier(copier *DeepCopier) Option {
	return &useCopier{copier: copier}
}

func (o OptionInstance) InterfaceBitwiseCopy() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.InterfaceStrategy = inner.InterfaceStrategyBitwiseCopy
	})
}

func (o OptionInstance) InterfaceSetNil() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.InterfaceStrategy = inner.InterfaceStrategySetNil
	})
}

func (o OptionInstance) MapBitwiseCopyKey() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.MapStrategy = inner.MapStrategyBitwiseCopyKey
	})
}

func (o OptionInstance) MapDeepCopyKey() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.MapStrategy = inner.MapStrategyDeepCopyKey
	})
}

func (u OptionUnsafeInstance) InterfaceDeepCopy() Option {
	return OptionFunc(func(conf *inner.Config) {
		conf.InterfaceStrategy = inner.InterfaceStrategyDeepCopyUnsafe
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
