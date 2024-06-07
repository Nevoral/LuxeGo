package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// FeGaussianBlur -
func FeGaussianBlur() *FeGaussianBlurTag {
	return &FeGaussianBlurTag{ComponentSvgTag: &ComponentSvgTag{Name: "fegaussianblur", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type FeGaussianBlurTag struct {
	*ComponentSvgTag
}

// In -
func (f *FeGaussianBlurTag) In(value string) *FeGaussianBlurTag {
	f.AddAttribute("in", value)
	return f
}

// StdDeviation -
func (f *FeGaussianBlurTag) StdDeviation(value string) *FeGaussianBlurTag {
	f.AddAttribute("stddeviation", value)
	return f
}

// EdgeMode -
func (f *FeGaussianBlurTag) EdgeMode(value string) *FeGaussianBlurTag {
	f.AddAttribute("edgemode", value)
	return f
}
