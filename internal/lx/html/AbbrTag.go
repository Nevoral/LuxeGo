package html

import (
	"LuxeGo/internal/lx"
)

//Abbr - 
func Abbr(tags ...lx.Content) *AbbrTag {
	return &AbbrTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "abbr", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type AbbrTag struct {
	*ComponentTag
}
