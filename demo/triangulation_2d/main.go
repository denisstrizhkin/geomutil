package main

import (
	"log"

	demo "github.com/denisstrizhkin/geomutil/demo"
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

func main() {
	points := []util.Point2D{
		util.NewPoint2D(0.0, 0.0),
		util.NewPoint2D(1.0, 0.0),
		util.NewPoint2D(0.0, 1.0),
		util.NewPoint2D(1.0, 1.0),
	}
	triangulation, _ := triangulation.NewTriangulation2D(points)

	d := demo.NewDemo(WINDOW_WIDTH, WINDOW_HEIGHT, "Triangulation 2D")
	defer d.Close()
	d.SetMouseSens(MOUSE_SENS)
	d.SetZoomSpeed(ZOOM_SPEED)

	cameraTarget, cameraZoom := demo.GetDefaultZoom(points)
	log.Println("target", cameraTarget)
	log.Println("zoom", cameraZoom)
	camera := d.Camera()
	camera.Target = cameraTarget
	camera.Zoom = cameraZoom

	btn := rl.NewRectangle(float32(rl.GetScreenWidth())-60, float32(rl.GetScreenHeight())-30, 60, 30)
	d.Run(func() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(*camera)

		triangles := triangulation.Triangles()
		demo.PlotTriangles(triangles, 2, rl.Green)
		demo.PlotPoints(points, 5, camera.Zoom, rl.Black)

		rl.EndMode2D()

		if rg.Button(btn, "Next") || rl.IsKeyPressed(rl.KeyN) {
			triangulation.Step()
		}

		rl.EndDrawing()
	})
}
