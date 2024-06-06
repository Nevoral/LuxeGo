package html

import (
	"LuxeGo/internal/lx"
)

// Figure -
func Figure(tags ...lx.Content) *FigureTag {
	return &FigureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figure", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FigureTag struct {
	*ComponentHtmlTag
}
