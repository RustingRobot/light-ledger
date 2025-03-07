package elements

import (
	"slices"
	"sort"
	"strings"

	"github.com/RustingRobot/light-ledger/data"
	"github.com/RustingRobot/light-ledger/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type QuickTags struct {
	x, y        int32
	color       rl.Color
	hovered_tag string
	data        *data.Data
	tag_manager *TagManager
}

func NewQuickTags(X, Y int32, color rl.Color, tag_manager *TagManager, data *data.Data) *QuickTags {
	return &QuickTags{x: X, y: Y, color: color, data: data, tag_manager: tag_manager}
}

func (r *QuickTags) Draw(ctx *ui.UiBundle) {
	ctx.Text_renderer.DrawText("tag history:", r.x+10, r.y, r.color)
	cur_x_pos := int32(0)
	type keyValue struct {
		key   string
		value int
	}

	var items []keyValue
	for k, v := range r.data.Tags {
		items = append(items, keyValue{k, v})
	}

	// Sort the slice: first by value (ascending), then by key (alphabetically)
	sort.Slice(items, func(i, j int) bool {
		if items[i].value == items[j].value {
			return items[i].key < items[j].key // Sort alphabetically if values are equal
		}
		return items[i].value > items[j].value // Sort by value if not equal
	})
	r.hovered_tag = ""
	for _, tag := range items {
		if !strings.HasPrefix(tag.key, r.tag_manager.GetText()) {
			continue
		}
		txt_width := int32(ctx.MeasureText(tag.key).X)
		color := rl.Gray
		// nasty update in draw function
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: float32(r.x + cur_x_pos + 135), Y: float32(r.y - 2), Width: float32(txt_width + 8), Height: float32(28)}) {
			color = rl.White
			r.hovered_tag = tag.key
		}

		rl.DrawRectangle(int32(r.x+cur_x_pos+115), int32(r.y-2), txt_width+10, 25, color)
		ctx.Text_renderer.DrawText(tag.key, r.x+cur_x_pos+120, r.y+3, rl.DarkGray)
		cur_x_pos += txt_width + 16
	}
}

func (r *QuickTags) Update(ctx *ui.UiBundle) {
	if r.hovered_tag != "" && rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !slices.Contains(r.tag_manager.added_tags, r.hovered_tag) {
		r.tag_manager.added_tags = append(r.tag_manager.added_tags, r.hovered_tag)
	}
}
