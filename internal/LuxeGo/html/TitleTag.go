package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// Title -
func Title(tags ...LuxeGo.Content) *TitleTag {
	return &TitleTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "title", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type TitleTag struct {
	*ComponentHtmlTag
}
