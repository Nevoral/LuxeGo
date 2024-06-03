package html

import (
	"LuxeGo/internal/lx"
)

//Hr - 
func Hr() *HrTag {
	return &HrTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "hr", Attributes: &lx.Attributes{}, Children: nil}}}
}

type HrTag struct {
	*ComponentTag
}
