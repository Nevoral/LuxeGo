package html

import (
	"LuxeGo/internal/lx"
)

//Dl - 
func Dl(tags ...lx.Content) *DlTag {
	return &DlTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "dl", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DlTag struct {
	*ComponentTag
}
