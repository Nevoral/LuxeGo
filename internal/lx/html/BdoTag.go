package html

import (
	"LuxeGo/internal/lx"
)

// Bdo -
func Bdo(tags ...lx.Content) *BdoTag {
	return &BdoTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "bdo", Attributes: &lx.Attributes{}, Children: &tags}}
}

type BdoTag struct {
	*ComponentHtmlTag
}
