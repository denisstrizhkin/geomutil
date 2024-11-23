package triangulation

import (
	util "github.com/denisstrizhkin/geomutil/util"
	"math"
)

type Triangle2D struct {
	A util.Point2D
	B util.Point2D
	C util.Point2D
}

func degreesToRadians(deg float32) float32 {
	return deg * math.Pi / 180.0
}

func getSuperTriangle(points []util.Point2D) Triangle2D {
	center := util.Point2DAvg(points)
	radius := float32(math.Inf(-1))
	for _, p := range points {
		dist := center.Distance(p)
		if dist > radius {
			radius = dist
		}
	}
	radius += 1.0
	rad := float64(degreesToRadians(30.0))
	half_median := radius / float32(math.Sin(rad))
	half_side := radius / float32(math.Cos(rad))
	return Triangle2D{
		util.NewPoint2D(center.X-half_side, center.Y-radius),
		util.NewPoint2D(center.X+half_side, center.Y-radius),
		util.NewPoint2D(center.X, center.Y+half_median),
	}
}

type Triangulation2D struct {
	triangles   []Triangle2D
	points      []util.Point2D
	edge_points []util.Point2D
}

func NewTriangulation2D(points []util.Point2D) Triangulation2D {
	superTriangle := getSuperTriangle(points)
	return Triangulation2D{
		triangles:   []Triangle2D{superTriangle},
		points:      points,
		edge_points: []util.Point2D{superTriangle.A, superTriangle.B, superTriangle.C},
	}
}
