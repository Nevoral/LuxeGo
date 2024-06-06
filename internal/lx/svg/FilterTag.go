package svg

import (
	"LuxeGo/internal/lx"
)

// Filter -
func Filter(tags ...lx.Content) *FilterTag {
	return &FilterTag{ComponentSvgTag: &ComponentSvgTag{Name: "filter", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FilterTag struct {
	*ComponentSvgTag
}

// X -
func (f *FilterTag) X(value string) *FilterTag {
	f.AddAttribute("x", value)
	return f
}

// Y -
func (f *FilterTag) Y(value string) *FilterTag {
	f.AddAttribute("y", value)
	return f
}

// Width -
func (f *FilterTag) Width(value string) *FilterTag {
	f.AddAttribute("width", value)
	return f
}

// Height -
func (f *FilterTag) Height(value string) *FilterTag {
	f.AddAttribute("height", value)
	return f
}

// FilterUnits -
func (f *FilterTag) FilterUnits(value string) *FilterTag {
	f.AddAttribute("filterunits", value)
	return f
}

// PrimitiveUnits -
func (f *FilterTag) PrimitiveUnits(value string) *FilterTag {
	f.AddAttribute("primitiveunits", value)
	return f
}

// FilterRes -
func (f *FilterTag) FilterRes(value string) *FilterTag {
	f.AddAttribute("filterres", value)
	return f
}
