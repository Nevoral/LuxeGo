package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Datalist -
func Datalist(tags ...LuxeGo.Content) *DatalistTag {
	return &DatalistTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "datalist", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DatalistTag struct {
	*ComponentHtmlTag
}
