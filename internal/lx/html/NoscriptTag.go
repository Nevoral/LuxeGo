package html

import (
	"LuxeGo/internal/lx"
)

//Noscript - 
func Noscript(tags ...lx.Content) *NoscriptTag {
	return &NoscriptTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "noscript", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type NoscriptTag struct {
	*ComponentTag
}
