package html

import (
	"LuxeGo/internal/lx"
)

//Pre - 
func Pre(tags ...lx.Content) *PreTag {
	return &PreTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "pre", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type PreTag struct {
	*ComponentTag
}
