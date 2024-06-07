package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Figure -
func Figure(tags ...LuxeGo.Content) *FigureTag {
	return &FigureTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figure", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type FigureTag struct {
	*ComponentHtmlTag
}
