package triangulation

import (
	"errors"
	"log"
	"slices"

	u "github.com/denisstrizhkin/geomutil/util"
)

func getBoundingTriangle(points []u.Point2D) u.Triangle2D {
	pMin := u.Point2DMin(points)
	pMax := u.Point2DMax(points)
	d := pMax.Subtract(pMin)
	dMax := 3 * max(d.X, d.Y)
	pCenter := pMin.Add(pMax).Scale(0.5)
	return u.NewTriangle2D(
		u.NewPoint2D(pCenter.X-0.866*dMax, pCenter.Y-0.5*dMax),
		u.NewPoint2D(pCenter.X+0.866*dMax, pCenter.Y-0.5*dMax),
		u.NewPoint2D(pCenter.X, pCenter.Y+dMax),
	)
}

type Triangulation2D struct {
	points    []u.Point2D
	triangles []u.Triangle2D
	bounding  u.Triangle2D
}

func NewTriangulation2D(points []u.Point2D) (Triangulation2D, error) {
	points = u.Point2DUnique(points)
	if len(points) < 3 {
		return Triangulation2D{}, errors.New("less than 3 unique points")
	}

	triangulation := Triangulation2D{points: points, triangles: make([]u.Triangle2D, 0)}
	triangulation.runBowyerWatson()

	return triangulation, nil
}

func (t *Triangulation2D) runBowyerWatson() {
	t.bounding = getBoundingTriangle(t.points)
	t.triangles = append(t.triangles, t.bounding)
	for _, point := range t.points {
		t.step(point)
	}
	for i := 0; i < len(t.triangles); {
		tri := t.triangles[i]
		if tri.HasPoint(t.bounding.A) || tri.HasPoint(t.bounding.B) || tri.HasPoint(t.bounding.C) {
			t.triangles = append(t.triangles[:i], t.triangles[i+1:]...)
		} else {
			i++
		}
	}
}

func (t *Triangulation2D) Points() []u.Point2D {
	return t.points
}

func (t *Triangulation2D) Triangles() []u.Triangle2D {
	return t.triangles
}

func (t *Triangulation2D) step(point u.Point2D) {
	badTriangles := make([]u.Triangle2D, 0)
	for _, triangle := range t.triangles {
		if triangle.IsInsideCircumcircle(point) {
			badTriangles = append(badTriangles, triangle)
		}
	}
	log.Printf("Bad triangles (%v): (%v)\n", len(badTriangles), badTriangles)
	edges := make(map[u.Edge2D]int, len(badTriangles)*3)
	for _, triangle := range badTriangles {
		for _, edge := range triangle.Edges() {
			rotated := edge.Rotate()
			if edges[rotated] != 0 {
				edges[rotated] += 1
			} else {
				edges[edge] += 1
			}
		}
	}
	log.Println("edges:", edges)
	polygon := make([]u.Edge2D, 0, len(badTriangles))
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
		t.triangles = append(t.triangles, u.NewTriangle2D(edge.A, edge.B, point))
	}
}
