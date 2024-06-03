package html

import (
	"LuxeGo/internal/lx"
)

//Col - 
func Col() *ColTag {
	return &ColTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "col", Attributes: &lx.Attributes{}, Children: nil}}}
}

type ColTag struct {
	*ComponentTag
}

//Span - 
func (c *ColTag) Span(value string) *ColTag {
	c.AddAttribute("span", value)
	return c
}
