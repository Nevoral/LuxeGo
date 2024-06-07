package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Data -
func Data(tags ...interface{}) *DataTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &DataTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "data", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type DataTag struct {
	*ComponentHtmlTag
}

// Value - Specifies the value of an <input> element.
func (d *DataTag) Value(value string) *DataTag {
	d.AddAttribute("value", value)
	return d
}
