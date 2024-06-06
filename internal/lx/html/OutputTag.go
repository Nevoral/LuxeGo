package html

import (
	"LuxeGo/internal/lx"
)

// Output -
func Output(tags ...lx.Content) *OutputTag {
	return &OutputTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "output", Attributes: &lx.Attributes{}, Children: &tags}}
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
