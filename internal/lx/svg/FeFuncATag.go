package svg

import (
	"LuxeGo/internal/lx"
)

//FeFuncA - 
func FeFuncA() *FeFuncATag {
	return &FeFuncATag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fefunca", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeFuncATag struct {
	*ComponentHtmlTag
}

//Type - 
func (f *FeFuncATag) Type(value string) *FeFuncATag {
	f.AddAttribute("type", value)
	return f
}

//TableValues - 
func (f *FeFuncATag) TableValues(value string) *FeFuncATag {
	f.AddAttribute("tablevalues", value)
	return f
}

//Slope - 
func (f *FeFuncATag) Slope(value string) *FeFuncATag {
	f.AddAttribute("slope", value)
	return f
}

//Intercept - 
func (f *FeFuncATag) Intercept(value string) *FeFuncATag {
	f.AddAttribute("intercept", value)
	return f
}

//Amplitude - 
func (f *FeFuncATag) Amplitude(value string) *FeFuncATag {
	f.AddAttribute("amplitude", value)
	return f
}

//Exponent - 
func (f *FeFuncATag) Exponent(value string) *FeFuncATag {
	f.AddAttribute("exponent", value)
	return f
}

//Offset - 
func (f *FeFuncATag) Offset(value string) *FeFuncATag {
	f.AddAttribute("offset", value)
	return f
}
