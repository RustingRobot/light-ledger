package elements

import (
	"strings"

	"slices"

	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TagManager struct {
	x, y, width, height int32
	color               rl.Color
	input               *TextBox
	add                 *Button
	added_tags          []string
	hovered_tag         int
}

func NewTagManager(X, Y, width, height int32, color rl.Color) *TagManager {
	tb := NewTextBox(X, Y, width-10-height, height, "tags", color)
	tm := &TagManager{x: X, y: Y, width: width, height: height, color: color, input: tb}
	tb.Enter_func = tm.addTag
	tm.add = NewButton(X+width-height, Y, height, height, "+", color, tm.addTag)
	return tm
}

func (r *TagManager) GetText() string {
	return r.input.Text
}

func (r *TagManager) EmptyTags() {
	r.added_tags = nil
}

func (r *TagManager) GetTags() []string {
	new_array := make([]string, len(r.added_tags))
	copy(new_array, r.added_tags)
	return new_array
}

func (r *TagManager) addTag() {
	if strings.Trim(r.input.Text, " ") == "" || slices.Contains(r.added_tags, r.input.Text) {
		return
	}
	r.added_tags = append(r.added_tags, strings.Trim(r.input.Text, " "))
	r.input.ClearText()
}

func (r *TagManager) Draw(ctx *ui.UiBundle) {
	r.input.Draw(ctx)
	r.add.Draw(ctx)
	cur_x_pos := int32(5)
	r.hovered_tag = -1
	for i, e := range r.added_tags {
		txt_width := int32(ctx.MeasureText(e).X)
		color := rl.Gray
		// nasty update in draw function
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x + r.width + cur_x_pos + 6), Y: float32(r.y), Width: float32(txt_width + 8), Height: float32(28)}) {
			color = rl.White
			r.hovered_tag = i
		}

		rl.DrawRectangle(r.x+r.width+cur_x_pos+6, r.y, txt_width+8, 28, color)
		ctx.Text_renderer.DrawText(e, r.x+r.width+cur_x_pos+10, r.y+5, rl.DarkGray)
		cur_x_pos += txt_width + 14
	}
}

func (r *TagManager) Update(ctx *ui.UiBundle) {
	if r.hovered_tag >= 0 && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		r.added_tags = slices.Delete(r.added_tags, r.hovered_tag, r.hovered_tag+1)
	}
	r.input.Update(ctx)
	r.add.Update(ctx)
}
