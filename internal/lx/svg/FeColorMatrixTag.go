package svg

import (
	"LuxeGo/internal/lx"
)

// FeColorMatrix -
func FeColorMatrix() *FeColorMatrixTag {
	return &FeColorMatrixTag{ComponentSvgTag: &ComponentSvgTag{Name: "fecolormatrix", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeColorMatrixTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeColorMatrixTag) In(value string) *FeColorMatrixTag {
	f.AddAttribute("in", value)
	return f
}

// Type -
func (f *FeColorMatrixTag) Type(value string) *FeColorMatrixTag {
	f.AddAttribute("type", value)
	return f
}

// Values -
func (f *FeColorMatrixTag) Values(value string) *FeColorMatrixTag {
	f.AddAttribute("values", value)
	return f
}
