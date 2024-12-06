package main

import (
	"math"

	triangulation "github.com/denisstrizhkin/geomutil/triangulation"
	u "github.com/denisstrizhkin/geomutil/util"

	// rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZOOM_SPEED    = 100
	MOUSE_SENS    = 100
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 450
)

func point2DToVector2(p u.Point2D) rl.Vector2 {
	return rl.NewVector2(p.X, -p.Y)
}

func getDefaultZoom(points []u.Point2D) (rl.Vector2, float32) {
	pMax := u.Point2DMax(points)
	pMin := u.Point2DMin(points)
	d := pMax.Subtract(pMin)
	center := pMin.Add(pMax).Scale(0.5)
	zoomX := float32(WINDOW_WIDTH) / d.X * 0.90
	zoomY := float32(WINDOW_HEIGHT) / d.Y * 0.90
	zoom := min(zoomX, zoomY)
	return point2DToVector2(center), zoom
}

func updateCamera(c *rl.Camera2D) {
	dt := rl.GetFrameTime()
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		mousePosDelta := rl.GetMouseDelta()
		mousePosDelta = rl.Vector2Scale(mousePosDelta, dt*MOUSE_SENS)
		c.Offset = rl.Vector2Add(c.Offset, mousePosDelta)
	}
	c.Zoom += dt * ZOOM_SPEED * rl.GetMouseWheelMove()
	if c.Zoom <= 0 {
		c.Zoom = 0
	}
}

func plotPoints(points []u.Point2D, radius float32, zoom float32, color rl.Color) {
	radius = radius / zoom
	for _, p := range points {
		rl.DrawCircleV(point2DToVector2(p), radius, color)
	}
}

func DrawLine(a u.Point2D, b u.Point2D, color rl.Color) {
	a_new := rl.Vector2SubtractValue(point2DToVector2(a), 0.5)
	b_new := rl.Vector2SubtractValue(point2DToVector2(b), 0.5)
	rl.DrawLineV(a_new, b_new, color)
}

func plotPolygon(points []u.Point2D, width float32, color rl.Color) {
	prevWidth := rl.GetLineWidth()
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(width)
	for i := 1; i < len(points); i++ {
		DrawLine(points[i-1], points[i], color)
	}
	DrawLine(points[len(points)-1], points[0], color)
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(prevWidth)
}

func plotTriangles(triangles []triangulation.Triangle2D, width float32, color rl.Color) {
	for _, triangle := range triangles {
		plotPolygon([]u.Point2D{triangle.A, triangle.B, triangle.C}, width, color)
	}
}

func drawQuarterEdge(e *triangulation.QuarterEdge, width float32, zoom float32, color rl.Color) {
	prevWidth := rl.GetLineWidth()
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(width)
	orig := e.Orig()
	dest := e.Dest()
	direction := dest.Negative().Add(orig).Negative()
	DrawLine(orig, dest, color)
	a := direction.Rotate(triangulation.DegreesToRadians(210)).Normalize().Scale(20 / zoom)
	a = dest.Add(a)
	b := direction.Rotate(triangulation.DegreesToRadians(150)).Normalize().Scale(20 / zoom)
	b = dest.Add(b)
	DrawLine(dest, a, color)
	DrawLine(dest, b, color)
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(prevWidth)
}

func main() {
	a1 := u.NewPoint2D(0, 0)
	b1 := u.NewPoint2D(1, 0)
	c1 := u.NewPoint2D(0.5, 0.5*float32(math.Tan(float64(triangulation.DegreesToRadians(60)))))
	a2 := u.NewPoint2D(0.4, 0.5)
	b2 := u.NewPoint2D(0.6, 0.5)
	c2 := u.NewPoint2D(0.5, 0.5-0.1*float32(math.Tan(float64(triangulation.DegreesToRadians(60)))))
	points := []u.Point2D{a1, b1, c1, a2, b2, c2}

	t1 := triangulation.MakeTriangle(a1, b1, c1)
	t2 := triangulation.MakeTriangle(a2, b2, c2)
	triangulation.Connect(t1.Sym(), t2)
	// triangulation.Connect(t1.Sym(), t2.Prev().Sym())
	// triangulation.Connect(t1.Next(), t2)
	// triangulation.Connect(t1.Next(), t2.Sym())
	e := t1

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "geomutil test")
	defer rl.CloseWindow()

	cameraTarget, cameraZoom := getDefaultZoom(points)
	cameraOffset := rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)
	camera := rl.NewCamera2D(cameraOffset, cameraTarget, 0, cameraZoom)
	// btn := rl.NewRectangle(float32(rl.GetScreenWidth())-60, float32(rl.GetScreenHeight())-30, 60, 30)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		updateCamera(&camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		plotPoints(points, 5, camera.Zoom, rl.Black)
		drawQuarterEdge(e, 3, camera.Zoom, rl.Red)

		rl.EndMode2D()

		switch rl.GetKeyPressed() {
		case rl.KeyW:
		case rl.KeyUp:
			e = e.Sym()
		case rl.KeyS:
		case rl.KeyDown:
			e = e.Next()
		case rl.KeyA:
		case rl.KeyLeft:
			e = e.Tor()
		case rl.KeyD:
		case rl.KeyRight:
			e = e.Rot()
		}

		// btn_clck := rg.Button(btn, "Next")
		// if btn_clck || rl.IsKeyPressed(rl.KeyN) {
		// 	triangulation.Step()
		// }

		rl.EndDrawing()
	}
}
