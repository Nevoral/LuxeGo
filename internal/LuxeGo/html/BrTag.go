package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Br -
func Br() *BrTag {
	return &BrTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "br", Attributes: &LuxeGo.Attributes{}, Children: nil}}
}

type BrTag struct {
	*ComponentHtmlTag
}
