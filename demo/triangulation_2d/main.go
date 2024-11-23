package main

import (
	"math"

	triangulation "github.com/denisstrizhkin/geomutil/triangulation"
	util "github.com/denisstrizhkin/geomutil/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZOOM_SPEED    = 20
	MOUSE_SENS    = 100
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
)

func pointToVector2(p util.Point2D) rl.Vector2 {
	return rl.NewVector2(p.X, -p.Y)
}

func getDefaultZoom(points []util.Point2D) float32 {
	xMax := math.Inf(-1)
	yMax := math.Inf(-1)
	center := util.Point2DAvg(points)
	for _, p := range points {
		d := p.Subtract(center)
		x := math.Abs(float64(d.X))
		y := math.Abs(float64(d.Y))
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	zoomX := float32(float64(WINDOW_WIDTH) / 2.0 / xMax * 0.95)
	zoomY := float32(float64(WINDOW_HEIGHT) / 2.0 / yMax * 0.95)
	if zoomX < zoomY {
		return zoomX
	}
	return zoomY
}

func updateCamera(c *rl.Camera2D) {
	dt := rl.GetFrameTime()
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		mousePosDelta := rl.GetMouseDelta()
		mousePosDelta = rl.Vector2Scale(mousePosDelta, dt*MOUSE_SENS)
		c.Offset = rl.Vector2Add(c.Offset, mousePosDelta)
	}
	c.Zoom += dt * ZOOM_SPEED * rl.GetMouseWheelMove()
}

func plotPoints(points []util.Point2D, radius float32, color rl.Color) {
	for _, p := range points {
		rl.DrawCircleV(pointToVector2(p), radius, color)
	}
}

func plotPolygon(points []util.Point2D, width float32, color rl.Color) {
	prevWidth := rl.GetLineWidth()
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(width)
	for i := 1; i < len(points); i++ {
		rl.DrawLineV(pointToVector2(points[i-1]), pointToVector2(points[i]), color)
	}
	rl.DrawLineV(pointToVector2(points[0]), pointToVector2(points[len(points)-1]), color)
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(prevWidth)
}

func plotTriangles(triangles []triangulation.Triangle2D, width float32, color rl.Color) {
	for _, triangle := range triangles {
		plotPolygon([]util.Point2D{triangle.A, triangle.B, triangle.C}, width, color)
	}
}

func main() {
	points := []util.Point2D{
		util.NewPoint2D(0.0, 0.0),
		util.NewPoint2D(0.5, 1.0),
		util.NewPoint2D(1.0, 2.0),
		util.NewPoint2D(1.5, 0.0),
	}
	triangulation := triangulation.NewTriangulation2D(points)

	pointsCenter := util.Point2DAvg(points)
	cameraTarget := pointToVector2(pointsCenter)
	cameraOffset := rl.NewVector2(WINDOW_WIDTH/2, WINDOW_HEIGHT/2)
	cameraZoom := getDefaultZoom(points)
	camera := rl.NewCamera2D(cameraOffset, cameraTarget, 0, cameraZoom)

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "geomutil test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		updateCamera(&camera)

		rl.BeginDrawing()
		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.RayWhite)
		plotPoints(points, 0.1, rl.Yellow)
		triangles := triangulation.Triangles()
		plotTriangles(triangles, 0.05, rl.Green)

		rl.EndDrawing()
	}
}
