package html

import (
	"LuxeGo/internal/lx"
)

//Tr - 
func Tr(tags ...lx.Content) *TrTag {
	return &TrTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "tr", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TrTag struct {
	*ComponentTag
}
