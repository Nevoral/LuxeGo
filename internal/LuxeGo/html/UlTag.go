package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Ul -
func Ul(tags ...LuxeGo.Content) *UlTag {
	return &UlTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ul", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type UlTag struct {
	*ComponentHtmlTag
}
