package geomutil

import (
	"image"
	"image/color"
	"math"
)

func DrawLine(img *image.RGBA, a, b image.Point, rgba color.RGBA) {
	y_transform := func(y int) int { return img.Bounds().Dy() - 1 - y }

	if a.X == b.X {
		if a.Y > b.Y {
			a, b = b, a
		}
		for y := max(0, a.Y); y <= min(img.Bounds().Dy()-1, b.Y); y++ {
			img.SetRGBA(a.X, y_transform(y), rgba)
		}
		return
	}

	if a.X > b.X {
		a, b = b, a
	}
	k := float64(b.Y-a.Y) / float64(b.X-a.X)
	m := float64(b.Y) - float64(b.X)*k
	for x := max(0, a.X); x <= min(img.Bounds().Dx()-1, b.X); x++ {
		y := int(math.Round(float64(x)*k + m))
		if y < 0 || y >= img.Bounds().Dy() {
			continue
		}
		img.SetRGBA(x, y_transform(y), rgba)
	}
}

func DrawRect(img *image.RGBA, rect image.Rectangle, rgba color.RGBA) {
	a := image.Point{rect.Min.X, rect.Min.Y}
	b := image.Point{rect.Min.X, rect.Max.Y}
	c := image.Point{rect.Max.X, rect.Max.Y}
	d := image.Point{rect.Max.X, rect.Min.Y}
	DrawLine(img, a, b, rgba)
	DrawLine(img, b, c, rgba)
	DrawLine(img, c, d, rgba)
	DrawLine(img, d, a, rgba)
}

func DrawSquare(img *image.RGBA, p image.Point, width int, rgba color.RGBA) {
	if width <= 0 {
		return
	}
	width--
	a := width / 2
	b := width - a
	min := image.Point{p.X - a, p.Y - a}
	max := image.Point{p.X + b, p.Y + b}
	DrawRect(img, image.Rectangle{min, max}, rgba)
}
