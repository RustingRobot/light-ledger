package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type uiElement interface {
	draw(ctx uiBundle)
	update()
}

type Button struct {
	X, Y          int32
	Width, Height int32
	On_click      func()
	Text          string
	Color         rl.Color
	hovered       bool
}

type uiBundle struct {
	text_renderer textRenderer
	ui_elements   []uiElement
}

func (r *uiBundle) Add(u uiElement) {
	r.ui_elements = append(r.ui_elements, u)
}

func (r uiBundle) Draw() {
	for _, ui_element := range r.ui_elements {
		ui_element.draw(r)
	}
}

func (r uiBundle) Update() {
	for _, ui_element := range r.ui_elements {
		ui_element.update()
	}
}

func (r *Button) draw(ctx uiBundle) {
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}, 2.0, r.Color)
	if r.hovered {
		rl.DrawRectangle(r.X, r.Y, r.Width, r.Height, rl.Color{R: r.Color.R, G: r.Color.G, B: r.Color.B, A: r.Color.A / 5})
	}

	text_size := rl.MeasureTextEx(ctx.text_renderer.font, r.Text, float32(ctx.text_renderer.font.BaseSize)/8, 1)
	ctx.text_renderer.DrawText(r.Text, int32(float32(r.X)+float32(r.Width/2)-float32(text_size.X/2)), int32(float32(r.Y)+float32(r.Height/2)-float32(text_size.Y/2)), r.Color)
}

func (r *Button) update() {

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.X), Y: float32(r.Y), Width: float32(r.Width), Height: float32(r.Height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			r.On_click()
		}
	} else {
		r.hovered = false
	}
}

func SetupBundle() uiBundle {
	out := uiBundle{text_renderer: getTextRenderer()}
	return out
}
