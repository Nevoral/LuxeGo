package svg

import (
	"LuxeGo/internal/lx"
)

//Text - 
func Text(tags ...lx.Content) *TextTag {
	return &TextTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "text", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TextTag struct {
	*ComponentHtmlTag
}

//X - 
func (t *TextTag) X(value string) *TextTag {
	t.AddAttribute("x", value)
	return t
}

//Y - 
func (t *TextTag) Y(value string) *TextTag {
	t.AddAttribute("y", value)
	return t
}

//Dx - 
func (t *TextTag) Dx(value string) *TextTag {
	t.AddAttribute("dx", value)
	return t
}

//Dy - 
func (t *TextTag) Dy(value string) *TextTag {
	t.AddAttribute("dy", value)
	return t
}

//Rotate - 
func (t *TextTag) Rotate(value string) *TextTag {
	t.AddAttribute("rotate", value)
	return t
}

//TextLength - 
func (t *TextTag) TextLength(value string) *TextTag {
	t.AddAttribute("textlength", value)
	return t
}

//LengthAdjust - 
func (t *TextTag) LengthAdjust(value string) *TextTag {
	t.AddAttribute("lengthadjust", value)
	return t
}