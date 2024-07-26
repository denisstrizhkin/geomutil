package geomutil

import (
	"image"
	"image/color"
	"math"
)

func DrawLine(img *image.RGBA, a, b Point, rgba color.RGBA) {
	y_transform := func(y int) int { return img.Bounds().Dy() - 1 - y }

	a_int := image.Point{int(math.Round(a.X)), int(math.Round(a.Y))}
	b_int := image.Point{int(math.Round(b.X)), int(math.Round(b.Y))}

	if a_int.X == b_int.X {
		if a_int.Y > b_int.Y {
			a_int, b_int = b_int, a_int
		}
		for y := max(0, a_int.Y); y <= min(img.Bounds().Dy()-1, b_int.Y); y++ {
			img.SetRGBA(a_int.X, y_transform(y), rgba)
		}
		return
	}

	if a.X > b.X {
		a, b = b, a
		a_int, b_int = b_int, a_int
	}
	k := (b.Y - a.Y) / (b.X - a.X)
	m := b.Y - b.X*k
	for x := max(0, a_int.X); x <= min(img.Bounds().Dx()-1, b_int.X); x++ {
		y := int(math.Round(float64(x)*k + m))
		if y < 0 || y >= img.Bounds().Dy() {
			continue
		}
		img.SetRGBA(x, y_transform(y), rgba)
	}
}

func DrawRect(img *image.RGBA, pmin, pmax Point, rgba color.RGBA) {
	a := Point{pmin.X, pmin.Y}
	b := Point{pmin.X, pmax.Y}
	c := Point{pmax.X, pmax.Y}
	d := Point{pmax.X, pmin.Y}
	DrawLine(img, a, b, rgba)
	DrawLine(img, b, c, rgba)
	DrawLine(img, c, d, rgba)
	DrawLine(img, d, a, rgba)
}

func DrawSquare(img *image.RGBA, p Point, width float64, rgba color.RGBA) {
	if width <= 0 {
		return
	}
	if width < 1 {
		width = 1
	} else {
		width = math.Round(width)
	}
	width--
	a := width / 2
	b := width - a
	min := Point{p.X - a, p.Y - a}
	max := Point{p.X + b, p.Y + b}
	DrawRect(img, min, max, rgba)
}
