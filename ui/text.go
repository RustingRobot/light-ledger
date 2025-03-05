package ui

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed font/PixelOperator-Bold.ttf
var font_embed []byte

type TextRenderer struct {
	Font rl.Font
	Size float32
}

func (t TextRenderer) DrawText(text string, x, y int32, color rl.Color) {
	rl.DrawTextEx(t.Font, text, rl.Vector2{X: float32(x), Y: float32(y)}, t.Size, 1, color)
}

func getTextRenderer() TextRenderer {
	var font = rl.LoadFontFromMemory(".ttf", font_embed, 16, nil)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterPoint)
	return TextRenderer{Font: font, Size: float32(font.BaseSize)}
}
