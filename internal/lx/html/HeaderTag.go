package html

import (
	"LuxeGo/internal/lx"
)

//Header - 
func Header(tags ...lx.Content) *HeaderTag {
	return &HeaderTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "header", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type HeaderTag struct {
	*ComponentTag
}
