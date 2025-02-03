package triangulation

import (
	"log"

	geomutil "github.com/denisstrizhkin/geomutil"
	u "github.com/denisstrizhkin/geomutil/util"
)

type AlphaShape2D struct {
	points    []u.Point2D
	triangles []u.Triangle2D
	shapes    []u.Shape2D
}

func NewAlphaShape2D(points []u.Point2D, alpha float32) (*AlphaShape2D, error) {
	triangulation, err := NewTriangulation2D(points)
	if err != nil {
		return nil, err
	}

	alpha_shape := AlphaShape2D{points: triangulation.Points(), triangles: triangulation.Triangles()}
	alpha_shape.prune(alpha)
	alpha_shape.shapes = alpha_shape.components()

	return &alpha_shape, nil
}

func (as *AlphaShape2D) Points() []u.Point2D {
	return as.points
}

func (as *AlphaShape2D) Triangles() []u.Triangle2D {
	return as.triangles
}

func (as *AlphaShape2D) Shapes() []u.Shape2D {
	return as.shapes
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

func (as *AlphaShape2D) components() []u.Shape2D {
	edges := make(map[u.Edge2D][]u.Triangle2D, len(as.triangles)*3)
	for _, tri := range as.triangles {
		for _, edge := range tri.Edges() {
			rotated := edge.Rotate()
			if len(edges[rotated]) > 0 {
				edges[rotated] = append(edges[rotated], tri)
			} else {
				edges[edge] = append(edges[edge], tri)
			}
		}
	}
	graph := make(map[u.Triangle2D][]u.Triangle2D, len(as.triangles))
	for _, triangles := range edges {
		if len(triangles) > 1 {
			graph[triangles[0]] = append(graph[triangles[0]], triangles[1])
			graph[triangles[1]] = append(graph[triangles[1]], triangles[0])
		}
	}
	return as.graphComponents(graph)
}

func (as *AlphaShape2D) graphComponents(graph map[u.Triangle2D][]u.Triangle2D) []u.Shape2D {
	visited := make(map[u.Triangle2D]bool, len(as.triangles))
	queue := geomutil.NewEventQueue[u.Triangle2D]()
	components := make([]u.Shape2D, 0)
	log.Println("before dfs")
	for _, tri := range as.triangles {
		if visited[tri] {
			continue
		}
		queue.Enqueue(tri)
		component := make([]u.Triangle2D, 0)
		for {
			tri, ok := queue.Dequeue()
			if !ok {
				break
			}
			visited[tri] = true
			log.Println("set visited to true")
			component = append(component, tri)
			for _, tri := range graph[tri] {
				if !visited[tri] {
					queue.Enqueue(tri)
				}
			}
		}
		components = append(components, u.NewShape2D(component))
	}
	return components
}
