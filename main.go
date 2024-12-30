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
	var font = rl.LoadFontEx("OpenSans-Medium.ttf", 512, nil, 0)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	var cur_filter = 0
	var text = "point"

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyPressed(rl.KeySpace) {
			cur_filter++
			if cur_filter > 5 {
				cur_filter = 0
			}
			switch cur_filter {
			case 0:
				rl.SetTextureFilter(font.Texture, rl.FilterPoint)
				text = "point"
			case 1:
				rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
				text = "bilinear"
			case 2:
				rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)
				text = "trilinear"
			case 3:
				rl.SetTextureFilter(font.Texture, rl.FilterAnisotropic4x)
				text = "anisotropic 4x"
			case 4:
				rl.SetTextureFilter(font.Texture, rl.FilterAnisotropic8x)
				text = "anisotropic 8x"
			case 5:
				rl.SetTextureFilter(font.Texture, rl.FilterAnisotropic16x)
				text = "anisotropic 16x"
			}
		}
		rl.DrawText(text, 20, 100, 20, rl.LightGray)
		rl.DrawTextEx(font, "custom text. How cool!", rl.Vector2{X: 20, Y: 170}, 15, 1, rl.Gray)
		rl.DrawTextEx(font, "custom text. How cool!", rl.Vector2{X: 20, Y: 200}, 20, 1, rl.Gray)
		rl.DrawTextEx(font, "custom text. How cool!", rl.Vector2{X: 20, Y: 230}, 40, 1, rl.Gray)
		rl.DrawTextEx(font, "custom text. How cool!", rl.Vector2{X: 20, Y: 260}, 60, 1, rl.Gray)

		rl.EndDrawing()
	}
}
