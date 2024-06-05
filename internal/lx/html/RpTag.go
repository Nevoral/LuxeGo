package html

import (
	"LuxeGo/internal/lx"
)

//Rp - 
func Rp(tags ...lx.Content) *RpTag {
	return &RpTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rp", Attributes: &lx.Attributes{}, Children: &tags}}
}

type RpTag struct {
	*ComponentHtmlTag
}
