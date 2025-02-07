package elements

import (
	"fmt"

	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TagManager struct {
	x, y, width, height int32
	color               rl.Color
	input               *TextBox
	add                 *Button
	addedTags           []string
}

func NewTagManager(X, Y, width, height int32, color rl.Color) *TagManager {
	tb := NewTextBox(X, Y, width-10-height, height, "tags", color)
	tm := &TagManager{x: X, y: Y, width: width, height: height, color: color, input: tb}
	tm.add = NewButton(X+width-height, Y, height, height, "+", color, tm.addTag)
	return tm
}

func (r *TagManager) addTag() {
	if r.input.Text == "" {
		return
	}
	r.addedTags = append(r.addedTags, r.input.Text)
	r.input.ClearText()
	fmt.Println(r.addedTags)
}

func (r *TagManager) Draw(ctx *ui.UiBundle) {
	r.input.Draw(ctx)
	r.add.Draw(ctx)
	cur_x_pos := int32(0)
	for _, e := range r.addedTags {
		txt_width := int32(ctx.MeasureText(e).X)
		rl.DrawRectangle(r.x+r.width+cur_x_pos+6, r.y, txt_width+8, 28, rl.Gray)
		ctx.Text_renderer.DrawText(e, r.x+r.width+cur_x_pos+10, r.y+5, rl.DarkGray)
		cur_x_pos += txt_width + 14
	}
}

func (r *TagManager) Update(ctx *ui.UiBundle) {
	r.input.Update(ctx)
	r.add.Update(ctx)
}
