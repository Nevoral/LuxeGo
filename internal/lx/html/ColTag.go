package html

import (
	"LuxeGo/internal/lx"
)

// Col -
func Col() *ColTag {
	return &ColTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "col", Attributes: &lx.Attributes{}, Children: nil}}
}

type ColTag struct {
	*ComponentHtmlTag
}

// Span -
func (c *ColTag) Span(value string) *ColTag {
	c.AddAttribute("span", value)
	return c
}
