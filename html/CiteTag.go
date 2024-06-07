package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Cite -
func Cite(tags ...LuxeGo.Content) *CiteTag {
	return &CiteTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "cite", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type CiteTag struct {
	*ComponentHtmlTag
}
