package html

import (
	"LuxeGo/internal/lx"
)

// Mark -
func Mark(tags ...lx.Content) *MarkTag {
	return &MarkTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "mark", Attributes: &lx.Attributes{}, Children: &tags}}
}

type MarkTag struct {
	*ComponentHtmlTag
}
