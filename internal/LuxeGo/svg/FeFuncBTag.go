package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeFuncB -
func FeFuncB() *FeFuncBTag {
	return &FeFuncBTag{ComponentSvgTag: &ComponentSvgTag{Name: "fefuncb", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeFuncBTag struct {
	*ComponentSvgTag
}

// Type -
func (f *FeFuncBTag) Type(value string) *FeFuncBTag {
	f.AddAttribute("type", value)
	return f
}

// TableValues -
func (f *FeFuncBTag) TableValues(value string) *FeFuncBTag {
	f.AddAttribute("tablevalues", value)
	return f
}

// Slope -
func (f *FeFuncBTag) Slope(value string) *FeFuncBTag {
	f.AddAttribute("slope", value)
	return f
}

// Intercept -
func (f *FeFuncBTag) Intercept(value string) *FeFuncBTag {
	f.AddAttribute("intercept", value)
	return f
}

// Amplitude -
func (f *FeFuncBTag) Amplitude(value string) *FeFuncBTag {
	f.AddAttribute("amplitude", value)
	return f
}

// Exponent -
func (f *FeFuncBTag) Exponent(value string) *FeFuncBTag {
	f.AddAttribute("exponent", value)
	return f
}

// Offset -
func (f *FeFuncBTag) Offset(value string) *FeFuncBTag {
	f.AddAttribute("offset", value)
	return f
}
