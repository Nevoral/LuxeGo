package html

import (
	"LuxeGo/internal/lx"
)

//Br - 
func Br() *BrTag {
	return &BrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "br", Attributes: &lx.Attributes{}, Children: nil}}
}

type BrTag struct {
	*ComponentHtmlTag
}
