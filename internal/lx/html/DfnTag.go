package html

import (
	"LuxeGo/internal/lx"
)

//Dfn - 
func Dfn(tags ...lx.Content) *DfnTag {
	return &DfnTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "dfn", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DfnTag struct {
	*ComponentTag
}
