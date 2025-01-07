package elements

import (
	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBox struct {
	x, y             int32
	width, height    int32
	placeholder_text string
	text             string
	color            rl.Color
	hovered          bool
	cursor_pos       int32
}

func NewTextBox(X int32, Y int32, Width int32, Height int32, Placeholder_text string, Color rl.Color) *TextBox {
	return &TextBox{x: X, y: Y, width: Width, height: Height, placeholder_text: Placeholder_text, color: Color, cursor_pos: 0}
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
		cursor_x := r.x + int32(rl.MeasureTextEx(ctx.Text_renderer.Font, r.text[:r.cursor_pos], ctx.Text_renderer.Size, 1).X) + 12
		rl.DrawLine(cursor_x, r.y+5, cursor_x, r.y+r.height-5, r.color)
		//rl.DrawLine(r.x, r.y+15, r.x+int32(rl.MeasureTextEx(ctx.Text_renderer.Font, r.text, ctx.Text_renderer.Size, 1).X), r.y+15, r.color)
	}
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}, 2.0, temp_color)
	if r.hovered {
		rl.DrawRectangle(r.x, r.y, r.width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}

	if r.text == "" {
		ctx.Text_renderer.DrawText(r.placeholder_text, r.x+10, r.y, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	} else {
		ctx.Text_renderer.DrawText(r.text, r.x+10, r.y, r.color)
	}
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
		if rl.IsKeyPressed(rl.KeyBackspace) && r.cursor_pos > 0 {
			r.text = r.text[:r.cursor_pos-1] + r.text[r.cursor_pos:]
			r.cursor_pos--
		}

		key := rl.GetCharPressed()
		for key > 0 {
			if (key >= 32) && (key <= 125) {
				r.text = r.text[:r.cursor_pos] + string(key) + r.text[r.cursor_pos:]
				r.cursor_pos++
			}
			key = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyLeft) && r.cursor_pos > 0 {
			r.cursor_pos--
		}

		if rl.IsKeyPressed(rl.KeyRight) && int(r.cursor_pos) < len(r.text) {
			r.cursor_pos++
		}
	}
}
