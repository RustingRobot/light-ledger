package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// OpenSans licensed under the SIL Open Font License Version 1.1
	// https://openfontlicense.org/open-font-license-official-text/
	var font = rl.LoadFontEx("OpenSans-Regular.ttf", 212, nil, 0)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)

	var shader = rl.LoadShader("", "sdf.fs")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		var color = rl.White
		var text = "By using this repository, you can significantly improve the efficiency of text rendering in your OpenGL projects."

		if rl.IsKeyDown(rl.KeySpace) {
			rl.DrawTextEx(font, "shader on", rl.Vector2{X: 20, Y: 40}, 20, 1, color)
			rl.BeginShaderMode(shader)
			rl.DrawTextEx(font, "shader on", rl.Vector2{X: 20, Y: 10}, 20, 1, color)
		}

		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 170}, 15, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 200}, 20, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 230}, 40, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 260}, 60, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 300}, 120, 1, color)

		if rl.IsKeyDown(rl.KeySpace) {
			rl.EndShaderMode()
		}

		rl.EndDrawing()
	}
}
