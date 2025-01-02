package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type uiElement interface {
	draw()
	update()
}

type button struct {
	x, y          int32
	width, height int32
	on_click      func()
	text          string
	color         rl.Color
}

type uiBundle struct {
	text_renderer textRenderer
	ui_elements   []uiElement
}

func (r uiBundle) add(u uiElement) {
	r.ui_elements = append(r.ui_elements, u)
}

func (r uiBundle) draw() {
	for _, ui_element := range r.ui_elements {
		ui_element.draw()
	}
}

func (r uiBundle) update() {
	for _, ui_element := range r.ui_elements {
		ui_element.update()
	}
}

func (r button) draw() {
	rl.DrawRectangle(r.x, r.y, r.width, r.height, r.color)
}

func (r button) update() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}) {
		fmt.Println("pressed!")
	}
}

func setup() /* uiBundle */ {
	fmt.Println("test")
}
