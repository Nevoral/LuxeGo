package html

import (
	"LuxeGo/internal/lx"
)

//U - 
func U(tags ...lx.Content) *UTag {
	return &UTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "u", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type UTag struct {
	*ComponentTag
}
