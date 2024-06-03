package html

import (
	"LuxeGo/internal/lx"
)

//Svg - 
func Svg(tags ...lx.Content) *SvgTag {
	return &SvgTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "svg", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SvgTag struct {
	*ComponentTag
}
