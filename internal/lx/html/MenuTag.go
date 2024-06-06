package html

import (
	"LuxeGo/internal/lx"
)

// Menu -
func Menu(tags ...lx.Content) *MenuTag {
	return &MenuTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "menu", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MenuTag struct {
	*ComponentHtmlTag
}
