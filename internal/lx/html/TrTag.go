package html

import (
	"LuxeGo/internal/lx"
)

// Tr -
func Tr(tags ...lx.Content) *TrTag {
	return &TrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tr", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TrTag struct {
	*ComponentHtmlTag
}
