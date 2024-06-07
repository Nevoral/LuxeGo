package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Col -
func Col() *ColTag {
	return &ColTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "col", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type ColTag struct {
	*ComponentHtmlTag
}

// Span -
func (c *ColTag) Span(value string) *ColTag {
	c.AddAttribute("span", value)
	return c
}
