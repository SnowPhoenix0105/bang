package caster

import (
	"reflect"
	"unsafe"
)

type FieldNameStrategy uint8

const (
	FieldNameStrategyStrict     = 0
	FieldNameStrategyIgnore     = 1
	FieldNameStrategyIgnoreCase = 2
)

type Config struct {
	CastFunc             CastFunc
	FieldNameStrategy    FieldNameStrategy
	EnablePtrToUintptr   bool
	EnablePtrToUnsafePtr bool
}

func NewDefaultConfig() *Config {
	return &Config{
		CastFunc: DefaultCastFunc,
	}
}

type CastFunc func(outType reflect.Type, addr unsafe.Pointer) interface{}

func DefaultCastFunc(outType reflect.Type, addr unsafe.Pointer) interface{} {
	return reflect.NewAt(outType, addr).Elem().Interface()
}
