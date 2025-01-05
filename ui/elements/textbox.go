package elements

import (
	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBox struct {
	X, Y             int32
	Width, Height    int32
	Placeholder_Text string
	Color            rl.Color
	hovered          bool
}

func (r *TextBox) Draw(ctx *ui.UiBundle) {
	temp_color := r.Color
	if ctx.Selected == r {
		temp_color = rl.Red
		if r.hovered {
			rl.SetMouseCursor(rl.MouseCursorIBeam)
		} else {
			rl.SetMouseCursor(rl.MouseCursorDefault)
		}
	}
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}, 2.0, temp_color)
	if r.hovered {
		rl.DrawRectangle(r.X, r.Y, r.Width, r.Height, rl.Color{R: r.Color.R, G: r.Color.G, B: r.Color.B, A: r.Color.A / 5})
	}

	ctx.Text_renderer.DrawText(r.Placeholder_Text, r.X+10, r.Y, r.Color)
}

func (r *TextBox) Update(ctx *ui.UiBundle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			ctx.Selected = r
		}
	} else {
		r.hovered = false
	}

	if ctx.Selected == r {
		if rl.IsKeyPressed(rl.KeyBackspace) && r.Placeholder_Text != "" {
			r.Placeholder_Text = r.Placeholder_Text[:len(r.Placeholder_Text)-1]
		}

		key := rl.GetCharPressed()
		for key > 0 {
			if (key >= 32) && (key <= 125) {
				r.Placeholder_Text += string(key)
			}
			key = rl.GetCharPressed()
		}
	}
}
