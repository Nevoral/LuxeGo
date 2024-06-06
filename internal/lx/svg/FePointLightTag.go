package svg

import (
	"LuxeGo/internal/lx"
)

// FePointLight -
func FePointLight() *FePointLightTag {
	return &FePointLightTag{ComponentSvgTag: &ComponentSvgTag{Name: "fepointlight", Attributes: &lx.Attributes{}, Children: nil}}
}

type FePointLightTag struct {
	*ComponentSvgTag
}

// X -
func (f *FePointLightTag) X(value string) *FePointLightTag {
	f.AddAttribute("x", value)
	return f
}

// Y -
func (f *FePointLightTag) Y(value string) *FePointLightTag {
	f.AddAttribute("y", value)
	return f
}

// Z -
func (f *FePointLightTag) Z(value string) *FePointLightTag {
	f.AddAttribute("z", value)
	return f
}
