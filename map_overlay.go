package geomutil

import (
// "sort"
// "fmt"
)

// func FindIntersections(edges []Edge) []Point {
// 	sort.Sort(ByPointA(edges))
// 	Q := edges
// 	for len(Q) != 0 {

// 	}
// 	return
// }

type ByPointA []Edge

func (p ByPointA) Len() int {
	return len(p)
}

func (p ByPointA) Less(i, j int) bool {
	return p[i].A.Y < p[j].A.Y
}

func (p ByPointA) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
