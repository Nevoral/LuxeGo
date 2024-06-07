package svg

import (
	"LuxeGo/internal/LuxeGo"
)

// FeBlend -
func FeBlend() *FeBlendTag {
	return &FeBlendTag{ComponentSvgTag: &ComponentSvgTag{Name: "feblend", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeBlendTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeBlendTag) In(value string) *FeBlendTag {
	f.AddAttribute("in", value)
	return f
}

// In2 -
func (f *FeBlendTag) In2(value string) *FeBlendTag {
	f.AddAttribute("in2", value)
	return f
}

// Mode -
func (f *FeBlendTag) Mode(value string) *FeBlendTag {
	f.AddAttribute("mode", value)
	return f
}
