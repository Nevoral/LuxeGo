package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Wbr -
func Wbr() *WbrTag {
	return &WbrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "wbr", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type WbrTag struct {
	*ComponentHtmlTag
}
