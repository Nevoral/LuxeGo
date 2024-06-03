package html

import (
	"LuxeGo/internal/lx"
)

//Nav - 
func Nav(tags ...lx.Content) *NavTag {
	return &NavTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "nav", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type NavTag struct {
	*ComponentTag
}
