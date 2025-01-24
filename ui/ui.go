package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type UiElement interface {
	Draw(ctx *UiBundle)
	Update(ctx *UiBundle)
}

type UiBundle struct {
	Text_renderer TextRenderer
	ui_elements   []UiElement
	Selected      UiElement
}

func (r *UiBundle) Add(u UiElement) {
	r.ui_elements = append(r.ui_elements, u)
}

func (r *UiBundle) Draw() {
	for _, ui_element := range r.ui_elements {
		ui_element.Draw(r)
	}
}

func (r *UiBundle) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		r.Selected = nil
	}

	for _, ui_element := range r.ui_elements {
		ui_element.Update(r)
	}
}

func NewBundle() *UiBundle {
	return &UiBundle{Text_renderer: getTextRenderer()}
}

func (r *UiBundle) MeasureText(text string) rl.Vector2 {
	return rl.Vector2{X: float32(rl.MeasureText(text, 20)), Y: 0}
	//return rl.MeasureTextEx(r.Text_renderer.Font, text, r.Text_renderer.Size, 1)
}
