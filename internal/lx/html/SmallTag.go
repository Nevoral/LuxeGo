package html

import (
	"LuxeGo/internal/lx"
)

//Small - 
func Small(tags ...lx.Content) *SmallTag {
	return &SmallTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "small", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SmallTag struct {
	*ComponentTag
}
