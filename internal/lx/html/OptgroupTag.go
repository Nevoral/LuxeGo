package html

import (
	"LuxeGo/internal/lx"
)

//Optgroup - 
func Optgroup(tags ...lx.Content) *OptgroupTag {
	return &OptgroupTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "optgroup", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type OptgroupTag struct {
	*ComponentTag
}

//Disabled - Indicates that the user cannot interact with the element.
func (o *OptgroupTag) Disabled() *OptgroupTag {
	o.AddAttribute("disabled", "")
	return o
}

//Label - 
func (o *OptgroupTag) Label(value string) *OptgroupTag {
	o.AddAttribute("label", value)
	return o
}
