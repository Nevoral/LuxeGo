package html

import (
	"LuxeGo/internal/lx"
)

//Details - 
func Details(tags ...lx.Content) *DetailsTag {
	return &DetailsTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "details", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DetailsTag struct {
	*ComponentHtmlTag
}

//Open - 
func (d *DetailsTag) Open() *DetailsTag {
	d.AddAttribute("open", "")
	return d
}
