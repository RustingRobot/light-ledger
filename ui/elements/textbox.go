package elements

import (
	"math"

	"github.com/RustingRobot/light-ledger/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBox struct {
	x, y                 int32
	width, height        int32
	placeholder_text     string
	Text                 string
	color                rl.Color
	hovered              bool
	cursor_pos           int
	selection_pos        int
	Enter_func, Tab_func func()
}

func NewTextBox(X int32, Y int32, Width int32, Height int32, Placeholder_text string, Color rl.Color) *TextBox {
	return &TextBox{x: X, y: Y, width: Width, height: Height, placeholder_text: Placeholder_text, color: Color, cursor_pos: 0}
}

func (r *TextBox) ClearText() {
	r.Text = ""
	r.cursor_pos = 0
	r.selection_pos = 0
}

func (r *TextBox) Draw(ctx *ui.UiBundle) {
	rl.BeginScissorMode(r.x, r.y, r.width, r.height)
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
				pos1 = r.x + 12 + int32(ctx.MeasureText(r.Text[:r.cursor_pos]).X)
				pos2 = int32(ctx.MeasureText(r.Text[r.cursor_pos:r.selection_pos]).X)
			} else {
				pos1 = r.x + 12 + int32(ctx.MeasureText(r.Text[:r.selection_pos]).X)
				pos2 = int32(ctx.MeasureText(r.Text[r.selection_pos:r.cursor_pos]).X)
			}
			rl.DrawRectangle(pos1, r.y, pos2, r.height, rl.Blue)
		}
		// cursor line
		cursor_x := r.x + int32(ctx.MeasureText(r.Text[:r.cursor_pos]).X) + 12
		rl.DrawLine(cursor_x, r.y+5, cursor_x, r.y+r.height-5, r.color)
	}
	// outline
	rl.DrawRectangleLinesEx(rl.Rectangle{X: float32(r.x), Y: float32(r.y), Width: float32(r.width), Height: float32(r.height)}, 1.0, temp_color)
	if r.hovered {
		// highlight
		rl.DrawRectangle(r.x, r.y, r.width, r.height, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	}

	if r.Text == "" {
		ctx.Text_renderer.DrawText(r.placeholder_text, r.x+10, r.y+4, rl.Color{R: r.color.R, G: r.color.G, B: r.color.B, A: r.color.A / 5})
	} else {
		ctx.Text_renderer.DrawText(r.Text, r.x+10, r.y+4, r.color)
	}
	rl.EndScissorMode()
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

				temp_length = int(ctx.MeasureText(r.Text[:char_nr]).X)
				for temp_length < mouse_pos && len(r.Text) > char_nr {
					char_nr++
					last_length = temp_length
					temp_length = int(ctx.MeasureText(r.Text[:char_nr]).X)
				}
				if char_nr == 0 {
					r.cursor_pos = 0
				} else if len(r.Text) == char_nr && temp_length < int(rl.GetMousePosition().X)-int(r.x)-10 {
					r.cursor_pos = len(r.Text)
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
				r.selection_pos = r.cursor_pos
			}
		}
	} else {
		r.hovered = false
	}

	if ctx.Selected == r {
		// backspace
		if rl.IsKeyPressed(rl.KeyBackspace) {
			if r.cursor_pos > r.selection_pos {
				r.Text = r.Text[:r.selection_pos] + r.Text[r.cursor_pos:]
				r.cursor_pos = r.selection_pos
			} else if r.cursor_pos < r.selection_pos {
				r.Text = r.Text[:r.cursor_pos] + r.Text[r.selection_pos:]
				r.selection_pos = r.cursor_pos
			} else if r.cursor_pos > 0 {
				r.Text = r.Text[:r.cursor_pos-1] + r.Text[r.cursor_pos:]
				r.cursor_pos--
				r.selection_pos--
			}
		}

		//text input
		key := rl.GetCharPressed()
		for key > 0 {
			if (key >= 32) && (key <= 125) {
				if r.cursor_pos == r.selection_pos {
					r.Text = r.Text[:r.cursor_pos] + string(key) + r.Text[r.cursor_pos:]
					r.cursor_pos++
					r.selection_pos++
				} else if r.cursor_pos > r.selection_pos {
					r.Text = r.Text[:r.selection_pos] + string(key) + r.Text[r.cursor_pos:]
					r.selection_pos++
					r.cursor_pos = r.selection_pos
				} else {
					r.Text = r.Text[:r.cursor_pos] + string(key) + r.Text[r.selection_pos:]
					r.cursor_pos++
					r.selection_pos = r.cursor_pos
				}
			}
			key = rl.GetCharPressed()
		}

		// move with arrow keys
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

		if rl.IsKeyPressed(rl.KeyRight) && int(r.cursor_pos) < len(r.Text) {
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
			r.selection_pos = len(r.Text)
		}

		// select all ctrl + a
		if rl.IsKeyPressed(rl.KeyA) && ctrlDown() {
			r.selection_pos = 0
			r.cursor_pos = len(r.Text)
		}

		// paste ctrl + v
		if rl.IsKeyPressed(rl.KeyV) && ctrlDown() {
			cb_text := rl.GetClipboardText()
			if r.cursor_pos == r.selection_pos {
				r.Text = r.Text[:r.cursor_pos] + cb_text + r.Text[r.cursor_pos:]
				r.cursor_pos += len(cb_text)
				r.selection_pos = r.cursor_pos
			} else if r.cursor_pos > r.selection_pos {
				r.Text = r.Text[:r.selection_pos] + cb_text + r.Text[r.cursor_pos:]
				r.selection_pos += len(cb_text)
				r.cursor_pos = r.selection_pos
			} else {
				r.Text = r.Text[:r.cursor_pos] + cb_text + r.Text[r.selection_pos:]
				r.cursor_pos += len(cb_text)
				r.selection_pos = r.cursor_pos
			}
		}

		// copy ctrl + c
		if rl.IsKeyPressed(rl.KeyC) && ctrlDown() {
			if r.cursor_pos > r.selection_pos {
				rl.SetClipboardText(r.Text[r.selection_pos:r.cursor_pos])
			} else {
				rl.SetClipboardText(r.Text[r.cursor_pos:r.selection_pos])
			}
		}

		// cut ctrl + x
		if rl.IsKeyPressed(rl.KeyX) && ctrlDown() {
			if r.cursor_pos > r.selection_pos {
				rl.SetClipboardText(r.Text[r.selection_pos:r.cursor_pos])
				r.Text = r.Text[:r.selection_pos] + r.Text[r.cursor_pos:]
				r.cursor_pos = r.selection_pos
			} else if r.cursor_pos < r.selection_pos {
				rl.SetClipboardText(r.Text[r.cursor_pos:r.selection_pos])
				r.Text = r.Text[:r.cursor_pos] + r.Text[r.selection_pos:]
				r.selection_pos = r.cursor_pos
			}
		}

		if rl.IsKeyPressed(rl.KeyTab) && r.Tab_func != nil {
			r.Tab_func()
		}

		if rl.IsKeyPressed(rl.KeyEnter) && r.Enter_func != nil {
			r.Enter_func()
		}
	}
}

func shiftDown() bool {
	return rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
}

func ctrlDown() bool {
	return rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl)
}
