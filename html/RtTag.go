package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Rt -
func Rt(tags ...LuxeGo.Content) *RtTag {
	return &RtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rt", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type RtTag struct {
	*ComponentHtmlTag
}
