package main

import (
	_ "embed"
	"fmt"

	"github.com/RustingRobot/light-ledger/ui"
	"github.com/RustingRobot/light-ledger/ui/elements"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tidwall/sjson"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	root := ui.NewBundle()
	tab1 := elements.NewContainer()
	descButton := elements.NewTextBox(100, 200, 300, 28, "description", rl.White)
	costButton := elements.NewTextBox(410, 200, 100, 28, "cost", rl.White)
	root.Add(elements.NewButton(520, 200, 300, 28, "add", rl.White, func() { saveToFile(descButton.Text, costButton.Text, tab1) }))
	root.Add(descButton)
	root.Add(costButton)

	tab1.Add(elements.NewText(100, 100, "This is in a different container", rl.White))
	root.Add(tab1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		root.Draw()
		root.Update()
		rl.EndDrawing()
	}
}

var data = `{"expenses":{"desc":[],"cost":[]}}`

func saveToFile(desc string, cost string, tab *elements.Container) {
	tab.Active = !tab.Active
	data, _ = sjson.Set(data, "expenses.desc.-1", desc)
	data, _ = sjson.Set(data, "expenses.cost.-1", cost)
	fmt.Println(data)
}
