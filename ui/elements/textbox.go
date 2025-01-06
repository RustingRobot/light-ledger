package elements

import (
	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBox struct {
	x, y             int32
	width, height    int32
	placeholder_Text string
	color            rl.Color
	hovered          bool
	cursor_pos       int32
}

func NewTextBox(X int32, Y int32, Width int32, Height int32, Placeholder_text string, Color rl.Color) *TextBox {
	return &TextBox{x: X, y: Y, width: Width, height: Height, placeholder_Text: Placeholder_text, color: Color, cursor_pos: 0}
}

func (r *TextBox) Draw(ctx *ui.UiBundle) {
	temp_color := r.color
	if ctx.Selected == r {
		temp_color = rl.Red
		if r.hovered {
			rl.SetMouseCursor(rl.MouseCursorIBeam)
		} else {
			rl.SetMouseCursor(rl.MouseCursorDefault)
		}
	}
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}, 2.0, temp_color)
	if r.hovered {
		rl.DrawRectangle(r.x, r.y, r.width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}

	ctx.Text_renderer.DrawText(r.placeholder_Text, r.x+10, r.y, r.color)
}

func (r *TextBox) Update(ctx *ui.UiBundle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			ctx.Selected = r
		}
	} else {
		r.hovered = false
	}

	if ctx.Selected == r {
		if rl.IsKeyPressed(rl.KeyBackspace) && r.placeholder_Text != "" {
			r.placeholder_Text = r.placeholder_Text[:len(r.placeholder_Text)-1]
		}

		key := rl.GetCharPressed()
		for key > 0 {
			if (key >= 32) && (key <= 125) {
				r.placeholder_Text += string(key)
			}
			key = rl.GetCharPressed()
		}
	}
}
