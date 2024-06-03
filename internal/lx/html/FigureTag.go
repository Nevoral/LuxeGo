package html

import (
	"LuxeGo/internal/lx"
)

//Figure - 
func Figure(tags ...lx.Content) *FigureTag {
	return &FigureTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "figure", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type FigureTag struct {
	*ComponentTag
}
