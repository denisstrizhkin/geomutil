package triangulation

import (
	u "github.com/denisstrizhkin/geomutil/util"
)

type QuadEdge struct {
	refs [4]*QuarterEdge
}

type QuarterEdge struct {
	current       QuadEdge
	current_index int

	next *QuarterEdge

	vertex u.Point2D
}

// Origin vertex
func (q *QuarterEdge) Orig() u.Point2D {
	return q.vertex
}

// Destination vertex
func (q *QuarterEdge) Dest() u.Point2D {
	return q.Sym().vertex
}

func (q *QuarterEdge) Rot() *QuarterEdge {
	return q.current.refs[(q.current_index+1)%4]
}

func (q *QuarterEdge) Sym() *QuarterEdge {
	return q.current.refs[(q.current_index+2)%4]
}

func (q *QuarterEdge) Tor() *QuarterEdge {
	return q.current.refs[(q.current_index+3)%4]
}

// Next edge around the origin
func (q *QuarterEdge) ONext() *QuarterEdge {
	return q.next
}

func (q *QuarterEdge) SetONext(e *QuarterEdge) {
	q.next = e
}

// Next edge around the left face
func (q *QuarterEdge) LNext() *QuarterEdge {
	return q.Tor().ONext().Rot()
}

// Next edge around the right face
func (q *QuarterEdge) RNext() *QuarterEdge {
	return q.Rot().ONext().Tor()
}

// Previous edge around the origin
func (q *QuarterEdge) OPrev() *QuarterEdge {
	return q.Rot().ONext().Rot()
}

func MakeQuadEdge(start u.Point2D, end u.Point2D) *QuarterEdge {
	start_end := &QuarterEdge{}
	right_left := &QuarterEdge{}
	end_start := &QuarterEdge{}
	left_right := &QuarterEdge{}

	quad_edge := QuadEdge{
		[4]*QuarterEdge{start_end, right_left, end_start, left_right},
	}

	start_end.vertex = start
	start_end.current = quad_edge
	start_end.current_index = 0
	start_end.next = start_end

	right_left.current = quad_edge
	right_left.current_index = 1
	right_left.next = left_right

	end_start.vertex = end
	end_start.current = quad_edge
	end_start.current_index = 2
	end_start.next = end_start

	left_right.current = quad_edge
	left_right.current_index = 3
	left_right.next = right_left

	return start_end
}

func Splice(a *QuarterEdge, b *QuarterEdge) {
	SwapNexts(a.ONext().Rot(), b.ONext().Rot())
	SwapNexts(a, b)
}

func SwapNexts(a *QuarterEdge, b *QuarterEdge) {
	a.next, b.next = b.next, a.next
}

func MakeTriangle(a u.Point2D, b u.Point2D, c u.Point2D) *QuarterEdge {
	ab := MakeQuadEdge(a, b)
	bc := MakeQuadEdge(b, c)
	ca := MakeQuadEdge(c, a)

	Splice(ab.Sym(), bc)
	Splice(bc.Sym(), ca)
	Splice(ca.Sym(), ab)

	return ab
}

func Connect(a *QuarterEdge, b *QuarterEdge) *QuarterEdge {
	new_edge := MakeQuadEdge(a.Dest(), b.vertex)
	Splice(new_edge, a.LNext())
	Splice(new_edge.Sym(), b)
	return new_edge
}

func Sever(e *QuarterEdge) {
	Splice(e, e.OPrev())
	Splice(e.Sym(), e.Sym().OPrev())
}

func InsertPoint(e *QuarterEdge, p u.Point2D) *QuarterEdge {
	first_spoke := MakeQuadEdge(e.vertex, p)
	Splice(first_spoke, e)
	spoke := first_spoke
	for {
		spoke = Connect(e, spoke.Sym())
		e = spoke.OPrev()
		if e.LNext() != first_spoke {
			break
		}
	}
	return first_spoke
}
