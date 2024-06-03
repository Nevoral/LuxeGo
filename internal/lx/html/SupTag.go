package html

import (
	"LuxeGo/internal/lx"
)

//Sup - 
func Sup(tags ...lx.Content) *SupTag {
	return &SupTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "sup", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SupTag struct {
	*ComponentTag
}
