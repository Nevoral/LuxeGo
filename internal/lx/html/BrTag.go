package html

import (
	"LuxeGo/internal/lx"
)

//Br - 
func Br() *BrTag {
	return &BrTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "br", Attributes: &lx.Attributes{}, Children: nil}}}
}

type BrTag struct {
	*ComponentTag
}
