package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// FeOffset -
func FeOffset() *FeOffsetTag {
	return &FeOffsetTag{ComponentSvgTag: &ComponentSvgTag{Name: "feoffset", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeOffsetTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeOffsetTag) In(value string) *FeOffsetTag {
	f.AddAttribute("in", value)
	return f
}

// Dx -
func (f *FeOffsetTag) Dx(value string) *FeOffsetTag {
	f.AddAttribute("dx", value)
	return f
}

// Dy -
func (f *FeOffsetTag) Dy(value string) *FeOffsetTag {
	f.AddAttribute("dy", value)
	return f
}
