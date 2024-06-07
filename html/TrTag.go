package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Tr -
func Tr(tags ...LuxeGo.Content) *TrTag {
	return &TrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tr", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TrTag struct {
	*ComponentHtmlTag
}
