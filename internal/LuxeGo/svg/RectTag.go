package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Rect -
func Rect() *RectTag {
	return &RectTag{ComponentSvgTag: &ComponentSvgTag{Name: "rect", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type RectTag struct {
	*ComponentSvgTag
}

// X -
func (r *RectTag) X(value string) *RectTag {
	r.AddAttribute("x", value)
	return r
}

// Y -
func (r *RectTag) Y(value string) *RectTag {
	r.AddAttribute("y", value)
	return r
}

// Width -
func (r *RectTag) Width(value string) *RectTag {
	r.AddAttribute("width", value)
	return r
}

// Height -
func (r *RectTag) Height(value string) *RectTag {
	r.AddAttribute("height", value)
	return r
}

// Rx -
func (r *RectTag) Rx(value string) *RectTag {
	r.AddAttribute("rx", value)
	return r
}

// Ry -
func (r *RectTag) Ry(value string) *RectTag {
	r.AddAttribute("ry", value)
	return r
}

// PathLength -
func (r *RectTag) PathLength(value string) *RectTag {
	r.AddAttribute("pathlength", value)
	return r
}
