package html

import (
	"LuxeGo/internal/lx"
)

// Header -
func Header(tags ...lx.Content) *HeaderTag {
	return &HeaderTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "header", Attributes: &lx.Attributes{}, Children: &tags}}
}

type HeaderTag struct {
	*ComponentHtmlTag
}
