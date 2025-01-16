package elements

import (
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tabs struct {
	x, y         int32
	height       int32
	titles       []string
	tab_elements []*Container
	color        rl.Color
	hovered      int32
	selected     int32
}

func NewTabs(X int32, Y int32, height int32, titles []string, tab_elements []*Container, color rl.Color) *Tabs {
	for i, tab := range tab_elements {
		if i != 0 {
			tab.Active = false
		}
	}
	return &Tabs{x: X, y: Y, height: height, titles: titles, tab_elements: tab_elements, color: color}
}

func (r *Tabs) Draw(ctx *ui.UiBundle) {
	if r.hovered >= 0 {
		rl.DrawRectangle(r.x+100*r.hovered, r.y, int32(ctx.MeasureText(r.titles[r.hovered]).X), r.height, r.color)
	}
	for i, title := range r.titles {
		ctx.Text_renderer.DrawText(title, r.x+int32(100*i), r.y, r.color)
	}
}

func (r *Tabs) Update(ctx *ui.UiBundle) {
	r.hovered = -1
	for i, title := range r.titles {
		text_width := ctx.MeasureText(title).X
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x + int32(100*i)), Y: float32(r.y), Width: text_width, Height: float32(r.height)}) {
			r.hovered = int32(i)
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				r.tab_elements[r.selected].Active = false
				r.selected = int32(i)
				r.tab_elements[r.selected].Active = true
			}
		}
	}
}
