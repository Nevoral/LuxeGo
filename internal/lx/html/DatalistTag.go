package html

import (
	"LuxeGo/internal/lx"
)

//Datalist - 
func Datalist(tags ...lx.Content) *DatalistTag {
	return &DatalistTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "datalist", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DatalistTag struct {
	*ComponentTag
}
