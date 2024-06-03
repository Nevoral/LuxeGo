package html

import (
	"LuxeGo/internal/lx"
)

//Ul - 
func Ul(tags ...lx.Content) *UlTag {
	return &UlTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "ul", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type UlTag struct {
	*ComponentTag
}
