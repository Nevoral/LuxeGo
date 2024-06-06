package html

import (
	"LuxeGo/internal/lx"
)

// Tfoot -
func Tfoot(tags ...lx.Content) *TfootTag {
	return &TfootTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tfoot", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TfootTag struct {
	*ComponentHtmlTag
}
