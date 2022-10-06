package value

import "github.com/snowphoenix0105/bang/internal/types/equal"

func Zero[T any]() (res T) {
	return
}

func IsNotZero[T any](val T) bool {
	return !IsZero(val)
}

func IsZero[T any](val T) bool {
	return equal.Of(Zero[T](), val)
}
