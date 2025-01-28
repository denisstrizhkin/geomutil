package demo

import (
	"log"

	u "github.com/denisstrizhkin/geomutil/util"

	tri "github.com/denisstrizhkin/geomutil/triangulation"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Demo struct {
	camera     rl.Camera2D
	zoom_speed float32
	mouse_sens float32
}

func (d *Demo) updateCamera() {
	dt := rl.GetFrameTime()
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		mousePosDelta := rl.GetMouseDelta()
		mousePosDelta = rl.Vector2Scale(mousePosDelta, dt*d.mouse_sens)
		d.camera.Offset = rl.Vector2Add(d.camera.Offset, mousePosDelta)
	}
	d.camera.Zoom += dt * d.zoom_speed * rl.GetMouseWheelMove()
	if d.camera.Zoom <= 0 {
		d.camera.Zoom = 0
	}
}

func NewDemo(width, height int32, name string) *Demo {
	rl.InitWindow(width, height, name)
	cameraOffset := rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)
	camera := rl.NewCamera2D(cameraOffset, rl.NewVector2(1.0, 1.0), 0, 1.0)
	return &Demo{
		camera: camera,
	}
}

func (d *Demo) Run(callback func()) {
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		d.updateCamera()
		callback()
	}
}

func (d *Demo) Close() {
	rl.CloseWindow()
}

func (d *Demo) Camera() *rl.Camera2D {
	return &d.camera
}

func (d *Demo) SetMouseSens(sens float32) {
	d.mouse_sens = sens
}

func (d *Demo) SetZoomSpeed(speed float32) {
	d.zoom_speed = speed
}

func Point2DToVector2(p u.Point2D) rl.Vector2 {
	return rl.NewVector2(p.X, -p.Y)
}

func GetDefaultZoom(points []u.Point2D) (rl.Vector2, float32) {
	pMax := u.Point2DMax(points)
	pMin := u.Point2DMin(points)
	d := pMax.Subtract(pMin)
	center := pMin.Add(pMax).Scale(0.5)
	zoomX := float32(rl.GetScreenWidth()) / d.X * 0.90
	zoomY := float32(rl.GetScreenHeight()) / d.Y * 0.90
	log.Println("zoomXY", zoomX, zoomY)
	zoom := min(zoomX, zoomY)
	return Point2DToVector2(center), zoom
}

func PlotPoints(points []u.Point2D, radius float32, zoom float32, color rl.Color) {
	radius = radius / zoom
	for _, p := range points {
		rl.DrawCircleV(Point2DToVector2(p), radius, color)
	}
}

func DrawLine(a u.Point2D, b u.Point2D, color rl.Color) {
	a_new := rl.Vector2SubtractValue(Point2DToVector2(a), 0.5)
	b_new := rl.Vector2SubtractValue(Point2DToVector2(b), 0.5)
	rl.DrawLineV(a_new, b_new, color)
}

func PlotPolygon(points []u.Point2D, width float32, color rl.Color) {
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

func PlotTriangles(triangles []tri.Triangle2D, width float32, color rl.Color) {
	for _, triangle := range triangles {
		PlotPolygon([]u.Point2D{triangle.A, triangle.B, triangle.C}, width, color)
	}
}
