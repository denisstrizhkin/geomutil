package geomutil

import (
	"fmt"
)

type EventQueue[T any] struct {
	items []T
}

func NewEventQueue[T any]() *EventQueue[T] {
	return &EventQueue[T]{items: make([]T, 0)}
}

func (eq *EventQueue[T]) Enqueue(item T) {
	eq.items = append(eq.items, item)
}

func (eq *EventQueue[T]) Peek() (T, bool) {
	if len(eq.items) == 0 {
		var zero T
		return zero, false
	}
	return eq.items[len(eq.items)-1], true
}

func (eq *EventQueue[T]) Dequeue() (T, bool) {
	if len(eq.items) == 0 {
		var zero T
		return zero, false
	}
	item, _ := eq.Peek()
	eq.items = eq.items[:len(eq.items)-1]
	return item, true
}

func (eq *EventQueue[T]) String() string {
	return fmt.Sprint(eq.items)
}
