package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// ForeignObject -
func ForeignObject(tags ...LuxeGo.Content) *ForeignObjectTag {
	return &ForeignObjectTag{ComponentSvgTag: &ComponentSvgTag{Name: "foreignobject", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ForeignObjectTag struct {
	*ComponentSvgTag
}

// X -
func (f *ForeignObjectTag) X(value string) *ForeignObjectTag {
	f.AddAttribute("x", value)
	return f
}

// Y -
func (f *ForeignObjectTag) Y(value string) *ForeignObjectTag {
	f.AddAttribute("y", value)
	return f
}

// Width -
func (f *ForeignObjectTag) Width(value string) *ForeignObjectTag {
	f.AddAttribute("width", value)
	return f
}

// Height -
func (f *ForeignObjectTag) Height(value string) *ForeignObjectTag {
	f.AddAttribute("height", value)
	return f
}