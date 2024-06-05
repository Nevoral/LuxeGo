package html

import (
	"LuxeGo/internal/lx"
)

//Dl - 
func Dl(tags ...lx.Content) *DlTag {
	return &DlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dl", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DlTag struct {
	*ComponentHtmlTag
}
