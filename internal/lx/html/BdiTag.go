package html

import (
	"LuxeGo/internal/lx"
)

// Bdi -
func Bdi(tags ...lx.Content) *BdiTag {
	return &BdiTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "bdi", Attributes: &lx.Attributes{}, Children: &tags}}
}

type BdiTag struct {
	*ComponentHtmlTag
}
