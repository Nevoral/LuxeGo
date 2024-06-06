package html

import (
	"LuxeGo/internal/lx"
)

// Noscript -
func Noscript(tags ...lx.Content) *NoscriptTag {
	return &NoscriptTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "noscript", Attributes: &lx.Attributes{}, Children: &tags}}
}

type NoscriptTag struct {
	*ComponentHtmlTag
}
