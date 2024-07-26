package geomutil

import (
	"errors"
)

type Triangle struct {
	A, B, C Point
}

type Vertice struct {
	A, B Point
}

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

type CDT struct {
	triangles *Stack[Triangle]
}

func NewCDT() *CDT {
	return &CDT{triangles: NewStack[Triangle]()}
}

func (cdt *CDT) AddPoint(p Point) {
	ntriangles := NewStack[Triangle]()
	// find triangle that contains p
	// devide it into t1, t2, t3
	var t1, t2, t3 Triangle
	ntriangles.Push(t1)
	ntriangles.Push(t2)
	ntriangles.Push(t3)
	for ntriangles.Length() != 0 {
		t, err := ntriangles.Pop()
		if err != nil {
			break
		}
		var t_op Triangle
		// t_op = cdt.OpposedTriangle(t, p)
		if true {
			ntriangles.Push(t)
			ntriangles.Push(t_op)
		}
	}
}

func (cdt *CDT) Triangulate(points []Point) {
	for _, p := range points {
		cdt.AddPoint(p)
	}
}

func (cdt *CDT) Vertices() []Vertice {
	vertices := make([]Vertice, 0)
	return vertices
}
