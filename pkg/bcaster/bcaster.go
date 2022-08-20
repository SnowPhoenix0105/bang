package bcaster

import (
	"github.com/snowphoenix0105/bang/internal/caster"
)

var globalConfig caster.Config

func init() {
	globalConfig = *caster.NewDefaultConfig()
}

type Func[InType any, OutType any] func(InType) OutType

func MustNew[InType any, OutType any](options ...Option) Func[InType, OutType] {
	fn, err := New[InType, OutType](options...)
	if err != nil {
		panic(err)
	}
	return fn
}

func New[InType any, OutType any](options ...Option) (Func[InType, OutType], error) {
	cpyConfig := globalConfig

	for _, opt := range options {
		opt.Apply(&cpyConfig)
	}

	fn, err := caster.New[InType, OutType](&cpyConfig)
	return Func[InType, OutType](fn), err
}
