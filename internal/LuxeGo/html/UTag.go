package html

import (
	"github.com/Nevoral/LuxeGo/internal/LuxeGo"
)

// U -
func U(tags ...LuxeGo.Content) *UTag {
	return &UTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "u", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type UTag struct {
	*ComponentHtmlTag
}
