package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// FeComposite -
func FeComposite() *FeCompositeTag {
	return &FeCompositeTag{ComponentSvgTag: &ComponentSvgTag{Name: "fecomposite", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeCompositeTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeCompositeTag) In(value string) *FeCompositeTag {
	f.AddAttribute("in", value)
	return f
}

// In2 -
func (f *FeCompositeTag) In2(value string) *FeCompositeTag {
	f.AddAttribute("in2", value)
	return f
}

// Operator -
func (f *FeCompositeTag) Operator(value string) *FeCompositeTag {
	f.AddAttribute("operator", value)
	return f
}

// K1 -
func (f *FeCompositeTag) K1(value string) *FeCompositeTag {
	f.AddAttribute("k1", value)
	return f
}

// K2 -
func (f *FeCompositeTag) K2(value string) *FeCompositeTag {
	f.AddAttribute("k2", value)
	return f
}

// K3 -
func (f *FeCompositeTag) K3(value string) *FeCompositeTag {
	f.AddAttribute("k3", value)
	return f
}

// K4 -
func (f *FeCompositeTag) K4(value string) *FeCompositeTag {
	f.AddAttribute("k4", value)
	return f
}
