package main

import (
	_ "embed"
	"fmt"

	"github.com/RustingRobot/light-ledger/ui"
	"github.com/RustingRobot/light-ledger/ui/elements"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var uiBundle = ui.SetupBundle()
	uiBundle.Add(&elements.Button{X: 100, Y: 100, Width: 300, Height: 50, On_click: func() { fmt.Println("click") }, Text: "click me", Color: rl.White})
	uiBundle.Add(&elements.TextBox{X: 100, Y: 170, Width: 300, Height: 28, Placeholder_Text: "one", Color: rl.White})
	uiBundle.Add(&elements.TextBox{X: 100, Y: 218, Width: 300, Height: 28, Placeholder_Text: "two", Color: rl.White})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		uiBundle.Draw()
		uiBundle.Update()
		rl.EndDrawing()
	}
}
