package html

import (
	"LuxeGo/internal/lx"
)

//Search - 
func Search(tags ...lx.Content) *SearchTag {
	return &SearchTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "search", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SearchTag struct {
	*ComponentHtmlTag
}
