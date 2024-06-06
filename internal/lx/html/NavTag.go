package html

import (
	"LuxeGo/internal/lx"
)

// Nav -
func Nav(tags ...lx.Content) *NavTag {
	return &NavTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "nav", Attributes: &lx.Attributes{}, Children: &tags}}
}

type NavTag struct {
	*ComponentHtmlTag
}
