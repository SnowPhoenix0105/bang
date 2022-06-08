package deepcopy

import "fmt"

type Option struct {
	InterfaceStrategy InterfaceStrategy
}

func (o *Option) CheckValid() error {
	switch o.InterfaceStrategy {
	case InterfaceStrategyBitwiseCopy, InterfaceStrategySetNil, InterfaceStrategyDeepCopyUnsafe:
		break
	default:
		return fmt.Errorf("InterfaceStrategy[%d] is not valid: %w", o.InterfaceStrategy, ErrOptionNotValid)
	}

	return nil
}

func NewDefaultOption() *Option {
	return &Option{}
}

type InterfaceStrategy uint8

const (
	InterfaceStrategyBitwiseCopy    InterfaceStrategy = 0
	InterfaceStrategySetNil         InterfaceStrategy = 1
	InterfaceStrategyDeepCopyUnsafe InterfaceStrategy = 2
)
