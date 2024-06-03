package html

import (
	"LuxeGo/internal/lx"
)

//I - 
func I(tags ...lx.Content) *ITag {
	return &ITag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "i", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type ITag struct {
	*ComponentTag
}
