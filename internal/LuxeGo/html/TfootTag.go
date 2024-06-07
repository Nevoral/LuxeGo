package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Tfoot -
func Tfoot(tags ...LuxeGo.Content) *TfootTag {
	return &TfootTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tfoot", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TfootTag struct {
	*ComponentHtmlTag
}
