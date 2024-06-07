package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Abbr -
func Abbr(tags ...LuxeGo.Content) *AbbrTag {
	return &AbbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "abbr", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type AbbrTag struct {
	*ComponentHtmlTag
}