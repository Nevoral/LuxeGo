package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Rp -
func Rp(tags ...LuxeGo.Content) *RpTag {
	return &RpTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rp", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type RpTag struct {
	*ComponentHtmlTag
}
