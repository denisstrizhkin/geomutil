package main

import (
	// rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
	geomutil "github.com/denisstrizhkin/geomutil"
)

func main() {
	PATH := "./a.txt"
	// fmt.Println(geomutil.ReadPoints(PATH))
	gh := geomutil.NewConvexHull(geomutil.ReadPoints(PATH))
	fmt.Println(gh.Points)
	geomutil.SavePoints(gh.Points, "conv_out.txt")
	// geomutil.SavePoints(hull, "points.txt")
	// rl.InitWindow(800, 450, "My convex hull")
	// defer rl.CloseWindow()

	// rl.SetTargetFPS(60)

	// for !rl.WindowShouldClose() {
	// 	rl.BeginDrawing()

	// 	rl.ClearBackground(rl.RayWhite)
	// 	// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	// 	for _, i := range points {
	// 		rl.DrawCircleV(i, 10, MAROON)
	// 	}

	// 	rl.EndDrawing()
	// }
}
