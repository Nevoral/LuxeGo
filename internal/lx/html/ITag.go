package html

import (
	"LuxeGo/internal/lx"
)

// I -
func I(tags ...lx.Content) *ITag {
	return &ITag{ComponentHtmlTag: &ComponentHtmlTag{Name: "i", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ITag struct {
	*ComponentHtmlTag
}
