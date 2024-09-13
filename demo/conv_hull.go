package main

import (
	"log"
	"math"

	geomutil "github.com/denisstrizhkin/geomutil"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZOOM_SPEED    = 20
	MOUSE_SENS    = 100
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450

	MODE_MOVE = 0
	MODE_ADD  = 1
)

func pointToVector2(p geomutil.Point) rl.Vector2 {
	return rl.NewVector2(float32(p.X), -float32(p.Y))
}

func pointAvg(points []geomutil.Point) geomutil.Point {
	avg := geomutil.Point{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(float64(1) / float64(len(points)))
}

func getDefaultZoom(points []geomutil.Point) float32 {
	xMax := -math.MaxFloat64
	yMax := -math.MaxFloat64
	center := pointAvg(points)
	for _, p := range points {
		x := math.Abs(p.X - center.X)
		y := math.Abs(p.Y - center.Y)
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	zoomX := float32(float64(WINDOW_WIDTH)/2/xMax) * 0.95
	zoomY := float32(float64(WINDOW_HEIGHT)/2/yMax) * 0.95
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

func plotPoints(points []geomutil.Point, radius float32, color rl.Color) {
	for _, p := range points {
		rl.DrawCircleV(pointToVector2(p), radius, color)
	}
}

func plotPolygon(points []geomutil.Point, width float32, color rl.Color) {
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

func main() {
	PATH := "./input.txt"
	points := geomutil.ReadPoints(PATH)
	gh := geomutil.NewConvexHull(points)

	pointsCenter := pointAvg(gh.Points)
	cameraTarget := pointToVector2(pointsCenter)
	cameraOffset := rl.NewVector2(WINDOW_WIDTH/2, WINDOW_HEIGHT/2)
	cameraZoom := getDefaultZoom(points)
	camera := rl.NewCamera2D(cameraOffset, cameraTarget, 0, cameraZoom)

	mode := uint8(MODE_MOVE)
	mode_text := ""

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "geomutil test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		switch mode {
		case MODE_MOVE:
			mode_text = "MOVE"
			updateCamera(&camera)
		case MODE_ADD:
			mode_text = "ADD"
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				pos := rl.GetMousePosition()
				log.Print(pos)
				pos = rl.Vector2Scale(pos, 1/camera.Zoom)
				log.Print(pos)
				points = append(
					points, geomutil.Point{X: float64(pos.X), Y: float64(pos.Y)},
				)
				gh = geomutil.NewConvexHull(points)
			}
		default:
			log.Fatal("Unknown mode:", mode)
		}

		rl.BeginDrawing()
		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(mode_text, 5, 5, 20, rl.Black)
		plotPoints(points, 2, rl.Yellow)
		plotPoints(gh.Points, 2, rl.Red)
		plotPolygon(gh.Points, 4, rl.Green)

		rl.EndDrawing()
	}
}
