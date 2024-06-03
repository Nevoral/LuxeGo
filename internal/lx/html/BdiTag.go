package html

import (
	"LuxeGo/internal/lx"
)

//Bdi - 
func Bdi(tags ...lx.Content) *BdiTag {
	return &BdiTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "bdi", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type BdiTag struct {
	*ComponentTag
}
