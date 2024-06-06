package svg

import (
	"LuxeGo/internal/lx"
)

// FeDiffuseLighting -
func FeDiffuseLighting() *FeDiffuseLightingTag {
	return &FeDiffuseLightingTag{ComponentSvgTag: &ComponentSvgTag{Name: "fediffuselighting", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeDiffuseLightingTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeDiffuseLightingTag) In(value string) *FeDiffuseLightingTag {
	f.AddAttribute("in", value)
	return f
}

// SurfaceScale -
func (f *FeDiffuseLightingTag) SurfaceScale(value string) *FeDiffuseLightingTag {
	f.AddAttribute("surfacescale", value)
	return f
}

// DiffuseConstant -
func (f *FeDiffuseLightingTag) DiffuseConstant(value string) *FeDiffuseLightingTag {
	f.AddAttribute("diffuseconstant", value)
	return f
}

// KernelUnitLength -
func (f *FeDiffuseLightingTag) KernelUnitLength(value string) *FeDiffuseLightingTag {
	f.AddAttribute("kernelunitlength", value)
	return f
}
