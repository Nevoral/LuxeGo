package html

import (
	"LuxeGo/internal/lx"
)

//Base - 
func Base() *BaseTag {
	return &BaseTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "base", Attributes: &lx.Attributes{}, Children: nil}}}
}

type BaseTag struct {
	*ComponentTag
}

//Href - Specifies the URL of the page the link goes to.
func (b *BaseTag) Href(value string) *BaseTag {
	b.AddAttribute("href", value)
	return b
}

//Target - 
func (b *BaseTag) Target(value string) *BaseTag {
	b.AddAttribute("target", value)
	return b
}
