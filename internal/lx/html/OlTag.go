package html

import (
	"LuxeGo/internal/lx"
)

// Ol -
func Ol(tags ...lx.Content) *OlTag {
	return &OlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ol", Attributes: &lx.Attributes{}, Children: &tags}}
}

type OlTag struct {
	*ComponentHtmlTag
}

// Reversed -
func (o *OlTag) Reversed() *OlTag {
	o.AddAttribute("reversed", "")
	return o
}

// Start -
func (o *OlTag) Start(value string) *OlTag {
	o.AddAttribute("start", value)
	return o
}

// Type - Specifies the type of an <input> element.
func (o *OlTag) Type(value string) *OlTag {
	o.AddAttribute("type", value)
	return o
}
