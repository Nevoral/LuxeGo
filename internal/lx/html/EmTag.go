package html

import (
	"LuxeGo/internal/lx"
)

// Em -
func Em(tags ...lx.Content) *EmTag {
	return &EmTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "em", Attributes: &lx.Attributes{}, Children: &tags}}
}

type EmTag struct {
	*ComponentHtmlTag
}
