package html

import (
	"LuxeGo/internal/lx"
)

//Search - 
func Search(tags ...lx.Content) *SearchTag {
	return &SearchTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "search", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SearchTag struct {
	*ComponentTag
}
