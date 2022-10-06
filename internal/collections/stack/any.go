package stack

import (
	"github.com/snowphoenix0105/bang/internal/collections/base"
	"github.com/snowphoenix0105/bang/internal/types/slice"
)

type Stack[T any] struct {
	container base.BackCollection[T]
}

func New[T any]() *Stack[T] {
	return NewWithCollection[T](slice.NewWithCap[T](defaultCapacity))
}

func NewWithCollection[T any](container base.BackCollection[T]) *Stack[T] {
	return &Stack[T]{container}
}

func (s *Stack[T]) Depth() int {
	return s.container.Len()
}

func (s *Stack[T]) Push(elem T) {
	s.container.PushBack(elem)
}

func (s *Stack[T]) Peak() T {
	return s.container.GetBack()
}

func (s *Stack[T]) Pop() {
	s.container.PopBack()
}
