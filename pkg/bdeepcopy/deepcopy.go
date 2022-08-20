package bdeepcopy

import (
	inner "github.com/snowphoenix0105/bang/internal/deepcopy"
)

var globalConfig = *inner.NewDefaultConfig()

type DeepCopier struct {
	config inner.Config
}

func NewDeepCopier(options ...Option) *DeepCopier {
	return &DeepCopier{config: *buildConfig(globalConfig, options)}
}

func (c *DeepCopier) InterfaceOf(origin interface{}, options ...Option) interface{} {
	return inner.ProduceDeepCopyInterface(buildConfig(c.config, options), origin)
}

func Of[T any](origin T, options ...Option) T {
	return InterfaceOf(origin, options...).(T)
}

func InterfaceOf(origin interface{}, options ...Option) interface{} {
	return inner.ProduceDeepCopyInterface(buildConfig(globalConfig, options), origin)
}
