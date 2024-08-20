package geomutil

import "errors"

type Stack[T any] struct {
	items []T
}

var ErrEmptyStack = errors.New("stack is empty")

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Pop() (T, error) {
	l := len(s.items)
	var t T
	if l == 0 {
		return t, ErrEmptyStack
	}
	t = s.items[l-1]
	s.items = s.items[:l-1]
	return t, nil
}

func (s *Stack[T]) Peek() (T, error) {
	l := len(s.items)
	var t T
	if l == 0 {
		return t, ErrEmptyStack
	}
	return s.items[l-1], nil
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}
