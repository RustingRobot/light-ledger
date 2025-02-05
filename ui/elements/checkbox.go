package elements

import (
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Checkbox struct {
	x, y    int32
	label   string
	color   rl.Color
	hovered bool
	Checked bool
}

func NewCheckbox(X int32, Y int32, label string, color rl.Color) *Checkbox {
	return &Checkbox{x: X, y: Y, label: label, color: color}
}

func (r *Checkbox) Draw(ctx *ui.UiBundle) {
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: 24, Height: 24}, 2.0, r.color)
	if r.hovered {
		rl.DrawRectangle(r.x, r.y, 25, 25, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}
	if r.Checked {
		ctx.Text_renderer.DrawText("x", r.x+7, r.y+1, rl.White)
	}
	ctx.Text_renderer.DrawText(r.label, r.x+35, r.y+2, r.color)
}

func (r *Checkbox) Update(ctx *ui.UiBundle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: ctx.MeasureText(r.label).X + 35, Height: 25}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			r.Checked = !r.Checked
		}
	} else {
		r.hovered = false
	}
}
