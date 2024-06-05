package svg

import (
	"LuxeGo/internal/lx"
)

//FeOffset - 
func FeOffset() *FeOffsetTag {
	return &FeOffsetTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "feoffset", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeOffsetTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeOffsetTag) In(value string) *FeOffsetTag {
	f.AddAttribute("in", value)
	return f
}

//Dx - 
func (f *FeOffsetTag) Dx(value string) *FeOffsetTag {
	f.AddAttribute("dx", value)
	return f
}

//Dy - 
func (f *FeOffsetTag) Dy(value string) *FeOffsetTag {
	f.AddAttribute("dy", value)
	return f
}
