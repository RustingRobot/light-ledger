package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"time"

	d "github.com/RustingRobot/light-ledger/data"
	"github.com/RustingRobot/light-ledger/ui"
	e "github.com/RustingRobot/light-ledger/ui/elements"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var true_data d.Data = d.Data{}

func main() {

	if content, err := os.ReadFile("db.json"); err != nil {
		data, _ := json.Marshal(true_data)
		os.WriteFile("db.json", data, 0666)
	} else {
		json.Unmarshal(content, &true_data)
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
	descTextbox := e.NewTextBox(10, 100, 400, 28, "description", rl.White)
	costTextbox := e.NewTextBox(420, 100, 90, 28, "cost", rl.White)
	dateTextbox := e.NewTextBox(520, 100, 135, 28, "", rl.White)
	t := time.Now()
	dateTextbox.Text = fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	/* 	useCurrentTimeBox := e.NewCheckbox(100, 238, "use current date", rl.White)
	   	useCurrentTimeBox.Checked = true */

	tabs := e.NewTabs(10, 50, 28, 150, []string{"add value", "data table", "calendar", "visualization"}, []*e.Container{addTab, tableTab, calendarTab, visualizeTab}, rl.White)
	addTab.Add(e.NewButton(10, 176, 300, 28, "add", rl.White, func() { addEntry(descTextbox, costTextbox) }))
	dirText := e.NewText(200, 10, db_location, rl.Gray)
	root.Add(dirText)
	root.Add(e.NewText(5, 10, "current database:", rl.LightGray))
	addTab.Add(descTextbox)
	addTab.Add(costTextbox)
	addTab.Add(dateTextbox)
	addTab.Add(e.NewTagManager(10, 138, 200, 28, rl.White))
	root.Add(tabs)

	tableTab.Add(e.NewTable(10, 100, &true_data, rl.White))
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

func addEntry(desc *e.TextBox, cost *e.TextBox) {
	true_data.Expenses.Cost = append(true_data.Expenses.Cost, cost.Text)
	true_data.Expenses.Description = append(true_data.Expenses.Description, desc.Text)
	desc.ClearText()
	cost.ClearText()
	d.SaveToFile(true_data)
}
