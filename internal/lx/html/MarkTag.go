package html

import (
	"LuxeGo/internal/lx"
)

//Mark - 
func Mark(tags ...lx.Content) *MarkTag {
	return &MarkTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "mark", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type MarkTag struct {
	*ComponentTag
}
