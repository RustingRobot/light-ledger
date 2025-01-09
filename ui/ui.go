package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type uiElement interface {
	Draw(ctx *UiBundle)
	Update(ctx *UiBundle)
}

type UiBundle struct {
	Text_renderer TextRenderer
	Ui_elements   []uiElement
	Selected      uiElement
}

func (r *UiBundle) Add(u uiElement) {
	r.Ui_elements = append(r.Ui_elements, u)
}

func (r *UiBundle) Draw() {
	for _, ui_element := range r.Ui_elements {
		ui_element.Draw(r)
	}
}

func (r *UiBundle) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		r.Selected = nil
	}

	for _, ui_element := range r.Ui_elements {
		ui_element.Update(r)
	}
}

func SetupBundle() UiBundle {
	out := UiBundle{Text_renderer: getTextRenderer()}
	return out
}

func (r *UiBundle) MeasureText(text string) rl.Vector2 {
	return rl.MeasureTextEx(r.Text_renderer.Font, text, r.Text_renderer.Size, 1)
}
