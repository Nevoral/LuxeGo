package html

import (
	"LuxeGo/internal/lx"
)

//P - 
func P(tags ...lx.Content) *PTag {
	return &PTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "p", Attributes: &lx.Attributes{}, Children: &tags}}
}

type PTag struct {
	*ComponentHtmlTag
}
