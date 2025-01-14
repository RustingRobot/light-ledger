package ui

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed font/OpenSans-Medium.ttf
var font_embed []byte

//go:embed font/sdf.fs
var sdf_embed []byte

type TextRenderer struct {
	Font   rl.Font
	Shader rl.Shader
	Size   float32
}

func (t TextRenderer) DrawText(text string, x, y int32, color rl.Color) {
	rl.BeginShaderMode(t.Shader)
	rl.DrawTextEx(t.Font, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(t.Font.BaseSize)/8, 1, color)
	rl.EndShaderMode()
}

func getTextRenderer() TextRenderer {
	var font = rl.LoadFontFromMemory(".ttf", font_embed, 200, nil)
	rl.GenTextureMipmaps(&font.Texture)
	rl.SetTextureFilter(font.Texture, rl.FilterTrilinear)
	var shader = rl.LoadShaderFromMemory("", string(sdf_embed))
	size := float32(font.BaseSize) / 8
	return TextRenderer{Font: font, Shader: shader, Size: size}
}
