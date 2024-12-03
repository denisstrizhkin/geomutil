package triangulation

import (
	u "github.com/denisstrizhkin/geomutil/util"
	"math"
	"slices"
)

type Triangle2D struct {
	A u.Point2D
	B u.Point2D
	C u.Point2D
}

func (t *Triangle2D) isInsideCircumcircle(p u.Point2D) bool {
	det := t.A.X*(t.B.Y*t.B.X*t.B.X+t.C.Y*t.C.X*t.C.X+p.Y*p.X*p.X) -
		t.A.Y*(t.B.X*t.B.X+t.C.X*t.C.X+p.X*p.X) +
		(t.B.X*t.B.Y*t.B.X + t.C.X*t.C.Y*t.C.X + p.X*p.Y) -
		(t.B.X*t.B.Y*t.C.X*t.C.Y + t.C.X*t.C.Y*t.B.X*p.Y + p.X*t.B.Y*t.C.Y)
	return det > 0
}

func (t *Triangle2D) Edges() [3]u.Edge2D {
	return [3]u.Edge2D{
		u.NewEdge2D(t.A, t.B),
		u.NewEdge2D(t.B, t.C),
		u.NewEdge2D(t.C, t.A),
	}
}

func (t *Triangle2D) hasEdge(e u.Edge2D) bool {
	for _, edge := range t.Edges() {
		if e == edge {
			return true
		}
	}
	return false
}

func degreesToRadians(deg float32) float32 {
	return deg * math.Pi / 180.0
}

func getSuperTriangle(points []u.Point2D) Triangle2D {
	center := u.Point2DAvg(points)
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
		u.NewPoint2D(center.X-half_side, center.Y-radius),
		u.NewPoint2D(center.X+half_side, center.Y-radius),
		u.NewPoint2D(center.X, center.Y+half_median),
	}
}

type Triangulation2D struct {
	triangles   []Triangle2D
	points      []u.Point2D
	edge_points []u.Point2D
}

func NewTriangulation2D(points []u.Point2D) Triangulation2D {
	superTriangle := getSuperTriangle(points)
	return Triangulation2D{
		triangles:   []Triangle2D{superTriangle},
		points:      points,
		edge_points: []u.Point2D{superTriangle.A, superTriangle.B, superTriangle.C},
	}
}

func (t *Triangulation2D) Triangles() []Triangle2D {
	return t.triangles
}

func (t *Triangulation2D) Step() {
	point := t.points[0]
	t.points = t.points[1:]
	badTriangles := make([]Triangle2D, 0)
	for _, triangle := range t.triangles {
		if triangle.isInsideCircumcircle(point) {
			badTriangles = append(badTriangles, triangle)
		}
	}
	polygon := make([]u.Edge2D, 0)
	for _, triangle := range badTriangles {
		for _, edge := range triangle.Edges() {
			countShared := 0
			for _, triangle := range badTriangles {
				if triangle.hasEdge(edge) {
					countShared += 1
				}
			}
			if countShared == 1 {
				polygon = append(polygon, edge)
			}
		}
	}
	for _, triangle := range badTriangles {
		i := slices.Index(t.triangles, triangle)
		t.triangles = append(t.triangles[:i], t.triangles[i+1:]...)
	}
	for _, edge := range polygon {
		t.triangles = append(t.triangles, Triangle2D{edge.A, edge.B, point})
	}
}
