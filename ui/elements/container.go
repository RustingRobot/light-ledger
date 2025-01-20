package elements

import (
	"github.com/RustingRobot/light-ledger/ui"
)

type Container struct {
	Active      bool
	ui_elements []ui.UiElement
}

func NewContainer() *Container {
	return &Container{Active: true}
}

func (r *Container) Add(u ui.UiElement) {
	r.ui_elements = append(r.ui_elements, u)
}

func (r *Container) Draw(ctx *ui.UiBundle) {
	if !r.Active {
		return
	}
	for _, ui_element := range r.ui_elements {
		ui_element.Draw(ctx)
	}
}

func (r *Container) Update(ctx *ui.UiBundle) {
	if !r.Active {
		return
	}
	for _, ui_element := range r.ui_elements {
		ui_element.Update(ctx)
	}
}
