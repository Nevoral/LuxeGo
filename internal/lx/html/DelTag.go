package html

import (
	"LuxeGo/internal/lx"
)

//Del - 
func Del(tags ...lx.Content) *DelTag {
	return &DelTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "del", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DelTag struct {
	*ComponentTag
}

//Cite - Specifies a reference to the source of a quote or information.
func (d *DelTag) Cite(value string) *DelTag {
	d.AddAttribute("cite", value)
	return d
}

//Datetime - Specifies the date and time.
func (d *DelTag) Datetime(value string) *DelTag {
	d.AddAttribute("datetime", value)
	return d
}
