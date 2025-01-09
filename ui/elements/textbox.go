package elements

import (
	"math"

	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBox struct {
	x, y             int32
	width, height    int32
	placeholder_text string
	text             string
	color            rl.Color
	hovered          bool
	cursor_pos       int
	selection_pos    int
}

func NewTextBox(X int32, Y int32, Width int32, Height int32, Placeholder_text string, Color rl.Color) *TextBox {
	return &TextBox{x: X, y: Y, width: Width, height: Height, placeholder_text: Placeholder_text, color: Color, cursor_pos: 0}
}

func (r *TextBox) Draw(ctx *ui.UiBundle) {
	temp_color := r.color
	if ctx.Selected == r {
		temp_color = rl.Red
		if r.hovered {
			rl.SetMouseCursor(rl.MouseCursorIBeam)
		} else {
			rl.SetMouseCursor(rl.MouseCursorDefault)
		}
		// text selection
		if r.cursor_pos != r.selection_pos {
			var pos1, pos2 int32
			if r.cursor_pos < r.selection_pos {
				pos1 = r.x + 12 + int32(ctx.MeasureText(r.text[:r.cursor_pos]).X)
				pos2 = int32(ctx.MeasureText(r.text[r.cursor_pos:r.selection_pos]).X)
			} else {
				pos1 = r.x + 12 + int32(ctx.MeasureText(r.text[:r.selection_pos]).X)
				pos2 = int32(ctx.MeasureText(r.text[r.selection_pos:r.cursor_pos]).X)
			}
			rl.DrawRectangle(pos1, r.y, pos2, r.height, rl.Blue)
		}
		// cursor line
		cursor_x := r.x + int32(ctx.MeasureText(r.text[:r.cursor_pos]).X) + 12
		rl.DrawLine(cursor_x, r.y+5, cursor_x, r.y+r.height-5, r.color)
	}
	// outline
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}, 2.0, temp_color)
	if r.hovered {
		// highlight
		rl.DrawRectangle(r.x, r.y, r.width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}

	if r.text == "" {
		ctx.Text_renderer.DrawText(r.placeholder_text, r.x+10, r.y, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	} else {
		ctx.Text_renderer.DrawText(r.text, r.x+10, r.y, r.color)
	}
}

func (r *TextBox) Update(ctx *ui.UiBundle) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}) {
		r.hovered = true
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) || rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			if ctx.Selected == r {
				// set cursor pos
				temp_length := 10
				last_length := temp_length
				mouse_pos := int(rl.GetMousePosition().X) - int(r.x) - 10
				char_nr := 0

				temp_length = int(ctx.MeasureText(r.text[:char_nr]).X)
				for temp_length < mouse_pos && len(r.text) > char_nr {
					char_nr++
					last_length = temp_length
					temp_length = int(ctx.MeasureText(r.text[:char_nr]).X)
				}
				if char_nr == 0 {
					r.cursor_pos = 0
				} else if len(r.text) == char_nr && temp_length < int(rl.GetMousePosition().X)-int(r.x)-10 {
					r.cursor_pos = len(r.text)
				} else {
					if math.Abs(float64(temp_length)-float64(mouse_pos)) < math.Abs(float64(last_length)-float64(mouse_pos)) {
						r.cursor_pos = char_nr
					} else {
						r.cursor_pos = char_nr - 1
					}
				}
				if !shiftDown() && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					r.selection_pos = r.cursor_pos
				}

			} else {
				ctx.Selected = r
			}
		}
	} else {
		r.hovered = false
	}

	if ctx.Selected == r {
		if rl.IsKeyPressed(rl.KeyBackspace) && r.cursor_pos > 0 {
			r.text = r.text[:r.cursor_pos-1] + r.text[r.cursor_pos:]
			r.cursor_pos--
			r.selection_pos--
		}

		key := rl.GetCharPressed()
		for key > 0 {
			if (key >= 32) && (key <= 125) {
				r.text = r.text[:r.cursor_pos] + string(key) + r.text[r.cursor_pos:]
				r.cursor_pos++
				r.selection_pos++
			}
			key = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyLeft) && r.cursor_pos > 0 {
			if shiftDown() {
				r.cursor_pos--
			} else if r.cursor_pos > r.selection_pos {
				r.cursor_pos = r.selection_pos
			} else if r.cursor_pos < r.selection_pos {
				r.selection_pos = r.cursor_pos
			} else {
				r.selection_pos--
				r.cursor_pos--
			}
		} else if rl.IsKeyPressed(rl.KeyLeft) && !shiftDown() {
			r.selection_pos = 0
		}

		if rl.IsKeyPressed(rl.KeyRight) && int(r.cursor_pos) < len(r.text) {
			if shiftDown() {
				r.cursor_pos++
			} else if r.cursor_pos < r.selection_pos {
				r.cursor_pos = r.selection_pos
			} else if r.cursor_pos > r.selection_pos {
				r.selection_pos = r.cursor_pos
			} else {
				r.selection_pos++
				r.cursor_pos++
			}
		} else if rl.IsKeyPressed(rl.KeyRight) && !shiftDown() {
			r.selection_pos = len(r.text)
		}
	}
}

func shiftDown() bool {
	return rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
}
