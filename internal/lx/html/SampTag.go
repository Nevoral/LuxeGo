package html

import (
	"LuxeGo/internal/lx"
)

//Samp - 
func Samp(tags ...lx.Content) *SampTag {
	return &SampTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "samp", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SampTag struct {
	*ComponentTag
}
