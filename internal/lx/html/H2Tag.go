package html

import (
	"LuxeGo/internal/lx"
)

//H2 - 
func H2(tags ...lx.Content) *H2Tag {
	return &H2Tag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "h2", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type H2Tag struct {
	*ComponentTag
}
