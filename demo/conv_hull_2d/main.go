package main

import (
	"log"

	"github.com/denisstrizhkin/geomutil"
	"github.com/denisstrizhkin/geomutil/demo"
	"github.com/denisstrizhkin/geomutil/util"
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

func main() {
	points, err := util.Point2DFromFile("../points.json")
	if err != nil {
		log.Fatalln(err)
	}
	gh := geomutil.NewConvexHull(points)

	d := demo.NewDemo(WINDOW_WIDTH, WINDOW_HEIGHT, "Convex hull 2D")
	d.SetMouseSens(MOUSE_SENS)
	d.SetZoomSpeed(ZOOM_SPEED)
	defer d.Close()

	cameraTarget, cameraZoom := demo.GetDefaultZoom(points)
	camera := d.Camera()
	camera.Target = cameraTarget
	camera.Zoom = cameraZoom * 0.5

	mode := uint8(MODE_MOVE)
	mode_text := ""

	d.Run(func() {
		if rl.IsKeyPressed(rl.KeySpace) {
			log.Print("switching modes")
			switch mode {
			case MODE_MOVE:
				mode = MODE_ADD
			case MODE_ADD:
				mode = MODE_MOVE
			default:
				log.Fatal("Unknown mode:", mode)
			}
		}

		switch mode {
		case MODE_MOVE:
			mode_text = "MOVE"
			d.UpdateCamera()
		case MODE_ADD:
			mode_text = "ADD"
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				pos := d.MousePositionPoint2D()
				points = append(
					points, util.NewPoint2D(pos.X, pos.Y),
				)
				gh = geomutil.NewConvexHull(points)
			}
		default:
			log.Fatal("Unknown mode:", mode)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(*camera)

		demo.PlotPolygon(gh.Points(), 2, rl.Green)
		demo.PlotPoints(points, 5, camera.Zoom, rl.Yellow)
		demo.PlotPoints(gh.Points(), 5, camera.Zoom, rl.Red)

		rl.EndMode2D()

		rl.DrawText(mode_text, 5, 5, 20, rl.Black)

		rl.EndDrawing()
	})
}
