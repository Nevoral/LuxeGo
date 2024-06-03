package html

import (
	"LuxeGo/internal/lx"
)

//Div - 
func Div(tags ...lx.Content) *DivTag {
	return &DivTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "div", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DivTag struct {
	*ComponentTag
}
