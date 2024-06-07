package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// FeMorphology -
func FeMorphology() *FeMorphologyTag {
	return &FeMorphologyTag{ComponentSvgTag: &ComponentSvgTag{Name: "femorphology", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeMorphologyTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeMorphologyTag) In(value string) *FeMorphologyTag {
	f.AddAttribute("in", value)
	return f
}

// Operator -
func (f *FeMorphologyTag) Operator(value string) *FeMorphologyTag {
	f.AddAttribute("operator", value)
	return f
}

// Radius -
func (f *FeMorphologyTag) Radius(value string) *FeMorphologyTag {
	f.AddAttribute("radius", value)
	return f
}
