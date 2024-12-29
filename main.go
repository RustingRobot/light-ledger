package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(800, 450, "Light Ledger")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// OpenSans licensed under the SIL Open Font License Version 1.1
	// https://openfontlicense.org/open-font-license-official-text/
	var better_font = rl.LoadFontEx("OpenSans-Regular.ttf", 300, nil, 256)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("welcome world!", 190, 200, 20, rl.LightGray)
		rl.DrawTextEx(better_font, "custom text. How cool!", rl.Vector2{X: 190, Y: 300}, 50, 1, rl.Gray)

		rl.DrawRectangle(20, 20, 300, 50, rl.Red)

		rl.EndDrawing()
	}
}
