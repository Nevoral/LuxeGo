package html

import (
	"LuxeGo/internal/lx"
)

//Bdo - 
func Bdo(tags ...lx.Content) *BdoTag {
	return &BdoTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "bdo", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type BdoTag struct {
	*ComponentTag
}
