package geomutil

type ConvexHull struct{
	points []Point
}

func NewConvexHull() *ConvexHull {
	return &ConvexHull{points: make([]Point, 0)}
}

func (ch *ConvexHull) FindConvexHull(points []Point) {
	sort.Sort(ByPointX(points))
	L_up = make([]Point, 0)
	L_up = append(L_up, points[1], points[2])
	for i := 1; i < len(points); i++ {
		
	}
}

