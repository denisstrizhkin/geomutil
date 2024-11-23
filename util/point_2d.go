package util

import (
	"math"
)

type Point2D struct {
	X, Y float32
}

func NewPoint2D(X float32, Y float32) Point2D {
	return Point2D{X: X, Y: Y}
}

func (p Point2D) Length() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

func (p Point2D) Scale(a float32) Point2D {
	return Point2D{p.X * a, p.Y * a}
}

func (p Point2D) Add(q Point2D) Point2D {
	return Point2D{p.X + q.X, p.Y + q.Y}
}

func (p Point2D) Subtract(q Point2D) Point2D {
	return Point2D{p.X - q.X, p.Y - q.Y}
}

func (p Point2D) Multiply(q Point2D) Point2D {
	return Point2D{p.X * q.Y, p.X * q.Y}
}

func (p Point2D) Distance(q Point2D) float32 {
	d := p.Subtract(q)
	return d.Length()
}

func Point2DAvg(points []Point2D) Point2D {
	avg := Point2D{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(1.0 / float32(len(points)))
}
