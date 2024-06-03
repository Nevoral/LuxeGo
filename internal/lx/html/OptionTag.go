package html

import (
	"LuxeGo/internal/lx"
)

//Option - 
func Option(tags ...lx.Content) *OptionTag {
	return &OptionTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "option", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type OptionTag struct {
	*ComponentTag
}

//Disabled - Indicates that the user cannot interact with the element.
func (o *OptionTag) Disabled() *OptionTag {
	o.AddAttribute("disabled", "")
	return o
}

//Label - 
func (o *OptionTag) Label(value string) *OptionTag {
	o.AddAttribute("label", value)
	return o
}

//Selected - 
func (o *OptionTag) Selected() *OptionTag {
	o.AddAttribute("selected", "")
	return o
}

//Value - Specifies the value of an <input> element.
func (o *OptionTag) Value(value string) *OptionTag {
	o.AddAttribute("value", value)
	return o
}
