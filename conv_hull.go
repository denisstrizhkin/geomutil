package geomutil

import (
	"fmt"
	"sort"
)

type ConvexHull struct {
	Points []Point
}

func NewConvexHull(points []Point) *ConvexHull {
	sort.Sort(ByPointX(points))
	fmt.Println(points)
	is_left_turn := func(a, b, c Point) bool {
		fmt.Println(a.Subtract(b).VectMult(b.Subtract(c)))
		return a.Subtract(b).VectMult(b.Subtract(c)) <= 0
	}
	L_up := []Point{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		L_up = append(L_up, points[i])
		for len(L_up) > 2 && is_left_turn(L_up[len(L_up)-1], L_up[len(L_up)-2], L_up[len(L_up)-3]) {
			L_up = append(L_up[:len(L_up)-2], L_up[len(L_up)-1])
		}
		fmt.Println(points)
	}
	L_low := []Point{points[len(points)-1], points[len(points)-2]}
	for i := len(points) - 3; i >= 0; i-- {
		L_low = append(L_low, points[i])
		for len(L_low) > 2 && is_left_turn(L_low[len(L_low)-1], L_low[len(L_low)-2], L_low[len(L_low)-3]) {
			L_low = append(L_low[:len(L_low)-2], L_low[len(L_low)-1])
		}
	}
	return &ConvexHull{append(L_up, L_low[1:len(L_low)-1]...)}
}
