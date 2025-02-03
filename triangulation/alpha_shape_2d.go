package triangulation

import (
	u "github.com/denisstrizhkin/geomutil/util"
)

type AlphaShape2D struct {
	points    []u.Point2D
	triangles []u.Triangle2D
}

func NewAlphaShape2D(points []u.Point2D, alpha float32) (*AlphaShape2D, error) {
	triangulation, err := NewTriangulation2D(points)
	if err != nil {
		return nil, err
	}

	alpha_shape := AlphaShape2D{triangulation.Points(), triangulation.Triangles()}
	alpha_shape.prune(alpha)

	return &alpha_shape, nil
}

func (as *AlphaShape2D) Points() []u.Point2D {
	return as.points
}

func (as *AlphaShape2D) Triangles() []u.Triangle2D {
	return as.triangles
}

func (as *AlphaShape2D) prune(alpha float32) {
	for i := 0; i < len(as.triangles); {
		tri := as.triangles[i]
		if tri.CircumcircleRadiusSquared() > (1/alpha)*(1/alpha) {
			as.triangles = append(as.triangles[:i], as.triangles[i+1:]...)
		} else {
			i++
		}
	}
}
