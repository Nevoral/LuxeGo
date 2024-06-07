package html

import (
	"LuxeGo/internal/LuxeGo"
)

// Data -
func Data(tags ...LuxeGo.Content) *DataTag {
	return &DataTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "data", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DataTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (d *DataTag) Value(value string) *DataTag {
	d.AddAttribute("value", value)
	return d
}
