package triangulation

type QuadEdge struct {
	refs [4]QuarterEdge
}

type QuarterEdge struct {
	current       *QuadEdge
	current_index int

	next       *QuadEdge
	next_index int
}

func (q *QuarterEdge) Rot() *QuarterEdge {
	return &q.current.refs[(q.current_index+1)%4]
}

func (q *QuarterEdge) Sym() *QuarterEdge {
	return &q.current.refs[(q.current_index+2)%4]
}

func (q *QuarterEdge) Tor() *QuarterEdge {
	return &q.current.refs[(q.current_index+3)%4]
}

func (q *QuarterEdge) Next() *QuarterEdge {
	return &q.next.refs[q.next_index]
}
