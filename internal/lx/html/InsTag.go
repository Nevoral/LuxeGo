package html

import (
	"LuxeGo/internal/lx"
)

//Ins - 
func Ins(tags ...lx.Content) *InsTag {
	return &InsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "ins", Attributes: &lx.Attributes{}, Children: &tags}}
}

type InsTag struct {
	*ComponentHtmlTag
}

//Cite - Specifies a reference to the source of a quote or information.
func (i *InsTag) Cite(value string) *InsTag {
	i.AddAttribute("cite", value)
	return i
}

//Datetime - Specifies the date and time.
func (i *InsTag) Datetime(value string) *InsTag {
	i.AddAttribute("datetime", value)
	return i
}
