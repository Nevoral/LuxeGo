package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Object -
func Object(tags ...LuxeGo.Content) *ObjectTag {
	return &ObjectTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "object", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ObjectTag struct {
	*ComponentHtmlTag
}

// Data -
func (o *ObjectTag) Data(value string) *ObjectTag {
	o.AddAttribute("data", value)
	return o
}

// Form -
func (o *ObjectTag) Form(value string) *ObjectTag {
	o.AddAttribute("form", value)
	return o
}

// Height -
func (o *ObjectTag) Height(value string) *ObjectTag {
	o.AddAttribute("height", value)
	return o
}

// Name -
func (o *ObjectTag) Name(value string) *ObjectTag {
	o.AddAttribute("name", value)
	return o
}

// Type - Specifies the type of an <input> element.
func (o *ObjectTag) Type(value string) *ObjectTag {
	o.AddAttribute("type", value)
	return o
}

// Usemap -
func (o *ObjectTag) Usemap(value string) *ObjectTag {
	o.AddAttribute("usemap", value)
	return o
}

// Width -
func (o *ObjectTag) Width(value string) *ObjectTag {
	o.AddAttribute("width", value)
	return o
}
