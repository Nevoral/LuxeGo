package html

import (
	"LuxeGo/internal/lx"
)

//Abbr - 
func Abbr(tags ...lx.Content) *AbbrTag {
	return &AbbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "abbr", Attributes: &lx.Attributes{}, Children: &tags}}
}

type AbbrTag struct {
	*ComponentHtmlTag
}
