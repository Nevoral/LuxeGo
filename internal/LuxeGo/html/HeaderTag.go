package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Header -
func Header(tags ...LuxeGo.Content) *HeaderTag {
	return &HeaderTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "header", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type HeaderTag struct {
	*ComponentHtmlTag
}
