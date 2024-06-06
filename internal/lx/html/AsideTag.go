package html

import (
	"LuxeGo/internal/lx"
)

// Aside -
func Aside(tags ...lx.Content) *AsideTag {
	return &AsideTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "aside", Attributes: &lx.Attributes{}, Children: &tags}}
}

type AsideTag struct {
	*ComponentHtmlTag
}
