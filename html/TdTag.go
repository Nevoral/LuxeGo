package html

import (
	"fmt"
	"github.com/Nevoral/LuxeGo"
)

// Td -
func Td(tags ...interface{}) *TdTag {
	var children []LuxeGo.Content
	for _, tag := range tags {
		switch v := tag.(type) {
		case string:
			children = append(children, FreeStr(v))
		case LuxeGo.Content:
			children = append(children, v)
		case []LuxeGo.Content:
			children = append(children, v...)
		default:
			// Handle unexpected types if necessary
			panic(fmt.Sprintf("unexpected type %T", v))
		}
	}
	return &TdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "td", Attributes: &LuxeGo.Attributes{}, Children: &children}}
}

type TdTag struct {
	*ComponentHtmlTag
}

// Colspan - Specify how many columns a cell should span in a table.
func (t *TdTag) Colspan(value string) *TdTag {
	t.AddAttribute("colspan", value)
	return t
}

// Headers -
func (t *TdTag) Headers(value string) *TdTag {
	t.AddAttribute("headers", value)
	return t
}

// Rowspan - Specify how many rows a cell should span in a table.
func (t *TdTag) Rowspan(value string) *TdTag {
	t.AddAttribute("rowspan", value)
	return t
}
