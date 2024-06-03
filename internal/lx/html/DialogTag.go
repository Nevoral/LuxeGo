package html

import (
	"LuxeGo/internal/lx"
)

//Dialog - 
func Dialog(tags ...lx.Content) *DialogTag {
	return &DialogTag{Component: &ComponentTag{WebComponent: &lx.WebComponent{Name: "dialog", Attributes: &lx.Attributes{}, Children: &tags}}}
}

type DialogTag struct {
	*ComponentTag
}

//Open - 
func (d *DialogTag) Open() *DialogTag {
	d.AddAttribute("open", "")
	return d
}
