package html

import (
	"LuxeGo/internal/lx"
)

//Wbr - 
func Wbr() *WbrTag {
	return &WbrTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "wbr", Attributes: &lx.Attributes{}, Children: nil}}}
}

type WbrTag struct {
	*ComponentTag
}
