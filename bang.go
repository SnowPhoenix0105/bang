package bang

import (
	"github.com/snowphoenix0105/bang/internal/warning"
	"github.com/snowphoenix0105/bang/pkg/bcaster"
	"github.com/snowphoenix0105/bang/pkg/bdeepcopy"
)

func SetOnWarning(fn func(msg string)) {
	warning.SetOnWarning(fn)
}

// bcaster

type Caster[InType any, OutType any] func(InType) OutType

func NewCaster[InType any, OutType any](options ...bcaster.Option) (Caster[InType, OutType], error) {
	fn, err := bcaster.New[InType, OutType](options...)
	return Caster[InType, OutType](fn), err
}

func MustNewCaster[InType any, OutType any](options ...bcaster.Option) Caster[InType, OutType] {
	fn, err := NewCaster[InType, OutType](options...)
	if err != nil {
		panic(err)
	}
	return fn
}

// bdeepcopy

func NewDeepCopier(options ...bdeepcopy.Option) *bdeepcopy.DeepCopier {
	return bdeepcopy.NewDeepCopier(options...)
}

func DeepCopyOf[T any](origin T, options ...bdeepcopy.Option) T {
	return bdeepcopy.Of(origin, options...)
}
