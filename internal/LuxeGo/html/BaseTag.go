package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Base -
func Base() *BaseTag {
	return &BaseTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "base", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type BaseTag struct {
	*ComponentHtmlTag
}

// Href - Specifies the URL of the page the link goes to.
func (b *BaseTag) Href(value string) *BaseTag {
	b.AddAttribute("href", value)
	return b
}

// Target -
func (b *BaseTag) Target(value string) *BaseTag {
	b.AddAttribute("target", value)
	return b
}
