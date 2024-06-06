package html

import (
	"LuxeGo/internal/lx"
)

// Colgroup -
func Colgroup(tags ...lx.Content) *ColgroupTag {
	return &ColgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "colgroup", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ColgroupTag struct {
	*ComponentHtmlTag
}

// Span -
func (c *ColgroupTag) Span(value string) *ColgroupTag {
	c.AddAttribute("span", value)
	return c
}
