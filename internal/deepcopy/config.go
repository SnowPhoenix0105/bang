package deepcopy

import (
	"fmt"

	"github.com/snowphoenix0105/bang/internal/deepcopy/with"
)

type Config struct {
	InterfaceStrategy InterfaceStrategy
	MapStrategy       MapStrategy
}

func NewZeroConfig() *Config {
	return &Config{}
}

func (c *Config) ApplyOptions(optionList []with.DeepCopyOption) {
	for _, option := range optionList {
		switch option {

		case with.InterfaceBitwiseCopy:
			c.InterfaceStrategy = InterfaceStrategyBitwiseCopy
		case with.InterfaceSetNil:
			c.InterfaceStrategy = InterfaceStrategySetNil
		case with.InterfaceDeepCopyUnsafe:
			c.InterfaceStrategy = InterfaceStrategyDeepCopyUnsafe

		case with.MapBitwiseCopyKey:
			c.MapStrategy = MapStrategyBitwiseCopyKey
		case with.MapDeepCopyKey:
			c.MapStrategy = MapStrategyDeepCopyKey

		}
	}
}

func (c *Config) CheckValid() error {
	switch c.InterfaceStrategy {
	case InterfaceStrategyBitwiseCopy, InterfaceStrategySetNil, InterfaceStrategyDeepCopyUnsafe:
		break
	default:
		return fmt.Errorf("InterfaceStrategy[%d] is not valid: %w", c.InterfaceStrategy, ErrConfigNotValid)
	}

	switch c.MapStrategy {
	case MapStrategyDeepCopyKey, MapStrategyBitwiseCopyKey:
		break
	default:
		return fmt.Errorf("MapStrategy[%d] is not valid: %w", c.MapStrategy, ErrConfigNotValid)
	}

	return nil
}

type InterfaceStrategy uint8

const (
	InterfaceStrategyBitwiseCopy    InterfaceStrategy = 0
	InterfaceStrategySetNil         InterfaceStrategy = 1
	InterfaceStrategyDeepCopyUnsafe InterfaceStrategy = 2
)

type MapStrategy uint8

const (
	MapStrategyBitwiseCopyKey MapStrategy = 0
	MapStrategyDeepCopyKey    MapStrategy = 1
)
