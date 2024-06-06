package html

import (
	"LuxeGo/internal/lx"
)

// Dt -
func Dt(tags ...lx.Content) *DtTag {
	return &DtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dt", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DtTag struct {
	*ComponentHtmlTag
}
