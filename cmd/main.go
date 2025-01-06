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
	uiBundle.Add(elements.NewButton(100, 100, 300, 50, "click me!", rl.White, func() { fmt.Println("click") }))
	uiBundle.Add(elements.NewTextBox(100, 170, 300, 28, "one", rl.White))
	uiBundle.Add(elements.NewTextBox(100, 218, 300, 28, "two", rl.White))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		uiBundle.Draw()
		uiBundle.Update()
		rl.EndDrawing()
	}
}
