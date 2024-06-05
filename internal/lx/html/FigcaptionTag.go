package html

import (
	"LuxeGo/internal/lx"
)

//Figcaption - 
func Figcaption(tags ...lx.Content) *FigcaptionTag {
	return &FigcaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figcaption", Attributes: &lx.Attributes{}, Children: &tags}}
}

type FigcaptionTag struct {
	*ComponentHtmlTag
}
