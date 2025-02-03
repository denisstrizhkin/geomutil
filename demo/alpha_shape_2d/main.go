package main

import (
	"fmt"
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
	WINDOW_WIDTH  = 1600
	WINDOW_HEIGHT = 900

	MODE_MOVE = 0
	MODE_ADD  = 1
)

func GetRandomColor() rl.Color {
	const threshold = 250
	var color rl.Color
	for {
		// Generate random RGB values
		color.R = uint8(rand.Intn(256))
		color.G = uint8(rand.Intn(256))
		color.B = uint8(rand.Intn(256))
		color.A = 255 // Fully opaque
		if color.R < threshold && color.G < threshold && color.B < threshold {
			break
		}
	}

	return color
}

func GetColors(n int) []rl.Color {
	colors := make([]rl.Color, n)
	for i := range colors {
		colors[i] = GetRandomColor()
	}
	return colors
}

func main() {
	// Slider control extended, returns selected value and has text
	// int GuiSlider(Rectangle bounds, const char *textLeft, const char *textRight, float *value, float minValue, float maxValue)
	// {
	//     return GuiSliderPro(bounds, textLeft, textRight, value, minValue, maxValue, GuiGetStyle(SLIDER, SLIDER_WIDTH));
	// }

	// points := []util.Point2D{
	// 	util.NewPoint2D(0.0, 0.0),
	// 	util.NewPoint2D(0.0, 1.0),
	// 	util.NewPoint2D(1.0, 1.0),
	// 	util.NewPoint2D(1.0, 0.0),
	// 	util.NewPoint2D(0.5, 0.25),
	// 	util.NewPoint2D(0.5, 0.75),
	// 	util.NewPoint2D(0.25, 0.5),
	// 	util.NewPoint2D(0.75, 0.5),
	// }
	points, err := util.Point2DFromFile("../points_alpha.json")
	if err != nil {
		log.Fatalln(err)
	}

	d := demo.NewDemo(WINDOW_WIDTH, WINDOW_HEIGHT, "Triangulation 2D")
	defer d.Close()
	d.SetMouseSens(MOUSE_SENS)
	d.SetZoomSpeed(ZOOM_SPEED)

	// shape, _ := tri.NewTriangulation2D(points)

	cameraTarget, cameraZoom := demo.GetDefaultZoom(points)
	camera := d.Camera()
	camera.Target = cameraTarget
	camera.Zoom = cameraZoom * 0.5

	var value float32 = 0.0
	value_old := value
	shape, _ := tri.NewAlphaShape2D(points, value)
	triangles := shape.Triangles()
	colors := GetColors(len(triangles))

	move_camera := true
	slider_boundaries := rl.NewRectangle(800, 50, 200, 50)

	d.Run(func() {
		mouse_pos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mouse_pos, slider_boundaries) && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			move_camera = false
		} else if !move_camera && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			move_camera = true
		}
		if move_camera {
			d.UpdateCamera()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(*camera)

		for i, triangle := range triangles {
			demo.DrawTriangle(triangle, colors[i])
		}
		demo.PlotPoints(points, 5, camera.Zoom, rl.Black)

		rl.EndMode2D()

		value = rg.Slider(slider_boundaries, "alpha", fmt.Sprintf("%.3f", value), value, 0, 0.2)
		if value != value_old {
			value_old = value
			shape, _ = tri.NewAlphaShape2D(points, value)
			triangles = shape.Triangles()
			colors = GetColors(len(triangles))
		}

		rl.EndDrawing()
	})
}
