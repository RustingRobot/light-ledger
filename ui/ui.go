package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UiElement interface {
	Draw(ctx *UiBundle)
	Update(ctx *UiBundle)
}

type UiBundle struct {
	Text_renderer TextRenderer
	ui_elements   []UiElement
	Selected      UiElement
	label         map[string][]int
}

func (r *UiBundle) Add(u UiElement) {
	r.ui_elements = append(r.ui_elements, u)
}

func (r *UiBundle) AddLabeled(u UiElement, label string) {
	r.ui_elements = append(r.ui_elements, u)
	r.label[label] = append(r.label[label], len(r.ui_elements))
}

func (r *UiBundle) RemoveLabeled(label string) {
	for _, index := range r.label[label] {
		r.ui_elements[index] = r.ui_elements[len(r.ui_elements)-1]
		r.ui_elements = r.ui_elements[:len(r.ui_elements)-1]
		return
	}
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
	bundle := UiBundle{Text_renderer: getTextRenderer()}
	bundle.label = make(map[string][]int)
	return &bundle
}

func (r *UiBundle) MeasureText(text string) rl.Vector2 {
	return rl.Vector2{X: float32(rl.MeasureText(text, int32(r.Text_renderer.Size))), Y: 0}
}
