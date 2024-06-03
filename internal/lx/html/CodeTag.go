package html

import (
	"LuxeGo/internal/lx"
)

//Code - 
func Code(tags ...lx.Content) *CodeTag {
	return &CodeTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "code", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type CodeTag struct {
	*ComponentTag
}
