package html

import (
	"LuxeGo/internal/lx"
)

// U -
func U(tags ...lx.Content) *UTag {
	return &UTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "u", Attributes: &lx.Attributes{}, Children: &tags}}
}

type UTag struct {
	*ComponentHtmlTag
}
