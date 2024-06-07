package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Output -
func Output(tags ...LuxeGo.Content) *OutputTag {
	return &OutputTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "output", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type OutputTag struct {
	*ComponentHtmlTag
}

// For -
func (o *OutputTag) For(value string) *OutputTag {
	o.AddAttribute("for", value)
	return o
}

// Form -
func (o *OutputTag) Form(value string) *OutputTag {
	o.AddAttribute("form", value)
	return o
}

// Name -
func (o *OutputTag) Name(value string) *OutputTag {
	o.AddAttribute("name", value)
	return o
}
