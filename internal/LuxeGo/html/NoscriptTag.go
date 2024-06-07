package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Noscript -
func Noscript(tags ...LuxeGo.Content) *NoscriptTag {
	return &NoscriptTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "noscript", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type NoscriptTag struct {
	*ComponentHtmlTag
}
