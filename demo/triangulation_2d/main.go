package main

import (
	"log"

	triangulation "github.com/denisstrizhkin/geomutil/triangulation"
	util "github.com/denisstrizhkin/geomutil/util"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZOOM_SPEED    = 100
	MOUSE_SENS    = 100
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
)

func point2DToVector2(p util.Point2D) rl.Vector2 {
	return rl.NewVector2(p.X, -p.Y)
}

func getDefaultZoom(points []util.Point2D) (rl.Vector2, float32) {
	pMax := util.Point2DMax(points)
	pMin := util.Point2DMin(points)
	d := pMax.Subtract(pMin)
	center := pMin.Add(pMax).Scale(0.5)
	zoomX := float32(WINDOW_WIDTH) / d.X * 0.95
	zoomY := float32(WINDOW_HEIGHT) / d.Y * 0.95
	return point2DToVector2(center), min(zoomX, zoomY)
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

func plotPoints(points []util.Point2D, radius float32, zoom float32, color rl.Color) {
	radius = radius / zoom
	for _, p := range points {
		rl.DrawCircleV(rl.Vector2AddValue(point2DToVector2(p), 0.5), radius, color)
	}
}

func plotPolygon(points []util.Point2D, width float32, color rl.Color) {
	prevWidth := rl.GetLineWidth()
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(width)
	for i := 1; i < len(points); i++ {
		rl.DrawLineV(point2DToVector2(points[i-1]), point2DToVector2(points[i]), color)
	}
	rl.DrawLineV(point2DToVector2(points[0]), point2DToVector2(points[len(points)-1]), color)
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

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "geomutil test")
	defer rl.CloseWindow()

	cameraTarget, cameraZoom := getDefaultZoom(points)
	log.Println(cameraTarget, cameraZoom)
	cameraOffset := rl.NewVector2(300, 125)
	camera := rl.NewCamera2D(cameraOffset, cameraTarget, 0, cameraZoom)
	btn := rl.NewRectangle(float32(rl.GetScreenWidth())-60, float32(rl.GetScreenHeight())-30, 60, 30)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		updateCamera(&camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		triangles := triangulation.Triangles()
		plotTriangles(triangles, 2, rl.Green)
		plotPoints(points, 5, camera.Zoom, rl.Black)

		rl.DrawCircleV(rl.Vector2AddValue(cameraTarget, 0.5), 0.1, rl.Gray)
		rl.EndMode2D()

		btn_clck := rg.Button(btn, "Next")
		if btn_clck || rl.IsKeyPressed(rl.KeyN) {
			triangulation.Step()
		}

		rl.EndDrawing()
	}
}
