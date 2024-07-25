package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/denisstrizhkin/geomutil"
)

func main() {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{100, 100}})
	geomutil.DrawLine(img, image.Point{0, 0}, image.Point{50, 50}, color.RGBA{213, 63, 119, 255})

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
