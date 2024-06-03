package html

import (
	"LuxeGo/internal/lx"
)

//Menu - 
func Menu(tags ...lx.Content) *MenuTag {
	return &MenuTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "menu", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type MenuTag struct {
	*ComponentTag
}
