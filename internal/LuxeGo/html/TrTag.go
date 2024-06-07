package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Tr -
func Tr(tags ...LuxeGo.Content) *TrTag {
	return &TrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tr", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TrTag struct {
	*ComponentHtmlTag
}
