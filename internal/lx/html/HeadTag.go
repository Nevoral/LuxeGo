package html

import (
	"LuxeGo/internal/lx"
)

//Head - 
func Head(tags ...lx.Content) *HeadTag {
	return &HeadTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "head", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type HeadTag struct {
	*ComponentTag
}
