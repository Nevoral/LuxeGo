package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeSpecularLighting -
func FeSpecularLighting() *FeSpecularLightingTag {
	return &FeSpecularLightingTag{ComponentSvgTag: &ComponentSvgTag{Name: "fespecularlighting", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeSpecularLightingTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeSpecularLightingTag) In(value string) *FeSpecularLightingTag {
	f.AddAttribute("in", value)
	return f
}

// SurfaceScale -
func (f *FeSpecularLightingTag) SurfaceScale(value string) *FeSpecularLightingTag {
	f.AddAttribute("surfacescale", value)
	return f
}

// SpecularConstant -
func (f *FeSpecularLightingTag) SpecularConstant(value string) *FeSpecularLightingTag {
	f.AddAttribute("specularconstant", value)
	return f
}

// SpecularExponent -
func (f *FeSpecularLightingTag) SpecularExponent(value string) *FeSpecularLightingTag {
	f.AddAttribute("specularexponent", value)
	return f
}

// KernelUnitLength -
func (f *FeSpecularLightingTag) KernelUnitLength(value string) *FeSpecularLightingTag {
	f.AddAttribute("kernelunitlength", value)
	return f
}
