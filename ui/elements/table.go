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
	x, y  int32
	data  *string
	color rl.Color
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
	ctx.RemoveLabeled("table")
	for index, entry := range data.Expenses.Description {
		ctx.Text_renderer.DrawText(entry, r.x+40, r.y+22*int32(index), r.color)
		ctx.Text_renderer.DrawText(data.Expenses.Cost[index], r.x+240, r.y+22*int32(index), r.color)
		ctx.AddLabeled(NewButton(r.x, r.y+22*int32(index), 20, 20, "X", rl.White, func() { fmt.Println(entry) }), "table")
	}
}

func (r *Table) Update(ctx *ui.UiBundle) {}
