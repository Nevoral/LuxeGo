package html

import (
	"LuxeGo/internal/lx"
)

//Strong - 
func Strong(tags ...lx.Content) *StrongTag {
	return &StrongTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "strong", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type StrongTag struct {
	*ComponentTag
}
