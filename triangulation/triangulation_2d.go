package triangulation

import (
	"errors"
	"slices"

	u "github.com/denisstrizhkin/geomutil/util"
)

func Circumcenter(a, b, c u.Point2D) u.Point2D {
	D := 2 * (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y))

	Ux := ((a.X*a.X+a.Y*a.Y)*(b.Y-c.Y) +
		(b.X*b.X+b.Y*b.Y)*(c.Y-a.Y) +
		(c.X*c.X+c.Y*c.Y)*(a.Y-b.Y)) / D
	Uy := ((a.X*a.X+a.Y*a.Y)*(c.X-b.X) +
		(b.X*b.X+b.Y*b.Y)*(a.X-c.X) +
		(c.X*c.X+c.Y*c.Y)*(b.X-a.X)) / D

	return u.NewPoint2D(Ux, Uy)
}

type Triangle2D struct {
	A             u.Point2D
	B             u.Point2D
	C             u.Point2D
	Center        u.Point2D
	RadiusSquared float32
}

func NewTriangle2D(a, b, c u.Point2D) Triangle2D {
	center := Circumcenter(a, b, c)
	radiusSquared := center.DistanceSquared(a)
	return Triangle2D{a, b, c, center, radiusSquared}
}

func (t *Triangle2D) isInsideCircumcircle(p u.Point2D) bool {
	d := t.Center.DistanceSquared(p)
	return d <= t.RadiusSquared
}

func (t *Triangle2D) edges() [3]u.Edge2D {
	return [3]u.Edge2D{
		u.NewEdge2D(t.A, t.B),
		u.NewEdge2D(t.B, t.C),
		u.NewEdge2D(t.C, t.A),
	}
}

func getBoundingTriangle(points []u.Point2D) Triangle2D {
	pMin := u.Point2DMin(points)
	pMax := u.Point2DMax(points)
	d := pMax.Subtract(pMin)
	dMax := 3 * max(d.X, d.Y)
	pCenter := pMin.Add(pMax).Scale(0.5)
	return NewTriangle2D(
		u.NewPoint2D(pCenter.X-0.866*dMax, pCenter.Y-0.5*dMax),
		u.NewPoint2D(pCenter.X+0.866*dMax, pCenter.Y-0.5*dMax),
		u.NewPoint2D(pCenter.X, pCenter.Y+dMax),
	)
}

type Triangulation2D struct {
	points    []u.Point2D
	triangles []Triangle2D
	step      int
}

func NewTriangulation2D(points []u.Point2D) (Triangulation2D, error) {
	points = u.Point2DUnique(points)
	if len(points) < 3 {
		return Triangulation2D{}, errors.New("less than 3 unique points")
	}

	triangulation := Triangulation2D{points, make([]Triangle2D, 0), 0}
	triangulation.setup()

	return triangulation, nil
}

func (t *Triangulation2D) setup() {
	boundingTriangle := getBoundingTriangle(t.points)
	t.triangles = append(t.triangles, boundingTriangle)
}

func (t *Triangulation2D) Points() []u.Point2D {
	return t.points
}

func (t *Triangulation2D) Triangles() []Triangle2D {
	return t.triangles
}

func (t *Triangulation2D) Step() {
	point := t.points[t.step]

	badTriangles := make([]Triangle2D, 0)
	for _, triangle := range t.triangles {
		if triangle.isInsideCircumcircle(point) {
			badTriangles = append(badTriangles, triangle)
		}
	}
	edges := make(map[u.Edge2D]int, len(badTriangles)*3)
	for _, triangle := range badTriangles {
		for _, edge := range triangle.edges() {
			edges[edge] += 1
			rotated := edge.Rotate()
			if _, exists := edges[rotated]; exists {
				edges[rotated] += 1
			}
		}
	}
	polygon := make([]u.Edge2D, len(badTriangles))
	for edge, count := range edges {
		if count == 1 {
			polygon = append(polygon, edge)
		}
	}

	for _, triangle := range badTriangles {
		i := slices.Index(t.triangles, triangle)
		t.triangles = append(t.triangles[:i], t.triangles[i+1:]...)
	}
	for _, edge := range polygon {
		t.triangles = append(t.triangles, NewTriangle2D(edge.A, edge.B, point))
	}

	t.step += 1
}
