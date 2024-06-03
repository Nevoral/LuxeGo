package html

import (
	"LuxeGo/internal/lx"
)

//S - 
func S(tags ...lx.Content) *STag {
	return &STag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "s", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type STag struct {
	*ComponentTag
}
