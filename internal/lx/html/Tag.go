package html

import (
	"LuxeGo/internal/lx"
)

// - 
func (tags ...lx.Content) *Tag {
	return &Tag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type Tag struct {
	*ComponentTag
}
