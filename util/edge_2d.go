package util

type Edge2D struct {
	A Point2D
	B Point2D
}

func NewEdge2D(a Point2D, b Point2D) Edge2D {
	return Edge2D{a, b}
}
