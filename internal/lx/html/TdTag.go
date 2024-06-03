package html

import (
	"LuxeGo/internal/lx"
)

//Td - 
func Td(tags ...lx.Content) *TdTag {
	return &TdTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "td", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type TdTag struct {
	*ComponentTag
}

//Colspan - Specify how many columns a cell should span in a table.
func (t *TdTag) Colspan(value string) *TdTag {
	t.AddAttribute("colspan", value)
	return t
}

//Headers - 
func (t *TdTag) Headers(value string) *TdTag {
	t.AddAttribute("headers", value)
	return t
}

//Rowspan - Specify how many rows a cell should span in a table.
func (t *TdTag) Rowspan(value string) *TdTag {
	t.AddAttribute("rowspan", value)
	return t
}
