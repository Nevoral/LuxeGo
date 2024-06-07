package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Search -
func Search(tags ...LuxeGo.Content) *SearchTag {
	return &SearchTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "search", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SearchTag struct {
	*ComponentHtmlTag
}
