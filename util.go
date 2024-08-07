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
