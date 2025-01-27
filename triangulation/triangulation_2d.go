package triangulation

import (
	"cmp"
	"errors"
	"math"

	u "github.com/denisstrizhkin/geomutil/util"
)

const BOUNDING_POINT_LEFT = -2
const BOUNDING_POINT_RIGHT = -1

type Triangle2D struct {
	A u.Point2D
	B u.Point2D
	C u.Point2D
}

type triangle2DNode struct {
	// Vertices of the triangle
	Ai int
	Bi int
	Ci int

	// Child triangles
	ChildA int
	ChildB int
	ChildC int

	// Adjacent triangles
	AdjA int
	AdjB int
	AdjC int
}

func newTriangle2DNode(ai, bi, ci int) triangle2DNode {
	return triangle2DNode{
		Ai: ai, Bi: bi, Ci: ci,
		ChildA: -1, ChildB: -1, ChildC: -1,
		AdjA: -1, AdjB: -1, AdjC: -1,
	}
}

type Triangulator2D struct {
	points []u.Point2D
	nodes  []triangle2DNode
}

func NewTriangulator2D(points []u.Point2D) (*Triangulator2D, error) {
	points = u.Point2DUnique(points)
	if len(points) < 3 {
		return nil, errors.New("less than 3 unique points")
	}
	return &Triangulator2D{points, make([]triangle2DNode, 0)}, nil
}

func (t *Triangulator2D) Triangulate() Triangulation2D {
	t.insertBoundingTriangle()
	t.runBowyerWatson()
	return Triangulation2D{}
}

func (t *Triangulator2D) insertBoundingTriangle() {
	highestPoint := t.findHighestPoint()
	t.nodes = append(t.nodes, newTriangle2DNode(BOUNDING_POINT_LEFT, BOUNDING_POINT_RIGHT, highestPoint))
}

func (t *Triangulator2D) findHighestPoint() int {
	highest := 0
	for i := range len(t.points) {
		if t.cmpPoints(i, highest) > 0 {
			highest = i
		}
	}
	return highest
}

func (t *Triangulator2D) cmpPoints(ai int, bi int) int {
	if ai == BOUNDING_POINT_LEFT || bi == BOUNDING_POINT_LEFT || ai == BOUNDING_POINT_RIGHT || bi == BOUNDING_POINT_RIGHT {
		return cmp.Compare(ai, bi)
	}
	a := t.points[ai]
	b := t.points[bi]
	diff := cmp.Compare(a.Y, b.Y)
	if diff == 0 {
		return cmp.Compare(a.X, b.X)
	}
	return diff
}

func (t *Triangulator2D) runBowyerWatson() {
	for pi := range len(t.points) {
		if pi == t.nodes[0].Ci {
			continue
		}

		// Index of the containing triangle
		trii := t.findTriangleNode(pi)
		tri := &t.nodes[trii]

		// indices of the newly created triangles.
		new_tri_ai := len(t.nodes)
		new_tri_bi := new_tri_ai + 1
		new_tri_ci := new_tri_bi + 2

		// The new triangles! All in CCW order
		new_tri_a := newTriangle2DNode(pi, tri.Ai, tri.Bi)
		new_tri_b := newTriangle2DNode(pi, tri.Bi, tri.Ci)
		new_tri_c := newTriangle2DNode(pi, tri.Ci, tri.Ai)

		// Setting the adjacency triangle references.  Only way to make
		// sure you do this right is by drawing the triangles up on a
		// piece of paper.
		new_tri_a.AdjA = tri.AdjC
		new_tri_b.AdjA = tri.AdjA
		new_tri_c.AdjA = tri.AdjB

		new_tri_a.AdjB = new_tri_bi
		new_tri_b.AdjB = new_tri_ci
		new_tri_c.AdjB = new_tri_ai

		new_tri_a.AdjC = new_tri_ci
		new_tri_b.AdjC = new_tri_ai
		new_tri_c.AdjC = new_tri_bi

		// The new triangles are the children of the old one.
		tri.ChildA = new_tri_ai
		tri.ChildB = new_tri_bi
		tri.ChildC = new_tri_ci

		t.nodes = append(t.nodes, new_tri_a)
		t.nodes = append(t.nodes, new_tri_b)
		t.nodes = append(t.nodes, new_tri_c)

		// if (nt0.A0 != -1) LegalizeEdge(nti0, nt0.A0, pi, p0, p1);
		// if (nt1.A0 != -1) LegalizeEdge(nti1, nt1.A0, pi, p1, p2);
		// if (nt2.A0 != -1) LegalizeEdge(nti2, nt2.A0, pi, p2, p0);
	}
}

func (t *Triangulator2D) findTriangleNode(_ int) int {
	return -1
}

type Triangulation2D struct {
	points    []u.Point2D
	triangles []Triangle2D
}

func (t *Triangulation2D) Points() []u.Point2D {
	return t.points
}

func (t *Triangulation2D) Triangles() []Triangle2D {
	return t.triangles
}

// func (t *Triangle2D) isInsideCircumcircle(p u.Point2D) bool {
// 	det := t.A.X*(t.B.Y*t.B.X*t.B.X+t.C.Y*t.C.X*t.C.X+p.Y*p.X*p.X) -
// 		t.A.Y*(t.B.X*t.B.X+t.C.X*t.C.X+p.X*p.X) +
// 		(t.B.X*t.B.Y*t.B.X + t.C.X*t.C.Y*t.C.X + p.X*p.Y) -
// 		(t.B.X*t.B.Y*t.C.X*t.C.Y + t.C.X*t.C.Y*t.B.X*p.Y + p.X*t.B.Y*t.C.Y)
// 	return det > 0
// }

func DegreesToRadians(deg float32) float32 {
	return deg * math.Pi / 180.0
}
