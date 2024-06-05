package html

import (
	"LuxeGo/internal/lx"
)

//Svg - 
func Svg(tags ...lx.Content) *SvgTag {
	return &SvgTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "svg", Attributes: &lx.Attributes{}, Children: &tags}}
}

type SvgTag struct {
	*ComponentHtmlTag
}
