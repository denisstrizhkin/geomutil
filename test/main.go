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
	mint := color.RGBA{162, 228, 184, 255}

	for y := 0; y < 100; y += 25 {
		for x := 0; x < 100; x += 25 {
			geomutil.DrawLine(img, image.Point{0, y}, image.Point{99, y}, magenta)
			geomutil.DrawLine(img, image.Point{x, 0}, image.Point{x, 99}, magenta)
		}
	}

	geomutil.DrawRect(img, image.Rectangle{image.Point{25, 25}, image.Point{75, 50}}, mint)
	geomutil.DrawSquare(img, image.Point{37, 62}, 26, mint)

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
