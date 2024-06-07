package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Td -
func Td(tags ...LuxeGo.Content) *TdTag {
	return &TdTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "td", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
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
