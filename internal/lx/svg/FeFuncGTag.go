package svg

import (
	"LuxeGo/internal/lx"
)

// FeFuncG -
func FeFuncG() *FeFuncGTag {
	return &FeFuncGTag{ComponentSvgTag: &ComponentSvgTag{Name: "fefuncg", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeFuncGTag struct {
	*ComponentSvgTag
}

// Type -
func (f *FeFuncGTag) Type(value string) *FeFuncGTag {
	f.AddAttribute("type", value)
	return f
}

// TableValues -
func (f *FeFuncGTag) TableValues(value string) *FeFuncGTag {
	f.AddAttribute("tablevalues", value)
	return f
}

// Slope -
func (f *FeFuncGTag) Slope(value string) *FeFuncGTag {
	f.AddAttribute("slope", value)
	return f
}

// Intercept -
func (f *FeFuncGTag) Intercept(value string) *FeFuncGTag {
	f.AddAttribute("intercept", value)
	return f
}

// Amplitude -
func (f *FeFuncGTag) Amplitude(value string) *FeFuncGTag {
	f.AddAttribute("amplitude", value)
	return f
}

// Exponent -
func (f *FeFuncGTag) Exponent(value string) *FeFuncGTag {
	f.AddAttribute("exponent", value)
	return f
}

// Offset -
func (f *FeFuncGTag) Offset(value string) *FeFuncGTag {
	f.AddAttribute("offset", value)
	return f
}
