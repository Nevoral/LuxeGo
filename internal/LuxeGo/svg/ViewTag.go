package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// View -
func View(tags ...LuxeGo.Content) *ViewTag {
	return &ViewTag{ComponentSvgTag: &ComponentSvgTag{Name: "view", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
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
