package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Menu -
func Menu(tags ...LuxeGo.Content) *MenuTag {
	return &MenuTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "menu", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type MenuTag struct {
	*ComponentHtmlTag
}
