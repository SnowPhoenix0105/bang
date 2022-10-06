package slice

import (
	"github.com/snowphoenix0105/bang/internal/types/equal"
	"github.com/snowphoenix0105/bang/internal/types/value"
)

type Slice[T any] []T

func (s *Slice[T]) Raw() []T {
	return *s
}

func (s *Slice[T]) RawPtr() *[]T {
	return (*[]T)(s)
}

func (s *Slice[T]) IndexIsValid(index int) bool {
	return index >= 0 && index < s.Len()
}

func (s *Slice[T]) Get(index int) T {
	return (*s)[index]
}

func (s *Slice[T]) TryGet(index int) (T, bool) {
	if !s.IndexIsValid(index) {
		return value.Zero[T](), false
	}

	return s.Get(index), true
}

func (s *Slice[T]) GetOrDefault(index int) T {
	return s.GetOr(index, value.Zero[T]())
}

func (s *Slice[T]) GetOr(index int, defaultValue T) T {
	if !s.IndexIsValid(index) {
		return defaultValue
	}

	return s.Get(index)
}

func (s *Slice[T]) Set(index int, elem T) {
	(*s)[index] = elem
}

func (s *Slice[T]) TrySet(index int, elem T) bool {
	if !s.IndexIsValid(index) {
		return false
	}

	s.Set(index, elem)
	return true
}

func (s *Slice[T]) SetLen(length int) bool {
	if length < 0 || length > s.Len() {
		return false
	}

	*s = (*s)[:length]
	return true
}

func (s *Slice[T]) Append(elems ...T) {
	*s = append(*s, elems...)
}

func (s *Slice[T]) ContainsKey(index int) bool {
	return s.IndexIsValid(index)
}

func (s *Slice[T]) ContainsValue(elem T) bool {
	return s.Any(func(e T) bool {
		return equal.Of(e, elem)
	})
}

func (s *Slice[T]) Contains(elem T) bool {
	return s.ContainsValue(elem)
}

func (s *Slice[T]) All(fn func(T) bool) bool {
	return s.All2(func(_ int, e T) bool {
		return fn(e)
	})
}

func (s *Slice[T]) All2(fn func(int, T) bool) bool {
	return !s.Any2(func(i int, e T) bool {
		return !fn(i, e)
	})
}

func (s *Slice[T]) Any(fn func(T) bool) bool {
	return s.Any2(func(_ int, e T) bool {
		return fn(e)
	})
}

func (s *Slice[T]) Any2(fn func(int, T) bool) bool {
	for i, e := range *s {
		if fn(i, e) {
			return true
		}
	}

	return false
}

func (s *Slice[T]) Find(elem T) (int, bool) {
	return s.FindFn(func(e T) bool {
		return equal.Of(e, elem)
	})
}

func (s *Slice[T]) FindFn(fn func(T) bool) (int, bool) {
	for i, e := range *s {
		if fn(e) {
			return i, true
		}
	}

	return 0, false
}

func (s *Slice[T]) Len() int {
	return len(*s)
}

func (s *Slice[T]) Cap() int {
	return cap(*s)
}

func (s *Slice[T]) PushBack(elem T) {
	*s = append(*s, elem)
}

func (s *Slice[T]) PopBack() {
	s.SetLen(s.Len() - 1)
}

func (s *Slice[T]) GetBack() T {
	return s.GetOrDefault(s.Len() - 1)
}

func (s *Slice[T]) ForAll(fn func(T)) {
	s.ForAll2(func(_ int, e T) {
		fn(e)
	})
}

func (s *Slice[T]) ForAll2(fn func(int, T)) {
	s.ForEach2(func(i int, e T) bool {
		fn(i, e)
		return true
	})
}

func (s *Slice[T]) ForEach(fn func(T) bool) {
	s.ForEach2(func(_ int, e T) bool {
		return fn(e)
	})
}

func (s *Slice[T]) ForEach2(fn func(int, T) bool) {
	for i, e := range *s {
		if !fn(i, e) {
			break
		}
	}
}
