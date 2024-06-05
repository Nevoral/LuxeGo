package svg

import (
	"LuxeGo/internal/lx"
)

//FeSpecularLighting - 
func FeSpecularLighting() *FeSpecularLightingTag {
	return &FeSpecularLightingTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fespecularlighting", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeSpecularLightingTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeSpecularLightingTag) In(value string) *FeSpecularLightingTag {
	f.AddAttribute("in", value)
	return f
}

//SurfaceScale - 
func (f *FeSpecularLightingTag) SurfaceScale(value string) *FeSpecularLightingTag {
	f.AddAttribute("surfacescale", value)
	return f
}

//SpecularConstant - 
func (f *FeSpecularLightingTag) SpecularConstant(value string) *FeSpecularLightingTag {
	f.AddAttribute("specularconstant", value)
	return f
}

//SpecularExponent - 
func (f *FeSpecularLightingTag) SpecularExponent(value string) *FeSpecularLightingTag {
	f.AddAttribute("specularexponent", value)
	return f
}

//KernelUnitLength - 
func (f *FeSpecularLightingTag) KernelUnitLength(value string) *FeSpecularLightingTag {
	f.AddAttribute("kernelunitlength", value)
	return f
}
