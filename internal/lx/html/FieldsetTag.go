package html

import (
	"LuxeGo/internal/lx"
)

//Fieldset - 
func Fieldset(tags ...lx.Content) *FieldsetTag {
	return &FieldsetTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "fieldset", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type FieldsetTag struct {
	*ComponentTag
}

//Disabled - Indicates that the user cannot interact with the element.
func (f *FieldsetTag) Disabled() *FieldsetTag {
	f.AddAttribute("disabled", "")
	return f
}

//Form - 
func (f *FieldsetTag) Form(value string) *FieldsetTag {
	f.AddAttribute("form", value)
	return f
}

//Name - 
func (f *FieldsetTag) Name(value string) *FieldsetTag {
	f.AddAttribute("name", value)
	return f
}
