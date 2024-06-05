package html

import (
	"LuxeGo/internal/lx"
)

//Wbr - 
func Wbr() *WbrTag {
	return &WbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "wbr", Attributes: &lx.Attributes{}, Children: nil}}
}

type WbrTag struct {
	*ComponentHtmlTag
}
