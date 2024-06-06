package html

import (
	"LuxeGo/internal/lx"
)

// Strong -
func Strong(tags ...lx.Content) *StrongTag {
	return &StrongTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "strong", Attributes: &lx.Attributes{}, Children: &tags}}
}

type StrongTag struct {
	*ComponentHtmlTag
}
