package svg

import (
	"github.com/Nevoral/LuxeGo"
)

// Circle -
func Circle() *CircleTag {
	return &CircleTag{ComponentSvgTag: &ComponentSvgTag{Name: "circle", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type CircleTag struct {
	*ComponentSvgTag
}

// Cx -
func (c *CircleTag) Cx(value string) *CircleTag {
	c.AddAttribute("cx", value)
	return c
}

// Cy -
func (c *CircleTag) Cy(value string) *CircleTag {
	c.AddAttribute("cy", value)
	return c
}

// R -
func (c *CircleTag) R(value string) *CircleTag {
	c.AddAttribute("r", value)
	return c
}

// PathLength -
func (c *CircleTag) PathLength(value string) *CircleTag {
	c.AddAttribute("pathlength", value)
	return c
}
