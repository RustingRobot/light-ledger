package elements

import (
	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	X, Y          int32
	Width, Height int32
	On_click      func()
	Text          string
	Color         rl.Color
	hovered       bool
}

func (r *Button) Draw(ctx *ui.UiBundle) {
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}, 2.0, r.Color)
	if r.hovered {
		rl.DrawRectangle(r.X, r.Y, r.Width, r.Height, rl.Color{R: r.Color.R, G: r.Color.G, B: r.Color.B, A: r.Color.A / 5})
	}

	text_size := rl.MeasureTextEx(ctx.Text_renderer.Font, r.Text, float32(ctx.Text_renderer.Font.BaseSize)/8, 1)
	ctx.Text_renderer.DrawText(r.Text, int32(float32(r.X)+float32(r.Width/2)-float32(text_size.X/2)), int32(float32(r.Y)+float32(r.Height/2)-float32(text_size.Y/2)), r.Color)
}

func (r *Button) Update(ctx *ui.UiBundle) {

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			r.On_click()
		}
	} else {
		r.hovered = false
	}
}
