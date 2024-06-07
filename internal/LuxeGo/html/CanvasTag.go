package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Canvas -
func Canvas(tags ...LuxeGo.Content) *CanvasTag {
	return &CanvasTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "canvas", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type CanvasTag struct {
	*ComponentHtmlTag
}

// Height -
func (c *CanvasTag) Height(value string) *CanvasTag {
	c.AddAttribute("height", value)
	return c
}

// Width -
func (c *CanvasTag) Width(value string) *CanvasTag {
	c.AddAttribute("width", value)
	return c
}
