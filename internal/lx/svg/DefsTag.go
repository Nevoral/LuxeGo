package svg

import (
	"LuxeGo/internal/lx"
)

//Defs - 
func Defs(tags ...lx.Content) *DefsTag {
	return &DefsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "defs", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DefsTag struct {
	*ComponentHtmlTag
}
