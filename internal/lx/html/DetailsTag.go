package html

import (
	"LuxeGo/internal/lx"
)

//Details - 
func Details(tags ...lx.Content) *DetailsTag {
	return &DetailsTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "details", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DetailsTag struct {
	*ComponentTag
}

//Open - 
func (d *DetailsTag) Open() *DetailsTag {
	d.AddAttribute("open", "")
	return d
}
