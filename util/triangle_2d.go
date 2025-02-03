package util

type Triangle2D struct {
	A Point2D
	B Point2D
	C Point2D
}

func NewTriangle2D(a, b, c Point2D) Triangle2D {
	return Triangle2D{a, b, c}
}

func (t *Triangle2D) Circumcenter() Point2D {
	a := t.A
	b := t.B
	c := t.C
	D := 2 * (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y))

	Ux := ((a.X*a.X+a.Y*a.Y)*(b.Y-c.Y) +
		(b.X*b.X+b.Y*b.Y)*(c.Y-a.Y) +
		(c.X*c.X+c.Y*c.Y)*(a.Y-b.Y)) / D
	Uy := ((a.X*a.X+a.Y*a.Y)*(c.X-b.X) +
		(b.X*b.X+b.Y*b.Y)*(a.X-c.X) +
		(c.X*c.X+c.Y*c.Y)*(b.X-a.X)) / D

	return NewPoint2D(Ux, Uy)
}

func (t *Triangle2D) CircumcircleRadiusSquared() float32 {
	return t.A.DistanceSquared(t.Circumcenter())
}

func (t *Triangle2D) IsInsideCircumcircle(p Point2D) bool {
	d := t.Circumcenter().DistanceSquared(p)
	return d <= t.CircumcircleRadiusSquared()
}

func (t *Triangle2D) HasPoint(p Point2D) bool {
	return t.A == p || t.B == p || t.C == p
}

func (t *Triangle2D) Edges() [3]Edge2D {
	return [3]Edge2D{
		NewEdge2D(t.A, t.B),
		NewEdge2D(t.B, t.C),
		NewEdge2D(t.C, t.A),
	}
}

func (t *Triangle2D) Volume() float32 {
	edges := t.Edges()
	a := edges[0].Length()
	b := edges[1].Length()
	c := edges[2].Length()
	s := 0.5 * (a + b + c)
	return Sqrt(s * (s - a) * (s - b) * (s - c))
}
