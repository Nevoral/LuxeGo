package html

import (
	"github.com/Nevoral/LuxeGo"
)

// Del -
func Del(tags ...LuxeGo.Content) *DelTag {
	return &DelTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "del", Attributes: &LuxeGo.Attributes{}, Children: &tags}}
}

type DelTag struct {
	*ComponentHtmlTag
}

// Cite - Specifies a reference to the source of a quote or information.
func (d *DelTag) Cite(value string) *DelTag {
	d.AddAttribute("cite", value)
	return d
}

// Datetime - Specifies the date and time.
func (d *DelTag) Datetime(value string) *DelTag {
	d.AddAttribute("datetime", value)
	return d
}
