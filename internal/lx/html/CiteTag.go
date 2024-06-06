package html

import (
	"LuxeGo/internal/lx"
)

// Cite -
func Cite(tags ...lx.Content) *CiteTag {
	return &CiteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "cite", Attributes: &lx.Attributes{}, Children: &tags}}
}

type CiteTag struct {
	*ComponentHtmlTag
}
