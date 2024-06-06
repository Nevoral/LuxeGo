package html

import (
	"LuxeGo/internal/lx"
)

// Ul -
func Ul(tags ...lx.Content) *UlTag {
	return &UlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ul", Attributes: &lx.Attributes{}, Children: &tags}}
}

type UlTag struct {
	*ComponentHtmlTag
}
