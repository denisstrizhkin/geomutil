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

func (eq *EventQueue[T]) Put(item T) {
	eq.items = append(eq.items, item)
}

func (eq *EventQueue[T]) Pick() T {
	return eq.items[len(eq.items)-1]
}

func (eq *EventQueue[T]) Pop() T {
	item := eq.Pick()
	eq.items = eq.items[:len(eq.items)-1]
	return item
}

func (eq *EventQueue[T]) String() string {
	return fmt.Sprint(eq.items)
}
