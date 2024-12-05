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

	next       QuadEdge
	next_index int

	vertex u.Point2D
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

func (q *QuarterEdge) Next() *QuarterEdge {
	return q.next.refs[q.next_index]
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
	start_end.next = quad_edge
	start_end.next_index = 0

	right_left.current = quad_edge
	right_left.current_index = 1
	right_left.next = quad_edge
	right_left.next_index = 3

	end_start.vertex = end
	end_start.current = quad_edge
	end_start.current_index = 2
	end_start.next = quad_edge
	end_start.next_index = 2

	left_right.current = quad_edge
	left_right.current_index = 3
	left_right.next = quad_edge
	left_right.next_index = 1

	return start_end
}
