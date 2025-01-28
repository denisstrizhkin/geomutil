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
	log.Print(check)
	return check <= 0
}

func upperBoundary(points []util.Point2D) []util.Point2D {
	upper := []util.Point2D{points[0], points[1]}
	for i := 2; i < len(points); i++ {
		upper = append(upper, points[i])
		for len(upper) > 2 {
			ai := len(upper) - 1
			bi := len(upper) - 2
			ci := len(upper) - 3
			if !isLeftTurn(upper[ai], upper[bi], upper[ci]) {
				break
			}
			upper = append(upper[:bi], upper[ai])
		}
		log.Print(points)
	}
	return upper
}

func lowerBoundary(points []util.Point2D) []util.Point2D {
	lower := []util.Point2D{points[len(points)-1], points[len(points)-2]}
	for i := len(points) - 3; i >= 0; i-- {
		lower = append(lower, points[i])
		for len(lower) > 2 {
			ai := len(lower) - 1
			bi := len(lower) - 2
			ci := len(lower) - 3
			if !isLeftTurn(lower[ai], lower[bi], lower[ci]) {
				break
			}
			lower = append(lower[:bi], lower[ai])
		}
		log.Print(points)
	}
	return lower
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
	log.Print(points)
	upper := upperBoundary(points)
	lower := lowerBoundary(points)
	return &ConvexHull{append(upper, lower[1:len(lower)-1]...)}
}
