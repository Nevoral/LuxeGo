package html

import (
	"LuxeGo/internal/lx"
)

//Cite - 
func Cite(tags ...lx.Content) *CiteTag {
	return &CiteTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "cite", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type CiteTag struct {
	*ComponentTag
}
