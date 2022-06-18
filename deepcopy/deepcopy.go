package deepcopy

import (
	"github.com/snowphoenix0105/bang/deepcopy/with"
)

var defaultConfig Config = *NewZeroConfig()

func SetDefaultOption(option *Config) {
	defaultConfig = *option
}

/*
DefaultConfig returns the copy of default-config (witch can be changed by SetDefaultOption()).
If you need a zero-value Config, use NewZeroConfig().
*/
func DefaultConfig() Config {
	ret := defaultConfig
	return ret
}

type DeepCopier struct {
	config Config
}

func (c *DeepCopier) InterfaceOf(origin interface{}, options ...with.DeepCopyOption) interface{} {
	configCopy := c.config
	configCopy.ApplyOptions(options)
	return produceInterface(&configCopy, origin)
}

func Of[T any](origin T, options ...with.DeepCopyOption) T {
	return InterfaceOf(origin, options...).(T)
}

func InterfaceOf(origin interface{}, options ...with.DeepCopyOption) interface{} {
	var config *Config
	if len(options) == 0 {
		config = &defaultConfig
	} else {
		defaultConfigCopy := defaultConfig
		defaultConfigCopy.ApplyOptions(options)
		config = &defaultConfigCopy
	}

	return produceInterface(config, origin)
}
