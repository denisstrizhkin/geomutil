package util

import (
	"encoding/json"
	"fmt"
	"os"
)

type Point3D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func NewPoint3D(X float32, Y float32, Z float32) Point3D {
	return Point3D{X: X, Y: Y, Z: Z}
}

func Point3DFromFile(path string) ([]Point3D, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %v: %v", path, err)
	}
	defer file.Close()

	var points []Point3D
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&points)
	if err != nil {
		return nil, fmt.Errorf("decoding JSON: %v", err)
	}

	return points, nil
}

func (p Point3D) Length() float32 {
	return Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p Point3D) Min(q Point3D) Point3D {
	return NewPoint3D(min(p.X, q.X), min(p.Y, q.Y), min(p.Z, q.Z))
}

func (p Point3D) Max(q Point3D) Point3D {
	return NewPoint3D(max(p.X, q.X), max(p.Y, q.Y), max(p.Z, q.Z))
}

func (p Point3D) Scale(a float32) Point3D {
	return NewPoint3D(p.X*a, p.Y*a, p.Z*a)
}

func (p Point3D) Add(q Point3D) Point3D {
	return NewPoint3D(p.X+q.X, p.Y+q.Y, p.Z+q.Z)
}

func (p Point3D) AddValue(a float32) Point3D {
	return NewPoint3D(p.X+a, p.Y+a, p.Z+a)
}

func (p Point3D) Subtract(q Point3D) Point3D {
	return Point3D{p.X - q.X, p.Y - q.Y, p.Z - q.Z}
}

func (p Point3D) VectMult(q Point3D) Point3D {
	return Point3D{p.Y*q.Z - p.Z*q.Y, p.X*q.Z - p.Z*q.X, p.X*q.Y - p.Y*q.X}
}

func (p Point3D) SubtractValue(a float32) Point3D {
	return NewPoint3D(p.X-a, p.Y-a, p.Z-a)
}

func (p Point3D) Multiply(q Point3D) Point3D {
	return NewPoint3D(p.X*q.X, p.Y*q.Y, p.Z*p.Z)
}

func (p Point3D) Distance(q Point3D) float32 {
	d := p.Subtract(q)
	return d.Length()
}

func (p Point3D) DistanceSquared(q Point3D) float32 {
	d := p.Subtract(q)
	return d.X*d.X + d.Y*d.Y + d.Z*d.Z
}

func (p Point3D) Rotate(ax Point3D, angle float32) Point3D {
	angle = DegToRad(angle)
	sina := Sin(angle)
	cosa := 1 - Cos(angle)
	x := (1-cosa*(ax.Z*ax.Z+ax.Y*ax.Y))*p.X +
		(-ax.Z*sina+cosa*ax.Y*ax.Z)*p.Y +
		(ax.Y*sina+cosa*ax.X*ax.Z)*p.Z
	y := (ax.Z*sina+cosa*ax.X*ax.Y)*p.X +
		(1-cosa*(ax.Z*ax.Z+ax.X*ax.X))*p.Y +
		(-ax.X*sina+cosa*ax.X*ax.Y)*p.Z
	z := (-ax.Y*sina+cosa*ax.X*ax.Z)*p.X +
		(ax.X*sina+cosa*ax.Y*ax.Z)*p.Y +
		(1-cosa*(ax.X*ax.X+ax.Y*ax.Y))*p.Z
	return NewPoint3D(x, y, z)
}

func (p Point3D) Negative() Point3D {
	return p.Scale(-1)
}

func (p Point3D) Normalize() Point3D {
	len := p.Length()
	if len > 0 {
		return p.Scale(1 / len)
	}
	return p
}

func Point3DUnique(points []Point3D) []Point3D {
	pMap := make(map[Point3D]bool, len(points))
	unique := make([]Point3D, 0, len(points))
	for _, point := range points {
		if !pMap[point] {
			unique = append(unique, point)
			pMap[point] = true
		}
	}
	return unique
}

func Point3DAvg(points []Point3D) Point3D {
	avg := Point3D{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(1.0 / float32(len(points)))
}

func Point3DMin(points []Point3D) Point3D {
	pMin := NewPoint3D(Inf(1), Inf(1), Inf(1))
	for _, p := range points {
		pMin = pMin.Min(p)
	}
	return pMin
}

func Point3DMax(points []Point3D) Point3D {
	pMax := NewPoint3D(Inf(-1), Inf(-1), Inf(-1))
	for _, p := range points {
		pMax = pMax.Max(p)
	}
	return pMax
}

type ByPoint3DX []Point3D

func (p ByPoint3DX) Len() int {
	return len(p)
}

func (p ByPoint3DX) Less(i, j int) bool {
	return p[i].X < p[j].X
}

func (p ByPoint3DX) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
