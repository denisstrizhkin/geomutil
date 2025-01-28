package geomutil

import (
	"log"
	"sort"

	util "github.com/denisstrizhkin/geomutil/util"
)

func checkTurn(p, q util.Point2D) float32 {
	return p.X*q.Y - p.Y*q.X
}

func isLeftTurn(a, b, c util.Point2D) bool {
	ab := a.Subtract(b)
	bc := b.Subtract(c)
	check := checkTurn(ab, bc)
	log.Println(check)
	return check <= 0
}

type ConvexHull struct {
	points []util.Point2D
}

func (ch *ConvexHull) Points() []util.Point2D {
	return ch.points
}

func NewConvexHull(points []util.Point2D) *ConvexHull {
	points = util.Point2DUnique(points)
	sort.Sort(util.ByPoint2DX(points))
	log.Println(points)
	L_up := []util.Point2D{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		L_up = append(L_up, points[i])
		for len(L_up) > 2 && isLeftTurn(L_up[len(L_up)-1], L_up[len(L_up)-2], L_up[len(L_up)-3]) {
			L_up = append(L_up[:len(L_up)-2], L_up[len(L_up)-1])
		}
		log.Println(points)
	}
	L_low := []util.Point2D{points[len(points)-1], points[len(points)-2]}
	for i := len(points) - 3; i >= 0; i-- {
		L_low = append(L_low, points[i])
		for len(L_low) > 2 && isLeftTurn(L_low[len(L_low)-1], L_low[len(L_low)-2], L_low[len(L_low)-3]) {
			L_low = append(L_low[:len(L_low)-2], L_low[len(L_low)-1])
		}
	}
	return &ConvexHull{append(L_up, L_low[1:len(L_low)-1]...)}
}
