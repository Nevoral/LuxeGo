package svg

import (
	"LuxeGo/internal/lx"
)

//FeGaussianBlur - 
func FeGaussianBlur() *FeGaussianBlurTag {
	return &FeGaussianBlurTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fegaussianblur", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeGaussianBlurTag struct {
	*ComponentHtmlTag
}

//In - 
func (f *FeGaussianBlurTag) In(value string) *FeGaussianBlurTag {
	f.AddAttribute("in", value)
	return f
}

//StdDeviation - 
func (f *FeGaussianBlurTag) StdDeviation(value string) *FeGaussianBlurTag {
	f.AddAttribute("stddeviation", value)
	return f
}

//EdgeMode - 
func (f *FeGaussianBlurTag) EdgeMode(value string) *FeGaussianBlurTag {
	f.AddAttribute("edgemode", value)
	return f
}