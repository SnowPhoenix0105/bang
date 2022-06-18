package walk

import (
	"reflect"
)

type NodeCallback func(ctx *NodeContext) error

type NodeContext struct {
	Path    string
	Address uintptr
	Type    reflect.Type
	Reflect reflect.Value
}
