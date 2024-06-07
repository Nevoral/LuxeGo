package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Rt -
func Rt(tags ...LuxeGo.Content) *RtTag {
	return &RtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rt", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type RtTag struct {
	*ComponentHtmlTag
}
