package html

import (
	"LuxeGo/internal/lx"
)

//Var - 
func Var(tags ...lx.Content) *VarTag {
	return &VarTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "var", Attributes: &lx.Attributes{}, Children: &tags}}
}

type VarTag struct {
	*ComponentHtmlTag
}
