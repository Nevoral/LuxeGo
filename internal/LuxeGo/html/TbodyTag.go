package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Tbody -
func Tbody(tags ...LuxeGo.Content) *TbodyTag {
	return &TbodyTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "tbody", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TbodyTag struct {
	*ComponentHtmlTag
}
