package elements

import (
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tabs struct {
	x, y              int32
	height, tab_width int32
	titles            []string
	tab_elements      []*Container
	color             rl.Color
	hovered           int32
	selected          int32
}

func NewTabs(X int32, Y int32, height int32, tab_width int32, titles []string, tab_elements []*Container, color rl.Color) *Tabs {
	for i, tab := range tab_elements {
		if i != 0 {
			tab.Active = false
		}
	}
	return &Tabs{x: X, y: Y, height: height, tab_width: tab_width, titles: titles, tab_elements: tab_elements, color: color}
}

func (r *Tabs) Draw(ctx *ui.UiBundle) {
	rl.DrawRectangle(0, r.y+r.height, int32(rl.GetScreenWidth()), 1, rl.White)
	if r.hovered >= 0 {
		rl.DrawRectangle(r.x+r.tab_width*r.hovered, r.y, r.tab_width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}
	rl.DrawRectangle(r.x, r.y+r.height, r.tab_width*int32(len(r.tab_elements)), 1, rl.White)

	for i, title := range r.titles {
		text_color := r.color
		if i == int(r.selected) {
			text_color = rl.DarkGray
			rl.DrawRectangle(r.x+r.tab_width*r.selected, r.y, r.tab_width, r.height, rl.White)
		}
		ctx.Text_renderer.DrawText(title, r.x+r.tab_width*int32(i)+10, r.y+4, text_color)
	}
}

func (r *Tabs) Update(ctx *ui.UiBundle) {
	r.hovered = -1
	for i, _ := range r.titles {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x + r.tab_width*int32(i)), Y: float32(r.y), Width: float32(r.tab_width), Height: float32(r.height)}) {
			r.hovered = int32(i)
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				r.tab_elements[r.selected].Active = false
				r.selected = int32(i)
				r.tab_elements[r.selected].Active = true
			}
		}
	}
}
