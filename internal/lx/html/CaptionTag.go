package html

import (
	"LuxeGo/internal/lx"
)

//Caption - 
func Caption(tags ...lx.Content) *CaptionTag {
	return &CaptionTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "caption", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type CaptionTag struct {
	*ComponentTag
}
