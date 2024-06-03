package html

import (
	"LuxeGo/internal/lx"
)

//H6 - 
func H6(tags ...lx.Content) *H6Tag {
	return &H6Tag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "h6", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type H6Tag struct {
	*ComponentTag
}
