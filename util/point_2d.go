package util

type Point2D struct {
	X, Y float32
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

func Point2DAvg(points []Point2D) Point2D {
	avg := Point2D{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(1.0 / float32(len(points)))
}
