package svg

import (
	"LuxeGo/internal/lx"
)

//FeFuncR - 
func FeFuncR() *FeFuncRTag {
	return &FeFuncRTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fefuncr", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeFuncRTag struct {
	*ComponentHtmlTag
}

//Type - 
func (f *FeFuncRTag) Type(value string) *FeFuncRTag {
	f.AddAttribute("type", value)
	return f
}

//TableValues - 
func (f *FeFuncRTag) TableValues(value string) *FeFuncRTag {
	f.AddAttribute("tablevalues", value)
	return f
}

//Slope - 
func (f *FeFuncRTag) Slope(value string) *FeFuncRTag {
	f.AddAttribute("slope", value)
	return f
}

//Intercept - 
func (f *FeFuncRTag) Intercept(value string) *FeFuncRTag {
	f.AddAttribute("intercept", value)
	return f
}

//Amplitude - 
func (f *FeFuncRTag) Amplitude(value string) *FeFuncRTag {
	f.AddAttribute("amplitude", value)
	return f
}

//Exponent - 
func (f *FeFuncRTag) Exponent(value string) *FeFuncRTag {
	f.AddAttribute("exponent", value)
	return f
}

//Offset - 
func (f *FeFuncRTag) Offset(value string) *FeFuncRTag {
	f.AddAttribute("offset", value)
	return f
}
