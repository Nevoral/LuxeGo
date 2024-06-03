package html

import (
	"LuxeGo/internal/lx"
)

//Data - 
func Data(tags ...lx.Content) *DataTag {
	return &DataTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "data", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DataTag struct {
	*ComponentTag
}

//Value - Specifies the value of an <input> element.
func (d *DataTag) Value(value string) *DataTag {
	d.AddAttribute("value", value)
	return d
}
