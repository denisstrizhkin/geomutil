package geomutil

import (
	"sort"
)

type ConvexHull struct {
	points []Point
}

func (p Point) Subtract(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (vec1 Point) VectMult(vec2 Point) float64 {
	return (vec1.X*vec2.Y - vec1.Y*vec2.X)
}

func NewConvexHull(points []Point) *ConvexHull {
	sort.Sort(ByPointX(points))
	L_up := make([]Point, 0)
	L_up = append(L_up, points[0], points[1])
	for i := 2; i < len(points); i++ {
		L_up = append(L_up, points[i])
		right_turn := false
		for len(L_up) > 2 && right_turn == false {
			right_turn = true
			if points[i].Subtract(points[i-1]).VectMult(points[i-1].Subtract(points[i-2])) <= 0 {
				right_turn = false
				L_up = append(L_up[:i-2], L_up[i])
			}
		}
	}
	L_low := make([]Point, 0)
	L_low = append(L_low, points[len(points)-1], points[len(points)-2])
	for i := len(points) - 3; i >= 0; i-- {
		L_low = append(L_low, points[i])
		right_turn := false
		for len(L_low) > 2 && right_turn == false {
			right_turn = true
			if points[i].Subtract(points[i+1]).VectMult(points[i+1].Subtract(points[i+2])) <= 0 {
				right_turn = false
				L_low = append(L_low[:i-2], L_low[i])
			}
		}
	}
	L := append(L_up, L_low[1:len(L_low)-1]...)
	return &ConvexHull{L}
}
