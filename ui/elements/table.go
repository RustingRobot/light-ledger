package elements

import (
	"github.com/RustingRobot/light-ledger/data"
	d "github.com/RustingRobot/light-ledger/data"
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Table struct {
	x, y    int32
	data    *data.Data
	color   rl.Color
	buttons []ui.UiElement
}

func NewTable(X int32, Y int32, data *data.Data, color rl.Color) *Table {
	return &Table{x: X, y: Y, data: data, color: color}
}

func (r *Table) Draw(ctx *ui.UiBundle) {
	for index, entry := range r.data.Expenses.Description {
		if index%2 == 0 {
			rl.DrawRectangle(r.x, r.y+22*int32(index), 500, 20, rl.Gray)
		}
		ctx.Text_renderer.DrawText(entry, r.x+40, r.y+22*int32(index), r.color)
		ctx.Text_renderer.DrawText(r.data.Expenses.Cost[index], r.x+240, r.y+22*int32(index), r.color)
	}

	for _, btn := range r.buttons {
		btn.Draw(ctx)
	}
	r.buttons = nil
}

func (r *Table) Update(ctx *ui.UiBundle) {
	for index, _ := range r.data.Expenses.Description {
		btn := NewButton(r.x, r.y+22*int32(index), 20, 20, "X", rl.White, func() { r.deleteEntry(r.data, index) })
		r.buttons = append(r.buttons, btn)
		btn.Update(ctx)
	}
}

func (r *Table) deleteEntry(data *data.Data, index int) {
	data.Expenses.Cost = append(data.Expenses.Cost[:index], data.Expenses.Cost[index+1:]...)
	data.Expenses.Description = append(data.Expenses.Description[:index], data.Expenses.Description[index+1:]...)
	d.SaveToFile(*data)
}
