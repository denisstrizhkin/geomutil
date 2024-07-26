package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/denisstrizhkin/geomutil"
)

func main() {
	img := image.NewRGBA(
		image.Rectangle{image.Point{0, 0}, image.Point{100, 100}},
	)
	magenta := color.RGBA{213, 63, 119, 255}
	//mint := color.RGBA{162, 228, 184, 255}

	draw_at := func(p geomutil.Point) { geomutil.DrawSquare(img, p, 3, magenta) }
	points := geomutil.ReadPoints("./a.txt")
	for _, p := range points {
		draw_at(p)
	}

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
