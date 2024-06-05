package html

import (
	"LuxeGo/internal/lx"
)

//Dd - 
func Dd(tags ...lx.Content) *DdTag {
	return &DdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dd", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DdTag struct {
	*ComponentHtmlTag
}
