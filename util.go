package geomutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

type Edge struct {
	A, B Point
}

func NewEdge(a, b Point) *Edge {
	if a.Y > b.Y {
		return &Edge{A: a, B: b}
	} else {
		return &Edge{A: b, B: a}
	}
}

func (p Point) Scale(a float64) Point {
	return Point{p.X * a, p.Y * a}
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Subtract(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type ByPointX []Point

func (p ByPointX) Len() int {
	return len(p)
}

func (p ByPointX) Less(i, j int) bool {
	return p[i].X < p[j].X
}

func (p ByPointX) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByPointY []Point

func (p ByPointY) Len() int {
	return len(p)
}

func (p ByPointY) Less(i, j int) bool {
	return p[i].Y < p[j].Y
}

func (p ByPointY) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ParseFloat(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return f64
}

func ReadPoints(path string) []Point {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	points := make([]Point, 0)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		tokens := strings.Split(line, " ")
		points = append(
			points, Point{X: ParseFloat(tokens[0]), Y: ParseFloat(tokens[1])},
		)
	}
	return points
}

func SavePoints(points []Point, path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("can't open: %s - %v", path, err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, p := range points {
		_, err := w.WriteString(fmt.Sprintf("%.4f %.4f\n", p.X, p.Y))
		if err != nil {
			log.Fatalf("can't write: %s - %v", path, err)
		}
	}
}
