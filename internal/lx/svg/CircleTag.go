package svg

import (
	"LuxeGo/internal/lx"
)

//Circle - 
func Circle() *CircleTag {
	return &CircleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "circle", Attributes: &lx.Attributes{}, Children: nil}}
}

type CircleTag struct {
	*ComponentHtmlTag
}

//Cx - 
func (c *CircleTag) Cx(value string) *CircleTag {
	c.AddAttribute("cx", value)
	return c
}

//Cy - 
func (c *CircleTag) Cy(value string) *CircleTag {
	c.AddAttribute("cy", value)
	return c
}

//R - 
func (c *CircleTag) R(value string) *CircleTag {
	c.AddAttribute("r", value)
	return c
}

//PathLength - 
func (c *CircleTag) PathLength(value string) *CircleTag {
	c.AddAttribute("pathlength", value)
	return c
}
