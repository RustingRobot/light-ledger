package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	d "github.com/RustingRobot/light-ledger/data"
	"github.com/RustingRobot/light-ledger/ui"
	e "github.com/RustingRobot/light-ledger/ui/elements"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var data = ""
var true_data d.Data = d.Data{}

func main() {
	if content, err := os.ReadFile("db.json"); err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile("db.json", []byte(`{"expenses":{"desc":[],"cost":[]}}`), 0666); err != nil {
				fmt.Println("ERROR")
			}
		}
	} else {
		json.Unmarshal(content, &true_data)
		//data = string(content)
	}
	path, _ := os.Getwd()
	db_location := path + "/db.json"

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	root := ui.NewBundle()
	addTab := e.NewContainer()
	tableTab := e.NewContainer()
	calendarTab := e.NewContainer()
	visualizeTab := e.NewContainer()
	descButton := e.NewTextBox(100, 200, 300, 28, "description", rl.White)
	costButton := e.NewTextBox(410, 200, 100, 28, "cost", rl.White)
	tabs := e.NewTabs(10, 50, 28, 150, []string{"add value", "data table", "calendar", "visualization"}, []*e.Container{addTab, tableTab, calendarTab, visualizeTab}, rl.White)
	addTab.Add(e.NewButton(520, 200, 300, 28, "add", rl.White, func() { addEntry(descButton.Text, costButton.Text) }))
	dirText := e.NewText(200, 5, db_location, rl.Gray)
	root.Add(dirText)
	root.Add(e.NewText(5, 5, "current database:", rl.LightGray))
	addTab.Add(descButton)
	addTab.Add(costButton)
	root.Add(tabs)

	addTab.Add(e.NewText(100, 100, "This is tab 1", rl.White))
	tableTab.Add(e.NewTable(100, 100, &true_data, rl.White))
	calendarTab.Add(e.NewText(100, 100, "This is tab 3", rl.White))
	root.Add(addTab)
	root.Add(tableTab)
	root.Add(calendarTab)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		if rl.IsFileDropped() {
			dirText.SetText(rl.LoadDroppedFiles()[0])
		}

		root.Update()
		root.Draw()
		rl.EndDrawing()
	}
}

func addEntry(desc string, cost string) {
	true_data.Expenses.Cost = append(true_data.Expenses.Cost, cost)
	true_data.Expenses.Description = append(true_data.Expenses.Description, desc)
	d.SaveToFile(true_data)
}
