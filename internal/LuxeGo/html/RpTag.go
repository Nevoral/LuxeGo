package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Rp -
func Rp(tags ...LuxeGo.Content) *RpTag {
	return &RpTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rp", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type RpTag struct {
	*ComponentHtmlTag
}
