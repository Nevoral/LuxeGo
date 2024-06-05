package svg

import (
	"LuxeGo/internal/lx"
)

//FeImage - 
func FeImage() *FeImageTag {
	return &FeImageTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "feimage", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeImageTag struct {
	*ComponentHtmlTag
}

//Href - 
func (f *FeImageTag) Href(value string) *FeImageTag {
	f.AddAttribute("href", value)
	return f
}

//PreserveAspectRatio - 
func (f *FeImageTag) PreserveAspectRatio(value string) *FeImageTag {
	f.AddAttribute("preserveaspectratio", value)
	return f
}
