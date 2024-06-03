package html

import (
	"LuxeGo/internal/lx"
)

//Thead - 
func Thead(tags ...lx.Content) *TheadTag {
	return &TheadTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "thead", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TheadTag struct {
	*ComponentTag
}
