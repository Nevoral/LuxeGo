package html

import (
	"LuxeGo/internal/lx"
)

//Select - 
func Select(tags ...lx.Content) *SelectTag {
	return &SelectTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "select", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SelectTag struct {
	*ComponentTag
}

//Autocomplete - Specifies whether a form or an input field should have autocomplete enabled.
func (s *SelectTag) Autocomplete(value string) *SelectTag {
	s.AddAttribute("autocomplete", value)
	return s
}

//Autofocus - Specifies that a form element should automatically get focus when the page loads.
func (s *SelectTag) Autofocus(value string) *SelectTag {
	s.AddAttribute("autofocus", value)
	return s
}

//Disabled - Indicates that the user cannot interact with the element.
func (s *SelectTag) Disabled() *SelectTag {
	s.AddAttribute("disabled", "")
	return s
}

//Form - 
func (s *SelectTag) Form(value string) *SelectTag {
	s.AddAttribute("form", value)
	return s
}

//Multiple - Indicates that multiple values can be entered or selected.
func (s *SelectTag) Multiple() *SelectTag {
	s.AddAttribute("multiple", "")
	return s
}

//Name - 
func (s *SelectTag) Name(value string) *SelectTag {
	s.AddAttribute("name", value)
	return s
}

//Required - Specifies that an input field must be filled out before submitting the form.
func (s *SelectTag) Required() *SelectTag {
	s.AddAttribute("required", "")
	return s
}

//Size - 
func (s *SelectTag) Size(value string) *SelectTag {
	s.AddAttribute("size", value)
	return s
}
