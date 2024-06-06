package html

import (
	"LuxeGo/internal/lx"
)

// Div -
func Div(tags ...lx.Content) *DivTag {
	return &DivTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "div", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DivTag struct {
	*ComponentHtmlTag
}
