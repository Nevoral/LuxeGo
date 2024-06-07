package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeFlood -
func FeFlood() *FeFloodTag {
	return &FeFloodTag{ComponentSvgTag: &ComponentSvgTag{Name: "feflood", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeFloodTag struct {
	*ComponentSvgTag
}

// FloodColor -
func (f *FeFloodTag) FloodColor(value string) *FeFloodTag {
	f.AddAttribute("flood-color", value)
	return f
}

// FloodOpacity -
func (f *FeFloodTag) FloodOpacity(value string) *FeFloodTag {
	f.AddAttribute("flood-opacity", value)
	return f
}
