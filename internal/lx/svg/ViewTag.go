package svg

import (
	"LuxeGo/internal/lx"
)

// View -
func View(tags ...lx.Content) *ViewTag {
	return &ViewTag{ComponentSvgTag: &ComponentSvgTag{Name: "view", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ViewTag struct {
	*ComponentSvgTag
}

// ViewBox -
func (v *ViewTag) ViewBox(value string) *ViewTag {
	v.AddAttribute("viewbox", value)
	return v
}

// PreserveAspectRatio -
func (v *ViewTag) PreserveAspectRatio(value string) *ViewTag {
	v.AddAttribute("preserveaspectratio", value)
	return v
}
