package svg

import (
	"LuxeGo/internal/lx"
)

// FeImage -
func FeImage() *FeImageTag {
	return &FeImageTag{ComponentSvgTag: &ComponentSvgTag{Name: "feimage", Attributes: &lx.Attributes{}, Children: nil}}
}

type FeImageTag struct {
	*ComponentSvgTag
}

// Href -
func (f *FeImageTag) Href(value string) *FeImageTag {
	f.AddAttribute("href", value)
	return f
}

// PreserveAspectRatio -
func (f *FeImageTag) PreserveAspectRatio(value string) *FeImageTag {
	f.AddAttribute("preserveaspectratio", value)
	return f
}
