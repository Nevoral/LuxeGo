package html

import (
	"LuxeGo/internal/lx"
)

//S - 
func S(tags ...lx.Content) *STag {
	return &STag{ComponentHtmlTag: &ComponentHtmlTag{Name: "s", Attributes: &lx.Attributes{}, Children: &tags}}
}

type STag struct {
	*ComponentHtmlTag
}
