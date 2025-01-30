package util

import (
	"fmt"
	"testing"
)

func TestPoint3DUnique(t *testing.T) {
	points := []Point3D{NewPoint3D(1.0, 2.0, 3.0), NewPoint3D(3.0, 4.0, 5.0), NewPoint3D(5.0, 6.0, 7.0), NewPoint3D(1.0, 2.0, 3.0)}
	result := Point3DUnique(points)
	points_str := fmt.Sprint(points[:3])
	result_str := fmt.Sprint(result)
	if points_str != result_str {
		t.Errorf("got %s; want %s", result_str, points_str)
	}
}

func TestRotation3D(t *testing.T) {
	vec1 := NewPoint3D(1.0, 0.0, 0.0)
	res1 := vec1.Rotate(NewPoint3D(0.0, 0.0, 1.0), 90)
	point_vec1 := fmt.Sprint("{0 1 0}")
	result_vec1 := fmt.Sprint(res1)
	if point_vec1 != result_vec1 {
		t.Errorf("got %s; want %s", result_vec1, point_vec1)
	}
}
