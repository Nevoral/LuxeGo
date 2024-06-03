package html

import (
	"LuxeGo/internal/lx"
)

//Figcaption - 
func Figcaption(tags ...lx.Content) *FigcaptionTag {
	return &FigcaptionTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "figcaption", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type FigcaptionTag struct {
	*ComponentTag
}
