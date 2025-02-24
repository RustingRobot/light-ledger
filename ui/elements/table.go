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

const header_margin = 1.5

func NewTable(X int32, Y int32, data *data.Data, color rl.Color) *Table {
	return &Table{x: X, y: Y, data: data, color: color}
}

func (r *Table) Draw(ctx *ui.UiBundle) {
	ctx.Text_renderer.DrawText("description", r.x+40, r.y, r.color)
	ctx.Text_renderer.DrawText("cost", r.x+440, r.y, r.color)
	ctx.Text_renderer.DrawText("date", r.x+540, r.y, r.color)
	ctx.Text_renderer.DrawText("tags", r.x+680, r.y, r.color)

	for index, entry := range r.data.Expenses {
		if index%2 != 0 {
			rl.DrawRectangle(r.x, r.y+22*int32(index+1)+5, 800, 20, rl.Gray)
		}
		ctx.Text_renderer.DrawText(entry.Description, r.x+40, r.y+22*int32(index+1)+5, r.color)
		ctx.Text_renderer.DrawText(entry.Cost, r.x+440, r.y+22*int32(index+1)+5, r.color)
		ctx.Text_renderer.DrawText(entry.Date, r.x+540, r.y+22*int32(index+1)+5, r.color)

		cur_x_pos := int32(5)
		for _, e := range entry.Tags {
			txt_width := int32(ctx.MeasureText(e).X)
			rl.DrawRectangle(cur_x_pos+r.x+680+6, r.y+22*int32(index+1)+5, txt_width+8, 20, rl.White)
			ctx.Text_renderer.DrawText(e, cur_x_pos+r.x+680+10, r.y+22*int32(index+1)+5, rl.DarkGray)
			cur_x_pos += txt_width + 14
		}
	}
	rl.DrawRectangle(r.x, r.y+22, 800, 2, rl.Red)
	for _, btn := range r.buttons {
		btn.Draw(ctx)
	}
	r.buttons = nil
}

func (r *Table) Update(ctx *ui.UiBundle) {
	for index, _ := range r.data.Expenses {
		btn := NewButton(r.x, r.y+22*int32(index+1)+5, 20, 20, "x", rl.White, func() { r.deleteEntry(r.data, index) })
		btn.Y_offset = -2
		r.buttons = append(r.buttons, btn)
		btn.Update(ctx)
	}
}

func (r *Table) deleteEntry(data *data.Data, index int) {
	data.Expenses = append(data.Expenses[:index], data.Expenses[index+1:]...)
	d.SaveToFile(*data, data.Expenses[index].Tags, false)
}
