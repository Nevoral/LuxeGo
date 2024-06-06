package svg

import (
	"LuxeGo/internal/lx"
)

// FeDropShadow -
func FeDropShadow() *FeDropShadowTag {
	return &FeDropShadowTag{ComponentSvgTag: &ComponentSvgTag{Name: "fedropshadow", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeDropShadowTag struct {
	*ComponentSvgTag
}

// Dx -
func (f *FeDropShadowTag) Dx(value string) *FeDropShadowTag {
	f.AddAttribute("dx", value)
	return f
}

// Dy -
func (f *FeDropShadowTag) Dy(value string) *FeDropShadowTag {
	f.AddAttribute("dy", value)
	return f
}

// StdDeviation -
func (f *FeDropShadowTag) StdDeviation(value string) *FeDropShadowTag {
	f.AddAttribute("stddeviation", value)
	return f
}
