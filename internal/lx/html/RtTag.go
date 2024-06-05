package html

import (
	"LuxeGo/internal/lx"
)

//Rt - 
func Rt(tags ...lx.Content) *RtTag {
	return &RtTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "rt", Attributes: &lx.Attributes{}, Children: &tags}}
}

type RtTag struct {
	*ComponentHtmlTag
}
