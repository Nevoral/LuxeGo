package html

import (
	"LuxeGo/internal/lx"
)

//Var - 
func Var(tags ...lx.Content) *VarTag {
	return &VarTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "var", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type VarTag struct {
	*ComponentTag
}
