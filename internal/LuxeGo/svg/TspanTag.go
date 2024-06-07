package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Tspan -
func Tspan() *TspanTag {
	return &TspanTag{ComponentSvgTag: &ComponentSvgTag{Name: "tspan", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type TspanTag struct {
	*ComponentSvgTag
}

// X -
func (t *TspanTag) X(value string) *TspanTag {
	t.AddAttribute("x", value)
	return t
}

// Y -
func (t *TspanTag) Y(value string) *TspanTag {
	t.AddAttribute("y", value)
	return t
}

// Dx -
func (t *TspanTag) Dx(value string) *TspanTag {
	t.AddAttribute("dx", value)
	return t
}

// Dy -
func (t *TspanTag) Dy(value string) *TspanTag {
	t.AddAttribute("dy", value)
	return t
}

// Rotate -
func (t *TspanTag) Rotate(value string) *TspanTag {
	t.AddAttribute("rotate", value)
	return t
}

// TextLength -
func (t *TspanTag) TextLength(value string) *TspanTag {
	t.AddAttribute("textlength", value)
	return t
}

// LengthAdjust -
func (t *TspanTag) LengthAdjust(value string) *TspanTag {
	t.AddAttribute("lengthadjust", value)
	return t
}
