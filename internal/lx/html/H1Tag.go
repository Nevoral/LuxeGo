package html

import (
	"LuxeGo/internal/lx"
)

//H1 - 
func H1(tags ...lx.Content) *H1Tag {
	return &H1Tag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "h1", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type H1Tag struct {
	*ComponentTag
}
