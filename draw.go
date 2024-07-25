package geomutil

import (
	"image"
	"image/color"
)

func DrawLine(img *image.Image, a, b image.Point, color color.Color) {
	if a.X > b.X {
		a, b = b, a
	}

}
