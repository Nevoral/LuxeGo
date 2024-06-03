package html

import (
	"LuxeGo/internal/lx"
)

//Canvas - 
func Canvas(tags ...lx.Content) *CanvasTag {
	return &CanvasTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "canvas", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type CanvasTag struct {
	*ComponentTag
}

//Height - 
func (c *CanvasTag) Height(value string) *CanvasTag {
	c.AddAttribute("height", value)
	return c
}

//Width - 
func (c *CanvasTag) Width(value string) *CanvasTag {
	c.AddAttribute("width", value)
	return c
}
