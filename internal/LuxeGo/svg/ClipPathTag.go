package svg

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// ClipPath -
func ClipPath(tags ...LuxeGo.Content) *ClipPathTag {
	return &ClipPathTag{ComponentSvgTag: &ComponentSvgTag{Name: "clippath", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type ClipPathTag struct {
	*ComponentSvgTag
}

// ClipPathUnits -
func (c *ClipPathTag) ClipPathUnits(value string) *ClipPathTag {
	c.AddAttribute("clippathunits", value)
	return c
}

// Transform -
func (c *ClipPathTag) Transform(value string) *ClipPathTag {
	c.AddAttribute("transform", value)
	return c
}
