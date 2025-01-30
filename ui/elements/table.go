package elements

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RustingRobot/light-ledger/data"
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Table struct {
	x, y    int32
	data    *string
	color   rl.Color
	buttons []ui.UiElement
}

func NewTable(X int32, Y int32, data *string, color rl.Color) *Table {
	return &Table{x: X, y: Y, data: data, color: color}
}

func (r *Table) Draw(ctx *ui.UiBundle) {
	var data data.Data
	err := json.Unmarshal([]byte(*r.data), &data)
	if err != nil {
		log.Fatal(err)
	}
	for index, entry := range data.Expenses.Description {
		ctx.Text_renderer.DrawText(entry, r.x+40, r.y+22*int32(index), r.color)
		ctx.Text_renderer.DrawText(data.Expenses.Cost[index], r.x+240, r.y+22*int32(index), r.color)
	}

	for _, btn := range r.buttons {
		btn.Draw(ctx)
	}
	r.buttons = nil
}

func (r *Table) Update(ctx *ui.UiBundle) {
	var data data.Data
	err := json.Unmarshal([]byte(*r.data), &data)
	if err != nil {
		log.Fatal(err)
	}
	for index, entry := range data.Expenses.Description {
		btn := NewButton(r.x, r.y+22*int32(index), 20, 20, "X", rl.White, func() { fmt.Println(entry) })
		r.buttons = append(r.buttons, btn)
		btn.Update(ctx)
	}
}
