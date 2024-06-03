package html

import (
	"LuxeGo/internal/lx"
)

//Tfoot - 
func Tfoot(tags ...lx.Content) *TfootTag {
	return &TfootTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "tfoot", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TfootTag struct {
	*ComponentTag
}
