package html

import (
	"LuxeGo/internal/lx"
)

//Em - 
func Em(tags ...lx.Content) *EmTag {
	return &EmTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "em", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type EmTag struct {
	*ComponentTag
}
