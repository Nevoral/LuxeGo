package html

import (
	"LuxeGo/internal/lx"
)

// Tbody -
func Tbody(tags ...lx.Content) *TbodyTag {
	return &TbodyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tbody", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TbodyTag struct {
	*ComponentHtmlTag
}
