package util

type Shape2D struct {
	triangles []Triangle2D
}

func NewShape2D(triangles []Triangle2D) Shape2D {
	return Shape2D{triangles}
}

func (s *Shape2D) Triangles() []Triangle2D {
	return s.triangles
}

func (s *Shape2D) Volume() float32 {
	vol := float32(0)
	for _, tri := range s.triangles {
		vol += tri.Volume()
	}
	return vol
}

func (s *Shape2D) Perimeter() []Point2D {
	return nil
}
