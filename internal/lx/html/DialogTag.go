package html

import (
	"LuxeGo/internal/lx"
)

// Dialog -
func Dialog(tags ...lx.Content) *DialogTag {
	return &DialogTag{ComponentHtmlTag: &ComponentHtmlTag{Name: "dialog", Attributes: &lx.Attributes{}, Children: &tags}}
}

type DialogTag struct {
	*ComponentHtmlTag
}

// Open -
func (d *DialogTag) Open() *DialogTag {
	d.AddAttribute("open", "")
	return d
}
