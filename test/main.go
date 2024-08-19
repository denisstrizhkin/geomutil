package main

import (
	"fmt"
	geomutil "github.com/denisstrizhkin/geomutil"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZOOM_SPEED    = 10
	MOUSE_SENS    = 100
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
)

func pointToVector2(p geomutil.Point) rl.Vector2 {
	return rl.NewVector2(float32(p.X), float32(p.Y))
}

func pointAvg(points []geomutil.Point) geomutil.Point {
	avg := geomutil.Point{}
	for _, p := range points {
		avg = avg.Add(p)
	}
	return avg.Scale(float64(1) / float64(len(points)))
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
	for i := 1; i < len(points); i++ {
		rl.DrawLineV(pointToVector2(points[i-1]), pointToVector2(points[i]), color)
	}
	rl.DrawLineV(pointToVector2(points[0]), pointToVector2(points[len(points)-1]), color)
}

func main() {
	PATH := "./input.txt"
	gh := geomutil.NewConvexHull(geomutil.ReadPoints(PATH))
	fmt.Println(gh.Points)

	pointsCenter := pointToVector2(pointAvg(gh.Points))
	camera := rl.NewCamera2D(rl.NewVector2(WINDOW_WIDTH/2, WINDOW_HEIGHT/2), pointsCenter, 0, 1)
	camera.Target = pointsCenter

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "geomutil test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		updateCamera(&camera)

		rl.ClearBackground(rl.RayWhite)
		plotPoints(gh.Points, 2, rl.Red)
		plotPolygon(gh.Points, 2, rl.Green)

		rl.EndDrawing()
	}
}
