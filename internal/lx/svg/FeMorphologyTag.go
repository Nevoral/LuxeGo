package svg

import (
	"LuxeGo/internal/lx"
)

//FeMorphology - 
func FeMorphology() *FeMorphologyTag {
	return &FeMorphologyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "femorphology", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeMorphologyTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeMorphologyTag) In(value string) *FeMorphologyTag {
	f.AddAttribute("in", value)
	return f
}

//Operator - 
func (f *FeMorphologyTag) Operator(value string) *FeMorphologyTag {
	f.AddAttribute("operator", value)
	return f
}

//Radius - 
func (f *FeMorphologyTag) Radius(value string) *FeMorphologyTag {
	f.AddAttribute("radius", value)
	return f
}
