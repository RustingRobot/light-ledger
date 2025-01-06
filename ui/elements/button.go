package elements

import (
	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	x, y          int32
	width, height int32
	on_click      func()
	text          string
	color         rl.Color
	hovered       bool
}

func NewButton(X int32, Y int32, Width int32, Height int32, Text string, Color rl.Color, On_click func()) *Button {
	return &Button{x: X, y: Y, width: Width, height: Height, text: Text, color: Color, on_click: On_click}
}

func (r *Button) Draw(ctx *ui.UiBundle) {
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}, 2.0, r.color)
	if r.hovered {
		rl.DrawRectangle(r.x, r.y, r.width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}

	text_size := rl.MeasureTextEx(ctx.Text_renderer.Font, r.text, float32(ctx.Text_renderer.Font.BaseSize)/8, 1)
	ctx.Text_renderer.DrawText(r.text, int32(float32(r.x)+float32(r.width/2)-float32(text_size.X/2)), int32(float32(r.y)+float32(r.height/2)-float32(text_size.Y/2)), r.color)
}

func (r *Button) Update(ctx *ui.UiBundle) {

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			r.on_click()
		}
	} else {
		r.hovered = false
	}
}
