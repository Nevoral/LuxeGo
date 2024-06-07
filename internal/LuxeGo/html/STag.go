package html

import (
	"LuxeGo/internal/LuxeGo"
)

// S -
func S(tags ...LuxeGo.Content) *STag {
	return &STag{ComponentHtmlTag: &ComponentHtmlTag{Name: "s", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type STag struct {
	*ComponentHtmlTag
}
