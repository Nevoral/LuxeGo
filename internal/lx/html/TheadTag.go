package html

import (
	"LuxeGo/internal/lx"
)

// Thead -
func Thead(tags ...lx.Content) *TheadTag {
	return &TheadTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "thead", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TheadTag struct {
	*ComponentHtmlTag
}
