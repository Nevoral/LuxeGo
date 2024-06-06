package html

import (
	"LuxeGo/internal/lx"
)

// Small -
func Small(tags ...lx.Content) *SmallTag {
	return &SmallTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "small", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SmallTag struct {
	*ComponentHtmlTag
}
