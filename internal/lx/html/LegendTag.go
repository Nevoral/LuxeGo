package html

import (
	"LuxeGo/internal/lx"
)

//Legend - 
func Legend(tags ...lx.Content) *LegendTag {
	return &LegendTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "legend", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type LegendTag struct {
	*ComponentTag
}
