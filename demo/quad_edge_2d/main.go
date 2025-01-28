package main

import (
	"log"
	"math"

	tri "github.com/denisstrizhkin/geomutil/triangulation"
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

func drawQuarterEdge(e *tri.QuarterEdge, width float32, zoom float32, color rl.Color) {
	prevWidth := rl.GetLineWidth()
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(width)
	orig := e.Orig()
	dest := e.Dest()
	direction := dest.Negative().Add(orig).Negative()
	DrawLine(orig, dest, color)
	a := direction.Rotate(tri.DegreesToRadians(210)).Normalize().Scale(20 / zoom)
	a = dest.Add(a)
	b := direction.Rotate(tri.DegreesToRadians(150)).Normalize().Scale(20 / zoom)
	b = dest.Add(b)
	DrawLine(dest, a, color)
	DrawLine(dest, b, color)
	rl.DrawRenderBatchActive()
	rl.SetLineWidth(prevWidth)
}

func main() {
	a1 := u.NewPoint2D(0, 0)
	b1 := u.NewPoint2D(1, 0)
	c1 := u.NewPoint2D(0.5, 0.5*float32(math.Tan(float64(tri.DegreesToRadians(60)))))
	a2 := u.NewPoint2D(0.4, 0.5)
	b2 := u.NewPoint2D(0.6, 0.5)
	c2 := u.NewPoint2D(0.5, 0.5-0.1*float32(math.Tan(float64(tri.DegreesToRadians(60)))))
	points := []u.Point2D{a1, b1, c1, a2, b2, c2}

	a1b1 := tri.MakeTriangle(a1, b1, c1)
	b1c1 := a1b1.LNext()
	c1a1 := b1c1.LNext()
	a2b2 := tri.MakeTriangle(a2, b2, c2)
	c2a2 := a2b2.RNext()
	b2c2 := c2a2.RNext()
	a1a2 := tri.Connect(c1a1, a2b2)
	a1c2 := tri.Connect(a1a2.Sym(), c2a2)
	c1b2 := tri.Connect(b1c1, b2c2)
	c1a2 := tri.Connect(c1b2.Sym(), a2b2)
	b1c2 := tri.Connect(a1b1, c2a2)
	b1b2 := tri.Connect(b1c2.Sym(), b2c2)
	c2a2.SetONext(a1c2.Sym())
	a1c2.Sym().SetONext(b1c2.Sym())
	b1c2.Sym().SetONext(b2c2.Sym())

	e := b1b2
	log.Println(a1c2, b1b2, c1a2)

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
			e = e.ONext()
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
