package deepcopy

import (
	"fmt"
)

type Config struct {
	InterfaceStrategy InterfaceStrategy
	MapStrategy       MapStrategy
}

func NewDefaultConfig() *Config {
	return &Config{}
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
