package html

import (
	"LuxeGo/internal/lx"
)

//Table - 
func Table(tags ...lx.Content) *TableTag {
	return &TableTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "table", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TableTag struct {
	*ComponentTag
}
