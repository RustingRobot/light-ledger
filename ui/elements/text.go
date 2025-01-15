package elements

import (
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	x, y  int32
	text  string
	color rl.Color
}

func NewText(X int32, Y int32, text string, color rl.Color) *Text {
	return &Text{x: X, y: Y, text: text, color: color}
}

func (r *Text) Draw(ctx *ui.UiBundle) {
	ctx.Text_renderer.DrawText(r.text, r.x+10, r.y, r.color)
}

func (r *Text) Update(ctx *ui.UiBundle) {}
