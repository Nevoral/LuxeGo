package html

import (
	"LuxeGo/internal/lx"
)

//Datalist - 
func Datalist(tags ...lx.Content) *DatalistTag {
	return &DatalistTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "datalist", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DatalistTag struct {
	*ComponentHtmlTag
}
