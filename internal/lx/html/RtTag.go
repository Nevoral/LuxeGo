package html

import (
	"LuxeGo/internal/lx"
)

//Rt - 
func Rt(tags ...lx.Content) *RtTag {
	return &RtTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "rt", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type RtTag struct {
	*ComponentTag
}
