package html

import (
	"LuxeGo/internal/lx"
)

// B -
func B(tags ...lx.Content) *BTag {
	return &BTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "b", Attributes: &lx.Attributes{}, Children: &tags}}
}

type BTag struct {
	*ComponentHtmlTag
}
