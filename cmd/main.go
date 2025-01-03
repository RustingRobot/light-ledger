package main

import (
	_ "embed"
	"fmt"
	"lightledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(900, 450, "Light Ledger")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var uiBundle = ui.SetupBundle()
	uiBundle.Add(&ui.Button{X: 100, Y: 100, Width: 300, Height: 50, On_click: func() { fmt.Println("click") }, Text: "click me", Color: rl.White})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		uiBundle.Draw()
		uiBundle.Update()
		rl.EndDrawing()
	}
}
