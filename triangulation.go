package geomutil

import (
	"errors"
)

type Triangle struct {
	A, B, C Point
}

type Edge struct {
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

func (cdt *CDT) Vertices() []Edge {
	vertices := make([]Edge, 0)
	return vertices
}

type DelaunayTriangulation struct {
	triangles []Triangle
}

func NewDelaunayTriangulation() *DelaunayTriangulation {
	return &DelaunayTriangulation{triangles: make([]Triangle, 0)}
}

func (dt *DelaunayTriangulation) FindHighestPoint(points []Point) {
	p0 = points[0]
	for _, p := range points {
		if p.Y > p0.Y {
			p0 = p
		} else if (p.Y == p0.Y) && (p.X < p0.X){
			p0 = p
		}
	}
	return p0
}


func (dt *DelaunayTriangulation) Triangulate(points []Point) {
	// find p0 in points (remove)
	// calculate points p(-1) and p(-2)
	// p0,p(-1),p(-2) form initial triangle containing all points
	p0 = FindHighestPoint(points)
	var p_bot, p_top Point
	for _, pr := range points {
		// find triangle pi,pj,pk containing pr
		if pr == p0 {
			continue
		}
		
		if true { // if pr lies in the interior of pi,pj,pk
			dt.legalizeEdge(pr, pi, pj)
			dt.legalizeEdge(pr, pj, pk)
			dt.legalizeEdge(pr, pk, pi)
		} else { // if pr lies on an edge of pi,pj,pk
			// TODO
			// dt.legalizeEdge(pr, pk, pi)
		}
	}
	// remove p1,p2 and their edges from triangulation
}

func (dt *DelaunayTriangulation) legalizeEdge(p, a, b Point) {
	if true { // edge a,b is illegal
		// a,b,c - triangle adjacent p,a,b
		dt.legalizeEdge(p, a, c)
		dt.legalizeEdge(p, c, b)
	}
}

func (dt *DelaunayTriangulation) Edges() (edges []Edge) {
	edges = make([]Edge, 0)
	// get all edges of resulting triangulation
	return edges
}
