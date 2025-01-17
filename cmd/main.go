package main

import (
	_ "embed"
	"fmt"

	"github.com/RustingRobot/light-ledger/ui"
	e "github.com/RustingRobot/light-ledger/ui/elements"
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
	tab1 := e.NewContainer()
	tab2 := e.NewContainer()
	tab3 := e.NewContainer()
	tab4 := e.NewContainer()
	descButton := e.NewTextBox(100, 200, 300, 28, "description", rl.White)
	costButton := e.NewTextBox(410, 200, 100, 28, "cost", rl.White)
	tabs := e.NewTabs(100, 50, 28, 130, []string{"tab 1", "tab 2", "last tab", "extra last tab"}, []*e.Container{tab1, tab2, tab3, tab4}, rl.White)
	tab1.Add(e.NewButton(520, 200, 300, 28, "add", rl.White, func() { saveToFile(descButton.Text, costButton.Text, tab1) }))
	tab1.Add(descButton)
	tab1.Add(costButton)
	root.Add(tabs)

	tab1.Add(e.NewText(100, 100, "This is tab 1", rl.White))
	tab2.Add(e.NewText(100, 100, "This is tab 2", rl.White))
	tab3.Add(e.NewText(100, 100, "This is tab 3", rl.White))
	root.Add(tab1)
	root.Add(tab2)
	root.Add(tab3)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		root.Draw()
		root.Update()
		rl.EndDrawing()
	}
}

var data = `{"expenses":{"desc":[],"cost":[]}}`

func saveToFile(desc string, cost string, tab *e.Container) {
	tab.Active = !tab.Active
	data, _ = sjson.Set(data, "expenses.desc.-1", desc)
	data, _ = sjson.Set(data, "expenses.cost.-1", cost)
	fmt.Println(data)
}
