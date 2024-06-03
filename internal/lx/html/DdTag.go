package html

import (
	"LuxeGo/internal/lx"
)

//Dd - 
func Dd(tags ...lx.Content) *DdTag {
	return &DdTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "dd", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DdTag struct {
	*ComponentTag
}
