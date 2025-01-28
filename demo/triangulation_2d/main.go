package main

import (
	"log"
	"math/rand"

	demo "github.com/denisstrizhkin/geomutil/demo"
	tri "github.com/denisstrizhkin/geomutil/triangulation"
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

func GetRandomColor() rl.Color {
	var color rl.Color
	for {
		// Generate random RGB values
		color.R = uint8(rand.Intn(256))
		color.G = uint8(rand.Intn(256))
		color.B = uint8(rand.Intn(256))
		color.A = 255 // Fully opaque
		if color.R < 240 && color.G < 240 && color.B < 240 {
			break
		}
	}

	return color
}

func main() {
	// points := []util.Point2D{
	// 	util.NewPoint2D(0.0, 0.0),
	// 	util.NewPoint2D(1.0, 0.0),
	// 	util.NewPoint2D(0.0, 1.0),
	// 	util.NewPoint2D(1.0, 1.0),
	// }
	points, err := util.Point2DFromFile("../points_A.json")
	if err != nil {
		log.Fatalln(err)
	}

	d := demo.NewDemo(WINDOW_WIDTH, WINDOW_HEIGHT, "Triangulation 2D")
	defer d.Close()
	d.SetMouseSens(MOUSE_SENS)
	d.SetZoomSpeed(ZOOM_SPEED)

	triangulation, _ := tri.NewTriangulation2D(points)

	cameraTarget, cameraZoom := demo.GetDefaultZoom(points)
	camera := d.Camera()
	camera.Target = cameraTarget
	camera.Zoom = cameraZoom * 0.5

	triangles := triangulation.Triangles()
	colors := make([]rl.Color, 0)
	btn := rl.NewRectangle(float32(rl.GetScreenWidth())-60, float32(rl.GetScreenHeight())-30, 60, 30)
	d.Run(func() {
		d.UpdateCamera()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(*camera)

		for i, triangle := range triangles {
			for {
				if i < len(colors) {
					break
				}
				colors = append(colors, GetRandomColor())
			}
			demo.DrawTriangle(triangle, colors[i])
		}
		demo.PlotPoints(points, 5, camera.Zoom, rl.Black)

		rl.EndMode2D()

		if rg.Button(btn, "Next") || rl.IsKeyPressed(rl.KeyN) {
			log.Println("triangles count:", len(triangles))
			log.Println("Next step")
			triangulation.Step()
			triangles = triangulation.Triangles()
			log.Println("triangles count:", len(triangles))
		}

		rl.EndDrawing()
	})
}
