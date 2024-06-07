package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// FeDisplacementMap -
func FeDisplacementMap() *FeDisplacementMapTag {
	return &FeDisplacementMapTag{ComponentSvgTag: &ComponentSvgTag{Name: "fedisplacementmap", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeDisplacementMapTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeDisplacementMapTag) In(value string) *FeDisplacementMapTag {
	f.AddAttribute("in", value)
	return f
}

// In2 -
func (f *FeDisplacementMapTag) In2(value string) *FeDisplacementMapTag {
	f.AddAttribute("in2", value)
	return f
}

// Scale -
func (f *FeDisplacementMapTag) Scale(value string) *FeDisplacementMapTag {
	f.AddAttribute("scale", value)
	return f
}

// XChannelSelector -
func (f *FeDisplacementMapTag) XChannelSelector(value string) *FeDisplacementMapTag {
	f.AddAttribute("xchannelselector", value)
	return f
}

// YChannelSelector -
func (f *FeDisplacementMapTag) YChannelSelector(value string) *FeDisplacementMapTag {
	f.AddAttribute("ychannelselector", value)
	return f
}
