package svg

import (
	"LuxeGo/internal/lx"
)

//FeBlend - 
func FeBlend() *FeBlendTag {
	return &FeBlendTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "feblend", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeBlendTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeBlendTag) In(value string) *FeBlendTag {
	f.AddAttribute("in", value)
	return f
}

//In2 - 
func (f *FeBlendTag) In2(value string) *FeBlendTag {
	f.AddAttribute("in2", value)
	return f
}

//Mode - 
func (f *FeBlendTag) Mode(value string) *FeBlendTag {
	f.AddAttribute("mode", value)
	return f
}
