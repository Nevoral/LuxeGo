package html

import (
	"LuxeGo/internal/lx"
)

//B - 
func B(tags ...lx.Content) *BTag {
	return &BTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "b", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type BTag struct {
	*ComponentTag
}
