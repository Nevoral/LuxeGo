package html

import (
	"LuxeGo/internal/lx"
)

//Title - 
func Title(tags ...lx.Content) *TitleTag {
	return &TitleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "title", Attributes: &lx.Attributes{}, Children: &tags}}
}

type TitleTag struct {
	*ComponentHtmlTag
}
