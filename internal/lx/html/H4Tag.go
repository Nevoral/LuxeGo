package html

import (
	"LuxeGo/internal/lx"
)

//H4 - 
func H4(tags ...lx.Content) *H4Tag {
	return &H4Tag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "h4", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type H4Tag struct {
	*ComponentTag
}
