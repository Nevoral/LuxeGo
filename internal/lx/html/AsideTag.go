package html

import (
	"LuxeGo/internal/lx"
)

//Aside - 
func Aside(tags ...lx.Content) *AsideTag {
	return &AsideTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "aside", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type AsideTag struct {
	*ComponentTag
}
