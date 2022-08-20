package caster

import (
	stderror "errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/snowphoenix0105/bang/internal/collections/stack"
)

type FieldInfo struct {
	Name string
	Kind reflect.Kind
}

func (fi *FieldInfo) String() string {
	return fmt.Sprintf("(%s)%s", fi.Kind.String(), fi.Name)
}

type FieldCheckError struct {
	In            FieldInfo
	Out           FieldInfo
	MismatchError error
}

func (e *FieldCheckError) Unwrap() error {
	return e.MismatchError
}

func (e *FieldCheckError) Error() string {
	return fmt.Sprintf("field %s cannot cast to %s because %s", e.In.String(), e.Out.String(), e.MismatchError.Error())
}

var (
	MismatchType       = stderror.New("type mismatch")
	MismatchFieldName  = stderror.New("field-name mismatch")
	MismatchFieldCount = stderror.New("field-count mismatch")
)

func check(conf *Config, inPath, outPath *stack.StringStack, inType, outType reflect.Type) error {
	newError := func(err error) error {
		if inPath.Depth() == 0 {
			return err
		}
		return &FieldCheckError{
			In: FieldInfo{
				Name: strings.Join(inPath.Raw(), "."),
				Kind: inType.Kind(),
			},
			Out: FieldInfo{
				Name: strings.Join(outPath.Raw(), "."),
				Kind: outType.Kind(),
			},
			MismatchError: err,
		}
	}

	inKind := inType.Kind()
	outKind := outType.Kind()

	if inKind != outKind {
		if inKind == reflect.Pointer {
			if conf.EnablePtrToUintptr && outKind == reflect.Uintptr {
				return nil
			}
			if conf.EnablePtrToUnsafePtr && outKind == reflect.UnsafePointer {
				return nil
			}
		}

		return newError(MismatchType)
	}

	switch inKind {
	default:
		if inType == outType {
			return nil
		}
		return newError(MismatchType)

	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Complex64, reflect.Complex128,
		reflect.Float32, reflect.Float64,
		reflect.UnsafePointer, reflect.Uintptr:
		return nil

	case reflect.Pointer:
		inPath.Push("->")
		outPath.Push("->")
		ret := check(conf, inPath, outPath, inType.Elem(), outType.Elem())
		inPath.Pop()
		outPath.Pop()
		return ret

	case reflect.Slice, reflect.Array:
		inPath.Push("[]")
		outPath.Push("[]")
		ret := check(conf, inPath, outPath, inType.Elem(), outType.Elem())
		inPath.Pop()
		outPath.Pop()
		return ret

	case reflect.Struct:
		fieldNum := inType.NumField()
		if fieldNum != outType.NumField() {
			return newError(MismatchFieldCount)
		}

		for i := 0; i < fieldNum; i++ {
			inField := inType.Field(i)
			outField := outType.Field(i)

			switch conf.FieldNameStrategy {
			case FieldNameStrategyStrict:
				if inField.Name != outField.Name {
					return newError(MismatchFieldName)
				}
			case FieldNameStrategyIgnoreCase:
				if !strings.EqualFold(inField.Name, outField.Name) {
					return newError(MismatchFieldName)
				}
			case FieldNameStrategyIgnore:
				break
			}

			inPath.Push(inField.Name)
			outPath.Push(outField.Name)
			err := check(conf, inPath, outPath, inField.Type, outField.Type)
			inPath.Pop()
			outPath.Pop()
			if err != nil {
				return err
			}
		}

		return nil
	}
}
