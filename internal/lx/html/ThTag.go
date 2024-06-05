package html

import (
	"LuxeGo/internal/lx"
)

//Th - 
func Th(tags ...lx.Content) *ThTag {
	return &ThTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "th", Attributes: &lx.Attributes{}, Children: &tags}}
}

type ThTag struct {
	*ComponentHtmlTag
}

//Abbr - Specifies a shorter version of the content in a cell.
func (t *ThTag) Abbr(value string) *ThTag {
	t.AddAttribute("abbr", value)
	return t
}

//Colspan - Specify how many columns a cell should span in a table.
func (t *ThTag) Colspan(value string) *ThTag {
	t.AddAttribute("colspan", value)
	return t
}

//Headers - 
func (t *ThTag) Headers(value string) *ThTag {
	t.AddAttribute("headers", value)
	return t
}

//Rowspan - Specify how many rows a cell should span in a table.
func (t *ThTag) Rowspan(value string) *ThTag {
	t.AddAttribute("rowspan", value)
	return t
}

//Scope - 
func (t *ThTag) Scope(value string) *ThTag {
	t.AddAttribute("scope", value)
	return t
}
