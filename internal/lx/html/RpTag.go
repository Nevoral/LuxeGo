package html

import (
	"LuxeGo/internal/lx"
)

//Rp - 
func Rp(tags ...lx.Content) *RpTag {
	return &RpTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "rp", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type RpTag struct {
	*ComponentTag
}
