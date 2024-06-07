package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Fieldset -
func Fieldset(tags ...LuxeGo.Content) *FieldsetTag {
	return &FieldsetTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "fieldset", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type FieldsetTag struct {
	*ComponentHtmlTag
}

// Disabled - Indicates that the user cannot interact with the element.
func (f *FieldsetTag) Disabled() *FieldsetTag {
	f.AddAttribute("disabled", "")
	return f
}

// Form -
func (f *FieldsetTag) Form(value string) *FieldsetTag {
	f.AddAttribute("form", value)
	return f
}

// Name -
func (f *FieldsetTag) Name(value string) *FieldsetTag {
	f.AddAttribute("name", value)
	return f
}
