package util

type Edge2D struct {
	A Point2D
	B Point2D
}

func NewEdge2D(a Point2D, b Point2D) Edge2D {
	return Edge2D{a, b}
}

func (e *Edge2D) Rotate() Edge2D {
	return NewEdge2D(e.B, e.A)
}

func (e *Edge2D) Length() float32 {
	return e.A.Distance(e.B)
}
