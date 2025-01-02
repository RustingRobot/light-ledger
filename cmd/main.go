package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed font/OpenSans-Medium.ttf
var font_embed []byte

//go:embed font/sdf.fs
var sdf_embed []byte

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var font = rl.LoadFontFromMemory(".ttf", font_embed, 200, nil)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)
	var shader = rl.LoadShaderFromMemory("", string(sdf_embed))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		var color = rl.White
		var text = "The Quick Brown Fox Jumps Over The Lazy Dog"

		if rl.IsKeyDown(rl.KeySpace) {
			rl.DrawTextEx(font, "shader on", rl.Vector2{X: 20, Y: 40}, 20, 1, color)
			rl.BeginShaderMode(shader)
			rl.DrawTextEx(font, "shader on", rl.Vector2{X: 20, Y: 10}, 20, 1, color)
		}

		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 170}, 15, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 200}, 25, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 230}, 40, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 260}, 60, 1, color)
		rl.DrawTextEx(font, text, rl.Vector2{X: 20, Y: 300}, 120, 1, color)

		if rl.IsKeyDown(rl.KeySpace) {
			rl.EndShaderMode()
		}

		rl.EndDrawing()
	}
}
