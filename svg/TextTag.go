package svg

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Text -
func Text(tags ...interface{}) *TextTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &TextTag{ComponentSvgTag: &ComponentSvgTag{Name: "text", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TextTag struct {
	*ComponentSvgTag
}

// X -
func (t *TextTag) X(value string) *TextTag {
	t.AddAttribute("x", value)
	return t
}

// Y -
func (t *TextTag) Y(value string) *TextTag {
	t.AddAttribute("y", value)
	return t
}

// Dx -
func (t *TextTag) Dx(value string) *TextTag {
	t.AddAttribute("dx", value)
	return t
}

// Dy -
func (t *TextTag) Dy(value string) *TextTag {
	t.AddAttribute("dy", value)
	return t
}

// Rotate -
func (t *TextTag) Rotate(value string) *TextTag {
	t.AddAttribute("rotate", value)
	return t
}

// TextLength -
func (t *TextTag) TextLength(value string) *TextTag {
	t.AddAttribute("textlength", value)
	return t
}

// LengthAdjust -
func (t *TextTag) LengthAdjust(value string) *TextTag {
	t.AddAttribute("lengthadjust", value)
	return t
}
