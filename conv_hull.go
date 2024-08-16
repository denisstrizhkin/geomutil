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

func (p Point) VectMult(q Point) float64 {
	return p.X*q.Y - p.Y*q.X
}

func NewConvexHull(points []Point) *ConvexHull {
	sort.Sort(ByPointX(points))
	L_up := points[:2]
	for i := 2; i < len(points); i++ {
		L_up = append(L_up, points[i])
		for {
			v12 := points[i].Subtract(points[i-1])
			v23 := points[i-1].Subtract(points[i-2])
			if v12.VectMult(v23) > 0 {
				L_up = append(L_up[:i-2], L_up[i])
				if len(L_up) > 2 {
					break
				}
			}
		}
	}
	L_low := []Point{points[len(points)-1], points[len(points)-2]}
	for i := len(points) - 3; i >= 0; i-- {
		L_low = append(L_low, points[i])
		for {
			v12 := points[i].Subtract(points[i+1])
			v23 := points[i+1].Subtract(points[i+2])
			if v12.VectMult(v23) > 0 {
				L_low = append(L_low[:i-2], L_low[i])
				if len(L_low) > 2 {
					break
				}
			}
		}
	}
	L := append(L_up, L_low[1:len(L_low)-1]...)
	return &ConvexHull{L}
}
