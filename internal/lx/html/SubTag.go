package html

import (
	"LuxeGo/internal/lx"
)

//Sub - 
func Sub(tags ...lx.Content) *SubTag {
	return &SubTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "sub", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type SubTag struct {
	*ComponentTag
}
