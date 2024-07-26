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
	geomutil.DrawLine(
		img,
		image.Point{0, 0},
		image.Point{50, 33},
		magenta,
	)
	geomutil.DrawLine(
		img,
		image.Point{0, 33},
		image.Point{50, 66},
		magenta,
	)
	geomutil.DrawLine(
		img,
		image.Point{0, 66},
		image.Point{50, 99},
		magenta,
	)
	geomutil.DrawLine(
		img,
		image.Point{0, 33},
		image.Point{100, 33},
		mint,
	)
	geomutil.DrawLine(
		img,
		image.Point{0, 66},
		image.Point{100, 66},
		mint,
	)
	geomutil.DrawLine(
		img,
		image.Point{0, 99},
		image.Point{100, 99},
		mint,
	)

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
