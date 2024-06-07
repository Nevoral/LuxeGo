package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Optgroup -
func Optgroup(tags ...LuxeGo.Content) *OptgroupTag {
	return &OptgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "optgroup", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type OptgroupTag struct {
	*ComponentHtmlTag
}

// Disabled - Indicates that the user cannot interact with the element.
func (o *OptgroupTag) Disabled() *OptgroupTag {
	o.AddAttribute("disabled", "")
	return o
}

// Label -
func (o *OptgroupTag) Label(value string) *OptgroupTag {
	o.AddAttribute("label", value)
	return o
}
