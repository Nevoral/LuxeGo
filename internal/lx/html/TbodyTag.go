package html

import (
	"LuxeGo/internal/lx"
)

//Tbody - 
func Tbody(tags ...lx.Content) *TbodyTag {
	return &TbodyTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "tbody", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TbodyTag struct {
	*ComponentTag
}
