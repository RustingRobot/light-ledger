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

	var uiBundle = ui.SetupBundle()
	var descButton = elements.NewTextBox(100, 200, 300, 28, "description", rl.White)
	var costButton = elements.NewTextBox(410, 200, 100, 28, "cost", rl.White)
	uiBundle.Add(elements.NewButton(520, 200, 300, 28, "add", rl.White, func() { saveToFile(descButton.Text, costButton.Text) }))
	uiBundle.Add(descButton)
	uiBundle.Add(costButton)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		uiBundle.Draw()
		uiBundle.Update()
		rl.EndDrawing()
	}
}

var data = `{"expenses":{"desc":[],"cost":[]}}`

func saveToFile(desc string, cost string) {
	data, _ = sjson.Set(data, "expenses.desc.-1", desc)
	data, _ = sjson.Set(data, "expenses.cost.-1", cost)
	fmt.Println(data)
}
