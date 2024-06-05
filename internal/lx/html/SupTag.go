package html

import (
	"LuxeGo/internal/lx"
)

//Sup - 
func Sup(tags ...lx.Content) *SupTag {
	return &SupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sup", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SupTag struct {
	*ComponentHtmlTag
}
