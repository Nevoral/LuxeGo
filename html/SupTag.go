package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Sup -
func Sup(tags ...LuxeGo.Content) *SupTag {
	return &SupTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "sup", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type SupTag struct {
	*ComponentHtmlTag
}
