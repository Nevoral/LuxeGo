package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Div -
func Div(tags ...LuxeGo.Content) *DivTag {
	return &DivTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "div", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DivTag struct {
	*ComponentHtmlTag
}
