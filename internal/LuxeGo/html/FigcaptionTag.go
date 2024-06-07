package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Figcaption -
func Figcaption(tags ...LuxeGo.Content) *FigcaptionTag {
	return &FigcaptionTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "figcaption", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type FigcaptionTag struct {
	*ComponentHtmlTag
}