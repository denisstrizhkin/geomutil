package util

import (
	"fmt"
	"testing"
)

func TestPoint2DUnique(t *testing.T) {
	points := []Point2D{NewPoint2D(1.0, 2.0), NewPoint2D(3.0, 4.0), NewPoint2D(5.0, 6.0), NewPoint2D(1.0, 2.0)}
	result := Point2DUnique(points)
	points_str := fmt.Sprint(points[:3])
	result_str := fmt.Sprint(result)
	if points_str != result_str {
		t.Errorf("got %s; want %s", result_str, points_str)
	}
}
