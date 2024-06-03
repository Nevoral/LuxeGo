package html

import (
	"LuxeGo/internal/lx"
)

//P - 
func P(tags ...lx.Content) *PTag {
	return &PTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "p", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type PTag struct {
	*ComponentTag
}
