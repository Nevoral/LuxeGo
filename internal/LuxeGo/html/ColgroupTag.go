package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Colgroup -
func Colgroup(tags ...LuxeGo.Content) *ColgroupTag {
	return &ColgroupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "colgroup", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ColgroupTag struct {
	*ComponentHtmlTag
}

// Span -
func (c *ColgroupTag) Span(value string) *ColgroupTag {
	c.AddAttribute("span", value)
	return c
}
