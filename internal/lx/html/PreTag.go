package html

import (
	"LuxeGo/internal/lx"
)

//Pre - 
func Pre(tags ...lx.Content) *PreTag {
	return &PreTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "pre", Attributes: &lx.Attributes{}, Children: &tags}}
}

type PreTag struct {
	*ComponentHtmlTag
}
