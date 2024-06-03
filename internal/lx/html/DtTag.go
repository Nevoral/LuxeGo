package html

import (
	"LuxeGo/internal/lx"
)

//Dt - 
func Dt(tags ...lx.Content) *DtTag {
	return &DtTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "dt", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DtTag struct {
	*ComponentTag
}
