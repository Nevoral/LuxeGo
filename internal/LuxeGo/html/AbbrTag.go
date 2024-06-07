package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Abbr -
func Abbr(tags ...LuxeGo.Content) *AbbrTag {
	return &AbbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "abbr", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type AbbrTag struct {
	*ComponentHtmlTag
}
