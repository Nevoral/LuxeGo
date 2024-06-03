package html

import (
	"LuxeGo/internal/lx"
)

//Title - 
func Title(tags ...lx.Content) *TitleTag {
	return &TitleTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "title", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TitleTag struct {
	*ComponentTag
}
