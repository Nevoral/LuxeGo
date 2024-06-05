package svg

import (
	"LuxeGo/internal/lx"
)

//ClipPath - 
func ClipPath(tags ...lx.Content) *ClipPathTag {
	return &ClipPathTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "clippath", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ClipPathTag struct {
	*ComponentHtmlTag
}

//ClipPathUnits - 
func (c *ClipPathTag) ClipPathUnits(value string) *ClipPathTag {
	c.AddAttribute("clippathunits", value)
	return c
}

//Transform - 
func (c *ClipPathTag) Transform(value string) *ClipPathTag {
	c.AddAttribute("transform", value)
	return c
}
