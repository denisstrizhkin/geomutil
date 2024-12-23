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

func (p Point2D) Min(q Point2D) Point2D {
	return Point2D{min(p.X, q.X), min(p.Y, q.Y)}
}

func (p Point2D) Max(q Point2D) Point2D {
	return Point2D{max(p.X, q.X), max(p.Y, q.Y)}
}

func (p Point2D) Scale(a float32) Point2D {
	return Point2D{p.X * a, p.Y * a}
}

func (p Point2D) Add(q Point2D) Point2D {
	return Point2D{p.X + q.X, p.Y + q.Y}
}

func (p Point2D) AddValue(a float32) Point2D {
	return Point2D{p.X + a, p.Y + a}
}

func (p Point2D) Subtract(q Point2D) Point2D {
	return Point2D{p.X - q.X, p.Y - q.Y}
}

func (p Point2D) SubtractValue(a float32) Point2D {
	return Point2D{p.X - a, p.Y - a}
}

func (p Point2D) Multiply(q Point2D) Point2D {
	return Point2D{p.X * q.Y, p.X * q.Y}
}

func (p Point2D) Distance(q Point2D) float32 {
	d := p.Subtract(q)
	return d.Length()
}

func (p Point2D) Rotate(angle float32) Point2D {
	sina := float32(math.Sin(float64(angle)))
	cosa := float32(math.Cos(float64(angle)))
	return Point2D{
		cosa*p.X - sina*p.Y, sina*p.X + cosa*p.Y,
	}
}

func (p Point2D) Negative() Point2D {
	return p.Scale(-1)
}

func (p Point2D) Normalize() Point2D {
	len := p.Length()
	if len > 0 {
		p.X *= 1 / len
		p.Y *= 1 / len
	}
	return p
}

func Point2DAvg(points []Point2D) Point2D {
	avg := Point2D{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(1.0 / float32(len(points)))
}

func Point2DMin(points []Point2D) Point2D {
	pMin := Point2D{float32(math.Inf(1)), float32(math.Inf(1))}
	for _, p := range points {
		pMin = pMin.Min(p)
	}
	return pMin
}

func Point2DMax(points []Point2D) Point2D {
	pMax := Point2D{float32(math.Inf(-1)), float32(math.Inf(-1))}
	for _, p := range points {
		pMax = pMax.Max(p)
	}
	return pMax
}
